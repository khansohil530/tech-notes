# Program and Process 

Process are logical construct used to encapsulate a running program.
It's made up of all the metadata and content required by OS to manage and run a program,
like entrypoint of program, the shared libraries which needs to be loaded, user data, program instructions, etc.

But how does OS understand where to get these information? To understand this, we'll need to look into the layout
of program files.

## Program

*Program* is an executable file which contains instructions for kernel about how to execute it’s work. 
These instructions/codes are complied (1) and linked (2) for a CPU.
{.annotate}

1. *Complied* means the instruction should be converted into codes understandable by CPU (which are 0s and 1s).
2. *Linked* means combining all different libraries and source object files into a single *execution file*. This can be
   done statically or dynamically at runtime.

??? note "Dynamic and Static Linking"
    You might notice this difference in linking when copying games. When we copied just the executable of a game from
    some friend and try to execute it on our laptop we’d get an error saying some DLL files are missing.
    These DLL files are libraries in windows which are linked in the executable dynamically. We could also statically
    combine all these libraries in same executable, but that’d increase the size of `exe` drastically which is why 
    they're kept dynamic for ease of distribution.

This executable file has specific format, (for example Unix like OS uses ELF (1) layout for its binaries) which the 
OS understands, using which it can load the process with required content and metadata to execute provided instructions.
{.annotate}

1. **E**xecutable and **L**inked **F**ormat

### ELF 
ELF is the standard binary format used on Linux/Unix systems for, Executable programs, Shared libraries (`.so`),
Object files (`.o`) and Core dumps.  It describes how code, data, libraries, and metadata are stored so the OS
loader can load and run the program. At a high level, an ELF file has following structural layers:
```text
+-----------------------+
| ELF Header            |  ← describes the whole file
+-----------------------+
| Program Headers       |  ← used at runtime for loading
+-----------------------+
| Sections (e.g., .text, .data, .bss)   |
+-----------------------+
| Section Headers       |  ← used by linkers/debuggers
+-----------------------+
```

- **ELF Header** contains information like the magic number `0x7F 45 4C 46` (1), CPU architecture for which the file is compiled, type of file, entrypoint
  address of program, offsets for program and section header.
    {.annotate}
    
    1. `0x7F  'E'  'L'  'F'` in ASCII, used to identify the file as an ELF binary and allows the kernel loader to recognize
       and parse it.

- **Program Headers** describes various segments to be mapped into process memory by OS loader. Each segment specifies its 
  type, file offset, size in file, size in memory, and memory permission (r/w/x  operations). For example,
    - `PT_LOAD` -> used to load data and instructions into process memory. For example, mapping code segment using 
      `.text` and `.rodata` sections with `r-x` permission. Similarly, data segments are loaded from `.data`, `.bss` 
      with `rw-` permission.
    - `PT_INTERP` -> specifies the dynamic loader for dynamically linked executables. Like `/lib64/ld-linux-x86-64.so.2`
      for x86_64 Linux.
    - `PT_DYNAMIC` -> contains `.dynamic` section data used by dynamic loader. 

- **Sections** and **Section Headers** consists of file content used for linking and debugging purpose only. They hold content
  as raw data from various part of program, like 
    - `.text` -> executable machine code instructions
    - `.data` -> initialized global and static variables
    - `.bss` -> uninitialized global and static variables
    - `.rodata` -> read-only data
    - `.symtab` -> symbol table


Checkout below sequence diagram to understand complete flow of loading a program into process using ELF format:

```mermaid
sequenceDiagram
    participant User as User / Shell
    participant Kernel as Linux Kernel
    participant ELF as ELF Executable
    participant Loader as Dynamic Loader (ld.so)
    participant Proc as New Process

    User->>Kernel: execve("program", argv, envp)
    Kernel->>ELF: Read ELF Header<br/>Check magic: 0x7F 'E' 'L' 'F'
    Kernel->>ELF: Read Program Header Table
    Kernel->>Proc: Create new process<br/>Create address space

    loop For each PT_LOAD segment
        Kernel->>Proc: Map PT_LOAD segment<br/>into virtual memory (R/W/X)
    end

    alt ELF contains PT_INTERP?
        ELF->>Kernel: PT_INTERP = "/lib64/ld-linux-x86-64.so.2"
        Kernel->>Loader: Load dynamic loader<br/>Map its segments
        Kernel->>Loader: Transfer control<br/>to loader entry point (user mode)
        Loader->>Proc: Resolve shared libs<br/>Relocations & symbol binding
        Loader->>Proc: Jump to ELF entry point (_start)
    end
   
    Kernel->>Proc: Set up user stack<br/>argc, argv, envp, auxv
    Kernel->>Proc: Jump directly to<br/>ELF entry point (_start)
    Proc-->>User: Program runs (_start → main)

```

