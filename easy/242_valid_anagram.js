/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  if (s.length !== t.length) {
    return false;
  }
  const sorted1 = s.split('').sort().join(''), sorted2 = t.split('').sort().join('');
  return sorted1 === sorted2;
};

