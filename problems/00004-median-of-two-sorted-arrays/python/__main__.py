class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        numbers = nums1
        for n in nums2:
            ind = next((i for i, v in enumerate(numbers) if v > n), None)
            if ind is None: numbers.append(n)
            else: numbers.insert(ind, n)
        l = len(numbers)
        if l % 2 == 0:
            return (numbers[(l // 2) - 1] + numbers[l // 2]) / 2.0
        else:
            return numbers[l // 2]
