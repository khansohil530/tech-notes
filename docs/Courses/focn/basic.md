# Basics of Networking

??? note "Brief History of Internet"
    Early computers were huge, expensive machines so much that an entire labs would only have single machine to serve their
    needs. The key requirement was to allow multiple people to access this system at same time, which lead to development
    of ideas to share processing power, storage, printers and terminals across connection. This lead to development of
    projects like **`ARPANET`** which was the early ancestor of the internet. At this point, multiple computers within a network
    could communicate with each other. But as more networks appeared, each used their own implementation for communication
    which proved as a hurdle to communicate across different networks. This required standardization in communication 
    protocol which could be implemented independent of underlying hardware. This solution was provided by TCP/IP protocol
    which did two critical jobs:
    
    - IP moves packets across different networks
    - TCP ensures those packets arrive reliably and in the correct order
    
    This marked as the starting point for development of internet which further evolved using key technologies like 
    DNS, WWW, Browsers, etc.  

Due to lack of standardized way of communication, computer vendors invented their own standards which caused
huge problems when integrating different systems . **OSI Model** was an attempt to provide common standards which
could be implemented by any hardware regardless of its vendor. Its idea was to break down networking in different layers,
 where each layer handled specific tasks. But by the time OSI protocols were ready, TCP/IP was already working in real
networks and had huge adoption. Due to this reason OSI protocols never took off, however its conceptual model survived
because it explains networking much clearly.

## OSI Model

OSI (**O**pen **S**ystem **I**nterconnection) model uses 7 layers of abstraction where each layer represents a
networking component.

- Layer 7/**Application Layer**, handles protocol like application protocols like HTTP, FTP, gRPC. Usually these 
   protocols are mostly implemented using standard libraries of respective programming languages, so that you won't
   have to go into implementation specific. 
- Layer 6/**Presentation Layer**, handles serialization and deserialization of programming objects into byte strings 
   as per the encoding provided. This may include encrypting/decrypting (SSL/TLS), formatting or encoding/decoding data. 
- Layer 5/**Session Layer**, handles sessions/connections b/w application. Usually involves around opening/closing 
   connections, keeping sessions active and synchronizing data. 
- Layer 4/**Transport Layer**, handles end-to-end communication b/w two systems using protocols like **TCP/UDP**. 
   It involves segmenting data into chunks (packets/datagrams), assigning them ports to allow sending/reading them by
   other processes (on same/different computer across network). 
- Layer 3/**Network Layer**, handles routing packets across networks to destined party using **IP address**. 
- Layer 2/**Data Link**, creates a reliable link between two devices on the same network using **MAC address**. 
- Layer 1/**Physical**, moves raw bits (0s and 1s) across wired/wireless signals by translating/assembling digital 
   signal into/from analog signals(like (1)) .
    {.annotate}

     1. like electric, light or radio waves.

For better understanding, follow below sequence diagrams for journey of POST request from both sender and receiver
perspective.
```mermaid
--8<-- "docs/Courses/focn/diagram/osi_send.mmd"
```

## Basic network routing

When your data is transmitted over network, it goes through different devices which helps to rout it towards the fastest
and appropriate direction. This redirection requires peeking information within our request like IP addresses,
MAC addresses.

- The request is first send to **Switch** which allows you to communicate within your LAN. For this it needs to look up
  the MAC address of data (L2), to move data b/w right device on LAN. However, switch doesn't have any information about
  external networks.
- To communicate with external networks, you've to use a **Router** (also known as default gateway) which has routing
  information required to send the data towards appropriate network. This requires looking up the destination IP 
  address (L3). 
- From there, data travel through ISP's access network which aggregate multiple data frames (L2) before moving them to
  their routers where they can be forwarded across cities, countries or continents. Across these WAN, data just hops
  from one router to another. 
- Finally, when data reaches the destination router, it moves the data to LAN switch which looks up the MAC address
  and forwards the request to correct host.

This is a brief overview of how multiple networking components interact to allow host to host communication over world.
In following chapters, we'll dive deep into each of these components and understand their HOWS and WHYS.