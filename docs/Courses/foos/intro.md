# Need of OS

Operating System is the layer of software which abstracts interactions of other software with the underlying
hardware by providing high level APIs while hiding the complex implementation details. The complexity comes
due to support for variety of hardware for same component which needs to be mapped under same layers of APIs.

However, computing environment varies to achieve different goals with unique constraints and workload specific to them.
For example, Smartphones have low power and touch screen requirements which lead to development of iOS and Android.
Online servers requires higher throughput, reliability which uses Linux server. General pcs which requires flexibility
in apps usage uses Windows or Linux Desktop. Each such OS would be uniquely tuned for their compatible device,
where a general purpose OS wouldn't be able to support needs like lower latency and smaller memory usage due to its
decision to support huge numbers of hardware devices which makes the abstraction layer larger and less efficient.
Even with all these differences, there are few things which are common across all OS:

- **Abstraction** to hide the underlying hardware from the software over an API layer. However, hardware are always 
  evolving due new inventions and discoveries, and OS have to provide backward compatability so that existing system
  don't fail. To do this, the abstraction is divided into different levels such that you can extend the
  existing design in order to handle new design changes without discarding previous design. 
- **Scheduling** concurrent processes as per users requirements so that you can run multiple processes concurrently
  without blocking one another. There are different scheduling algorithms which vary depending on goals such as
  fairness, throughput, response time, and real-time guarantees. Few of the popularly known are Round Robin,
  First-Come First-Served, Shortest Job First, Priority Scheduling, etc. You can customize scheduling to 
  select the algorithm which works best for your workload.

## Core components and terminologies

- **Kernel**: core component of an OS which manages most of the things like drivers, memory, CPU, scheduling. Other
  components are tools which communicate with Kernel to perform their operations. You can directly work with kernel,
  but to avoid the hassle there are different linux distributions (distros) already build on top of same Kernel to
  provide ready to use features like GUI, software managers, etc.
- **CPU**: Most of the computation work is executed in CPU, hence it’s very important to have frequently used values
  near to CPU. There’s `L1` and `L2` cache which lives within the CPU core and the closest in respective order.
  Then there’s L3 cache which is shared between cores but still closer to CPU. Each core can execute a process 
  individually allowing Kernel to schedule multiple process at the same time. 
  Instruction executed by CPU are known as machine code, which are specific to CPUs. Programs are naturally compiled
  into machine code to execute them in CPU, as such their compiled version are specific to CPU 
  (limitation of compiled language). Interpreted languages like Python, Javascript can run same code on different CPUs
  because you’ve compiled runtimes specific to the CPU which executes the code.
- **Memory**: (or RAM) is a fast (slower than CPU caches) but volatile storage space where CPU can store information
  which it needs for the execution of instructions. RAM is called Random because in earlier day’s the memory are based
  on tapes which had to be sequentially skipped ahead to access the information at specific location. In case of RAM,
  you can directly point to the address of information to fetch it. CPU uses RAM for lots of use case like storage of
  process state, data, caching etc.  To manage memory, OS uses the concept of virtual memory where it abstracts all 
  the details behind the scene like allocating space to process, swapping memory to disk, etc. to efficiently use memory.
- **Storage**: (SSDs/HDDs) a persistent storage even after power cut to the component. It’s slower than RAM and provides
  sequential access to data along with some limitation like HDD have to store data in sectors (section of bytes) 
  even if we’re updating just single byte, or like SDDs which have to write on pages and can’t erase single page but
  whole block to free up storage. In earlier days, controller to work with storage lived in memory as a program,
  but it was later move within the disk itself. Now, the controller exposes an API to work with disk which the OS 
  integrates with itself and exposes API to application (kind of chaining APIs). This was done to avoid involvement
  of OS to update the controller whenever there were modification required within controller due to changes in disk.
- **Network**: also called NIC (Network Interface Controller),is one of the way to communicate with other hosts. 
  The Network controllers receives data/packets from internet as electrical (Ethernet/Fiber) or analog signal which 
  gets translated into bits  (layer 1) -> frames (layer 2) -> packets (layer 3) -> segments (layer 4).
  OS is responsible for dealing with Layer 3 and 4 also known as the communication protocol which is responsible for
  parsing these bytes into sensible data. For example, TCP protocol which is implemented within OS is responsible for
  providing features like parsing and scheduling packets, dealing with acknowledgement (ACK), etc.
- **File System**: Storage is mostly blocks of bytes and earlier kernel accessed storage directly based on the
  implementation of disk at the time. This was however bad because whenever the disk had to evolve, it’d need to make 
  the change in OS as well.  To avoid this, OS abstracted access to disk using an API known as **LBA**
  (Logical Block Addressing). It assumes that disks are array of blocks of fixed size and the LBA will translate any 
  logical address to physical address on disk. But block address weren’t directly consumable by users, so it was then
  abstracted using file system which used identifiable objects with name, headers, types, size to map specific blocks 
  in storage to their respective file such that it can be accessed conveniently. There are different kinds of file
  systems like XFS, NTFS, FAT32, EXT4, etc. which are build around their specific use case or for general purpose. 
  With this, disk are exposed as big arrays of LBAs which can be divided into separate sections using *partitions*.
  Each partition is formed in specific file system and the data store in it can be only access in its respective format
  of file system.
- **Program and Process**: Programs are executable file and Process are running instance of a program.
  These executable file have header layout specific to OS.
- **Process Management**: Kernel manages the processes like schedules/switches them in CPU, providing them access to
  resource. However, for security reasons it's done in different mode. So process operates in two modes — 
  *user mode and kernel mode*, user mode involves all the general stuff a process executes and kernel mode involves
  access to resource via kernel. Each of these mode is allocated their own space in memory of process. 
  For example, your browser lives in user space of memory where it can operate on its instruction, but when it needs 
  to receive a request it’ll ask the kernel to open a socket, for which it switches to kernel mode and then the kernel
  sends over the request over designated socket.
- **Device Drivers**: Software in kernel which knows how to communicate with pieces of hardware. It works through
  *Interrupt Service Routine*, for example if we press any key on keyboard, it’ll interrupt the CPU to immediately
  execute the code to read the keypress from keyboard buffer, and sends it over to whatever is reading it.
- **System calls**: bridge between user space and kernel space. Apps make system calls to jump from user to kernel 
  space. Some of the system calls are `read`, `write`, `malloc`. It causes a mode switch in CPU to get higher 
  privileges for executing instruction, that’s why It's important to remain isolated for security.
