/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
export function rotate(nums, k) {
    if (nums.length === 1) {
        return;
    }

    if (k > nums.length) {
        k = k - nums.length
    }

    const copy = nums.slice();
    for (let i = 0; i < nums.length; i++) {
        let j = Math.abs((nums.length - k + i) % nums.length);
        nums[i] = copy[j];
    }
};