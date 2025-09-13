# MSCS632 Assignment 4

This repo contains code for Assignment 4: Implementing Control Structures.

## Description

The task is to build a weekly shift planner.  
Each day has three shifts morning afternoon evening.  
Rules: one shift per person per day, max five days total, at least two per shift.  
If not enough volunteers, the program fills with random choices from those still eligible.  
Conflicts are handled by moving people to another shift on the same day or to later days.

## Languages

Python: `python/scheduler.py`  
Go: `go/scheduler.go`

## How To Run

### Python
```bash
python3 python/scheduler.py
```

### Go
```bash
cd go
go run scheduler.go
cd ..
```

## Layout

```plaintext
Assignment4/
  data/
    sample_input.json
  python/
    scheduler.py
  go/
    scheduler.go
  screenshots/
  README.md
```
