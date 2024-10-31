import { expect, test } from 'vitest';
import { merge } from './merge_sorted_array';

test('test case 1', () => {
    const nums1 = [1, 2, 3, 0, 0, 0];
    merge(nums1, 3, [2, 5, 6], 3);

    expect(nums1).toStrictEqual([1, 2, 2, 3, 5, 6]);
});

test('test case 2', () => {
    const nums1 = [1];
    merge(nums1, 1, [], 0);

    expect(nums1).toStrictEqual([1]);
});

test('test case 3', () => {
    const nums1 = [0];
    merge(nums1, 0, [1], 1);

    expect(nums1).toStrictEqual([1]);
});