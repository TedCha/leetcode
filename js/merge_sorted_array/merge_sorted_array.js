/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
export function merge(nums1, m, nums2, n) {
    if (n === 0) {
        return;
    }

    let endIdx = m + n - 1;
    while (n > 0 && m > 0) {
        if (nums2[n - 1] >= nums1[m - 1]) {
            nums1[endIdx] = nums2[n - 1];
            n--;
        } else {
            nums1[endIdx] = nums1[m - 1];
            m--;
        }
        endIdx--;
    }

    while (n > 0) {
        nums1[endIdx] = nums2[n - 1];
        n--;
        endIdx--;
    }
};