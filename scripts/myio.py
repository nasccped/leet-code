"""
Useful IO methods.
"""

def norm_green_str(obj):
    return f"\033[92m{obj}\033[0m"

def norm_cyan_str(obj):
    return f"\033[96m{obj}\033[0m"

def ital_black_str(obj):
    return f"\033[3;90m{obj}\033[0m"

def erase_n(n):
    """Erase `n` amount of lines. Nothing if `n` <= 0."""
    if n <= 0: return
    print("\r\033[K", end = "")
    print("\033[A\033[K" * (n - 1), end = "")

def print_table(arr):
    """Print tables in a fancy way. The `arr` variable should be a list"""
    for i, item in enumerate(arr):
        print(f"  {norm_green_str(i + 1)}. {item}")
