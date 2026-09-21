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
 * @return {number}
 */
var minDepth = function(root) {
    if(root===null) return 0
    var q = [root]
    var depth = 1
    while(q.length>0) {
        var snapshot=q.length
        for (var i=0;i<snapshot;i++) {
            var cur = q.shift()
            if(!cur.left && !cur.right) return depth
            if(cur.left) q.push(cur.left)
            if(cur.right) q.push(cur.right)
        }
        depth++
    }
    return depth
};