# CPU

CPU (Central Processing Unit) as understood from its name, is a central component of computers responsible for processing.
This processing can be executing instructions for software or performing arithmetics and logical calculations or
controlling other components in computer to operate as a single unit(1) essentially acting as brain for computers.
{.annotate}

1. Like a brain in human body controlling different parts of body.

## CPU Architecture

![CPU IC](static/CPU-IC.png){align=right width=256px}
CPUs are packaged as a single integrated chip as shown on right, which is connected via motherboard. Internally,
this single chip is divided into different components like cores, shared caches and more. 

**CPU core** is an individual processing unit which can independently execute instructions, allowing parallel processing
in multicore CPUs. All these instructions which are understood by CPU are provided using an interface called **ISA** 
(Instruction Set Architecture) which defines essentials(1) required by software to execute their work. There are many 
different types of ISAs (2) which determines kinds of programs you can run and how efficiently they run.
When compiling a software into program, it's compiled for a specific ISA, and it only runs natively on CPU which 
implements respective ISA. Most of these implementation can be generally categorized into either RISC or CISC based. 
{.annotate}

1. like the **instructions**, **data types**, **registers**, **addressing modes** for main memory, **virtual memory**,etc.
2. like **x86-64**, **ARM**, **RISC-V**


###  RISC vs CISC
**RISC** or **Reduced Instruction Set** ISAs (like ARM) provide fewer simple instructions to perform tasks. This keep execution
predictable since each instruction is executed in single clock cycle, but the programs would need to provide multiple
instructions for executing a simple task. For example, you've to add two numbers, you've the following instructions: 
save value to register, add value in two registers, load value from register.

- save first number in a register
  - save second number in other register
  - add values in both register
  - save result in register. 


**CISC** or **Complex Instruction Set** ISAs (like x86-64) provides complex instruction which can perform multiple
steps using single instruction. As such the numbers of instructions are more than RISC ISAs, since you need to provide
permutation of all these complex instruction whereas RISC ISAs would simply compose its simple instruction. The key
advantage of using CISC is program can be executed in fewer instructions reducing the overhead of instruction 
translation but the execution is unpredictable because different instruction can take different number of clock cycles.
For example, to add two numbers CISC architecture would have single instruction for add two numbers.
Only one instruction can be used, but this would still require multiple steps to execute the process behind the scene.

## Instruction Cycle

To execute any instruction, CPU processes it through different stages. To provide an overview for general execution, 
CPU executes instructions in following stages:

1. **Fetch**: CPU retrieves the instruction from memory for execution. The address to next executing instruction is always maintained
   by CPU using **PC** (**Program Counter**) register. This address is fed to **IFU** (**Instruction Fetch Unit**)
   which looks up the instruction in memory and returns the instruction bits. The lookup in memory is done in following order: 
   L1I (L1 Instruction Cache) -> L2 cache -> L3 cache -> RAM/memory. Finally, the PC is updated to next instruction.
2. **Decode**: Now we've raw instruction bits loaded into the CPU, but it doesn't know where to head next or make
   sense of the instruction. **Instruction Decoder** helps in decoding this information, like the kind of operation
   (ADD, LOAD, CMP, MOV), registers required to read/write, and does it need to access memory or any other execution 
   unit. Additionally, for CISC CPUs instructions might also get broken into micro-operations (uOps). 
3. **Execute**: Depending on the kind of instruction, the instruction is sent to respective execution unit to perform 
   the operation. 

This pipeline can have additional stages based on the kind of instruction which requires a different path, or depending
on the architecture of CPU which introduces new stages to optimize its performance. 

### Major Instruction Types

1. **Arithmetic and Logical Instructions** like `ADD`, `AND` uses **ALU**. The ALU 
   typically takes two inputs with a control signal and outputs the result. The type of computation operation is 
   determined by the opcode decoded in Decode stage.
