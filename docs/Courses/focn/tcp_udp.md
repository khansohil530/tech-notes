---
comments: false
---

# TCP and UDP

IP protocol was designed to only deliver packets over network to destination IP address, it doesn't involve anything
around:

- Guarantee on delivery of packets as they can be dropped
- No way to recover data malformed during transit.
- No way to reassemble request made of multiple packet if they reach out of order.
- Which application should receive the data as computer can have multiple applications communicating concurrently.

If applications had to deal with all these problems themselves, every app would need to reinvent the wheel.
This is why such responsibility is handed over to Transport Layer (L4), so that network layer can stay simple and fast
while applications can reuse the same functionality.
To discuss more about it, we'll look into the most popular protocol implemented on layer:
**TCP** (Transport Control Protocol) and **UDP** (User Datagram Protocol).

!!! note ""
    Other popular L4 protocol is **QUIC**, which is relatively newer and is designed for faster, more secure web performance.  

## UDP

UDP is the simplest transport protocol on Internet as it doesn't provide any fancy features like TCP's retries, 
ordering, congestion control or connection setup.
Its core responsibility only involves 3 features:

- Deliver messages from one application to another using **ports**
- Each UDP packet is a complete message (also called datagram), there's no segmentation.
- Use Checksum to detect corrupt data (corrupted packets are discarded)

This simplicity makes UDP an unreliable protocol but much faster and simpler protocol than TCP.
This tradeoff is required across various systems which prioritizes low latency over data loss. For example,
real-time systems live-streaming, voice calls can afford to loss few packets without causing failure, online
games can also afford to loss some data since each next message would update missed information anyway,
and protocols like DNS which needs faster delivery of small message.
Another advantage of simplicity is it's easier to extend on top of it, you can decide what matters for your application
and built it on top of UDP instead of subscribing to all the overhead from TCP stack.

UDP unit of data is called **datagram**, which encompasses of header and data section. 

```text
        ┌─────────────────────────────────────────────────────────────┐
        │                    IP Headers                               │
        │                    20-60 bytes                              │
        └─────────────────────────────────────────────────────────────┘
┌───────┌─────────────────────────────────────────────────────────────┐
│       │          Source Port        │         Destination Port      │     
8 bytes │           16 b              │            16 b               │     
UDP     ├─────────────────────────────────────────────────────────────┤
Headers │        Length               │           Checksum            │
│       │            16 b             │              16 b             │
└───────└─────────────────────────────────────────────────────────────┘
        ┌─────────────────────────────────────────────────────────────┐
        │                    Data                                     │
        │                                                             │
        └─────────────────────────────────────────────────────────────┘
```

- Since source and destination ports are of 16-bits each, computer can use almost 65536 ($2^{16}$) ports.
- Length indicates size of datagram, i.e. $2^{16} - 64 = 8kb$ of application data can be sent over single datagram. 

??? note "Demo"
    To demo UDP, we can build UDP servers which listens on specific port and processes incoming datagrams at that
    port. But when building this using a low level language like C, you’ve to handle small implementation 
    details abstracted by libraries when using high level language, details like creating a 
    listening socket, gathering result in memory buffer, etc.

    === "JS Server"
    
        ```js
        import dgram from 'dgram'
         
        const socket = dgram.createSocket('udp4')
        socket.bind(5500, '127.0.0.1')
        socket.on('message', (msg, info) => { 
            console.log(`My server got a datagram: ${msg}, from : ${info.address}:${info.port}`)
        })
        ```
    
    === "C Server"
    
        ```c
        #include <stdio.h>
        #include <stdlib.h>
        #include <string.h>
        #include <sys/socket.h>
        #include <sys/types.h>
        #include <netinet/in.h>
        #include <arpa/inet.h>
        
        int main(int argc, char **argv){
        
          int port = 5501;
          int sockfd;
          struct sockaddr_in myaddr, remoteAddr;
          char buffer[1024];
          socklen_t addr_size;
        
          sockfd = socket(AF_INET, SOCK_DGRAM, 0); // creates an UDP IPv4 socket
        
          memset(&myaddr, '\0', sizeof(myaddr)); // allocate memory to myaddr variable
          myaddr.sin_family = AF_INET;
          myaddr.sin_port = htons(port);
          myaddr.sin_addr.s_addr = inet_addr("127.0.0.1");
        
          bind(sockfd, (struct sockaddr*)&myaddr, sizeof(myaddr)); // bind this socket socket to my address
          addr_size = sizeof(remoteAddr);
          recvfrom(sockfd, buffer, 1024, 0, (struct sockaddr*)& remoteAddr, &addr_size); // receive data on given socket FD and put it
                                               // in given buffer area of size. Also save the sender info in remoteAddr variable of given size. 
          printf("got data from %s ", buffer);
          return 0;
        }
        ```
    To ping the server you can use `nc` in linux. For example, `nc -u 127.0.0.1:5500`

