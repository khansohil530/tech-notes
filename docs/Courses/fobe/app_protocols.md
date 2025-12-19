# Communication Protocol

Previously we'd discussed [backend communication patterns](comm_pattern.md), which describe when and why services 
communicate, while in this chapter we'll describe how that communication is technically carried out using specific
protocol. Backend communication protocols define specifications which can be implemented by two application to 
communicate with each other over network. These specifications involve details about network level, message formats, 
transport rules, and reliability guarantees so systems can exchange data consistently and securely. Without these
protocols, distributed backends could never interoperate, scale, or evolve independently. Some of the commonly used
backend protocol implementations [HTTP](../focn/http.md), [TLS](../focn/tls.md) were already discussed in the 
networking [notes](../focn/index.md) and fewer others are discussed below.

## gRPC

As systems at this time are moving towards distributed architecture, where microservices talk to each other across 
networks, often at high scale and low latency requirements. Within such architecture, traditional APIs which rely 
heavily on text-based formats like JSON and loosely defined contracts can lead to inefficiencies like larger
payloads, runtime errors due to schema mismatches, and duplicated client/server logic. To solve these, microservices
need faster communication, stronger API contracts, and better support for real-time and bidirectional interactions.

**gRPC (Google RPC)** addresses these issues by using Protocol Buffers (Protobuf) as a compact, binary, strongly typed
message format and HTTP/2 as its transport. 

- Protobuf enforces a clear schema and enables automatic code generation for clients and servers across multiple 
  languages, reducing boilerplate and integration errors. 
- HTTP/2 brings multiplexing, header compression, and persistent connections, allowing gRPC to support unary calls, 
  server streaming, client streaming, and full bidirectional streaming efficiently.

Implementing protobuf starts with a `.proto` file, which defines the schema of your communication specified using:

- The data structures (messages) with strongly typed fields
- The service interface (RPC methods) and their request/response types

This schema is the single source of truth and because it is language-agnostic, both client and server teams rely on the
exact same definition, eliminating ambiguity and undocumented behavior. To automate boilerplate code required to use 
given schema while also ensuring that the code adheres to latest schema change, gRPC uses Protobuf Complier (`protoc`).
Using `protoc` with configured language extension (like Python, Go) you can generate code which handles:

- Data classes for all messages (with serialization/deserialization logic)
- Client stubs (methods the client calls)
- Server interfaces or base classes (methods the server must implement)

On server side, developers implement the generated service interface. Each RPC method is filled in with business logic, 
just like implementing an abstract class or interface while the gRPC runtime handles networking implementation details
like receiving binary Protobuf messages, deserializing them into typed objects, invoking the correct handler function,
and serializing response. This allows developer to directly work on typed objects without manually parsing requests or
builds responses

While on the client side, the generated stub looks like a normal local method call. The client calls a function, 
passes a typed request object, and receives a typed response (or stream). Under the hood, the stub handles networking 
details like serializing the request to Protobuf, sending it over HTTP/2, handling retries, streaming and deserializing
the response back into objects. This makes remote calls feel similar to in-process function calls while still being 
network-safe.

Another issues with APIs, it that they inevitably evolve (1) over time. If these changes are not done careful, they can 
break existing clients which were still dependent on old schema, causing runtime failures in production. To avoid such
failures, you need to make sure the schema change are backward compatible, so that older clients can continue to work
even after the schema changes. gRPC provides backward compatability by using numeric tag to identify field instead of 
their name. As long as you're not changing or reusing existing numeric tags, this format is going to be backward 
compatible. Also, since the schema is strongly typed many errors related to schema are caught during build phase, when 
generating code from schema.  This is a key reason gRPC scales well in large organizations with many independently 
deployed services. 
{.annotate}

1. new fields are added, old fields are removed/renamed, data type changes


Still, gRPC is not always ideal for all APIs:

- Its binary format is not human-readable, making debugging and testing harder compared to REST with JSON. 
- Browser support is limited and often requires proxies(such as gRPC-Web). 
- It also has a steeper learning curve, introduces additional tooling and build steps, and may be overkill for simple
  CRUD APIs.


## WebRTC

