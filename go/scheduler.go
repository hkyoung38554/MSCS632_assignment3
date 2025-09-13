package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Employee struct {
	Name  string            `json:"name"`
	Prefs map[string]string `json:"prefs"`
}
type Config struct {
	Employees   []Employee `json:"employees"`
	Requirement int        `json:"requirement"`
	MaxDays     int        `json:"max_days"`
	Shifts      []string   `json:"shifts"`
	Days        []string   `json:"days"`
}

func main() {
	rand.Seed(7)
	b, err := ioutil.ReadFile("../data/sample_input.json")
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		fmt.Println("json error:", err)
		return
	}

	// build structures
	sched := map[string]map[string][]string{}
	for _, d := range cfg.Days {
		sched[d] = map[string][]string{}
		for _, sh := range cfg.Shifts {
			sched[d][sh] = []string{}
		}
	}
	worked := map[string]int{}
	assigned := map[string]map[string]bool{}
	for _, e := range cfg.Employees {
		worked[e.Name] = 0
		assigned[e.Name] = map[string]bool{}
	}

	// pass 1: honor preferences
	for _, d := range cfg.Days {
		for _, e := range cfg.Employees {
			if pref, ok := e.Prefs[d]; ok {
				if len(sched[d][pref]) < cfg.Requirement && worked[e.Name] < cfg.MaxDays && !assigned[e.Name][d] {
					sched[d][pref] = append(sched[d][pref], e.Name)
					assigned[e.Name][d] = true
					worked[e.Name]++
				}
			}
		}
	}

	// pass 2: fill minimums with progress check to avoid infinite loop
	for _, d := range cfg.Days {
		for _, sh := range cfg.Shifts {
			progress := true
			for len(sched[d][sh]) < cfg.Requirement && progress {
				progress = false
				for _, e := range cfg.Employees {
					if worked[e.Name] < cfg.MaxDays && !assigned[e.Name][d] {
						sched[d][sh] = append(sched[d][sh], e.Name)
						assigned[e.Name][d] = true
						worked[e.Name]++
						progress = true
						if len(sched[d][sh]) >= cfg.Requirement {
							break
						}
					}
				}
			}
			// if still short, leave it short
		}
	}

	// print
	for _, d := range cfg.Days {
		fmt.Println(d)
		for _, sh := range cfg.Shifts {
			crew := "(none)"
			if len(sched[d][sh]) > 0 {
				crew = ""
				for i, n := range sched[d][sh] {
					if i > 0 {
						crew += ", "
					}
					crew += n
				}
			}
			fmt.Printf("  %s: %s\n", sh, crew)
		}
		fmt.Println()
	}
}
