# https://leetcode.com/problems/longest-substring-without-repeating-characters/

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        
        max_substr_len = 0
        for i in range(len(s)):

            substr = s[i]

            if i + 1 < len(s):
                for c in s[i+1:]:
                    if c not in substr:
                        substr += c
                    else:
                        break
            
            if len(substr) > max_substr_len:
                max_substr_len = len(substr) 

        return max_substr_len
