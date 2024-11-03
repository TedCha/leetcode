/**
 * @param {number} n
 * @return {number}
 */
export function climbStairs(n) {
    const sub = [1, 1];
    for (let i = 0; i < n - 1; i++) {
        sub.push(sub[i] + sub[i + 1]);
    }

    return sub[n];
};