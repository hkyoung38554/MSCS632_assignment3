import json, random
from collections import defaultdict

random.seed(7)

def load(path):
    with open(path, "r", encoding="utf-8") as f:
        return json.load(f)

def main():
    cfg = load("data/sample_input.json")
    employees = cfg["employees"]
    days, shifts, req, MAX_DAYS = cfg["days"], cfg["shifts"], cfg["requirement"], cfg["max_days"]

    schedule = {d: {s: [] for s in shifts} for d in days}
    worked, assigned = defaultdict(int), defaultdict(set)

    # pass 1: honor preferences when possible
    for d in days:
        for emp in employees:
            e, pref = emp["name"], emp["prefs"].get(d)
            if pref and len(schedule[d][pref]) < req and worked[e] < MAX_DAYS and d not in assigned[e]:
                schedule[d][pref].append(e)
                assigned[e].add(d)
                worked[e] += 1

    # pass 2: fill minimums without hanging if impossible
    for d in days:
        for s in shifts:
            changed = True
            while len(schedule[d][s]) < req and changed:
                changed = False
                for emp in employees:
                    e = emp["name"]
                    if worked[e] < MAX_DAYS and d not in assigned[e]:
                        schedule[d][s].append(e)
                        assigned[e].add(d)
                        worked[e] += 1
                        changed = True
                        if len(schedule[d][s]) >= req:
                            break
            # if still underfilled here, leave it underfilled and move on

    # print result
    for d in days:
        print(d)
        for s in shifts:
            crew = ", ".join(schedule[d][s]) if schedule[d][s] else "(none)"
            print(f"  {s}: {crew}")
        print()

if __name__ == "__main__":
    main()
