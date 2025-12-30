class MinStack:

    def __init__(self):
        self.min_stk = []
        self.store_stk = []

    def push(self, val: int) -> None:
        if len(self.store_stk) == 0 or self.min_stk[-1] >= val:
            self.min_stk.append(val)
        self.store_stk.append(val)

    def pop(self) -> None:
        if len(self.store_stk) == 0:
            return

        val = self.store_stk.pop()
        if val == self.min_stk[-1]:
            self.min_stk.pop()

    def top(self) -> int:
        if len(self.store_stk) != 0:
            return self.store_stk[-1]

    def getMin(self) -> int:
        if len(self.store_stk) != 0:
            return self.min_stk[-1]
