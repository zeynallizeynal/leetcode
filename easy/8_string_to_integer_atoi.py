class Solution:
    def myAtoi(self, s: str) -> int:
        multiplier = 1
        max_digit = 57
        
        # remove whitespace
        s = s.strip(' ')
        if s[0] in "-+":
            if s[0] == '-':
                multiplier = -1
                
            # remove leading sign
            s = s[1:]
        
        
        # remove leading zero
        s = s.removeprefix("0")

        
        res:int = 0
        level:int = 1
        
        for i in range(len(s) - 1, -1, -1):
            ascii_val = ord(s[i])

            # is valid digit
            if (max_digit - ascii_val >=0) and (max_digit - ascii_val)<=9:
                res += (9-(max_digit - ascii_val)) * level
                level*=10
        
        return res * multiplier
    
print(Solution().myAtoi("-04dŷ2"))