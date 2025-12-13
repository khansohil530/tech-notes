---
comments: false
---

# Internet Protocol 

IP is a L3 (network) protocol whose main responsibility is to help data move towards its destination host through
a graph of hosts (or network). To understand IP, we need to understand **IP Address** which is used to uniquely 
identify a host while also providing efficient routing. 

## IP Address

IP Address helps solve IDing hosts and inefficient routing issues by organizing the address into `xxx.xxx.xxx.xxx` 
format, where each number can range from `0-255` (essentially using 4 bytes, i.e. IPv4).  To help routing data faster,
this continuous address is divided smaller, manageable segments, called **subnets** so that you can just traverse these
subnets to avoid the majority of network graph while moving towards destined IP address. This is done by dividing the
address into two parts:

1. Network portion, which uses some starting bits in the address to identify the **subnet**.
2. Host portion, which uses the remaining bits to identify the host in subnet. 

This split is represented using **CIDR** (Classless Inter-Domain Routing) notation, which defines subnet using 
`a.b.c.d/x` format. Here `x` bits in `a.b.c.d` belongs to network and rest to identify hosts within subnet. But this
notation is for human-readability, internally computers store this information separately, where network = `a.b.c.d` 
and `/x` is stored using subnet mask formed using 1s in first `x` bits of 32 bits. This way computers can AND this mask
with any IP address to determine if the address belongs to network.
For example, `192.168.254.9/24` indicates 

- first 24 bits belong to subnet -> Network = `192.168.254.0`
  - Subnet Mask = `11111111 11111111 11111111 00000000` -> `255.255.255.0`  
