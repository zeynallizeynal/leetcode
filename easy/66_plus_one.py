class Solution:
    def plusOne(self, digits: list[int]) -> list[int]:
        num = 0
        level = 1
        for i in range(len(digits) - 1, -1, -1):
            num += digits[i] * level
            level *= 10

        num += 1
        res:list[int] = []
        while num>0:
            res.insert(0, num%10)
            num//=10
        return res


print(Solution().plusOne([9]))
