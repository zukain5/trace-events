# trace-events
Trace-events parses event information from trace.out (output by runtime/trace) and outputs it to csv.

## Who is it for?
- Beginning Go users unfamiliar with goroutine
- People who have too much information in `go tool trace`

## Usage
1. Run your Go program and get trace.out (see [official documentation](https://pkg.go.dev/runtime/trace#example-package) for how to output trace.out).
2. Copy trace.out to the trace-events directory.
3. Execute main.go (`go run .` or `go run main.go`).
4. You will get output.csv.

## Output Csv Format
- The first column represents the timestamp.
- The second and subsequent columns represent event information for each goroutine.
- The first line shows the id of the goroutine.

## Caution
- This program has been tested only on go1.20.5.
- It is very likely that other versions will not work because this program uses a copy-paste of go's internal process.

## Example Output
```csv
,18,0,2,3,4,17,1,19
0,,GoCreate,,,,,Created,
2942,,GoCreate,Created,,,,,
3026,,,GoWaiting,,,,,
5635,,GoCreate,,Created,,,,
5663,,,,GoWaiting,,,,
5857,,GoCreate,,,Created,,,
5913,,,,,GoWaiting,,,
6107,,GoCreate,,,,Created,,
6163,,,,,,GoWaiting,,
6218,,ProcStart,,,,,,
6329,,,,,,,GoStart,
8023,,,,,,,Gomaxprocs,
31676,Created,,,,,,GoCreate,
39616,,,,,,,GoCreate,Created
43030,,,,,,,GoBlockRecv,
51276,,,,,,,,GoStart
51692,,,,,,,Unblocked,GoUnblock
58799,,,,,,,,GoEnd
59299,,,,,,,GoStart,
66878,,,,,,,HeapAlloc,
80398,,,,,,,GoSysCall,
81758,,ProcStart,,,,,,
87644,GoStart,,,,,,,
89226,GoSysCall,,,,,,,
102940,,ProcStart,,,,,,
104051,GoSysBlock,,,,,,,
104245,,ProcStop,,,,,,
105050,,ProcStop,,,,,,
117321,,,,,,,GoSched,
```

| |18|0|2|3|4|17|1|19|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|0| |GoCreate| | | | |Created| |
|2942| |GoCreate|Created| | | | | |
|3026| | |GoWaiting| | | | | |
|5635| |GoCreate| |Created| | | | |
|5663| | | |GoWaiting| | | | |
|5857| |GoCreate| | |Created| | | |
|5913| | | | |GoWaiting| | | |
|6107| |GoCreate| | | |Created| | |
|6163| | | | | |GoWaiting| | |
|6218| |ProcStart| | | | | | |
|6329| | | | | | |GoStart| |
|8023| | | | | | |Gomaxprocs| |
|31676|Created| | | | | |GoCreate| |
|39616| | | | | | |GoCreate|Created|
|43030| | | | | | |GoBlockRecv| |
|51276| | | | | | | |GoStart|
|51692| | | | | | |Unblocked|GoUnblock|
|58799| | | | | | | |GoEnd|
|59299| | | | | | |GoStart| |
|66878| | | | | | |HeapAlloc| |
|80398| | | | | | |GoSysCall| |
|81758| |ProcStart| | | | | | |
|87644|GoStart| | | | | | | |
|89226|GoSysCall| | | | | | | |
|102940| |ProcStart| | | | | | |
|104051|GoSysBlock| | | | | | | |
|104245| |ProcStop| | | | | | |
|105050| |ProcStop| | | | | | |
|117321| | | | | | |GoSched| |

- The content of the cell represents an event (the content of the event is defined in [internal/trace](https://pkg.go.dev/internal/trace))
- Information written in the goroutines that are the target of the event, such as "Created", "Unblocked", etc., is added by trace-events.

## License
MIT