- Reserved Broadcast address(1) -> `192.168.1.255`
  {.annotate}

    1. used as signal for broadcasting the request on subnet, like ARP Request discussed 
       [below](#intra-network-communication-using-arp)

- Hosts can use IP from `192.168.254.1` -> `192.168.254.254`.

??? note "Example, if an IP belongs to subnet"
    ```text
    Subnet = 10.1.32.0/20 -> mask = 11111111 11111111 11110000 00000000
    IncomingIP = 10.1.38.77 ->  AND 00001010 00000001 00100110 01001101
                                    -----------------------------------
                                    00001010 00000001 00100000 00000000
      same as network address    ->     10         1        32     0 
      IP belongs to subnet
    ```


Creating subnets simplifies a lot of networking:

- Routers use subnet information to make faster decisions, forwarding packets only to the necessary subnet, not the
  whole internet. Now, routers only need to know how to reach specific subnets, making routing tables smaller and more
  efficient.
- Subnets can be aggregated into larger blocks (supernets), further reducing routing table complexity for internet routers. 
- Subnets organize devices by location, department, or function, allowing for targeted routing policies and easier
  management.

??? note "How to Divide Subnet into Smaller Subnets?"
    If you want to divide this subnet into further smaller subnets for better organization, you can extend the bits used by
    network. For example,
    ```
    /24 -> /26
    
                         192.168.254.0/26
    192.168.254.9/24 ->  192.168.254.64/26
                         192.168.254.128/26
                         192.168.254.192/26
    ```


With this information, we can start explaining the journey of IP packets when moving through network.
Once the IP packets are assembled at Network Layer, sender checks if the IP address belong in its subnet or not.

- If the destination IP belong same subnet, it'll follow the 
  [Intra-network Communication](#intra-network-communication-using-arp) pathway.
- If the destination IP is outside its subnet, it'll follow the
  [Inter-network Communication](#inter-network-communication-using-ip) pathway.

## Intra-network communication using ARP

When destination belongs within same subnet, the communication entirely happens within LAN using just frames at 
Data Link Layer (L2). To develop frames, you'll need receiver's **MAC address** (1) to identify the host reliably, but
sender only know IP address of receiver (it got through DNS resolution). To get receiver's MAC address, it uses
**ARP** (Address Resolution Protocol). To understand ARP, let's take following example. Device A wants to send an 
IP packet to Device B within same network.
{.annotate}

1. which is a permanent address associated to your network hardware on device (like NIC)

```mermaid
--8<-- "docs/Courses/focn/diagram/arp.mmd"
```

??? note "ARP Poisoning"
    Since ARP request is broadcasted to every device in a subnet, bad host can respond with their own MAC
    address before the real host which leads to redirect of traffic to bad players. This attack is called ARP Poisoning.
    Modern systems uses strategies like dynamic ARP inspection, switch port security or keeping static ARP entries in 
    critical systems to avoid such attacks.

This entire networking infrastructure is handled by **Switch** (1) which stores each device's MAC address, and map it to 
respective physical port to forward the frames and **Ethernet Cables** which moves the frames from one device to 
another. In case of Wireless devices (like Wi-Fi), devices communicate through Access Point (AP) which then connects to
a switch.
{.annotate}

1. different from Router

This model of communication is simple and works fine for small LANs, but scale for WANs due to congestion caused by 
broadcasts. It's also easy to spoof (ARP poisoning), which makes it vulnerable for communicating with unknown device 
on internet. That's why when a device wants to communicate with hosts outside the LAN, it's done using a different
protocol called **IP**, which is used for inter-network communication. 


## Inter-network communication using IP

When destination belongs outside sender's subnet, it'll send the request to a special device within its subnet
called **Default gateway** (usually your router) (1). **Routers** are another core infrastructure in networking, whose main
responsibility is forwarding packets into right direction. This is done by looking at the destination IP address of
packet (L3, need to strip L2 headers) and searching it against **Routing Table**(2) to find best match(3) which will be 
either the next hop of packet towards destination or the destination subnet itself.
{.annotate}

1. It's named gateway literally based on its functioning, i.e. it acts as a gateway for subnet devices to help them 
   communicate with devices on other subnets.
2. data structure maintained across all routers, which stores routing information like destination network,	interface and
   next hop in a table.
3. the longest prefix match.

!!! note ""
    A key entry in routing table is `0.0.0.0/0`, this is used when router doesn’t recognize the destination network, 
    so it send the packet to its ISP router (default route when no match is found).

This terminology of "default gateway" is relative to each device, where if the device (router/computer) doesn't know
where to send the request next, it'll send it to default gateway device upstream having knowledge of wider network.
For PCs/phones this is default gateway/router, for your router this is default route to ISP's router in routing table,
for ISP's router this is some upstream router (like (1)). This way you can divide the internet across different levels 
of routers and every router always know where to send the packet next without knowing the whole internet. 
{.annotate}

1. Regional / National ISP routers for connecting local ISP to larger ISPs which defaults to tier-1 ISPs which forms backbone of 
   internet like Tata, AT&T.

But before your router sends your IP packet outside subnet, it performs another key operation called **NAT**.

### NAT

IPv4 address only allows the system to address ~$4.3$ billion ($2^{32}$) hosts which is insufficient at
this time due to high usage of devices like PCs, phones, servers, IoT. To solve this issue, there were two 
proposals:

- Use IPv6 address which can support upto ~$340$ undecillion ($2^{128}$) devices.
- Hiding many private devices behind one public IPv4 address using **NAT** (**Network Address Translation**). 

??? note "IPv6 adoption"
    IPv6 was the right solution, but it came too late for the crisis. IPv4 exhaustion became serious in the 
    mid-1990s while IPv6 specs were finalized around 1998. The immediate response to avert the crisis was NAT, which
    could be deployed immediately without waiting for support from OS, routers, applications. It worked with existing
    IPv4 and by just changing single router you could fix many devices. Since NAT was cheaper and faster backward-compatible
    option, it was adopted immediately. While IPv6 is still being implemented slowly even after two decades (since it's
    the right choice technically) but its adoption was delayed due to huge success of NAT.

NAT is implemented on routers, and it rewrites IP addresses (and usually ports) in packets so that multiple 
private devices can share one or a few public IP addresses. To differentiate b/w each of its private host, it uses
**NAT Table** which maps devices privateIP (and port) to the publicIP (and port). This entries in table are created only
when private host initiates an internet request, and usually times out when unused for some time.  For example,

```mermaid
--8<-- "docs/Courses/focn/diagram/pat.mmd"
```

The above example is one type of NAT, also known as **PAT** (Port Address Translation) since it also uses ports to map 
requests. Since single router can have $1000$s of ports, PAT allows the router to translate many devices making it most
commonly used. Other commonly used NAT implementation is **Static NAT** which maps public to private IP (1:1), usually used for 
servers.

With NAT, your devices could now be assigned IP address which would never be used on internet. This means, we can 
use same IP address for two devices on different private network without conflict. This idea leads to reserving 
chunks of specific IP address (1) for internal usage in private network (RFC 1918). Additionally, to make it more secure, 
routers would immediately drop any packet using private IP address.
{.annotate}

1. 10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16

NAT also have many other applications apart from public to private translations. Some of them are

- Port forwarding which can be used to expose local webservers publicly
- Load balancing at Layer 4, used by proxies like HAProxy which replaces the destIP to one of the servers grouped to 
  this destIP (*reverse proxying*).

---

With this, we can explain how out inter-network communication moves through different devices using IP protocol. 

??? note "End-To-End Overview of Internet Request"
    ```mermaid
    --8<-- "docs/Courses/focn/diagram/internet_req.mmd"
    ```

## Anatomy of IP Packet

IP Packet (IPv4 specifically) is the unit of data used for networking at Network Layer (L3) as per IP specification 
(also known as PDU (1)). Until now, we've only discussed its routing aspect through IP address, but there are few other
required fields defined by protocol. To understand each of them, let's go through the anatomy of IP packets.
{.annotate}

1. Protocol Data Units

An IP packet is continuous block of bytes organized into two major section:

1. Header, which stores metadata about IP Packet using first 20-60 bytes of packet. 
2. Body, which stores the actual segmented data (datagrams/segments) from Transport layer.

```text
            ┌──────────────4 bytes────────────────────────────────────────┐
┌───────────┌─────────────────────────────────────────────────────────────┐
│           │Version│  IHL  │  DSCP │ ECN │         Total Length          │
│           │ 4 b   │  4 b  │  6 b  │ 2 b │            16 b               │
│           ├─────────────────────────────────────────────────────────────┤
20 bytes    │        Identification       │Flags│   Fragment Offset       │
Required    │            16 b             │ 3 b │        13 b             │
│           ├─────────────────────────────────────────────────────────────┤
│           │  TTL          │   Protocol  │        Header Checksum        │
│           │ 8 b           │     8 b     │            16 b               │
│           ├─────────────────────────────────────────────────────────────┤
│           │                     Source IP Address                       │
│           │                         32 b                                │
│           ├─────────────────────────────────────────────────────────────┤
│           │                  Destination IP Address                     │
└───────────│                         32 b                                │
┌───────────├─────────────────────────────────────────────────────────────┤
40 bytes    │                    Options (if any) + Padding               │
Optional    │                  Variable (0–40 bytes)                      │
└───────────└─────────────────────────────────────────────────────────────┘
            ┌─────────────────────────────────────────────────────────────┐
            │                    Data                                     │
            │                                                             │
            └─────────────────────────────────────────────────────────────┘
```

- **Version**: First 4 bits are used to determine the Version of packet, if its IPv4 or IPv6.
- **IHL** (Internet Header Length): defines length of header in 4 byte lines. Default is 5 which reads 20 ($5 \times 4$) bits 
  for required headers.
- **DSCP** (Differentiated Service Code Point): defines the forwarding priority of packets. This is required because
  not all packets are equally important. For example, VoIP (Voice Over IP) packets are given more priority since they're
  used for voice call which is sensitive to delay. With the help of DSCP, you can mark your packets at the edge 
  (host, router, firewall) based on your policy and prioritize traffic as per it.
- **ECN** (Explicit Congestion Notification): used by L3 devices to notify upper layers if the network is facing
  congestion. We'll look into it when talking about Congestion Control in TCP protocol (L4).
- **Total Length**: Total length of packet (header+body) in bytes. Since it uses 16 bits, packets size range can range
  from $0$ - $2^{16}$ bytes (~64kB).
- Next line of header is used by **IP fragmentation**.

    ??? annotate "IP Fragmentation"
        Every network link is associate with an **MTU** (Maximum Transmission Unit) (1). A host may send a perfectly 
        valid IP packet that fits its outgoing interface (using TCP segmentation (2)) but might be too large for a 
        downstream link. This is why routers can't assume same MTU everywhere and that the sender knows the path MTU. 
        To help with this, **Fragmentation** was introduced which allow routers to split large IP packet into smaller 
        pieces (fragments) so it can traverse network links with a smaller MTU, with the fragments later reassembled at 
        the destination host using respective IP headers.
        
        - **Identification**: ID to reassemble fragments belonging to the same original IP packet.
        - **Flags**: three 1-bit flags,
            - bit 0: reserved for future compatability, must be set to 0.
            - bit 1: DF (Don't Fragment) flag, set this bit to tell routers to not fragment the packet and drop if it's too big. 
              This is used in Path MTU Discovery.
            - bit 2: MF (More Fragment) flag, set this bit to tell receiver about more incoming packets. Last fragment will have 
              MF = 0
        - **Fragment Offset**: Offset where this fragmented packet belongs in whole packet
      1. size of largest data packet link can handle without fragmentation
      2. This will be discussed more in TCP chapter.

- **TTL** (Time to Live): defines maximum number of hops a packet can survive before its discarded. This is done to 
  avoid packet roaming around the network infinitely (due to cycles). At each hop, the router must decrement this field
  and when any router encounters an IP packet where TTL=0, it'll discard the packet and return an ICMP(1) message stating 
  the reason back to client. 
  {.annotate}
    1. disused [below](#icmp)

- **Protocol**: identifies which upper-layer protocol (like ICMP, TCP, UDP) should receive the payload. Without this 
  field, the receiver wouldn’t know how to interpret the payload. 
- **Header Checksum**: protects header fields from corruption and tampering. This is required because headers like TTL
  are modified by routers inflight.   
- **Options**: providing future compatability without redefining the protocol. Sender can use this space to attach extra
  instructions or metadata to an IP packet.

## ICMP

ICMP (Internet Control Message Protocol) is another important Network Layer (L3) protocol which is used by network 
devices to communicate errors/operational information about IP packet delivery. ICMP was designed because IP doesn't define
a way to handle errors like delivery failures, routing problems or reachability and timing info back to sender. ICMP
message have the following format:

```text
┌─────────────────────────────────────────────────────────────┐
│     Type 8 b      │     Code 8 b     │   Checksum 16 b      │
├─────────────────────────────────────────────────────────────┤
│                Message-Specific Data 32 b                   │
├─────────────────────────────────────────────────────────────┤
│        Original IP Header + First 8 Bytes of Payload        │
│                 for reporting as logs                       │
└─────────────────────────────────────────────────────────────┘

```

- **Type**: defines the general category of ICMP message. For example, 8 - Echo Request, 0 - Echo Reply, 3 - Destination
  Unreachable, 11 - Time Exceeded.
- **Code**: adds more specific detail about the Type. For example, Type 3 (Destination Unreachable) can have Code 0 -
  network unreachable, 1 - Host unreachable, 3 - Port unreachable, 4 - Fragmentation needed.
- **Checksum**: ensures integrity of the ICMP message (header + data).
- Message-Specific Data to make ICMP flexible for different message type. For example, Destination Unreachable can 
  send the **MTU value** required for fragmentation.
- Original IP Header + First 8 Bytes of Payload, to allow sender to match the ICMP message to specific packet which 
  caused the error

!!! note ""

    To prevent boardcast storms ICMP errors aren't send for other ICMP errors or for broadcast message. 

The most common use case of ICMP are:

1. Error reporting: Routers and hosts send ICMP messages when a packet can’t be delivered.
2. Network diagnostics: You can build tools on ICMP to provide diagnostics on network. For example, `ping` uses ICMP
   Echo Request / Reply to ping an ip and `traceroute` uses ICMP Time Exceeded message to provide packet path taken from
   source to destination.
3. Path MTU Discovery (PMTUD): Start sending large packets with set DF flag. On the way, router would drop the oversized
   packet and sends ICMP "Fragmentation Needed". You can reduce the packet size and repeat the process until it reaches 
   destination. This way you can adjust your packet size while avoiding fragmentation.

