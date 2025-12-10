# Extra Topics

## Compiled vs Interpreted

Programs runs on Machine code which are specific to CPU architecture, like ARM, x86, M series of apple each of them 
operate on machine code instructions specific to their CPU. But developing programs in machine code wasn't efficient
and reliable for huge programs as it isn’t easily readable by humans. To solve this, high level languages were invented
which could be understood by humans and also translated into machine code for CPU to understand. 

One of the earlier and popular higher level language is **Assembly**. It's the closest language to machine code, in a 
way that each assembly instruction can be mapped to a specific machine code. But even some instructions in Assembly
can be specific to CPU. However, as programs became larger and more complex, the amount of assembly instructions
required also grew exponentially making assembly inconvenient for writing such programs. 

Higher Level Languages solve this by abstracting multiple low-level instructions into single high-Level 
instruction. But now we’ve to translate this high level instructions into machine code so that they can be executed by 
CPU. This process of translation is called **Compilation**. But to execute these translated machine code, we’ve to add more
stuff like EFI headers which tells metadata like the entrypoint of execution in the program, the heap, the stack, the
code area, import external files, etc. All these changes are done by Linker, and this process is called **Linking**. 

Compilation is done using a *compiler* program (like (1)), where it produces machine code as an object file for
each of the source files. To execute this object file, you’ll need to find and link all external object files required 
and create a single file containing all the code. Further, you also have  to add headers (ELF headers) required
by OS for executing program. For example, Linux uses *ELF* which contains all the metadata about program like where
is the text data, static data, heap data, stack area, entrypoint, etc. This step is also done by *linker*.
Linking is done using another program called *linker* (like (2)). This linked file is usually 
 the final executable which has all the instruction required to be executed against CPU.
{.annotate}

1. gcc, clang, rustc
2. gold linker, lld, mold

??? note "Bootstrapping complier"
    Compiler themselves are complex programs, so how do you create a compiler using your high level language before 
    even having the tool to translate your language to machine code? It's similar to "_Which came first, the Duck or the 
    Egg_" analogy. There are two commonly used approach:
    
    - Translate your language into another mature language which already have an implemented compiler (for example, Python).
    - Creating an initial compiler in another mature language, and then using it create compiler written in your own 
        language to make it **Self-hosted**. And the process is  called **Bootstrapping a  compiler**
    
    There are several approach to implement bootstrapping:
    
    1. The most commonly used approach is to write the first compiler in another language. Once this compiler exists, 
       you can rewrite the compiler in your language itself, and use the first version to compile the new one.
    2. Build an interpreter for your language in a host language and create the compiler program in your own language. Now,
      your compiler program can be compiled using the interpreter which can then be used to compile instruction in your
      langauge without the interpreter.
    3. Used historically, developers wrote the first compiler in assembly or even binary. Once you've minimal compiler, 
      used it to compile better versions. 
    
    `Go` wrote its host compiler in C, which was later rewritten in Go to self-host. There are several benefits for
     self-hosting a language:
    
    - further development is easier as the compiler can recompile itself.
    - removes dependence on another language or external tools.


Another key issue when developing programs was you've to maintain a different version of your program to support their
execution on different OS and CPU architecture. To solve this, **Interpreted Languages** were created which allowed you to
execute same code for every support environment. Such languages (like Python, JavaScript, Java) uses an intermediate 
program called **Runtime** which executes their instruction instead of directly executing them on CPU.
This way, we can just compile the Runtime specific to the platform (mac, window, linux) and it'll translate the 
intermediate code into respective machine code. The tradeoff is performance, since we’re adding an extra step for
executing code. Also, programs are only executable for environment in which the Runtime is available.

??? note "JIT Optimization"

    - **JIT** (Just In Time) Compilation is an optimization for interpreted languages. Interpreted Languages
    translates their intermediate code into machine code on the fly (when the program is running). This gives you faster startup
    but slower execution. 
    - **AOT** (Ahead of time) compilation turns all the code into machine code beforehand which makes
      is faster in execution but slower in startup. 
    
    *JIT* compilation is a hybrid of both these approaches, where it starts the program by running interpreted code 
    (which provides faster startup) and as the program continues, runtime monitors which path of code is getting executed 
    more frequently (hot path). These hot paths are then compiled into machine code on the fly and replaced which their
    interpreted version to avoid re-translation everytime. This gives it the performance of faster execution. 
    The issue however is the security risk of providing access to modify code on fly, which could be 
    used by malicious agents to execute hidden instructions. This is the reason AOT compiled process only loaded code on 
    startup and marks the area in memory as read only. JIT makes it less insecure by using static code segment for initial
    runtime and when it needs to switch with compiled instruction, it loads the code in heap, writes the changes and then
    makes it read only and executable so that program counter can execute these instructions.

