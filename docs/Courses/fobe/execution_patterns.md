# Execution Patterns

Backend execution typically begins with a **listener** that waits for and buffers incoming connections.
These connections are then accepted by an **acceptor**, which delegates the work to subcomponents such as readers, 
processors, and writers. These components are organized in different patterns because no single execution model can 
optimize latency, throughput, cost, reliability, and simplicity at the same time. For example, the classic 
listener–handler pattern breaks on large scale due to blocking I/O, long-running tasks, and unpredictable load. 
To handle such cases, you can organize components to separated concerns, like listeners focus on intake, acceptors on
control, handlers on logic, and workers on execution. Following are few production-tested execution patterns that can be
reused directly from their proven performance and trade-offs which allows faster development while reducing unknown risk
and preventing known scalability and reliability failures.

Most of these execution patterns revolve around a small  set of core components, with each pattern differing mainly in
how these components are arranged and coordinated:

1. **Listener**, detects incoming work like connections, messages, events, timers.
2. **Acceptor**, controls acceptance of work.
3. **Dispatcher/Router**, determines where accepted work should go.
4. **Queue/Buffer**, decouples arrival rate from processing rate.
5. **Worker/Executor**, provides the execution context like thread, process or event loop. 
6. **Handler/Processor**, implements business logic. 

Every execution pattern discussed below is essentially a different strategy for separating, combining, or constraining
these components to meet specific requirements.

## Single Threaded

The simplest possible backend, which uses same execution context (one thread, one call stack) to handle every 
responsible. Conceptually, the Listener accepts a request and immediately becomes the Handler that executes it from 
start to finish. For example,  

```mermaid
--8<-- "docs/Courses/fobe/diagram/exec_single_thread.mmd"
```

Above diagram showcases that the execution of two requests using this model are done sequentially, which blocks execution
of entire server. The backend would only process one request at a time, even if machine has multiple CPU cores, memory
which leads to underutilization of host resources. The throughput would also suffer even under modest concurrency, 
encountering issues like head-of-line blocking for fast request due to an ongoing slow request. 

## Single Listener + Single Worker

This model separates request acceptance from request execution, ensuring that incoming connections are accepted
immediately, even if execution is slow. This is achieved by separating responsibility of accepting requests and lining
them for execution using a separate thread/process called **Acceptor**, essentially decoupling the arrival rate from
execution rate. For example, 

```mermaid
--8<-- "docs/Courses/fobe/diagram/exec_single_thread_single_worker.mmd"
```

The execution is still single-threaded, but acceptance is no longer blocked by execution. This way, clients are no longer 
blocked/failed at connection time. Slow execution no longer prevents accepting new requests, and the backend can 
temporarily absorb some traffic spikes. However, using only one Worker to execute one request at a time can flood the 
queues under load, causing increased latency which linearly grows with load. 


## Single Listener + Multiple Workers

Using multiple workers to process enqueued requests seems the obvious way to relive load of execution from single 
worker. To implement this, backend now needs two new components:

- **Worker Pool**, multiple executing units (like threads/process) which can execute requests independently.  
- **Dispatcher**, to assign work to available worker in pool. This allows you to decouple work scheduling from work 
  execution. 

A simple execution example would look like following,

```mermaid
--8<-- "docs/Courses/fobe/diagram/exec_single_thread_multi_worker.mmd"
```

The system now separates concerns like acceptance, buffering, scheduling and execution of requests cleanly.
This allows the server to accept and execute multiple requests concurrently while utilizing machine resources efficiently
and increasing throughput for processing request as whole. Another key characteristics of this model is controlled 
concurrency, using queue and dispatcher which smoothen the load and enforce limits on concurrent tasks. 

## Event-Driven Execution

