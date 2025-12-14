# HTTP

During early days of internet, information was scattered across different computers and each system used its own way to
access files. **HTTP** (Hypertext Transfer Protocol) was invented to solve this problem by providing a standard way
to share documents across computers (regardless of the computer or operating system) using **HTML** format while also 
allowing users to follow links to different documents using **URL** (Uniform Resource Locator). 

This way of organizing documents was soon termed as **Web** because information was connected through interlinked 
documents, forming a structure like a spider's web. With this organization each page could link to many others, 
and users could move non-linearly from one piece of information to another instead of following a single, linear path.
But as internet expanded globally, the original idea of web was also evolved to **WWW (World Wide Web)** where 
information was accessible global, which required in turn required HTTP to address problems like scale, performance,
and reliability. This lead to creation of different version of HTTP, each of which was developed due to new use case 
on WWW which couldn't be solved by its precursors.

## HTTP/0.9

This was the initial version of HTTP which ran over TCP and was designed to only fetch HTML files using a 
request-response model over a small network. It worked since it was developed for small number of documents sharable
across trusted users. Client could easily access information with following steps:

```mermaid
--8<-- "docs/Courses/focn/diagram/http_09.mmd"
```

But this limited the server from sending only `index.html` file. There was no way to send any non-html file, or send an error, 
and you'd to host multiple websites on one server to share multiple documents. 

## HTTP/1.0

HTTP/1 transformed HTTP from a minimal experiment into a practical protocol for the early WWW. It introduced several 
features that made the web viable:

- _Request and Response Headers_ to allow exchanging metadata. This allowed clients to understand how to process 
  responses, which enabled sharing of non-HTMl content (like images, css, files) using `Content-Type` header. This
  `Content-Type` header turned the web into a general-purpose document system by enabling sharing of multimedia content.
- _Status code_ to allowed servers to communicate failure intelligently.
- _Multiple HTTP Methods_ like POST, HEAD to support early dynamic web applications.

The rest of the design, the way of communication was kept same as HTTP/0.9 which uses one TCP connection per HTTP 
request. This limited scalability of HTTP/1 due to high latency and increased load on server from frequent establishing
and destroying TCP connection. There was also no way to pipeline request, clients had to send one request at a time,
and wait for its response before sending next request. 

## HTTP/1.1

HTTP/1.1 fixed the performance and scalability problems that made HTTP/1.0 unsuitable for a growing web. Few of the key
features introduced in this version:

- It allowed _Persistent Connections_ by using `Connection: keep-alive` header which was made default. This eliminated
  repeated TCP handshakes, which reduced latency while improving throughput of HTTP.
- Allowed _Request Pipelining_, so that clients can send multiple requests without waiting for responses. 
  However, responses must be returned in order, which leads to **head-of-line blocking**.
- Made `Host` Header required for Virtual Hosting. This allowed multiple domains to share a single IP address, making
  shared hosting economically viable which enabled large-scale web hosting.
  
    ??? note "Virtual Hosting"
    
        An IP address identifies a server, not a website. Before HTTP/1.1, server had no way to identify more than one website
        which made single website to use single IP which made hosting website expensive and inefficient. But with `Host` header,
        server could identify which domain the client intended. Conceptually, web servers are configured like:
        ```text
        IP: 203.0.113.10
         ├── example.com → /var/www/example.com
         ├── example.org → /var/www/example.org
         └── blog.example.net → /var/www/blog
        ```
        
        The server reads the Host header, matches it to a virtual host configuration and serves content from the correct site.
        This is called name-based virtual hosting.

- Introduced `Transfer-Encoding: chunked` which allowed server to send data in chunks without knowing the total size
  upfront. This enabled streaming response and supports dynamically generated content.
- Expanded HTTP methods which enabled RESTful APIs
- Expanded status code

This made modern web possible but still had few structural issues which limited its scalability for asset rich websites
and low bandwidth devices like mobiles. Two of the major issues were:

- Head of Line Blocking, where a slow response blocks whole pipeline. A workaround browsers began using was 6–8
  parallel connections but this increased congestion, and server overhead.
- Sending plain text headers which are sent in full on every request even if they're often mostly identical. This lead
  to wasted bandwidth which specially impacted case like application using large cookies headers, or high-latency device
  like mobile.


## HTTP/2

HTTP/2 made changes to how data is transported over the network while keeping the same HTTP/1.1 semantics to make the
web faster, more efficient, and more responsive which is required for modern, asset-heavy websites. Few of the key 
limitation which were addressed are as follows:

