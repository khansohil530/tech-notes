class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        stk = []
        op = {"+", "-", "*", "/"}
        for token in tokens:
            if token not in op:
                stk.append(int(token))
            else:
                b= stk.pop()
                a = stk.pop()
                if token == "+":
                    stk.append(a+b)
                elif token == "-":
                    stk.append(a-b)
                elif token == "*":
                    stk.append(a*b)
                elif token == "/":
                    stk.append(int(a/b))
        return stk.pop()