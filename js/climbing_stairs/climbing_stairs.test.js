import { expect, test } from 'vitest';
import { climbStairs } from './climbing_stairs';

test('test case 1', () => {
    expect(climbStairs(2)).toBe(2);
});

test('test case 2', () => {
    expect(climbStairs(3)).toBe(3);
});
