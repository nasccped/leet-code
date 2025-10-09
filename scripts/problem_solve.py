import os
import myio
from abort import ABORT_TRIGGER
from directory import ProblemDirectory

def gen_python_solve(path: str):
    path = os.path.join(path, "python")
    try:
        os.makedirs(path)
    except:
        print(f"(Path already/Parent dir doesn't) exists ({path})")
        exit(1)

    with open(os.path.join(path, "__main__.py"), "w") as f:
        f.write("""class Solution(object):
    \"\"\"
    # Solution class:

    Place bellow the solution func:
    \"\"\"
    def solution_func(self, *inputs):
        pass"""
        )

def gen_rust_solve(path: str):
    path = os.path.join(path, "rust")
    try:
        os.makedirs(path)
    except:
        print(f"(Path already/Parent dir doesn't) exists ({path})")
        exit(1)

    with open(os.path.join(path, "Cargo.toml"), "w") as f:
        f.write("""[package]
name = \"rust\"
version = \"0.1.0\"
edition = \"2024\"

[dependencies]""")

    path = os.path.join(path, "src")
    try:
        os.makedirs(path)
    except:
        print(f"(Path already/Parent dir doesn't) exists ({path})")
        exit(1)

    with open(os.path.join(path, "main.rs"), "w") as f:
        f.write("""#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn solution_func() -> ??? {
        todo!("Code here...")
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}""")

def gen_go_solve(path: str):
    path = os.path.join(path, "go")
    try:
        os.makedirs(path)
    except:
        print(f"(Path already/Parent dir doesn't) exists ({path})")
        exit(1)

    with open(os.path.join(path, "main.go"), "w") as f:
        f.write("""package main

func solution() {
    
}""")


AVAILABLE_LANGUAGES = {
    "python": gen_python_solve,
    "rust": gen_rust_solve,
    "go": gen_go_solve
}

class ProblemSolve:
    """
    New problem solve + it's inner attributes.
    """

    def __init__(self):
        print(f"Creating a {myio.norm_cyan_str('new problem solve')} (use '{ABORT_TRIGGER}' to stop execution)")
        problem = ProblemDirectory()
        problem.try_fullpath_from_self_num()
        print(f"> Solve with: ({', '.join(l for l in AVAILABLE_LANGUAGES.keys())})")
        func = AVAILABLE_LANGUAGES[input("> ")]
        if func is None:
            print("Non available lang")
            print("aborting...")
            exit(1)
        func(problem.fullpath)
        print(f"Solve dir successfully generated at {problem.fullpath}.")
