import unittest
from longest_substring_without_repeating_characters import Solution


class TestLengthOfLongestSubstring(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_longest_substring_without_repeating_characters_1(self):
        actual = self.solution.lengthOfLongestSubstring("abcabcbb")
        expected = 3
        self.assertEqual(actual, expected)

    def test_longest_substring_without_repeating_characters_2(self):
        actual = self.solution.lengthOfLongestSubstring("bbbbb")
        expected = 1
        self.assertEqual(actual, expected)

    def test_longest_substring_without_repeating_characters_3(self):
        actual = self.solution.lengthOfLongestSubstring("pwwkew")
        expected = 3
        self.assertEqual(actual, expected)
    
    def test_longest_substring_without_repeating_characters_4(self):
        actual = self.solution.lengthOfLongestSubstring("dvdf")
        expected = 3
        self.assertEqual(actual, expected)

    def test_longest_substring_without_repeating_characters_5(self):
        actual = self.solution.lengthOfLongestSubstring("")
        expected = 0
        self.assertEqual(actual, expected)


if __name__ == '__main__':
    unittest.main()