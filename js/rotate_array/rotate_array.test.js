import { expect, test } from 'vitest';
import { rotate } from './rotate_array';

test('test case 1', () => {
    const nums = [1,2,3,4,5,6,7];
    rotate(nums, 3);

    expect(nums).toStrictEqual([5, 6, 7, 1, 2, 3, 4]);
});

test('test case 2', () => {
    const nums = [-1,-100,3,99];
    rotate(nums, 2);

    expect(nums).toStrictEqual([3,99,-1,-100]);
});

test('test case 3', () => {
    const nums = [1,2];
    rotate(nums, 3);

    expect(nums).toStrictEqual([2,1]);
});

test('test case 4', () => {
    const nums = [1,2];
    rotate(nums, 5);

    expect(nums).toStrictEqual([2,1]);
});

test('test case 5', () => {
    const nums = [1,2,3];
    rotate(nums, 4);

    expect(nums).toStrictEqual([3,1,2]);
});