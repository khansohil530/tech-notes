---
comments: false
hide:
  - footer
  - toc
---

# Foundations of Data System

Many applications today are *data-intensive*, as opposed to *compute-intensive*. Raw CPU power is rarely a limiting 
factor for these applications — bigger problems are usually the amount of data, the complexity of data, and the speed 
at which it’s changing. Data Intensive applications are commonly build using standard building blocks like *databases*
to store data, *cache* to speed up reads or store expensive results, *search indexes, stream processing, batch 
processing, etc*.

These data systems are such a successful abstractions that most of the time we don’t need to think about writing them
twice. But in reality, every application has different usecases which needs different requirements. As such these 
data systems evolves into different stream to support their specific use cases. So now as a software engineer, we need
this knowledge to figure out tools which solves our task at hand. Sometimes you might need to combine multiple tools to 
achieve the task, while sometimes single tool can suffice. 

This page discuss all these technicalities, exploring all these tools, what’s common among them and what distinguishes
them.

<div class="grid cards" markdown>
- :fontawesome-solid-circle-check:{ .status-pending } [1. Reliable, Scalable and Maintainable Applications](c1.md)
- :fontawesome-solid-circle-check:{ .status-pending } [2. Data Models and Query Languages]()
- :fontawesome-solid-circle-check:{ .status-pending } [3. Storage and Retrieval]()
- :fontawesome-solid-circle-check:{ .status-pending } [4. Encoding and Evolution]()
</div>