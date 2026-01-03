from typing import Optional

class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        copyList = dict()
        copyList[None] = None
        curr = head
        while curr:
            copyList[curr] = Node(curr.val)
            curr = curr.next

        copyCurr, curr = copyList[head], head
        while curr:
            copyCurr.next = copyList[curr.next]
            copyCurr.random = copyList[curr.random]
            curr = curr.next
            copyCurr = copyCurr.next
        return copyList[head]
