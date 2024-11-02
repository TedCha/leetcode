import { expect, test } from 'vitest';
import { removeDuplicates } from './remove_duplicates_from_sorted_array';

test('test case 1', () => {
    const nums1 = [1,1,2];
    removeDuplicates(nums1);

    expect(nums1.toSorted()).toStrictEqual([1, 2]);
});

test('test case 2', () => {
    const nums1 = [0,0,1,1,1,2,2,3,3,4];
    removeDuplicates(nums1);

    expect(nums1.toSorted()).toStrictEqual([0, 1, 2, 3, 4]);
});