## TCP

TCP is also a transport layer protocol which provides all the features of UDP plus features which are core to most of
the applications on internet, like 

- Reliably transferring packets
- Reassembling request by correctly ordering packets even if they reach out of order.
- Ensuring both network and receiver aren't overwhelmed by its communication (1).
  {.annotate}

    1. congestion and flow control


??? note "TCP vs UDP"
    You might think that both TCP and UDP are opposites of each other, TCP is reliable slow while UDP is unreliable 
    fast but this is misleading way to think. Instead, TCP and UDP both solve the transport problem. TCP additionally
    implements features required by apps around the internet out of box, making it closer to application layer while
    UDP just ensure data can be communicated b/w application making it closer to IP layer.
    
    TCP is also said to be slower protocol compared to UDP, which is somewhat true but only because UDP ignores congestion
    control allowing it to borrow additional bandwidth from network, but this can be hazardous and collapse entire 
    network if not managed properly.
    TCP, on the other hand is optimized for fairness and stability making sure both network and receiver aren't overwhelmed 
    by its communication rate. And it can be extremely fast on clean networks like private networks where you don't have
    to worry about congestion.  

To provide features mentioned above, TCP needs to store information on hosts. For example, it needs to maintain state on
how many packets will it receive for a request, how many packets have been received, what's the window size of packets
it can send to avoid any congestion, etc. This is why, it's called a stateful protocol and its initialized when 
establishing a connection. 

### Establish Connection 

