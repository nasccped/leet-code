class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        mapping = {}
        for i, n in enumerate(nums):
            temp = mapping.get(target - n)
            if not (temp is None):
                return [temp, i]
            else:
                mapping[n] = i

print("This program doesn't have testing since there's no suitable way of complex testing",)