Before WebRTC, real-time voice/video on the web was very fragmented and insecure. It was mostly controlled by
proprietary vendors and often required installing plugins to support real-time voice/video communication. Due to these 
problems, it wasn't possible to build native browser video calling. To eliminate these problems, Google initiated WebRTC
to make real-time, secure audio, video, and data communication a native capability of the web without using any plugins,
proprietary protocols, or vendor lock-in. This is why, most video calls apps like Zoom and Google Meets are able to run
in a browser. However, WebRTC isn't a communication protocol but a stack of different technologies like (1) to enables 
real-time features like (2) natively within browser.  
{.annotate}

1. Browser APIs, networking protocols (ICE, STUN, TURN), media protocols (RTP, SRTP) and security (DTLS)
2. video/voice calls, screen sharing, peer-to-peer file transfer, multiplayer games, live collaboration

It's high level architecture involves 3 phases:

### Phase 1: Signaling

Signaling phase is used to gather all the information required for exchanging any real video, audio, or data. 
WebRTC doesn't define any protocol for this process, so that WebRTC can be freely integrated into existing system
without adhering to a single protocol. With signaling channel decides, both peers must first agree on how real-time
communication will happen. This happens by exchanging SDP (Session Description Protocol) message, which describes either
peer's capabilities and intentions. It lists information like kinds of media the peer wants to send/receive, which 
codecs it supports, how encryption will be done, and what transport mechanisms are acceptable. 
This signaling process happens in following steps:


```mermaid
--8<-- "docs/Courses/fobe/diagram/webrtc_p1.mmd"
```

### Phase 2: Connection Establishment

Once both sides have exchanged and accepted each other’s session descriptions, they share a common understanding of the
media format, security parameters, and connection expectations. At this point, they know what they want to do, but don't 
know how to reach each other on the network (peer-to-peer). 

Most devices on internet are behind NATs or firewalls, so a peer usually doesn't know its own publicly reachable 
address. Using **ICE** (Interactive Connectivity Establishment), each peer gathers possible connection endpoints, 
known as **ICE candidates**. These may include local network addresses, public addresses discovered via **STUN** servers,
or relay addresses provided by **TURN** servers. Each of these candidates represents a possible path that the peers 
could use to communicate. 

??? note "NAT"
    Network Address Translation (NAT) allows devices in a private network to access the internet using one or more 
    public IP addresses. For outgoing traffic, the NAT device rewrites the packet’s source private IP and port to a 
    public IP and port, and stores this association in a NAT table. Incoming packets are translated back to the original
    private IP and port using this table. 
    
    Depending on how incoming packets are matched against the NAT table, NAT behavior is commonly classified into four types.
    
    - In a **Full-Cone (One-to-One) NAT**, once a private IP and port are mapped to a public IP and port, any external host 
        can send packets to that public IP and port, and they will be forwarded to the mapped private endpoint, 
        regardless of the sender’s address.
    - In an **Address-Restricted NAT**, incoming packets are forwarded only if their source IP address matches an IP 
        address that the internal host has previously sent packets to, while the destination public IP and port match
        an existing mapping.
    - In a **Port-Restricted NAT**, incoming packets are forwarded only if both the source IP address and source port 
        match a destination that the internal host has previously contacted, in addition to matching the public IP and port.
    - In a **Symmetric NAT**, each outbound connection to a unique destination IP and port results in a distinct public 
        IP and port mapping, and incoming packets are accepted only if they match the exact destination IP and port
        associated with that specific mapping.
    
??? note "STUN"
    Session Traversal Utilities for NAT (STUN) is a protocol that allows a device behind a NAT to discover the public 
    IP address and port that the NAT assigns to its outbound traffic, as observed from the public internet. The purpose
    of STUN is to enable a peer to advertise a reachable public endpoint to other peers.
    
    In a typical STUN interaction, 

    - the device sends a request to a STUN server on the internet. 
    - The STUN server responds with the source IP address and port from which it received the request.
    - The device can then share this public IP and port with other peers as a candidate for incoming connections.

    This approach works with cone NATs, including full-cone, address-restricted cone, and port-restricted cone NATs, 
    because once a private IP and port are mapped to a public IP and port, that mapping is reused for traffic to
    multiple destinations. As a result, the public endpoint learned via STUN is valid for communication with other peers.

    STUN fails with symmetric NATs because the public IP and port learned from the STUN server are specific to the STUN
    server’s IP and port. When the device attempts to communicate with a different peer, the NAT creates a new mapping
    with a different public port. Since STUN assumes that NAT mappings are stable and reusable across destinations, 
    symmetric NAT violates this assumption by creating mappings on a per-destination basis.
    
