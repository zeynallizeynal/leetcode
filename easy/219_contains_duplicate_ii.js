/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function(nums, k) {
    var m = new Map();
    for (var i = 0; i < nums.length; i++) {
        var hasVal = m.has(nums[i]);
        if (hasVal && Math.abs(i-m.get(nums[i]))<=k) {
            return true;
        }
        m.set(nums[i], i);
    }
    return false;
};