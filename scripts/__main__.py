import myio
from problem_page import ProblemPage
from problem_solve import ProblemSolve

MENU_OPTIONS = [
    ("new problem page", ProblemPage),
    ("new solve for a problem page", ProblemSolve),
    ("abort", exit)
]

if __name__ == "__main__":
    print("Choose an option:")
    myio.print_table([opt[0] for opt in MENU_OPTIONS])
    opt = input("\n> ")
    while opt not in [str(i + 1) for i in range(len(MENU_OPTIONS))]:
        myio.erase_n(2)
        opt = input("> ")
    func = MENU_OPTIONS[int(opt) - 1][1]
    func()