??? note "TURN"
    Traversal Using Relays around NAT (TURN) is a protocol that enables communication with devices behind restrictive 
    NATs, including symmetric NATs, by relaying traffic through a publicly reachable server. A peer establishes an 
    outbound connection to a TURN server, which allocates a relay IP address and port. Other peers send traffic to this
    relay address, and the TURN server forwards the packets over the existing NAT-friendly connection to the device. 
    Because all traffic flows through the relay, TURN works regardless of NAT type, at the cost of additional latency 
    and bandwidth usage.

??? note "ICE"
    Interactive Connectivity Establishment (ICE) is a framework that coordinates the process of finding the best way for
    two peers to connect. Since a peer may have multiple possible addresses—such as local addresses, server-reflexive 
    addresses discovered via STUN, and relayed addresses obtained from TURN. ICE gathers these as candidates and 
    exchanges them with the remote peer using signaling. ICE then performs connectivity checks on candidate pairs and 
    selects the best working path, preferring direct peer-to-peer connections and falling back to relayed connections
    only when necessary.
    

These ICE candidates are exchanged during signaling as they're discovered, so that other side can test them to see which
ones actually work. Eventually, ICE selects the best viable path, ideally a direct peer-to-peer route and, if that fails,
a relay through a TURN server. Only after this candidate exchange and selection does WebRTC establish the actual secure 
transport connection.

??? note "Summarized Phase 2"
    ```mermaid
    --8<-- "docs/Courses/fobe/diagram/webrtc_p2.mmd"
    ```


### Phase 3: Media & data transfer

Phase 3 begins after ICE has finished and a working network path between the peers has been selected. 
At this point, the peers already know how to reach each other and can start sending real-time data securely and 
efficiently.

Before any media/application data flows, the peers establish a **DTLS** (Datagram Transport Layer Security) session 
over the chosen network path. DTLS is similar to TLS (used in HTTPS), but designed for UDP. This step authenticates 
the peers and performs a cryptographic handshake to derive shared encryption keys. Importantly, WebRTC mandates this 
step and unencrypted communication is not allowed. Once the DTLS handshake completes, the connection splits logically
into two kinds of traffic: **media** and **data**, each using a protocol suited to its needs.

- **Audio/Video** are transmitted using **SRTP** (Secure Real-time Transport Protocol). RTP is a protocol designed for
  real-time media, where low latency matters more than perfect reliability. If a packet is lost, it is usually better
  to skip it than to wait for a retransmission, because late media is worse than missing media. SRTP is simply RTP with
  encryption and authentication added. The encryption keys used by SRTP are derived from the DTLS handshake, which is
  why the media is end-to-end encrypted without needing a separate key exchange mechanism.
- **Application data**, such as chat messages, game state updates, or file transfers, flows through **SCTP** over DTLS. 
  **SCTP** (Stream Control Transmission Protocol) is more flexible than TCP and better suited for WebRTC’s needs. 
  It supports multiple independent streams, configurable reliability, and optional ordering. This allows a data channel
  to behave like TCP when reliability is required, or more like UDP when low latency is preferred. SCTP packets are 
  encrypted by DTLS in the same way media packets are.

Even though both media and data often travel over the same underlying UDP connection, WebRTC keeps them logically 
separated and applies the correct transport semantics to each. Media uses SRTP for timing-sensitive delivery, while
data uses SCTP for structured and configurable messaging.

All this traffic is encrypted end-to-end, i.e. encryption happens in the browser/app itself, and decryption happens 
only at the receiving peer. Intermediate servers, including TURN servers relaying packets, cannot read or modify the 
contents. They only forward encrypted packets. This property is enforced by the WebRTC specification and is not optional.

??? note "Summarized Phase 3"
    ```mermaid
    --8<-- "docs/Courses/fobe/diagram/webrtc_p3.mmd"
    ```

---

To conclude, WebRTC excels in area which requires web native real-time audio, video, and data communication over a
low-latency, and secure channel. Its mandatory end-to-end encryption, efficient UDP-based transport, and built-in NAT 
traversal allow applications to work reliably across diverse networks and devices. However, these benefits come with 
notable limitations: 

- WebRTC is complex to implement and debug due to its multi-layered protocols and asynchronous behavior
- peer-to-peer connectivity is not always possible and often requires TURN servers, which add latency and ongoing
  bandwidth costs
- scaling beyond small groups requires additional infrastructure such as SFUs which increases architectural complexity
- browser-specific differences in codecs, APIs, and performance can introduce interoperability challenges.