## Garbage Collection

While managing memory manually if done right allows program to keep their memory usage low and clean, but getting it
right everytime requires developers to understand the lifetimes of data across every code path. You've to free the 
exact memory at the right time,

- freeing is done too early, you'll get **use-after-free** bug which could crashes your program or corrupt its state 
  leading to security vulnerabilities.
- freeing is done too late, it leads to memory leak. If it's completely forgotten, it'll leak into system reducing its
  memory capacity over time.
- and other disasters like free twice, freeing the wrong pointer, or returning early due to an error.

You can avoid all these failures and complexity if you can make memory management automatic, as a part of the
language/external program which can allocate memory as required and deallocate it when its no longer in use. This idea
of automatic deallocation was termed as **garbage collection**, which finds programming objects unreachable by code and
frees them from memory. The trick to make it feasible is "_how the GC finds the unreachable objects_".
There are several strategies and each has different trade-offs.

??? note "Mark-and-Sweep algorithm" 
    As suggested by name, it does garbage collection in two phases:

    - Mark phase: Start from "roots" (like global variables, stack variables) and follow all pointers, marking everything
     reachable. 
    - Sweep phase: Walk through all heap objects; any object not marked can be freed.
    
    This approach is used widely as it works reliably with arbitrary object graphs (like cycles) and simple to implement.
    The tradeoff begin it block the main program during GC to mark objects reliably. Also, it can easily fragment
    the heap memory, making it directly unusable for large chunks of allocation.

??? annotate "Generational GC"
    Used across Java, Go, JS - this approach relies on of a key observation in real program:
    "_Most objects die young_", because most objects are either created from temporary tasks (like (1)), and long-lived
    objects (like (2)) are quite rare. Empirically, profilers and GC log shows that 80–95% of objects die before a second GC cycle.
   
    We could exploit this information, by storing the short and long-lived object separately so that GC can focus more 
    on the area used for storing short-live objects and collect it frequently while collecting long-lived objects 
    rarely. Area for storing short-live object is termed as **Young generation** and long-lived as **Old generation**.
    When collecting young generation, living objects are copied into a new space making everything left behind as garbage.
    This uses **Cheney's algorithm**, and it helps in defragmenting the heap automatically.  
    
    This approach allows faster GC with fewer pauses and less wasted work, while causing very little fragmentation. The 
    tradeoff begin overhead from memory copying.
1. storing intermediate results, function parameters, iteration variables
2. configs, sessions, cache, object pools

??? note "Reference Counting"
    Used in CPython, Rust ARC - Each object internally keeps a count of how many references point to it. When a 
    reference is created/removed, the count is incremented/decremented respectively. When an objects reference count 
    reaches 0, it is immediately freed.
   
    This saves program from big pauses due to GC cycles. The tradeoffs being, more overhead when creating new reference, 
    since it also requires updating the reference count. It also doesn't handle cyclic cases where two objects points
    to each other can cause memory leak, as their reference count is always up by 1. To solve this, CPython uses a 
    **secondary cycle collector** which cleans up cycles periodically. Rust ARC avoid secondary cycle collector by allowing
    developers to use **weak reference** for relationships that shouldn’t imply ownership. Using weak references doesn't
    increase the reference count of pointed object, as such can be used to avoid cycles.

Other than freeing memory GCs also helps to reduce memory fragmentation by Compaction which allows

- faster allocations
- improves cache locality
- reduces the chances the system must ask the OS for more memory

Without compaction, a program might have free memory but in such small chunks that new objects can't fit.

## Virtualization vs Containerization

Hardware resource where very scarce and expensive during early days of computing. Companies only had
single mainframe, on which users had to take turns for running their workload. But with creation of **Virtualization**,
multiple users could run workloads on the same physical machine without interfering with each other. 

**Virtualization** allows one physical computer (the host) to run multiple isolated environments (the **virtual machines**
or **VMs**) at the same time. Each VM acts as a complete standalone computer, with its own 
CPU/Memory/Disk/OS but none of these are physical devices, they're virtualized by software.