- Removes head of line blocking by using **Multiplexing**. Client can now handle multiple requests/responses 
  simultaneously over a single TCP connection using streams, without waiting for other request to finish.

    ??? note "Multiplexing"
        The browser opens one TCP connection to the server and HTTP/2 creates multiple **streams** within the connection 
        where each stream represents one request/response pair. For example, Stream 1 -> `index.html`, 
        Stream 3 -> `styles.css`, Stream 5 → `app.js`. These streams are independent, bidirectional and identifiable by IDs.
        
        When sending multiple requests,they're broken into small binary chunks called **frames**. Each frame includes 
        metadata like Stream ID, Frame type, Length and some Flags to identify and parse each chunk for respective request.
        When transmitting packet over network, these frames are intermixed so that no single request gets blocked at HTTP 
        level.HTTP/2 also allows you to prioritize stream which allows server to send critical resource first.

- Saves bandwidth using **`HPACK`** compression over headers which also allows you to index duplicate headers, so that 
  they can be referenced in later request using the index key instead of whole header. 

    ??? note "`HPACK` Core Concept"
              
        `HPACK` reduces header size by doing following 3 things:
        
        - Replacing repeated headers with indexes
        - Compressing new values efficiently
        - Remembering headers across requests
        
        To implement this, it uses header table which are divided into two types:
        
        1. Static table: Build into HTTP/2 spec, this table stores common HTTP headers and values like `status=200`, `method=GET` 
           (any static header and value). Each entry in this table is given an index number. Now HTTP/2 can use this index 
           number to reference this header instead of writing the whole header field and value. For example, `:method: GET` -> `2`.
        2. Dynamic table: Stores dynamic request specific headers and values, like cookies, session id, etc. Since these headers
           are mostly exchanged during connection, its build dynamic and shared b/w both client and server.
        
        With this implementation, HTTP/2 can shrink huge headers into few bytes.

- Using Binary frames instead of plain text allowed faster parsing and fewer errors at transport layer.

    ??? note "Explanation why Binary frames are faster than plain text"
    
        Parsing plain text into bytes requires reading and splitting into chunks until a delimiter to extract information while
        also handling extra whitespaces, line folding, case-insensitive headers, etc. HTTP/2 replaced this free-form text with
        binary frames which have fixed well-defined layout.
        ```text
        ┌───────────────────────────────────────┐
        │ Length (24) | Type (8) | Flags (8)    │
        ├───────────────────────────────────────┤
        │ Stream Identifier (31)                │
        ├───────────────────────────────────────┤
        │ Frame Payload (Length bytes)          │
        └───────────────────────────────────────┘
        ```
        With binary framing, parser immediately knows the length of bytes to read, what field is represent by the bytes, and
        which stream does it belong to. This makes parsing an O(1) operations since field size are fixed. Also, you can 
        preallocate buffers since the size is already defined, making it easier on CPU.


While multiplexing solved the Head-of-line blocking at HTTP level by using streams, it didn't solve **Head-of-line 
blocking at TCP level**. If anyone TCP packet is missing, TCP buffers everything that arrives after it and waits for the
missing packet before passing any of that data upward. This problem stems from the implementation of TCP itself, where
it doesn't deliver data to the upper layer (HTTP/2) out of order. For example,
```text
TCP packets sent:
[ P1 ][ P2 ][ P3 ][ P4 ]

Receiver gets:
[ P1 ][    ][ P3 ][ P4 ]
```
As per TCP spec, it delivers P1 packet to HTTP/2 but buffers P3 and P4 packet because P2 is missing. It then waits
until P2 is retransmitted and only then delivers P2, P3, P4 to HTTP/2.

## HTTP/3

To get around TCP-head-of-line blocking in HTTP/2, HTTP/3 upgraded its transport layer protocol to [**QUIC**](quic.md).
QUIC which is build on top of UDP, which implements core features from TCP which guarantees reliability, ordering,
congestion control, and security in user space. The key difference is that it applies these rules to independent 
stream so that even if one stream faces packet loss, it doesn't block other streams. Now each HTTP request/response 
can be mapped to a QUIC stream to avoid TCP head-of-line blocking. Additionally, since QUIC is a much recent protocol 
compared to TCP, it implements other common features by default which were used on TCP as an extension.

- Provide built in Encryption using TLS1.3
- Faster handshakes to minimize RTT due to multiple back and forth setup request/response
- Allows Connection Migration so that either host device can switch connection in middle of transmission without 
  losing existing connection
- Header compression using `QPACK`

This fixed various issue which HTTP/2 could not while using same HTTP API. As such, its increasingly getting adopted
across various browsers like safari, chrome, firefox, most CDNs while using HTTP/2 as a fallback. And usually you don't
have to manually decide the protocol, its automatically handled by devices during negotiation via `Alt-Svc`.