2. **Memory Instructions** like `LOAD` to read from memory and `STORE` to write to memory involves **MMU** (Memory 
   Management Unit). The lookup in memory is done following hierarchy from Registers -> LCaches -> Memory -> Storage. 
   This ordering is decided based on the speed of transferring information to/from register into respective level.
   Any access from lower level of memory will also save the information in high levels, so that further access is fast.

    ??? note "Data on Memory Lookup and Size"
        | Level                 | Size           | Approx Access Time | Approx CPU Cycles                 |
        |-----------------------|----------------|--------------------|-----------------------------------|
        | **CPU Registers**     | ~1 KB total    | ~0.3 – 1 ns        | **~1 cycle**                      |
        | **L1 Cache**          | 32–64 KB       | ~1 ns              | **~4 cycles**                     |
        | **L2 Cache**          | 256 KB – 1 MB  | ~3–4 ns            | **~12 cycles**                    |
        | **L3 Cache**          | 4–64 MB        | ~10–15 ns          | **~40–50 cycles**                 |
        | **Main Memory (RAM)** | 8–128 GB       | ~70–100 ns         | **~200–300 cycles**               |
        | **SSD (NVMe)**        | 256 GB – 4 TB  | ~50–150 µs         | **~150,000–400,000 cycles**       |
        | **HDD**               | 500 GB – 10 TB | ~5–10 ms           | **~15,000,000–30,000,000 cycles** |

3. **Branch Instructions** like `JMP`, `CALL`, `RET` are used to move ahead to a specific instruction due to execution of control
   statements like `if/else`, `for-loops` and `return` statements. These instructions require special execution unit due
   to **pipelining** of instruction.
   
     Pipelining is done to maximize the utilization of different CPU components to reduce its idle time. For example,
     when executing instructions sequentially only the currently running execution unit would be actively working while
     rest of the components would be idle. To avoid this, CPU pipelines multiple instruction such that while current
     instruction is in later stage of execution, the next instruction has already started its processing.

    ??? note "Pipelining Visually"
    
        ```text
        Instruction Stages:   FETCH → DECODE → EXECUTE → WRITEBACK
        Time ->
        
        Without Pipelining
        ┌───────────────────┐     ┌───────────────────┐     ┌───────────────────┐
        │ F -> D -> E -> W   │ -> │ F -> D -> E -> W   │ -> │ F -> D -> E -> W   │
        └───────────────────┘     └───────────────────┘     └───────────────────┘
          Instr 1                      Instr 2                   Instr 3
        
        With Pipelining
        ┌───────────────────┐     
        │ F -> D -> E -> W   │ -> 
        └───────────────────┘     
          Instr 1                                         
              ┌───────────────────┐     
              │ F -> D -> E -> W   │ -> 
              └───────────────────┘     
                 Instr 2
                    ┌───────────────────┐
                    │ F -> D -> E -> W   │ ->
                    └───────────────────┘
                      Instr 3
        ```

    But with control statements the execution of instructions becomes branched, as such pipelining becomes complex since
   the branched code might not include the already pipelined instruction. This essential makes the pipelined instructions
   after branching invalid, which the CPU have to flush out and restart the pipeline which is very inefficient. To 
   minimize the flushing of pipelines, CPu uses **Branch Predictors** which guesses if a branch will be taken or not.
   At the time of execution, if the branch is mispredicted the CPU will restart the pipeline but this happens rarely as
   branch predictors have almost ~95% accuracy in predicting the right branches.

## Optimizations 

Over the time, CPUs added new architectural optimizations to maximize performance and efficiency, out of which we've
already discussed one above, **Branch prediction** which allows speculative execution and keeps pipelines full by
guessing future control flow and executing ahead. Other commonly known techniques are Out of Order Execution, Register
Renaming, Prefetching, using specialized hardware which are discussed below.

### Out of Order Execution

Programs often contain operations that depend on the results of previous ones. In a simple in-order CPU, 
if one instruction stalls (for example, waiting on memory), everything behind it must also wait—even independent work.
To avoid wasting these cycles, **Out-of-order execution** allows a CPU to execute instructions as soon as their input
data is ready, rather than strictly following the original program order. The final results always appear as if
instructions were executed in order, but internally, the CPU rearranges them to avoid stalls.

