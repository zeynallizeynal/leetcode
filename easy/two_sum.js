var twoSum = function(nums, target) {
  const map = {};
  for (let i = 0; i < nums.length; i++) {
    const curr = nums[i];
    const diff = target - curr;
    if (diff in map) {
      return [map[diff], i];
    }
    map[curr] = i;
  }
};

console.log(twoSum([2, 7, 11, 15], 9));