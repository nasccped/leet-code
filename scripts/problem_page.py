import os

PROBLEMS_DIR = "./problems"
DIFFICULTY_AND_COLORS = [
    ("Easy", "319148"),
    ("Medium", "916f31"),
    ("Hard", "913c31"),
    ("Abort", None)
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

class ProblemPage:
    """
    New problem page + it's inner attributes.
    """

    def __init__(self):
        self.topics = []

    def set_number(self, number: str):
        self.number: str = number

    def number_is_safe(self) -> bool:
        num: str = self.number
        is_not_zero = False
        if len(num) != 5:
            return False
        for c in num:
            if c < '0' or c > '9': return False
            is_not_zero |= c > '0'
        return (not any(
            entry.startswith(num) \
            for entry in os.listdir(PROBLEMS_DIR)
        )) and is_not_zero

    def set_title_name(self, title: str):
        self.title = " ".join(
            word.lower() \
            for word in title.split(" ") if word
        ).title()
        self.dir_name = self.title.lower().replace(" ", "-")

    def set_difficulty_and_color(self, difficulty: str, color: str):
        self.difficulty = difficulty
        self.color = color

    def print_topics(self):
        print("[{}]".format(", ".join("\033[92m" + str(t) + "\033[0m" for t in self.topics)))

    def push_topic(self, topic: str):
        if topic not in self.topics:
            self.topics.append(topic)

    def remove_topic(self, topic: str):
        if topic in self.topics:
            self.topics.remove(topic)

    def generate(self):
        dirname = f"{PROBLEMS_DIR}/{self.number}-{self.dir_name}"
        os.makedirs(dirname, exist_ok = True)
        with open(f"{dirname}/README.md", "w") as f:
            f.write(f"# {int(self.number)}. {self.title}\n\n")
            f.write(f"[![{self.difficulty}](https://img.shields.io/badge/{self.difficulty.replace(' ', '_')}-{self.color})](#)\n")
            for t in self.topics:
                f.write(f"[![{t}](https://img.shields.io/badge/{t.replace(' ', '_')}-302f33)](#)\n")
            f.write("\n")
            f.write("Place the content here...")
