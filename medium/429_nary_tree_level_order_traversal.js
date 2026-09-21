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
 * @return {number[][]}
 */
var levelOrder = function (root) {
  if (root === null) {
    return [];
  }

  var result = [];
  var queue = [root];
  while (queue.length > 0) {
    var size = queue.length;
    var level = [];
    while (size-- > 0) {
      var cur = queue.shift();
      level.push(cur.val);
      for (var c of cur.children) {
        q.push(c);
      }
    }
    result.push(level);
  }

  return result;
};