Last [model](#single-listener--multiple-workers) solves the problem of concurrency, however its worker still block their 
execution during an I/O operation (like DB calls, disk request). Under load, this causes issues worker starvation where
active workers idly waiting for completion of I/O which introduces artificial throughput limits. Developer might even 
over-provision workers and machines to hide this I/O wait.

To solve these issues, execution is split into I/O-bound and CPU-bound work. In this model, Workers are used only for 
CPU-heavy execution, while I/O-bound operations are handled in a non-blocking, event-driven manner using **Handler**
and **EventLoop** components. The Handler orchestrates request execution and explicitly decides whether a given step is 
CPU-bound or I/O-bound. 

- For I/O-bound operations, the Handler initiates a non-blocking I/O request and registers a callback or future 
  representing the continuation of execution. This registration is handled by the EventLoop, which reacts to I/O
  readiness notifications from the underlying system and schedules the associated continuation for execution. 
  After initiating and registering the I/O operation, the Handler yields control back to the EventLoop, allowing the 
  execution thread to process other ready tasks instead of blocking.
- For CPU-bound operations, the Handler submits the work to a Dispatcher, which is responsible for scheduling the task 
  onto an available Worker. Workers execute CPU-intensive code and return the result back to the Handler once 
  computation completes.

If execution does not require CPU-heavy processing, the Handler continues orchestrating control flow and lightweight
logic directly on the event loop thread. In this design, Workers never wait on I/O, and Handlers are resumed by the
EventLoop in response to events rather than being actively polled, enabling high concurrency with efficient CPU
utilization. For example, 


```mermaid
--8<-- "docs/Courses/fobe/diagram/exec_event_driven.mmd"
```

This way, the Handler is non-blocking and executes only lightweight orchestration logic on the event-loop thread. 
Any potentially blocking work—whether I/O-bound or CPU-bound—is offloaded to other components (the I/O subsystem or
Workers). After initiating such work, control is returned to the EventLoop, which drives execution forward by resuming 
handlers when the corresponding events or results become available.

But as a result, execution is no longer linear. Control flow is fragmented across handler invocations and event loop 
scheduling points, which makes debugging more complex and often requires tracing asynchronous boundaries. Additionally, 
backpressure must be carefully managed to prevent fast producers from overwhelming slower consumers or downstream 
systems.    


??? note "Does this makes the model single threaded?"
    At any given moment, a single handler executes on a single event-loop thread, which creates the illusion of 
    single-threaded execution. This design is intentional as single-threaded event loop eliminates concurrency hazards
    such as race conditions and the need for explicit synchronization within handlers.

    However, the server as a whole is not single-threaded. To scale, systems typically run multiple event loops, often 
    one per CPU core, process, or thread. In addition, Workers execute CPU-bound tasks in parallel, and the operating 
    system handles I/O concurrently. Together, these layers provide high concurrency and parallelism while preserving a 
    simple, single-threaded execution model within each event loop.

## Asynchronous Job Queue with Background Workers

[Above](#event-driven-execution) execution model above works and scales well for request-response tasks, but doesn't fit
for tasks whose execution happens later (like sending emails, video/image processing, batch jobs). Executing such task
under using request handler causes long response times, request timeouts and cascading failures under load. To solve this,
we can decouple request handling from task execution entirely.

The core idea is to allows request to enqueue work and returns immediately, while execution happens later, elsewhere.
To implement this, 

- **Handler** validates the incoming request, creates a job description which is enqueued on **JobQueue** 
  for executing later, and immediately returns a response for acceptance/rejection. 
- Since JobQueue now maintain state of each request, they must be durable to not loss any submitted jobs. Additionally,
  they must enable retries and failure recovery to execute jobs in case of errors.
- **Background Workers** can pull jobs from this queue and execute them independently of requests 
  (allows horizontal scaling)

An example of complete execution, follow below diagram:
```mermaid
--8<-- "docs/Courses/fobe/diagram/exec_job_queue.mmd"
```

This of decoupling of request handling from long-running execution provides us following benefits:

- Lower request latency since its no longer tied to job duration. Clients receive a fast, predictable response 
  regardless of how long the background task takes. 
- Failures are isolated to background workers, making them easier to retry without impacting live traffic. 
- Throughput scales independently, since the system can scale request handling to absorb incoming traffic while 
  separately scaling background workers to match processing capacity. You can buffer excess work in the queue instead
  of overwhelming execution resources or applying backpressure downstream.

However, every benefit in software engineering comes with some tradeoff:

- Since work completes asynchronously, results are eventually consistent, and clients may need to poll for status or
  register callbacks to observe completion. 
- Debugging becomes more complex because execution now spans multiple components—handlers, queues, and workers—often
  across different machines or processes. 
- Retries are a normal part of the model, idempotency becomes critical to prevent duplicate side effects.

But despite these costs, for real-world systems operating at scale, the resilience, scalability, and predictability
gained from this model far outweigh the added complexity.
