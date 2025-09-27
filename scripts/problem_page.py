import os
import myio
from directory import ProblemDirectory
from abort import Abort, ABORT_TRIGGER

DIFFICULTY_AND_COLORS = [
    ("Easy", "319148"),
    ("Medium", "916f31"),
    ("Hard", "913c31")
]
TOPICS = [
    "Array",
    "String",
    "Hash Table",
    "Dynamic Programming",
    "Math",
    "Sorting",
    "Greedy",
    "Depth-First Search",
    "Binary Search",
    "Database",
    "Matrix",
    "Tree",
    "Breadth-First Search",
    "Bit Manipulation",
    "Two Pointers",
    "Prefix Sum",
    "Heap (Priority Queue)",
    "Simulation",
    "Binary Tree",
    "Graph",
    "Stack",
    "Counting",
    "Sliding Window",
    "Design",
    "Enumeration",
    "Backtracking",
    "Union Find",
    "Linked List",
    "Number Theory",
    "Ordered Set",
    "Monotonic Stack",
    "Segment Tree",
    "Trie",
    "Combinatorics",
    "Bitmask",
    "Divide and Conquer",
    "Queue",
    "Recursion",
    "Geometry",
    "Binary Indexed Tree",
    "Memoization",
    "Hash Function",
    "Binary Search Tree",
    "Shortest Path",
    "String Matching",
    "Topological Sort",
    "Rolling Hash",
    "Game Theory",
    "Interactive",
    "Data Stream",
    "Monotonic Queue",
    "Brainteaser",
    "Doubly-Linked List",
    "Randomized",
    "Merge Sort",
    "Counting Sort",
    "Iterator",
    "Concurrency",
    "Probability and Statistics",
    "Quickselect",
    "Suffix Array",
    "Line Sweep",
    "Minimum Spanning Tree",
    "Bucket Sort",
    "Shell",
    "Reservoir Sampling",
    "Strongly Connected Component",
    "Eulerian Circuit",
    "Radix Sort",
    "Rejection Sampling",
    "Biconnected Component"
]

def generate(obj):
    dirname = obj.problem.fullpath
    num = obj.number
    tit = obj.title
    dif = obj.difficulty
    col = obj.color
    topics = obj.topics
    obj.problem.generate()
    with open(f"{dirname}/README.md", "w") as f:
        f.write(f"# {num}. {tit}\n\n")
        f.write(f"[![{dif}](https://img.shields.io/badge/{dif}-{col})](#)\n")
        for t in topics:
            f.write(f"[![{t}](https://img.shields.io/badge/{t.replace(' ', '_')}-302f33)](#)\n")
        f.write("\n")
        f.write("Place the content here...")

class ProblemPage:
    """
    New problem page + it's inner attributes.
    """

    def __init__(self):
        print(f"Creating a {myio.norm_cyan_str('new problem page')} (use '{ABORT_TRIGGER}' to stop execution)")
        self.topics = []
        problem = ProblemDirectory()
        Abort.check_and_execute(problem.number)
        while not problem.number_is_safe() or problem.number_already_exists():
            myio.erase_n(2)
            problem = ProblemDirectory()
            Abort.check_and_execute(problem.number)
        self.problem = problem
        self.number = int(problem.number)
        print(
            f"> Set the problem title: \033[3;90mCapitalized Name\033[0m",
            end = "\033[16D"
        )
        title = input()
        Abort.check_and_execute(title)
        self.title = " ".join([word for word in title.split(" ") if word]).title()
        self.problem.set_fullpath_with_problem_name(self.title)

        colored_diffs = [f"\033[38;2;{';'.join(str(int(hex_color[i:i+2],16)) for i in (0,2,4))}m{txt}\033[0m" for txt, hex_color in DIFFICULTY_AND_COLORS]
        print(f"> Set the problem difficulty:")
        for i, dif in enumerate(colored_diffs):
            print(f"  {i + 1}. {dif}")
        dif = input("> ")

        while dif not in [str(i + 1) for i in range(len(colored_diffs))]:
            myio.erase_n(2)
            dif = input("> ")
            Abort.check_and_execute(dif)

        Abort.check_and_execute(dif)
        ind = int(dif) - 1
        self.difficulty, self.color = DIFFICULTY_AND_COLORS[ind]

        self.topics = set()
        print("Chose the topics (+X | -X | Empty str to continue):")
        while True:
            print(">> [{}]".format(", ".join("\033[92m" + str(t) + "\033[0m" for t in self.topics)))
            top = " ".join(word.capitalize() for word in input("> ").strip().split(" ") if word)
            Abort.check_and_execute(top)
            if top == "": break
            elif top.startswith("+"): func = self.topics.add
            elif top.startswith("-"): func = self.topics.remove
            else: continue
            top = top[1:].strip().title()
            myio.erase_n(3)
            if top in TOPICS:
                func(top)

        generate(self)
        print(f"Successfully generated ({self.problem.fullpath})")
