class Solution(object):
    """
    # Solution class:

    Place bellow the solution func:
    """
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

# Provided inputs + expected outputs: (list/dictonaries are prefered)
INPUTS = [
    (([2, 7, 11, 15], 9), ([0, 1], [1, 0])),
    (([3, 2, 4], 6), ([1, 2], [2, 1])),
    (([2, 3], 6), ([0, 1], [1, 0]))
]

def assertion(current, input, *expecteds):
    """
    # Assertion function:

    Mimics the `assert` functionality with a more fancy printing.
    """
    print(
        f"Running test \033[96m{current}\033[0m:",
        end = " "
    )
    assertion_success = False
    for e in expecteds:
        try:
            assert input == e
            assertion_success = True
            break
        except:
            pass

    if assertion_success:
        print("[\033[92mok\033[0m]")
    else:
        print("[\033[91mfail\033[0m]")
        print(f"  input:       \033[93m{input}\033[0m")
        print(f"  expected(s): {' or '.join([str(e) for e in expecteds])}")

def main():
    solution = Solution()
    for (i, ((nums, target), (output_1, output_2))) in enumerate(INPUTS):
        result = solution.twoSum(nums, target)
        assertion(i + 1, result, output_1, output_2)

if __name__ == "__main__":
    main()