This introduces 3 new step in above specified Instruction Lifecycle. After Fetch and Encode,

1. **Register Renaming**: This is done to avoid false dependencies resulting from blocked registers. Actual 
   dependency would be when an instruction needs a result produced earlier (also known as _Read-After-Write_). While other
   two dependencies, 
        
      - _Write After Read_, when instruction writes to register that an earlier instruction will read from.
      - _Write After Write_, when two instructions write the same register.
   
      aren't really dependency if we can just rename the registers to used different ones. This renaming is done to 
   temporary registers also known as **physical registers**. These are different from architectural registers specified 
   by ISA, and they're present in larger pool (180 in x86-64).

2. **Instruction Scheduling**: , Instructions whose operands are ready goes into a 
   pool where a scheduler selects ready instructions and issues them to execution units—regardless of original order.
3. **Reorder Buffer (ROB) & In-Order Retirement**: Even though execution is out of order, results are written in order.
   This is ensured by a new hardware component **ROB** which holds results temporarily and update them in 
   registers/memory only when all older instructions have finished cleanly. Additionally, it ensures results from 
   branch misprediction are discarded. 

The benefit of using Out of Order Execution are improved throughput and performance. While the tradeoff being,
increase CPU complexity from components like ROB, dependency tracking, physical registers, higher power consumption 
and heat production which is why simpler CPUs (e.g., microcontrollers) use in-order execution.

### Prefetching

A CPU can execute billions of instructions per second, but a memory access can take hundreds of cycles.
Memory access is the most common reason for stalling pipeline. Prefetching reduces these stalls by predicting what
data will be needed soon, and it’s already waiting in a cache instead of main memory. Data needed is predicted using 
programs access behaviour. Prefetching can be either implemented automatically using hardware built in CPU or using 
software where programmer or compiler can use special instructions to hint CPU to load data. This is very useful in 
manually optimizing code and high-performance computing (HPC) workloads.

### Accelerators

CPUs are extremely flexible but aren't efficient at highly parallel tasks, repetitive math operations ,large matrix
computations, graphics and image processing and machine learning workloads. To optimize these tasks, we can use 
**Accelerators** which are specialized hardware designed to speed up specific computations like 

- **GPUs** (Graphics Processing Units) for rendering, machine learning, and video encoding.
- **TPUs** (Tensor Processing Units) for matrix multiplications, neural network training.
- **DSPs** (Digital Signal Processors) for signal-processing tasks like, audio, radar, telecommunications, compression. 
- **ASICs** (Application-Specific Integrated Circuits) for fixed tasks like video encoders/decoders (H.264, HEVC),
  crypto accelerators (AES, SHA), SSD controllers. 

They usually work alongside the CPU in a system, connected through either PCIe (GPUs, NICs), on-chip interconnects
(NPUs, image processors) or memory-mapped interfaces. Typical CPU sets up the task for accelerator which processes data
and returns the results while CPU continues main program flow. This offloading frees the CPU to handle logic, branching,
and system tasks. This allows CPU to achieve more parallelism and higher throughput making accelerators essential
for workloads like AI, graphics, data analytics, and networking demand.

## Conclusion

To conclude, let's look at the execution of following C program which will put all above topics together.

```c
for (i = 0; i < N; i++) {
    A[i] = B[i] * C + D;
}
```
This is compiled into following assembly code

```asm
L1: LOAD R1, [B + i*4]        ; load B[i]
    LOAD R2, C                ; load constant C
    MUL  R3, R1, R2           ; R3 = B[i] * C
    ADD  R4, R3, D            ; R4 = R3 + D
    STORE [A + i*4], R4       ; A[i] = R4
    ADD  i, i, 1              ; loop counter
    CMP  i, N                 
    JLT  L1                   ; loop branch

```

This code is executed in CPU as follows:

