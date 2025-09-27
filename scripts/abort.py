ABORT_TRIGGER = "!ABORT"

def is_abort_call(input):
    return input.lower().strip() == ABORT_TRIGGER.lower()

def abort_func():
    print("Aborting...")
    exit(0)

class Abort:

    @staticmethod
    def check_and_execute(input: str):
        if is_abort_call(input): abort_func()
