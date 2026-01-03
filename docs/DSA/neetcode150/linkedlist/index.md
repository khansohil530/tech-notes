---
tags:
  - Linked List
hide:
  - toc
---

# Linked List

## Floyd's Cycle Detection Algorithm

Also known as Tortoise and Hare algorithm, uses two pointers $slow$ and $fast$ to traverse the linked list at different
speed to detect cycle within a linked list. The key insight is that if a cycle exists, the $fast$ pointer will 
eventually "lap" the $slow$ pointer, causing them to meet at some point within the cycle. If there is no cycle, the 
fast pointer will reach the end of the list.

- Time Complexity: $O(n)$, as each node is visited a constant number of times.
- Space Complexity: $O(1)$, as it only uses two pointers and no extra data structures. 

??? note "Algorithm Pseudocode"
    ```mermaid
    flowchart TD
        A([Start]) --> B[Initialize slow = head<br/>fast = head]
        B --> C{fast != null<br/>AND fast.next != null?}
        C -- No --> D([No Cycle Detected])
        C -- Yes --> E[slow = slow.next<br/>fast = fast.next.next]
        E --> F{slow == fast?}
        F -- Yes --> G([Cycle Detected])
        F -- No --> C
    ```

Proof of argument "_Once both pointers are inside the cycle, the fast pointer gains one node per step on the slow 
pointer, so it must eventually lap and meet it._"

- Suppose, there is a cycle and both pointers enter the cycle when iterating. Let the length of cycle be $C$. 
- Inside the cycle, the $fast$ pointer closes the distance b/w $slow$ pointer by one node per iteration. 
- If $d$ is the initial distance b/w $fast$ and $slow$, the distance after $k$ iterations can be written as
  $(d-k) \mod C$.  

    !!! note "" 
        Inside cycle, node positions are determined modulo $C$ because even if both pointers are at the same node their 
        step counts differ by a multiple of $C$

- Since $k$ is incremented by $1$ per iteration, the distance $(d-k) \mod C$ will eventually reach $0$, i.e.
  $fast == slow$.    

---

"**_Why the meeting point is not necessarily the cycle entry?_**"

Let: 

- $\mu$ = number of nodes before the cycle starts
- $C$ = length of the cycle

When slow and fast meet, they are somewhere inside the cycle, but not necessarily at its start.

- Suppose $slow$ has taken $k$ steps  and $fast$ has taken $2k$  within cycle steps before meeting.
- Inside the cycle, positions repeat every $C$ steps. Since they meet at same point, the extra distance traveled by 
  fast must be an exact whole number of cycles.
- $=> distFast - distSlow = m \times C$ for some integer $m$ $=> 2k+\mu - k+\mu = k$ $=> k$ is a multiple of $C$.
- Now, the distance travelled by $slow = k = \mu + x$, where $x$ is how far into the cycle the meeting occurs.
- Since $\mu$ isn't necessarily multiple of $C$, $x$ isn't always guaranteed to be $0$, making the meeting point an
  offset from the entry of cycle.

## Detect the start of cycle

From earlier, $=> k = \mu + x = m \times C$, $=> \mu = m \times C − x$ which means distance from meeting point to 
cycle entry is same as $\mu$. Using this information, we can find the start of cycle 

- Using two pointers, one at meeting point and another at start of linked list
- Iterate them at same pace. The node at which they meet will be the start of cycle.

