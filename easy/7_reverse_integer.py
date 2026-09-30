class Solution:
    def reverse(self, x: int) -> int:
        INT_MAX = 2**31 - 1
        INT_MIN = -2**31
        
        sign = 1

        if x < 0:
            sign = -1
            x*=-1
        
        res = 0
        level = 1
        temp = x
        while temp // 10 > 0:
            level*=10
            temp//=10
            
        while x > 0:
            res+=x%10*level
            x//=10
            level//=10
            
        res*=sign
        if res > INT_MAX or res < INT_MIN:
            return 0
        return res


print(Solution().reverse(1))