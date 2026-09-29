class Solution:
    def myAtoi(self, s: str) -> int:
        multiplier = 1
        max_digit = 57
        max_32int = 2**31 - 1
        min_32int = -2**31
        
        # remove whitespace
        s = s.strip(' ')
        if s[0] in "-+":
            if s[0] == '-':
                multiplier = -1
                
            # remove leading sign
            s = s[1:]
        
        
        # remove leading zero
        s = s.removeprefix("0")

        
        if len(s) == 0:
            return 0
        
        res:int = 0
        level:int = 1
        valid_str = ""
        
        for i in range(0, len(s)):
            ascii_val = ord(s[i])

            # is valid digit
            if (max_digit - ascii_val >=0) and (max_digit - ascii_val)<=9:
                valid_str+=s[i]
            else:
                break
        
        for i in range(len(valid_str)-1, -1,-1):
            ascii_val = ord(s[i])
            res += (9-(max_digit - ascii_val)) * level
            level*=10
        
        res *=multiplier
        if res > max_32int:
            return max_32int
        elif res < min_32int:
            return min_32int
        return res
    
print(Solution().myAtoi("-04dŷ2"))