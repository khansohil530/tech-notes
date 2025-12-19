---
comments: false
---

# Backend Communication Patterns

During early days of web, most of the applications were just servers sharing resources with their clients through
APIs. Requirement for this mode of communication was easily fulfilled using the **Request-Response** model, where the client
requests for a resource, and server sends back the resource in response. But as web apps and their requirements 
evolved, request-response model wasn't sufficient anymore. For example, apps where client submits a task for server to
process and to get a update on completion status of task, client needs to repeatedly send requests until the task is 
completed. This is very inefficient, as server losses resources to reply the client requests which might even delay the
processing of its task. Due to such requirements, various communication patterns emerged as a need for different 
requirements which couldn't be just fulfilled by existing patterns.

## Request-Response Model

Even though the design of this model looks very simple from outside (client sends a request to the server and server 
responds to client for their request), it requires multiple steps underneath:


1. Client sends a request. This request is received as stream of bytes by the server, which is needs to parse into 
   programmable objects to make sense of request.
2. Server parses the request. It gathers all the individual segments of bytes in ordered sequence to form the complete
   request, which is passed onto application layer for processing. 
3. Server processes the request. This requires deserializing the content of request into programmable objects. The choice
   of message format can decide the latency for this operation, because parsing a text based format (JSON) is more 
   expensive for computers than a binary format (ProtoBuf). But text based format are human-readable which makes it 
   easier to debug any error, which binary format message aren't human-friendly.
4. Server sends the response. Once the object is parsed and processed, server serializes the object into bytes and hands
   it over to transport layer for transmitting it to client.
5. Client receives the response which is parsed and processed to consume.

These steps have two repetitive components which are used on both hosts before the application can consume the request,
i.e. ordering the bytes of request which is governed by transport protocol, and serial/deserialization of those bytes
into objects which is governed by the format of requested message. For example,

- HTTP protocol and JSON message format used in REST APIs excel in developing general purpose web APIs, and are easier
  to debug (human-readable message format).
- HTTP/2 protocol and Protobuf message format used in gRPC excel in microservice architecture which requires 
  high-performance network calls b/w different services.
- HTTPS protocol and GraphQL used in GraphQL queries allows client to flexibly fetch only what's required, which helps
  to save bandwidth and improves overall performance.

This way the model has been employed across different apps using protocols like HTTP, DNS, SSH, RPC, SQL and Database 
protocols, APIs (REST/SOAP/GraphQL). But the model fails in places where the client isn't sure if the requested 
resource is ready, like notification service, chatting apps where client has to constantly ping the server to check if
there’s any new notification/chat or places where client have to wait for long time to form the response.

## Push Model

Request-Response model fails for apps which needs real-time communication, which requires pushing new data to other side
as soon as its available. For example, live feeds (scoreboard/stocks), chatting apps, alerting systems, requires 
immediate response from server and this can't be scaled using request-response.

Due to this Push model was designed in which the server can push data to client to immediately notifying the 
designated change. The core idea involves a Client opening a connection to subscribe for
updates, on which Server pushes updates allowing Client to react to them in near real-time. The real implementation varies
across different technologies like following:

### Server-Sent Events (SSE)

SSE is a one-way push communication model where the server streams updates to the client over HTTP. It's used across 
apps where updates mostly flow one way (like live dashboards, notifications) and requires a simple-reliable protocol.
The implementation involves Browsers using a persistent HTTP connection(1) which Server can use to push its text-based
events whenever data changes.  
{.annotate}

1. Using EventSource browser interface

The key characteristics which defines SSE's niche use case are:

- **Simple** since its build on top of HTTP. Clients which are mostly browsers can just send HTTP requests with 
  `Accept: text/event-stream` and the browser automatically receives and processes updates, while the server can use
  the header to keep open connection and send plain text.
- **Auto reconnection** feature is built-in the EventSource API used by browsers to support SSE, and it is handled 
  natively by the browser.
- **Reliable** ordered delivery using metadata like `Event-ID`/`Last-Event-ID` which can be used to resume streaming 
  events after reconnection.
- **Text-only** message format makes it easy to inspect and debug using browser curl/devtools, however text-based 
  message uses more bandwidth and computation for parsing than binary formats.
- One-way communication keeps protocol simple and light as there's no need for coordination, however client would need
  separate connection for sending data. 

While these characteristics allow SSE to excel in scenarios where the server needs to continuously push one-way, real-time 
updates to many clients with minimal complexity, such as live dashboards, notifications, activity feeds.
It's limited for use cases that require high throughput, binary data, or bidirectional interaction like chat systems,
collaborative applications, online games, media streaming. This limitation primarily occurs because text parsing is 
less efficient than binary formats and clients must use separate requests to send data back to the server.

### Websockets

