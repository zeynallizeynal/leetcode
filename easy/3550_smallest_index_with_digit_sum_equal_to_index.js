/**
 * @param {number[]} nums
 * @return {number}
 */
var smallestIndex = function(nums) {
    for (var i = 0; i < nums.length; i++) {
        var cur = nums[i], sum = 0;
        while (cur>0) {
            sum +=(cur%10);
            cur = (cur / 10) | 0;
        }
        if (sum === i) {
            return i;
        }
    }
    return -1;
};