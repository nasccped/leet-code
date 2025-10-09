class Solution(object):
    open_allowed = ['(', '[', '{']
    close_allowed = [')', ']', '}']
    stack = None

    def pushOpen(self, char):
        """
        Turns the given char into the suitable parentheses
        representation and push it to the self.stack.

        The char safeness is checked within isValid function. Passing
        a non "()[]{}" char to this function can result into an
        unexpected result.
        """
        char = 'r' if char == '(' \
               else 's' if char == '[' \
               else 'c'
        self.stack.append(char)

    def extractAndCheck(self, char):
        """
        Pops the top item from the self.stack and compare with the
        char representation, returning a boolean (True -> ok,
        False -> not ok). The given char is expected to be an
        self.close_allowed item, otherwise this function will return True.
        """
        if not self.stack:
            return False
        extracted = self.stack.pop()
        if char == ')' and extracted != 'r':
            return False
        elif char == ']' and extracted != 's':
            return False
        elif char == '}' and extracted != 'c':
            return False
        return True

    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        self.stack = []
        for c in s:
            if c in self.open_allowed:
                self.pushOpen(c)
            elif c in self.close_allowed:
                if not self.extractAndCheck(c):
                    return False
            else:
                return False
        return not self.stack
