/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function (s) {
    var romans = {
        I: 1,
        V: 5,
        X: 10,
        L: 50,
        C: 100,
        D: 500,
        M: 1000,
    };

    const romanKeys = Object.keys(romans);
    let res = 0;
    for (let i = 0; i < s.length; i++) {
        const rIndex = romanKeys.indexOf(s[i])
        const nextRIndex = romanKeys.indexOf(s[i + 1])
        if (rIndex < nextRIndex) {
            res -= romans[s[i]]
        } else {
            res += romans[s[i]]
        }
    }
    return res
};