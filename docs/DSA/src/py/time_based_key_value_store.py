from collections import defaultdict


class TimeMap:

    def __init__(self):
        self.data = defaultdict(list)


    def set(self, key: str, value: str, timestamp: int) -> None:
        self.data[key].append((value, timestamp))

    def get(self, key: str, timestamp: int) -> str:
        items = self.data[key]
        if not items: return ""
        left, right = 0, len(items)
        while left < right:
            mid = left + (right-left)//2
            if timestamp < items[mid][1]:
                right = mid
            else:
                left = mid + 1
        return items[left - 1][0] if left > 0 else ""
