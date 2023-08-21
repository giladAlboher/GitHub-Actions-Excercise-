import datetime
import json

def create_log_file():
    current_time = datetime.datetime.now()
    log_filename = "logtimes.json"

    log_entry = {
        "timestamp": current_time.strftime('%Y-%m-%d %H:%M:%S'),
        "log_number": get_log_number(log_filename)
    }

    with open(log_filename, 'a') as log_file:
        json.dump(log_entry, log_file)
        log_file.write('\n')

    print(f"Log entry added to '{log_filename}'.")

def get_log_number(filename):
    try:
        with open(filename, 'r') as log_file:
            log_entries = [json.loads(line) for line in log_file if line.strip()]
        return len(log_entries) + 1
    except FileNotFoundError:
        return 1

if name == "main":
    create_log_file()