The limitations from SSE were covered in Websockets protocol which is designed to support persistent, bidirectional 
communication between a client and a server. The communication start as a standard HTTP request so they can traverse 
existing web infrastructure, but after the server accepts the `Upgrade: websocket` header, the connection switches to 
the WebSocket protocol. From that moment onward, the connection behaves like a long-lived TCP socket rather than 
discrete HTTP requests, removing the need for repeated handshakes and headers. Few key characteristics of Websockets
which allows it to excel in its niche:

- **Low Latency**: Once upgraded, the connection stays open and messages are sent as small, framed payloads instead of 
  full HTTP requests. This eliminates the repeated overhead of HTTP headers, routing, and connection setup.
- **Bidirectional Communication Channel**: Websockets allow both the client and server to send messages independently 
  and simultaneously over the same connection. This enables highly interactive flow where client actions immediately 
  affect server state and vice versa. For example, in a chat application, users can send messages while also receiving 
  messages from others without blocking or creating new connection.
- **Flexibility on message format**: WebSockets support both text and binary frames which lets applications choose the 
  right trade-off between human readability and performance, which is critical in high-throughput or latency-sensitive
  systems.

However, these benefits comes with a tradeoff:

- **Scaling Implications**: Because WebSocket connections remain open, the server must maintain state for each connected
  client, including connection metadata and session context. This statefulness complicates horizontal scaling, as 
  requests cannot be freely routed to any server instance without coordination. Load balancers often need sticky 
  sessions or shared state layers, and infrastructure must handle reconnects, dropped connections, and failover 
  scenarios. While solvable, these concerns increase architectural complexity compared to stateless HTTP APIs.
- **Increased Resource Usage**: Long-lived connections consume server resources such as memory, file descriptors, and 
  CPU for heartbeat and connection management. At large scale, this requires careful tuning and monitoring to avoid 
  exhaustion. Developers must also implement mechanisms for detecting broken connections, applying backpressure, and 
  preventing slow consumers from degrading system performance. 
- **Debugging Limitations**: After the HTTP upgrade, WebSocket traffic becomes opaque to many intermediaries. 
  Traditional HTTP tooling, caching, and logging no longer apply, making debugging and observability more challenging.
  Some proxies and firewalls may limit or terminate long-lived connections, especially in restrictive enterprise 
  environments. As a result, WebSockets can be less reliable in networks that are optimized for short-lived HTTP traffic.

These characteristics make WebSockets excel in applications that require continuous, real-time, bidirectional 
interaction, such as chat systems, collaborative editors, multiplayer games. However, they are often overkill for 
scenarios involving one-way updates, infrequent events, or simple data retrieval, where the added complexity and 
resource usage outweigh the benefits. In those cases, simpler push model or request–response models provide a better balance.

### Pub/Sub

Another common communication pattern involved systems which would emit a signal/message for happening of event
(producer) and this event can be used by other systems (receiver) to update their state. 
In traditional point-to-point communication, these producers must know who their consumers are and often wait for them, 
which makes systems brittle when new consumers are added, existing ones change, or workloads spike. To remove this 
strong coupling of producer and consumer, **Pub/Sub** (short for Publish/Subscribe) model was developed.

In **Pub/Sub** model, **producers** (publishers) emit events to a **broker/topic**, and **consumers** (subscribers)
receive those events without the publisher knowing who they are. Unlike direct request–response or socket-based 
communication, pub/sub introduces an intermediary that decouples senders from receivers. This decoupling allows systems
to scale and evolve independently, making pub/sub a foundational pattern for **event-driven architectures**. It's 
commonly implemented as follows:

- Subscribers express interest in a **topic/channel**, and the broker pushes messages to them as soon as they arrive. 
- Publishers simply send messages to the broker and move on, without waiting for acknowledgment from each consumer.

This enables the push model which allows events to propagate in near real time, and downstream services respond 
to the changes immediately. However, the key characteristic of pub/sub is the decoupling of publisher and subscriber
systems which provides them the isolation to work independently while also reducing coordination cost and independent
scaling b/w different system (**fan-out**(1))
{.annotate}

1. When a single published message may need to reach thousands or millions of subscribers. 

The implementation is often handled in a pub/sub systems (like (1)) by persisting message and delivering them to subscribers 
asynchronously to enable parallel processing on publisher. With this, the publishers are not slowed down by slow
consumers, and consumers can process messages at their own pace. Additionally, many pub/sub systems support configurable
delivery (like `at-most-once`, `at-least-once`, or `exactly-once`), persisted storage for message, etc. These options 
allow developers to finetune their systems as per requirement (like choosing reliability over speed for failure recovery).
{.annotate}

1. Brokers like Kafka, RabbitMQ


However, this comes at a cost which involves few tradeoff:

- Additional operational overhead from Broker. Developer must manage topic design, message schemas, retention policies,
  consumer groups, and monitoring of lag and throughput.
- Debugging becomes harder because message flows are asynchronous and indirect. You can enable tracing events across 
  multiple services to ease debugging, but managing these tracing events also adds their overhead. 

