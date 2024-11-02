/**
 * @param {number[]} nums
 * @return {number}
 */
export function removeDuplicates(nums) {
    let last = 0;
    for (let i = 0; i < nums.length; i++) {
        const x = nums[i];
        const y = nums[i + 1];
        if (y != null && x !== y) {
            nums[last+1] = y;
            last++;
        }
    }

    nums.splice(last+1, nums.length);
};
