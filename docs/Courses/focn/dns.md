# DNS

Computer IP addresses are random 32-bit numbers and dynamic, which made it very hard to for us to remember when 
visiting any online service. It's like when you'd to remember the phone number of every person you had contacts for, 
but that was solved using phone book. Similarly, IP addresses were provided names (called **domain**) and since 
computers only communicate using IP address, we need to translate that name, which was done by 
**DNS** (Domain Name System). 

The earliest implementation of this idea was to keep a single file which would map the names to IPs, but this approach 
wasn't scalable as you'd to manually update the file and share that with all others. Instead, using hierarchical design
where each level delegates responsibility to next proved much better scaling as it allows distributed control without
single bottleneck. This lead to development of domain names composed of identifiers from different level of hierarchy.
For example, `www.example.com` has `com`, `example` and `www` levels where each steps from right to left narrows down
the search space. The general hierarchy followed is as follows:

```mermaid
--8<-- "docs/Courses/focn/diagram/domain_hierarchy.mmd"
```

- **Root (.)**: Domain hierarchy starts from right to left, so `.` is marked at the right most part of domain but this
  is usually handled implicitly when translating domain (instead of explicitly showing it on domain).
- **TLD** (Top Level Domain): Highest visible category of domain like `.com`, `.org`. TLD are categorized for different
  use case like generic TLD like `.com` which is available for commercial use, country-code TLD like `.in` for 
  specifically india, or restricted TLD like `.edu` restricted for educational institutions.
- **Second-Level Domain**: registered domain name, for example `www.example.com` has `example`.
- **Subdomain**: used for logical or functional subdivision within domain owner.

!!! note ""
    In `www.example.com` example, `.com` is TLD, `example` is secnd level domain and `www` is subdomain.

To provide Domain-to-IP mapping globally, it's stored across hierarchy of distributed database. The official mapping 
lives on machine called **Authoritative Name Servers**, which stores the mapping as **DNS records** in **DNS Zone** 
files.  

??? note "Common DNS Record Types"
    - `A/AAAA`: used to map a domain name to a fixed IP address. `A` is for IPv4 and `AAAA` for IPv6. 
    - `CNAME`: canonical name, used to makes one name alias another
    - `NS`: defines authoritative name server for a domain


With this organization of information, we can resolve domain names into IP addresses, which is implemented using 
DNS protocol. DNS protocol is developed over UDP since DNS queries are small and fast, which could be retried if any
UDP packet is lost. But since UDP is stateless, how does lookup servers know where to reply back the response? Using
**Transaction ID** header which is kept common across all the query used for resolving a domain. This ID header is used
to identify which request the response belongs to. Finally, to understand how DNS lookup resolves domain name to IP 
address, look at following steps:

```mermaid
--8<-- "docs/Courses/focn/diagram/dns.mmd"
```

## DNS Packet

```text
        ┌───────────────────────────────────────────────────┐
        │                 IP Header                         │
        │                 20-60 bytes                       │
        └───────────────────────────────────────────────────┘
        ┌───────────────────────────────────────────────────┐
        │                  UDP Header                       │
        │                   8 bytes                         │
        └───────────────────────────────────────────────────┘
┌───────┌───────────────────────────────────────────────────┐
│       │            Transaction ID (ID)        │   16 bits │
│       ├───────────────────────────────────────────────────┤
│       │QR│ Opcode │AA│TC│RD│RA│ Z │   RCODE   │   16 bits │
│       │1 │  4 b   │1 │1 │1 │1 │3b │    4 b    │           │
│       ├───────────────────────────────────────────────────┤
DNS     │          QDCOUNT (Questions)          │   16 bits │
Headers ├───────────────────────────────────────────────────┤
12 B    │          ANCOUNT (Answer RRs)         │   16 bits │
│       ├───────────────────────────────────────────────────┤
│       │          NSCOUNT (Authority RRs)      │   16 bits │
│       ├───────────────────────────────────────────────────┤
│       │          ARCOUNT (Additional RRs)     │   16 bits │
└───────└───────────────────────────────────────────────────┘
        ┌───────────────────────────────────────────────────┐
        │                  Data                             │
        │                                                   │
        └───────────────────────────────────────────────────┘
```
  
- **Transaction ID**: ID chosen by the client to match responses to queries (since UDP is stateless).
- Different flags:
      - **QR** (Query Response) : 0->Query, 1->Response
      - **Opcode**: to specify type of DNS operation like 0 for standard query, 2 for server status. 
      - **AA** (Authoritative Answer): set to 1 if the responding server is authoritative.
      - **TC** (Truncated): set when the response is too large for UDP, client should retry with TCP.
      - **RD** (Recursion Desired): set by client to request recursive resolution. 
      - **RA** (Recursion Available): set by server if it supports recursion.
      - **Z** (Reserved): reused for extensions (e.g., AD, CD in DNSSEC)
      - **RCODE** (Response Code): indicates result of query, like 0 -> NOERROR, 1 -> FORMERR, 2 -> SERVFAIL.
- **QDCOUNT** (Question Count): Number of entries in the Question section, mostly 1 
- **ANCOUNT** (Answer Record Count): Number of resource records in the Answer section 
- **NSCOUNT** (Authority Record Count): Number of resource records in the Authority section, used for NS records.
- **ARCOUNT** (Additional Record Count): Number of resource records in the Additional section



??? note "DoT/DoH"
    DNS queries aren’t encrypted by default which means ISP knows every domain you’ve queried through it.
    To encrypt the data, there are solutions like DoT (DNS over TLS) and DoH (DNS over HTTPS) but these aren’t enabled
    by default.
