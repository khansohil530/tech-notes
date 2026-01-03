class ListNode:
    def __init__(self, key=0, val=0):
        self.key, self.val = key, val
        self.prev = self.next = None

class LRUCache:

    def __init__(self, capacity: int):
        self.cap = capacity
        self.data = dict()

        self.lru, self.mru = ListNode(), ListNode()
        self.lru.next, self.mru.prev = self.mru, self.lru

    def remove(self, node):
        prev, nxt = node.prev, node.next
        prev.next, nxt.prev = nxt, prev

    def insert(self, node):
        prev, nxt = self.mru.prev, self.mru
        prev.next = nxt.prev = node
        node.prev, node.next = prev, nxt

    def get(self, key: int) -> int:
        if key in self.data:
            self.remove(self.data[key])
            self.insert(self.data[key])
            return self.data[key].val
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.data:
            self.remove(self.data[key])

        self.data[key] = ListNode(key=key, val=value)
        self.insert(self.data[key])

        if len(self.data) > self.cap:
            lru = self.lru.next
            self.remove(lru)
            del self.data[lru.key]

