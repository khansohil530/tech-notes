from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def _getNode(self, curr, n):
        while curr and n > 0:
            curr = curr.next
            n-=1
        return curr

    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        group_prev = dummy
        while True:
            kth_node = self._getNode(group_prev, k)
            if not kth_node:
                break
            group_next = kth_node.next

            prev, curr = group_next, group_prev.next
            while curr != group_next:
                nxt = curr.next
                curr.next = prev
                prev = curr
                curr = nxt

            tmp = group_prev.next
            group_prev.next = kth_node
            group_prev = tmp
        return dummy.next