Now you have overview on how a program is loaded into memory to form a Process, let's understand how Process
executes instructions in program to perform the coded work.

## Process

A process is what a program becomes after the kernel loads its segments, sets up virtual memory (code, data, heap, stack),
prepares registers & CPU state, and begins executing at the entry point. It's footprint can be divided into 3 major
categories:

1. **User-Space Memory**, which is a private memory space assigned to process during its creation. It's a continuous chunk of
   virtual memory associated with a high and low memory address. You can check below memory layout used for typical linux process
    ``` 
    +-------------------------------+ High address
    |        Stack (grows ↓)        |   |
    +-------------------------------+   |
    |  Memory-mapped region (mmap)  |   |
    |  ← Shared libraries live here |   |
    |  ← also VDSO, ld.so, JIT code |   |
    +-------------------------------+   |
    |        Heap (grows ↑)         |   |
    +-------------------------------+   |
    |  .bss / .data / .text         |   ↓
    +-------------------------------+ Low address 
    ```
   
2. **CPU Execution Context**, the state of CPU needed to resume the process after a context switch. It includes values for various
   pointers like `pc`(1), `sp`(2), `bp`(3), etc. We'll discuss how these pointers are used later.
    {.annotate}
    
     1. Program Counter, points next instruction to execute
     2. Stack Pointer, top of user stack
     3. Base Pointer, points to start of currently executing function frame

3. **PCB** (1) which stores information about process like ids (PID, PPID, UID, GID(3)), scheduling info,
   MMU (2) structs, table of open file descriptors, process running status and threading info. This area of memory
   is only accessible to Kernel for security purposes.
   {.annotate}

    1. **P**rocess **C**ontrol **B**lock
    2. Memory Management Unit, describing the virtual memory mapping to physical. 
    3. Process ID, Parent Process ID, UserID, GroupID 

To get a simple understanding how process execution happens, go through below diagram which we'll continue to explore 
in depth.

```mermaid
sequenceDiagram
    autonumber

    participant Kernel
    participant Memory
    participant CPU

    %% 1. Process created and text loaded
    Kernel->>Memory: Loads and map process memory
    Kernel->>CPU:  Sets initial values to CPU registers (PC -> entrypoint, SP -> top of stack)
    Kernel->>CPU: Schedules process on CPU

    %% 3. CPU fetches instruction into IR
    CPU->>Memory: Fetch instruction at PC address
    Memory->>CPU: Load instruction set into IR  (instruction register)

    %% 4. Execute and increment PC
    CPU->>CPU: Execute instruction and Increment PC to point next instruction

    %% 5. Use cache for next instruction
    CPU->>CPU: Check LCaches for next instruction
    alt Instruction cached
        CPU->>CPU: Load instruction from cache
    else Not cached
        CPU->>Memory: Fetch instructions in burst
        Memory->>CPU: Fill cache line
        CPU->>CPU: Provide next instruction
    end

    %% 6. PC not saved unless context switch
    CPU-->>Kernel: Continue executing until preemption
    Kernel->>Memory: Save registers on stack only during context switch
    Memory->>CPU: Restore registers from stack when rescheduled

```

### Stack

Stack as seen above is part of User-Space Memory of process. The primary role of stack is to keep track of function calls
such that CPU can jump to previous function after completing current function. This is achieved by using various pointers
like `sp`(1), `bp`(2), `lr`(3), etc. Other roles includes storing functions local variables or temporary register values
onto stack.
{.annotate}

1. Stack Pointer, CPU register which points to end of current function frame.
2. Base Pointer, CPU register which points start of function frame.
3. Link Register, CPU register which stores the address of instruction after function call. 

