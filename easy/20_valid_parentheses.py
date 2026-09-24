def isValid(s: str) -> bool:
    if len(s) == 1:
        return False

    stack = []
    openings = ["(", "[", "{"]
    closings = [")", "]", "}"]

    if s[0] in closings:
        return False

    for i in s:
        if len(stack) == 0 and i in closings:
            return False
        if i in openings:
            stack.append(i)
            continue
        if i in closings and stack[-1] == openings[closings.index(i)]:
            stack.pop()
        else:
            return False

    return len(stack) == 0


print(isValid("(){}}{"))