The process of establishing a connection is also previously discussed in [OS](../foos/socket.md#establishing-a-connection)
notes, but there we primarily discussed it w.r.t to OS. Here we'll look into what information is exchanged during the 
3-way handshake and where is this information required.
??? note "Connection in OS briefly"
    Within operating system, this connection is a file descriptor identified from the combined hash of 
    `(srcIP, srcPort, destIP, destPort)` fields. This file descriptor is responsible for everything related to this
    connection, like storing state, sending/receiving data.

Before hosts can start sending/receiving data over TCP, they need to establish a connection to exchange information 
required by TCP. This is done using 3-way handshake as follows:

- **`Client -SYN-> Server`**: Client initiates the connection by sending an `SYN` request to synchronize with server,
sharing its **ISN** (Initial Sequence Number) and other TCP options like MSS, window scaling, etc.
- **`Client <-SYN/ACK- Server`**: Server on listening to this incoming `SYN` request creates a half-open connection 
to save client's TCP options and move it in its `SYN Queue`. After which, it replies back to client with `SYN/ACK`
to acknowledge client's synchronization request, along with sending its own TCP option (like Server ISN) for
synchronization . This message is tagged with acknowledgement number -> `C_ISN+1` (short for Client ISN).
- **`Client -ACK-> Server`**: Client on receiving this stores server's TCP options in its socket and replies with `ACK`
to acknowledge server's synchronization request. This message is similarly tagged with acknowledgement number ->
`S_ISN+1` (short for Server ISN). When Server receives this `ACK`, it creates full connection from the half-opened
connection in `SYN Queue` and moves it to `ACCEPT Queue`, marking connection ready for communication.

??? note "ISN & MSS"
    **ISN (Initial Sequence Number)** is the metadata used to make TCP reliable. It's the starting byte number which is 
    tagged to every segment. For every segment sent, the sequence number of the next segment increases by the number of
    data bytes sent in the previous segment. **MSS (Maximum Segment Size)** on the other hand is self-explanatory by its name. 
    
    For example, with ISN=1000 and MSS=500, sender will attach following sequence number for next 3 segments: 
    `(1000, 1500, 2000)`. Notice, `SYN`, `ACK` requests consumes 1 sequence number  even though they carry no data.
    It's a common misconception that segment increments the sequence number by 1 which is wrong, TCP only knows byte 
    offsets in a stream.

This ensures:

- Both hosts agree on starting sequence numbers so data can be ordered, tracked, and retransmitted correctly.
- Confirm both side can send and receive packets.
- New Sequence number ensures delayed or duplicate packets from earlier sessions aren’t mistaken for valid data.
- Negotiate other important connection parameters (like MSS, window scaling) before data starts flowing.

??? note "Why not 2-way handshake?"
    2 steps wouldn’t confirm that both sides can receive data. If last `ACK` isn't received, server could confirm that
    its `SYN/ACK` hadn't reached client.

### Sending and receiving data

Once the connection is established, both parties can send and receive data using `read()`/`write()` on the connection
socket. When application writes data over socket, TCP breaks the request into chunks (1), adds TCP headers to each segment
and places them in send buffer where kernel sends them over to network (2).
{.annotate}

1. based on factors like MSS, available congestion window. 
2. controlled by **Flow** and**Congestion control**.

Over the top this is how TCP sends data (very similar to UDP) but reality isn't as simple. TCP has to additionally 
provide two important guarantees which is reason its adopted across majority of internet: 

- Reliable delivery, as long as two devices can be connected over a network it'll make sure every byte of data 
  transmitted is received by other party.
- Control over the rate of transfer so that it:
      - doesn't overwhelm receiver system (Flow control)
      - doesn't cause congestion in network (Congestion control)



#### TCP's Reliability 

To make TCP delivery reliable, the protocol requires receiver to confirms the number of bytes received. This is done
by sending back `ACK` message tagged with acknowledgement number which indicates successfully received all bytes up to 
it and bytes which its expecting next (basically everything before it is assumed delivered correctly).
For example, sender sends 3 segments with sequence numbers (`1000, 1500, 2000`). If receiver `ACK`s `2000` means receiver
has received all the bytes until `2000` and is expecting data after `2000`.   

**Lost Segment**: Receiver only advances ACKs to the point before which all bytes have been received in order. So in any
case where a segment is missing, receiver will repeatedly ACK the same point (**duplicate ACKs**) to which it has 
received all the data. For example, sender sends (`1000, 1500, 2000`) segments but `1500` is lost on the way and 
receiver only gets (`1000-1499` and `2000-2499`) segments. So the receiver will only `ACK=1500` and keep repeating it
until it receives the expected (`1500`) sequenced segment. On the sender side, missing data is detected in two ways:

1. **Duplicate ACKs**: If the sender receives 3 duplicate ACKs then it assumes the expected segment was lost, but later data
   is arriving. So the sender immediately retransmits the missing segment by adjusting its **congestion window** 
   ([**fast recovery**](#congestion-control)). This is faster and efficient, as there's no waiting for a timeout. 
2. **Timeout (RTO)**: If sender doesn't get any ACK for a segment within **Retransmission Timeout** (RTO) (1) then it 
   assumes something when badly (2), so it retransmits the unacknowledged segment and aggressively resets its congestion
   window ([**slow start**](#congestion-control)). Timeouts are slower but handle cases multiple packets are lost or no duplicate ACKs can be 
   generated without worsening congestion. 
   {.annotate}
    
     1. $currentTime − sendTime(oldestUnackedByte) > RTO$
     2. The segment was lost, subsequent segments were also lost, or ACKs themselves were lost.


However, `ACK`ing individual segments seems hardly efficient as it'll waste network bandwidth by flooding it with
small packets while also increasing latency, as sender would now have to wait for `ACK` on previously sent 
segment before sending the next. So sending multiple segments should be the right choice, but how many segments? Apart
from sender host, this depends on two external factors:

1. Receiver, how many segments can receiver host process at a time without overwhelming the system? This is controlled 
   by **Flow Control**.
2. Network, how many segments can network handle at a time without causing congestion? This is controlled by 
   **Congestion Control**.
---
#### Flow Control

Flow control ensures the sender only transmits as much data as the receiver can handle at any given time. The core idea
of its implementation is to use a sliding window mechanism based on the receiver’s available buffer space. 

Receiver application reads data from this buffer at their own pace. To information about currently available space in buffer,
receiver sends the value (termed as **receive window** (`rwnd` (1))) along with every TCP ACK.
{.annotate}

1. how many more bytes the receiver can currently accept.

The sender may have multiple bytes "in flight" which haven't been ACKed. Before sending any additional data, TCP
always checks if bytes in flight $\leq rwnd$. If the receiver advertises a smaller window, the sender slows down.
If it advertises a larger window, the sender can speed up.

??? note "Zero window"
    If the receiver’s buffer is full, it advertises $rwnd = 0$, which stops the sender from sending any data. To resume
    transmission, sender periodically sends a window probe to check if receiver has freed up and only resumes once
    receiver advertises a non-zero $rwnd$.    

With this control, applications can communicate over TCP 

- reliably across devices with different speed and memory.
- without compromising receiver by overflowing their buffer.
- keep stable long-lived connections over file transfers, streaming, etc. 

However, this only considers receiver device when transmitting multiple segments. It's equally important if not more, 
to consider intermediate devices used across network for transmitting your data which also have limited memory and 
compute, and if overwhelmed can halt the entire network (**Congestion Collapse**) impacting all the users of that network.
To take this into account, TCP uses **Congestion Control**.  

#### Congestion Control

Congestion control allows TCP to prevent the network (routers and links) from being overloaded, which would otherwise
cause packet loss, high delay, and collapse of throughput. It's controlled by maintaining congestion window (`cwnd`)
by the sender which limits how much data can be "in flight" at a time. This makes the actual window size: 
Bytes in flight $≤ min(cwnd, rwnd)$ since sender need to consider both receiver and network. 
But how does sender actually know if it's causing a congestion in network? It's indirectly inferred from following 
signals: 

- Classical implementation was Packet loss detected from RTO or duplicate ACKs.
- Modern variant monitors increasing RTT (Round Trip Time) to lower its `cwnd`

In working, Congestion control happens as follows:

1. **Slow Start**: Sender starts with a small `cwnd` (usually 1 MSS) and increments its exponentially for each ACK 
   (roughly doubles every RTT). This allows the sender to rapidly discover available bandwidth, and it continues until 
   sender either detects packet loss or reaches **`ssthresh`** (Slow Start Threshold). 
2. **Congestion Avoidance**: Once past `ssthresh`, `cwnd` grows linearly (~ +1 MSS per RTT) allowing to probe available 
   bandwidth more cautiously. This helps avoid triggering any congestion while still increasing throughput.
3. **Detecting Congestion**: Sender infers congestion from two signals:
    - _Timeout_: If sender doesn't receive an ACK for a transmitted segment before the **RTO** (retransmission timeout)
      expires, it picks this signal as a major congestion problem and cuts `cwnd` aggressively by setting 
      $ssthresh = cwnd / 2$ and `cwnd = 1 MSS` entering slow start.
    - _Triple duplicate ACKs_: When sender receives 3 ACK with same acknowledgement numbers, it picks the signal as mild
      congestion indicating single packet loss, to which it immediately retransmits the lost packet without waiting on 
      timeout. Since the loss still implies congestion, it sets $ssthresh = cwnd / 2$, reduces `cwnd` (roughly by half)
      and moves to fast recovery phase (avoid unnecessary slow start).
4. **Fast Recovery**: Lost segment is retransmitted without waiting for timeout, `cwnd` is reduced (still relatively
   large) and moves sender to congestion avoidance.

??? note "Why timeout is a signal for severe congestion?"
    Timeout means either the data packet was lost, and the retransmitted packet or its ACK was also delayed or lost. 
    That usually means queues are overflowing somewhere in the network. When comparing with duplicate ACKs which suggests 
    that the network is congested, but still delivering traffic. 
    
    This timeout is measured using -> $RTO ~=$ Smoothed $RTT + 4 × RTT$ variation and RTT is continously measured for
    each ACKs. 


??? note "Congestion Notification"
    Modern routers can explicitly notify congestion to sender using **ECN** (Explicit Congestion Notification),
    but this wasn't implemented in classical TCP since it required additional network support. Today, it's enabled
    by negotiation b/w sender and receiver during the TCP handshake. If either side doesn’t support it, ECN is disabled.
    If enabled, sender sets ECN to `10/01` in the IP header, which signals routers that they're allowed to mark the 
    packet if congested. 
    
    When a router’s queue starts filling, it set ECN IP header to `11` instead of dropping them to mark congestion 
    experienced. When receiver reads this, it sets **ECE** (ECN Echo) TCP header in its ACK to signal sender about 
    congestion. The sender upon receiving set ECE signal takes appropriate steps for congestion control.
    And when sending its next segments, it marks their **CWR** (Congestion Window Reduced) TCP header to signal that it
    had reduced window size inorder to prevent repeated reactions to the same congestion event.


??? note "Small Packet Problem: Nagel's Algorithm"
    To reduces the number of small TCP segments on the network and combine small writes into larger segments, kernels
    implements Nagel's Algorithm. Which suggests: "If sender has unacknowledged data in flight, buffer small writes and 
    only send them unless an ACK arrives or enough data accumulates to fill an MSS". This lowers packet overhead and
    router processing load by limiting tiny packets, and as such reduces network congestion.
    
    However, using this algorithm can introduce artificial latency due to delayed ACK problem, where Nagle waits for
    ACKs and Delayed ACKs wait before sending ACKs. This is why, interactive apps like SSH, games, RPC which prioritizes
    low latency over anything disable Nagle (using `TCP_NODELAY`) but is preferable for bulk transfer apps like HTTP, 
    file transfer which can incur small delay over network degradation.

### Terminating Connection

When both hosts are done exchanging data with each other, they need to terminate the connection and free up the occupied
resources by it. This is done over a 4-way handshake so that each side can express their intent for connection 
termination while allowing opposite party to acknowledge that. The 4-way termination is done over following steps:

- **Client -FIN-> Server**: Client initiates its intent to terminate the connection by sending a `FIN` segment 
  (short for finish). On the client machine, this moves the socket into `FIN_WAIT_1` state.  
- **Client <-ACK- Server**: Server upon receiving this replies with `ACK` acknowledging the request. This moves the 
  server socket into `CLOSE_WAIT` and client socket in `FIN_WAIT_2` state. At this point, server can still send data to 
  client.
- **Client <-FIN- Server**: When server finishes sending its remaining data, it sends `FIN` and enters `LAST_ACK` state.
- **Client -ACK-> Server**: Client replies this with `ACK` and enters `TIME_WAIT` state, while server enters `CLOSED` 
  state. The `TIME_WAIT` state is moved to `CLOSED` only after the duration of `TIME_WAIT` expires (which is
  $2\times MSL$ (Maximum Segment Lifetime)).  This is done to prevent any (if possible) old duplicate segments from
  being misinterpreted by a new connection which occupies same connection IP+port hash.

Both connection sockets are only cleaned up after they enter `CLOSED` state. The machine that initiates the termination
is usually the one that ends up in TIME_WAIT state, because it has to send the last ACK and make sure that no
further segments are received on same socket. Both these side are also named based on which side initiates the 
termination,i.e. Active Closer for side which first sends the FIN and Passive Closer for side that receives first FIN.


### TCP Segment Anatomy

Since TCP uses much more complex mechanism of transport than UDP, its unit of data (called **segment**) is also much 
larger (20-60 bytes). Check below to understand what different metadata is associated with mechanisms we learned so far:

```text
         ┌──────────────────────────────────────────────────────────────┐
         │                    IP Headers                                │
         │                    20–60 bytes                               │
         └──────────────────────────────────────────────────────────────┘
┌────────┌──────────────────────────────────────────────────────────────┐
│        │          Source Port        │         Destination Port       │
│        │             16 b            │             16 b               │
│        ├──────────────────────────────────────────────────────────────┤
│        │                    Sequence Number                           │
│        │                         32 b                                 │
│        ├──────────────────────────────────────────────────────────────┤
│        │                 Acknowledgment Number                        │
│  TCP   │                         32 b                                 │
│Header  ├──────────────────────────────────────────────────────────────┤
│20–60 b │ Data Off │ Res │ NS │ CWR │ ECE │ URG │ ACK │ PSH │ RST │ FIN│
│        │  4 b     │3 b  │1 b │ 1 b │ 1 b │ 1 b │ 1 b │ 1 b │ 1 b │ 1 b│
│        ├──────────────────────────────────────────────────────────────┤
│        │                      Window Size                             │
│        │                         16 b                                 │
│        ├──────────────────────────────────────────────────────────────┤
│        │        Checksum             │        Urgent Pointer          │
│        │           16 b              │             16 b               │
│        ├──────────────────────────────────────────────────────────────┤
│        │                  Options + Padding                           │
│        │                     0–40 bytes                               │
└────────└──────────────────────────────────────────────────────────────┘
         ┌──────────────────────────────────────────────────────────────┐
         │                    Data                                      │
         │                                                              │
         └──────────────────────────────────────────────────────────────┘
```


- **Sequence Number**: Byte number of the first data byte in this segment, used for ordering and reliability.
- **Acknowledgment Number**: Next byte expected from the peer, while confirming receipt of all prior bytes.
- **Data Offset**: Length of the TCP header to find where data begins.
- **Reserved**: reserved future protocol extensions, unused till now.
- **Flags**: 9 control bits to manage connection state and behavior:
    - **SYN**: start connection 
    - **FIN**: close connection 
    - **RST**: abort connection 
    - **ACK**: acknowledgment valid 
    - **PSH**: push data to application 
    - **URG**: urgent pointer valid 
    - **ECE**: marked by receiver to signal sender about congestion. 
    - **CWR**: marked by sender to signal routers/receiver about reduced `cwnd`
    - **NS**: Nonce sum, deprecated now, used to secure ECN by verifying congestion signals, using random nonces.
- **Window Size**: Receiver’s advertised buffer space, not exactly `rcwd` due to window scaling.
- **Checksum**: Error detection for TCP header + data.
- **Urgent Pointer**: Indicates end of urgent data, used only if URG flag is set.
- **Options**: Optional features like MSS, window scaling, timestamps, SACK.

??? annotate "Window Scaling"
    Window Size field is 16 bits which limits TCP to $2^{16} − 1 = 65,535$ bytes ($64$KB) which was fine for early 
    internet. But on high-bandwidth, high-latency links, this size limit of TCP would severally limit the connection
    throughput (1). To increase the actual window size while keeping it backward compatible, TCP introduced 
    **Window Scaling**. During the 3-way handshake, both sides advertise a **scale factor**, which is used to calculate
    the effective window size using -> $EffectiveWindow = advertisedWindow × 2^{scale}$. 
1. $Throughput \approx WindowSize / RTT$