To reference variables stored in stack, you can use `sp` since its dynamic and keeps changing. As such another pointer
`bp` is used. Since `bp` constantly points to top of function frame, you can easily reference variable address relative 
to it, for example `a->bp`, `b->bp-4`, `c->bp-8` (where size of each variable is 4 bytes). With this information,
we can explain how stack is used when executing a process,

1. **When we call a new function**,
    ```mermaid
    sequenceDiagram
        autonumber
    
        participant CPU as CPU
        participant SP as SP (Stack Pointer)
        participant BP as BP (Base Pointer)
        participant Stack as Stack Memory
        
        note over CPU,Stack: Function call instruction is executed
        %% Function Call
        CPU->>Stack: push return address
        note right of SP: SP moves DOWN<br/>(decrement)
        SP-->>SP: SP = SP - addr_size
    
        CPU->>Stack: push old BP
        note right of SP: SP moves DOWN again
        SP-->>SP: SP = SP - addr_size
    
        CPU->>BP: BP = SP
        note right of BP: New frame base established
    
        %% Inside Function
        CPU->>Stack: allocate locals (SP = SP - frame_size)
        note right of SP: SP moves DOWN for locals
       
        note over CPU,Stack: Execution continues using locals<br/>until function prepares to return
    ```
   
    !!! note ""
        Some CPU architecture like ARM consists of `lr` which stores the return address for current frame. Others 
            like x86 don't have such registers, as such the return address is pushed onto stack. 

2. **While we're in a function**, we can store local function variable or temporary register values like `lr` or `bp` in stack
   locally and reference them w.r.t `bp` of frame.
   ```mermaid
   sequenceDiagram
       autonumber

       participant CPU as CPU (Executes Instructions)
       participant SP as SP (Stack Pointer)
       participant BP as BP (Base Pointer)
       participant Stack as Stack Memory

       %% Function has already been called
       note over CPU,Stack: We are now INSIDE the function<br/>Stack frame has been created

       CPU->>Stack: Reserve local variable <br/>e.g., int x

       %% Storing variables
       CPU->>Stack: Write value of local var x<br/>(at BP - offset)
       Stack-->>CPU: Store complete

       %% Using local variables
       CPU->>Stack: Read local x via (BP - offset)
       Stack-->>CPU: Return value of x

       %% Temporary values / spills
       CPU->>Stack: Spill register value to stack
       SP-->>SP: SP moves DOWN (push)
       Stack-->>CPU: Load spilled value later (pop)
       SP-->>SP: SP moves UP (pop)

       note over CPU,Stack: Execution continues using locals<br/>until function prepares to return
   ```

3. **When we return from a function**, 
    ```mermaid
    sequenceDiagram
        autonumber
    
        participant CPU as CPU (Executes Instructions)
        participant SP as SP (Stack Pointer)
        participant BP as BP (Base Pointer)
        participant Stack as Stack Memory
    
        note over CPU,Stack: Function has finished executing<br/>Now preparing to RETURN to caller
    
        %% Step 1 — Deallocate locals
        CPU->>SP: Move SP back to BP (SP = BP)
        note right of SP: SP jumps UP<br/>removing local variables
    
        %% Step 2 — Restore old BP
        CPU->>Stack: pop saved BP
        Stack-->>BP: Write old BP into BP register
        note right of BP: BP now points to caller's frame
    
        %% Step 3 — Load return address
        CPU->>Stack: pop return address
        Stack-->>CPU: Return address loaded
    
        %% Step 4 — Jump back to caller
        CPU->>CPU: RET instruction<br/>PC = return address
        note over CPU,Stack: CPU resumes executing caller function
    ```


Access to memory is very costly for CPU, but having them laid out next to each other helps a lot due to caching few next
required instruction/variable with single burst. Also, Memory allocation and deallocation in stack is managed using `sp`
is very fast, 

- to allocate new memory, you increase the `sp` and give the new space to required variables/functions.
- to clear up memory, you can decrement the `sp` to mark the memory outside it as garbage which can then be cleared or
  overwritten. 

!!! note "Best coding practices"
    Few takeaway from understanding this design of execution:

    - Function calls are expensive, as we’ve to move around data between register and stack memory. 
      So avoid using too many function without any cause. Compilers even optimize this by using *inlining*, where
      it inserts machine code of a function inline to where it was called but this also bloats the code if used too much.
    - Stack has limited space, which protects the process from infinite function calls in case of recursions.
    - Avoid using large local variables, every step which requires fetching value from memory is considered expensive in CPU.    


