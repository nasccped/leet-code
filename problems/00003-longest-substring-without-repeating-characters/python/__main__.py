class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        lon, start = 0, 0
        for i, c in enumerate(s):
            if c in s[start : i]:
                if i - start > lon:
                    lon = i - start
                while c in s[start : i]:
                    start += 1
        lon = max(len(s) - start, lon)
        return lon
