package trace

import (
	"encoding/csv"
	"os"
	"strconv"
)

func formatForCsv(events []*Event) [][]string {
	goroutines := make(map[string]struct{})
	for _, event := range events {
		goroutines[strconv.FormatUint(event.G, 10)] = struct{}{}
	}

	var goroutineIds []string
	for id := range goroutines {
		goroutineIds = append(goroutineIds, id)
	}
	headers := append([]string{""}, goroutineIds...)

	idPos := make(map[uint64]int)
	for i, id := range headers {
		uint_id, _ := strconv.ParseUint(id, 10, 16)
		idPos[uint_id] = i
	}

	colNum := len(goroutineIds) + 1
	rows := make([][]string, len(events))
	for i := 0; i < len(events); i++ {
		rows[i] = make([]string, colNum)
	}

	for i, event := range events {
		rows[i][0] = strconv.FormatInt(event.Ts, 10)
		desc := EventDescriptions[event.Type]
		rows[i][idPos[event.G]] = desc.Name
		switch event.Type {
		case EvGoCreate:
			object_id := event.Args[0]
			rows[i][idPos[object_id]] = "Created"
		case EvGoUnblock:
			object_id := event.Args[0]
			rows[i][idPos[object_id]] = "Unblocked"
		}
	}

	table := append([][]string{headers}, rows...)

	return table
}

func OutputCsv(events []*Event, path string) {
	f, _ := os.Create(path)
	w := csv.NewWriter(f)
	w.WriteAll(formatForCsv(events))
	w.Flush()
}
