/**
 * @param {number[]} nums
 * @return {number}
 */
export function majorityElement(nums) {
    let majority;
    let majorityCount;
    let counts = {};
    for (let i = 0; i < nums.length; i++) {
        if (counts[nums[i]] != null) {
            counts[nums[i]] += 1;
        } else {
            counts[nums[i]] = 1;
        }

        if (majority == null || counts[nums[i]] > majorityCount) {
            majority = nums[i];
            majorityCount = counts[nums[i]];
        }
    }

    return majority;
};