!!! note "How Virtualization works?"
    Virtualization at core uses a software layer called **Hypervisor** which sits between virtual machines and 
    physical hardware. Whenever a VM tries to execute an instruction as usual, the hypervisor quietly translates, schedules,
    and manages the requests over to real hardware. With time, modern CPUs added built-in virtualization features which
    allow VMs to run instructions almost directly on the CPU by providing special modes where the hypervisor can
    intercept dangerous or privileged operations. This made virtualization fast. Before these extensions, it relied on
    slow software tricks. There are two types of Hypervisor at this time:

    1. **Bare-Metal Hypervisors** which are installed directly on the hardware. For example, VMware ESXi, Microsoft
        Hyper-V, etc. These are often used in datacenters and clouds because they’re faster and more secure.
    2. **Hosted Hypervisors**: Run inside an existing OS as a software layer. For example, VirtualBox, VMware 
       Workstation. These are great for developers, testing, and personal use.


Even though hardware is cheaper and readily accessible now, virtualization still remains essential because of few
reasons:

1. Server Consolidation: run many servers on one machine.
2. Security Isolation: VMs are very strongly isolated, making them ideal for cloud customers and regulated industries 
   like banking, government. Multi-tenant clouds (AWS, Azure, GCP) rely heavily on VMs to safely isolate customers.
3. Legacy OS Support: run Windows XP, old Linux kernels, or any outdated system safely.
4. Snapshots: snapshot whole OS state for backup, migration or cloning.

---

As computing evolved, hardware become cheaper while applications grew larger and more complex. There were two commonly
faced issues when developing software during this period:

- Apps could behave differently on different machine due to conflicts in underlying dependencies of system. This made 
  development painful, as developer would've to often manually tweak OS to successfully run their application.
- Monolith architecture of application made them larger, complex, and harder to manage. Due to this, Microservice 
  architecture became popular which allowed developers to break single applications into smaller services which would
  operate with each other in coordination to become the fully functional application.

Both of the above problems were solved using **Containerization**. It allowed developers to package an application and
everything it needs (libraries, dependencies, runtime) into a single file called **image**. Images can be executed 
independently across different OS using a **container runtime** program (1) solving the conflict in dependencies of 
platform. The process created for running the application is termed **Container**. Containers have a control view of other 
processes and resources, so that they can work in an isolated environment without external interference. This solved the
second issue where developers can now use containers for creating their service which are isolated from other process 
and can be managed independently to coordinate with other process or scale. One of the implementation of container is
by using built-in Linux kernel features: **Namespace** and **Cgroups**.
{.annotate}

1. similar to how interpreted language use runtime to achieve platform independence

**Namespaces** are used to isolate what a process can see. Different types of namespaces are used to isolate different
  parts of  process. For example,

- PID namespace allows you to isolate process tree. This way containers can’t see or kill host processes, behaving 
  like its own mini-OS
- NET namespace allows network isolation, giving each container its own virtual network stack. 
- MNT namespace to control mount point visible to container 
- IPC namespace to isolate shared memory, 
- UTS namespace to provide container its own hostname and domain name.
- USER namespace for user ID isolation. 

**cgroups** (Control Groups) limits how much resource a process can consume, like CPU, Memory, Storage I/O and Process
  limits. Without cgroups, a container could hog the machine.

??? note "Linux Container"
    At this time, the term containers is synonymous to Linux containers which are one implementation of containers using
    built-in Linux Kernel features due to their dominance over modern software development. Since other platforms like
    macOS and Window doesn't provide Linux namespaces or cgroups support, they're supported non-natively using a
    lightweight Linux virtual machine.

Further, to make container images easy to distribute and faster to build, they're built in layers using **Union Filesystem**
(like (1)). Tools like OverlayFS merge these layers into
a single filesystem at runtime.
{.annotate}

1. base OS layer, dependency layers and application layers

To summarize, containers are essential for software development as they're

1. Portable: runs the same everywhere
2. Fast: spins up in milliseconds, making them perfect for scaling microservices, CI/CD workflows and serverless backends
3. Efficient: uses fewer resources. You can run dozens or hundreds of containers per host.
4. Reproducible builds: Dockerfiles let you define environments as code.
5. Fits perfectly with DevOps automated pipelines, Kubernetes and cloud-native tooling
