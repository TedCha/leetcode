/**
 * @param {number[]} nums
 * @return {number}
 */
export function removeDuplicates(nums) {
    let last;
    let count = 1;
    let offset = 1;
    for (let i = 0; i < nums.length; i++) {
        const curr = nums[i];

        if (last != null && curr === last && count < 2) {
            nums[offset] = curr;
            offset++;
            count++;
        } else if (last != null && curr !== last) {
            nums[offset] = curr;
            offset++;
            count = 1;
        }

        last = curr;
    }

    nums.length = offset;
};
