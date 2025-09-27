import os
from abort import is_abort_call, abort_func
from myio import ital_black_str

PROBLEMS_DIR = "problems"

class ProblemDirectory:
    """
    Intermediary Directory class to handle Directory related values.
    """

    def __init__(self):
        self.number = input(f"> Set the problem number: {ital_black_str('00000')}\033[5D")

    def set_fullpath_with_problem_name(self, name: str):
        problem = name.lower().replace(" ", "-")
        self.fullpath = os.path.join(".", PROBLEMS_DIR, f"{self.number}-{problem}")

    def generate(self):
        os.makedirs(self.fullpath)

    def number_already_exists(self):
        return any(dir.startswith(self.number) for dir in os.listdir(PROBLEMS_DIR))

    def try_fullpath_from_self_num(self):
        path = os.path.join(".", PROBLEMS_DIR)
        if not os.path.exists(path):
            print(f"Error: Problems path doesn't exists ({path})")
            exit(1)

        problem = [entry for entry in os.listdir(path) if entry.startswith(self.number)]

        if len(problem) == 0:
            print(f"Error: no problem directory for the given problem number ({self.number})")
            exit(1)
        elif len(problem) != 1:
            print(f"Error: more than one dir for the given problem number ({self.number})")
            exit(1)

        self.fullpath = os.path.join(".", PROBLEMS_DIR, problem[0])

    def exists(self):
        return os.path.exists(self.fullpath)

    def number_is_safe(self):
        res = False
        for c in self.number:
            if c < '0' or c > '9':
                return False
            res |= c > '0'
        return len(self.number) == 5 and res