1. **Prefetching & Instruction Fetch**: As the loop runs repeatedly, the CPU’s instruction prefetcher notices a sequential 
   pattern and fetches multiple future instructions before they’re needed. Prefetcher reads ahead into the I-cache. 
   Branch predictor predicts the loop branch as taken. This keeps instructions flowing without waiting for memory.
2. **Decode + Micro-Op Translation**: Instructions move to the decode stage, CISC instructions (like x86) are broken into 
   micro-ops. CPU identifies loads, stores, ALU ops, and branch ops. The branch predictor supplies the predicted next
   PC to keep the pipeline full.
3. **Register Renaming**: Architectural registers (R1, R2, R3, etc.) are mapped to physical registers for example, R1 ->P7, 
   R2 -> P12, R3 -> P9, R4 -> P14. Every new instruction that writes a register gets a new physical register. 
   This removes WAW/WAR dependencies so instructions can run in parallel.
4. **Dispatch into Reservation Stations & ROB**: Each instruction allocates a ROB entry (for in-order retirement) and 
   a reservation station entry (to wait until operands are ready). For example, ROB after decoding an iteration
   
    | Entry | Instruction | Physical Dest | Ready? |
    |-------|-------------|---------------|--------|
    | 0     | LOAD B[i]   | P7            | no     |
    | 1     | LOAD C      | P12           | yes    |
    | 2     | MUL         | P9            | no     |
    | 3     | ADD         | P14           | no     |
    | 4     | STORE       | —             | no     |

5. **Out-of-Order Execution (OOO)**: The CPU checks which instructions are ready,
    
       ```asm
        LOAD C        ;(executes immediately)
        ADD i, i, 1   ;(independent; executes early)
        CMP i, N      ;(also executes early)
        LOAD B[i]     ;(waiting on memory)
        MUL           ;(waits for B[i])
        ADD R4        ;(waits on MUL result)
        STORE         ;(waits on ADD)
       ```    

6. **Cache System & Prefetchers**: Because the loop accesses arrays B and A sequentially, the hardware data
   prefetcher recognizes a streaming pattern and starts fetching future `B[i+1]`, `B[i+2]` and `A[i+1]` cache lines. 
   So future iterations hit in L1 or L2, reducing stalls.
7. **Results Become Ready in Reservation Stations**: When the `LOAD B[i]` returns from memory, it broadcasts its result 
   via the **Common Data Bus** (CDB): Waiting MUL is awakened, MUL runs, ADD depends on MUL → executes afterward.
   Everything continues smoothly.
8. **In-Order Retirement Through ROB**: Even though execution was out-of-order, results become architecturally visible
   in program order. The ROB commits entries:
      - Commit `LOAD B[i]`
      - Commit `LOAD C`
      - Commit `MUL` 
      - Commit `ADD` 
      - Commit `STORE`

    If a branch misprediction occurred, the CPU flushes the ROB, rolls back register mappings  and restarts from 
   correct path

??? note "Analyzing CPU usage in Linux using top"
    Analyzing CPU usage for a programs give you better insight on its execution performance, and decide whether
    its CPU bound or I/O bound. Running `top` in terminal gives you information similar to following fields at top,
    ```shell
    %Cpu(s):  5.0 us,  2.0 sy,  0.0 ni, 90.0 id,  3.0 wa,  0.0 hi,  0.0 si,  0.0 st
    ```
    
    - **us** -> Time spent running *user* processes
      - **sy** ->Time spent in *kernel/system* processes
      - **id** -> Idle time
      - **wa** -> Time spent *waiting for I/O*, tells if a process is CPU bound or I/O bound
      - **hi/si** -> Hardware/software interrupts
      - **st** -> Steal time (virtualized environments)
    
    
    In the per process stats, we get following fields:
    ```shell
    PID USER  PR NI  VIRT  RES  SHR S %CPU %MEM  TIME+ COMMAND
    ```
    
    - **%CPU** tell	% of CPU the process is consuming
      - **S** is state of process, like Running (R), Sleeping (S), Uninterruptible IO sleep (D), etc.
      - **TIME+** tells total CPU time used
      - **COMMAND** is the name of process