These characteristics make pub/sub excel in event-driven systems, such as microservices reacting to domain events, 
analytics pipelines, notification systems, etc. However, it struggles in scenarios requiring low-latency bidirectional 
interaction, or simple request–response logic where the flexibility and scalability of pub/sub are outweighed
by its added complexity and indirectness.

??? note "Pub/sub mixes Push and Pull Model?"
    At a high level, pub/sub behaves like push because producers publish once and events are delivered to interested 
    consumers without polling. However, some systems may use pull on the consumer side to solve reliability and
    scalability problems. For example, Kafka, SQS use pull-based consumption so that subscriber can pull when ready
    in-contrast to systems like RabbitMQ which pushes events on subscriber as soon as they're available.

## Pull Model

Push Communication Model (discussed [above](#push-model)) allows servers to push updates as soon as they're available.
Since this is done without any control, it could overwhelm the consumer's system specially when consumers have limited
capacity, or the data surges unpredictably. **Pull model** are designed to solve these problems, where the consumer 
can request data from a producer. If there’s new data, it’s returned otherwise there's no new data. While producer does
not initiate communication, it passively tracks new available data so that they can be made available as soon as anyone
asks. To use this effectively, there are few conditions usually needed in place:

- **Accessible Data Source**: The producer must store or expose data in a way that can be queried later.
- **Consumer Awareness**: Consumers need to know where to ask and what to ask for.
- **Polling or Request Strategy**: The consumer must decide how often to check—too often wastes resources, too 
  infrequently increases latency.
- **State Management**: Consumers often track what they’ve already seen so they don’t process the same data repeatedly.

While keeping these basic, there are different approach for implementing Pull Model:

### Short Polling

Short polling involves client repeatedly sending requests to a server at fixed time intervals to check for available 
updates. After receiving the response, the client waits for a predefined interval before sending the next request.
This way it creates a continuous request–response loop where the client initiates a request, the server processes it and 
replies instantly, and the connection is closed. The key to understanding short polling, is its frequency of polling: 

- shorter intervals provide faster updates but increase server load and network traffic, 
- longer intervals reduce resource usage but introduce delays in receiving updates. 

Due to this, it's commonly used in systems where simplicity is more important than real-time responsiveness like web apps
for checking notifications, task status. Because it relies on standard request–response mechanisms, short polling works 
reliably across browsers, firewalls, and network configurations. Developers can implement it using basic HTTP requests
without additional infrastructure as it doesn't require any persistent connections, complex server logic, or specialized
protocols.

However, short polling this simple design comes at a cost:

- When no new data is available, requests still occur, wasting bandwidth and server resources. As the number of clients 
  increases, this inefficiency becomes more pronounced, leading to higher CPU usage and unnecessary database queries. 
- Short polling cannot provide true real-time updates, since changes are only detected at the next polling interval.

Due to this, it's best suited for low-frequency updates and small-scale systems, and is often replaced by more efficient
alternatives such as long polling or WebSockets when real-time communication or scalability is required.

### Long Polling 

Long polling allows client and a server to deliver near real-time updates while still using the traditional 
request–response model. It is an improvement over short polling and was designed to reduce unnecessary network traffic 
and server load. In long polling, the client sends a request to the server and the server does not respond immediately.
Instead, it keeps the request open until new data becomes available or a timeout occurs.

The working of long polling differs fundamentally from short polling. When the client sends a request, the server checks
whether new data is available. If data exists, the server responds immediately otherwise it holds the connection open 
and wait. As soon as new data arrives, the server sends the response and closes the connection. After receiving the 
response, the client immediately sends a new request to wait for the next update. This creates a continuous, 
event-driven communication cycle. This design choices allows long polling to have the following advantage:

- It can be used to achieve near real-time communication without requiring persistent connections such as WebSockets, 
  commonly found in chat applications, live notifications, but especially in environments where WebSockets are not 
  supported. 
- Because the server only responds when there is actual data, long polling significantly reduces the number of empty or 
  wasted requests compared to short polling. This also conserves network bandwidth as responses are only sent when 
  meaningful data exists. It also offers better responsiveness, as updates are delivered as soon as they occur rather 
  than waiting for the next fixed interval. 
- Additionally, long polling works over standard HTTP, making it compatible with existing infrastructure, proxies, and firewalls.

However, they come at following cost:

- Holding many open connections for long periods can increase memory usage and connection management complexity on the
  server. At large scale, this can affect performance and require careful tuning. 
- Timeouts, dropped connections, and retries must also be handled correctly to avoid missed updates or excessive reconnections.

To summarize, long polling is an optimized form of polling that balances simplicity and responsiveness. It reduces the 
inefficiencies of short polling while remaining compatible with traditional HTTP systems. However, for highly scalable, 
truly real-time applications, long polling is often replaced by more modern approaches such as WebSockets or server-sent
events.
