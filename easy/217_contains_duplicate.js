/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    var m = new Map();
    for (var i = 0; i < nums.length; i++) {
        if (m.has(nums[i])) {
            return true;
        }
        m.set(nums[i], i);
    }
    return false;
};