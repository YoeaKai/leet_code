from typing import List
from typing import operator


class Solution:
    def __init__(self):
        self.operators = {
            '+': lambda y, x: x + y,
            '-': lambda y, x: x - y,
            '*': lambda y, x: x * y,
            '/': lambda y, x: int(operator.truediv(x, y))
        }

    # Solution Using lambda.
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        for token in tokens:
            if token in self.operators:
                stack.append(self.operators[token](stack.pop(), stack.pop()))
            else:
                stack.append(int(token))

        return stack[0]

    # Intuitive solution.
    def evalRPN_2(self, tokens: List[str]) -> int:
        stack = []

        for token in tokens:
            if token not in "+-*/":
                stack.append(int(token))
            else:
                val1, val2 = stack.pop(), stack.pop()
                if token == "+":
                    stack.append(val1+val2)
                elif token == "-":
                    stack.append(val2-val1)
                elif token == "*":
                    stack.append(val1*val2)
                else:
                    stack.append(int(val2/val1))

        return stack.pop()
