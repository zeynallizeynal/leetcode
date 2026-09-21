/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var averageOfLevels = function(root) {
    var res = []
    if (root === null) {
        return res
    }

    var q = [root]
    while (q.length>0) {
        var snapshot = q.length
        var sum = 0
        for (let i = 0; i <snapshot; i++) {
            var cur = q.shift()
            sum+=cur.val
            if (cur.left) q.push(cur.left)
            if (cur.right) q.push(cur.right)
        }
        res.push(sum / snapshot)
    }
    return res
};