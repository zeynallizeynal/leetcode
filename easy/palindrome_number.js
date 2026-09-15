/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
  if (x < 0) return false;
  let div = 1;
  while (x >= 10 * div) {
    div *= 10;
  }

  while (x > 0) {
    let right = x % 10;
    let left = (x / div) | 0;
    if (left !== right) return false;
    x = (x % div / 10) | 0;
    div = div / 100;
  }
  return true;
};

console.log(isPalindrome(121));