/**
 * // Definition for a _Node.
 * function _Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {_Node|null} root
 * @return {number[]}
 */
var postorder = function(root) {
    if (root === null) {
        return []
    }

    var res = []
    for (var c of root.children) {
        res = [...res, ...postorder(c)]
    }
    res.push(root.val)
    return res
};