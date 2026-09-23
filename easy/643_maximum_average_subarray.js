/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function (nums, k) {
  var res = 0;
  for (var i = 0; i < k; i++) {
    res += nums[i];
  }

  var cur = res;
  for (var i = 0; i <= nums.length - k; i++) {
    cur-=nums[i];
    cur+=nums[i+k];
    if (cur > res) res = cur;
  }
  return res/k;
};