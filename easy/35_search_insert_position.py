class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        min = 0
        max = len(nums) - 1
        while min <= max:
            mid = (max + min) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                min = mid + 1
            else:
                max = mid - 1
        return min


# print(Solution().searchInsert([1,3,5,6], 5))
print(Solution().searchInsert([1, 3, 5, 6],2))
