import { expect, test } from 'vitest';
import { removeElement } from './remove_element';

test('test case 1', () => {
    const nums1 = [3, 2, 2, 3];
    removeElement(nums1, 3);

    expect(nums1.toSorted()).toStrictEqual([2, 2]);
});

test('test case 2', () => {
    const nums1 = [0, 1, 2, 2, 3, 0, 4, 2];
    removeElement(nums1, 2);

    expect(nums1.toSorted()).toStrictEqual([0, 0, 1, 3, 4]);
});