### Data Section

Fixed size section in memory layout of process which is responsible for storing program instructions, constants and 
global variables. Its size is determined by compiler during compilation using static analyses of code.
The section is further divided into following subsections:

- `.text` memory section stores program instructions, function bodies, CPU opcodes. It's only given read-execute 
  permission for security, so that any marlware can't edit code during its execution. Program Counter (`pc`) fetches
  instructions from this section, which are then decoded and executed by CPU. 
- `.rodata` stores read-only data like constant variables, string literals, etc. This separate memory section is created
  so that any constant values isn't modified accidentally. Another benefit can be ease of sharing same data with other processes.
- `.data` stores initialized global variables. Since these variables are available across all functions, CPU directly
  references them with ease. 


The variables stored in different sections are addressed using offset based on start of data section . And the offset
is calculated by compiler during compilation.

### Heap

Heap section is responsible for storing/referencing large dynamic variables in memory. However, the data needs to be
removed explicitly, and if not handled properly you might have memory leaks (1). It grows from lower to higher memory
address. Kernel provides you with 3 APIs to manage memory in heap: `malloc`, `free` and `new`. 
{.annotate}

1. When data in memory isn’t referenced by any function in stack.

To access data stored in heap, you’ve to use *Pointers*. Pointers are variables which stores memory address of first
byte of data stored in heap. And based on the type of pointer which tells us the size of data, we can fetch the 
required bytes to get complete data.

During memory allocation, we’ve to specifying the memory size required. But freeing memory doesn't require you to mention
the memory size. This is done using fixed size headers attached to pointer location which holds metadata on the
allocated memory. Kernel uses that information to determine how much memory to free. Also, whenever we ask Kernel to
allocate some memory, it’ll always return it in some multiple of memory page size and not the exact memory size asked.

Few things to know when using heap:

- **Memory leak**: When memory isn’t freed up, the Kernel will still keep the data in memory even if it isn’t used in 
  any function. This leads to unwanted memory growth known as memory leak. High level programming language uses garbage 
  collection algorithms to avoid this, one of which is `refcounting` which stores the number of references in use to 
  the data within the header. If the reference becomes 0, means nothing is pointing to this data and as such garbage
  collector can free up this place.
- **Dangling Pointers**: When the original data your pointer references is freed up (like in a downstream function 
  call) and when you try to access it, you’ll read random headers leading to errors like segfault. For example, 
  when you try to free a pointer twice, it leads to the crash of process.
- **Performance**: Heap is slower compared to stack because you’ve to go to allocate memory, read headers to fetch the
  data, and free memory while stack doesn’t involve such tedious process. Stack also has locality of related data which
  are cached when reading in burst while heap is unorganized. Stack space is limited but heap can grow.
    
    ??? note "Google TCP/IP Performance Boost"
          Google improved performance of TCP/IP stack in Linux Kernels by 40% just by re-ordering the variables in order 
          they were accessed by kernel. This significant improvement was due to the locality of data which resulted in 
          caching from burst. So at Kernel level, always try to cache things and not take memory for granted, going to
          memory might look fast for one instruction but over millions of instruction these things add up.
    
- **Escape Analysis**: Some languages (like Java, Go) allocate memory within stack itself whenever possible to avoid
  the cost of heap. They’ll create a pointer which points to a memory location within stack itself. This is mostly
  applicable in places where we don’t pass a pointer outside current function.
- **Program Break**: Older version used `brk` / `sbrk` functions to allocate/deallocate memory from heap which basically
  added a break at top of heap and whenever a section isn’t used this memory would be deallocated. This is inefficient
  because data is placed/freed randomly in heap, as such it's very rare for block to go completely free. 
  It was later updated to `MMAP` which resolved this issue.

!!! note "View process layout in Linux"
    In Linux, you can view the internals of process using the command `cat /proc/{PID}/`.
    This exposes API to view metadata of process which can be used to create tools like resource monitors. 
    Also `/proc` isn’t a physical file system present on disk, its only present in memory.

!!! warning ""
    CPU-Context and Kernel-Space for process will be discussed in following chapters to keep them in flow with respective
    topic.