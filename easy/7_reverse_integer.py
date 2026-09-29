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
            
        while x > 0 and level>0:
            x = x - (x//level * level)
            level//=10
            print(x)
            
        
        return res


print(Solution().reverse(-321))