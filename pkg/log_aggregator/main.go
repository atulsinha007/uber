package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	startStr   = "**START**"
	endStr     = "**END**"
	startState = "STARTED"
	endState   = "ENDED"
)

//PROCESS_ID:THREAD_--ID::THREAD_NAME LOGGED_TIME- LOG_MESSAGE

//
//f = open_file(file_name)
//start_pos = seek_position(f)
//
//for log_line in log_file:
//thread_id = get_thread_id(log_line)
//file_name = get_thread_file_name(thread_id)
//append_to_file(file_name, log_line)
//
//
//def append_to_file(file_name, log_line):
//for c in log_line:
//put(f, c)

type Conf struct {
	Count int
	State string
}

var pIdThreadIdStateMap = make(map[string]Conf)

func init() {
	// lru cache initialization
	//cache, _ := lru.New(128)

}

func main() {
	file, err := os.Open("pkg/log_aggregator/logfile.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string

	for scanner.Scan() {
		line = scanner.Text()
		//text = append(text, line)
		fmt.Println(line)
		fmt.Println("----------------")

		processLine(line)
	}
	file.Close()
}

type Log struct {
	Pid        string
	ThreadId   string
	ThreadName string
	Timestamp  string
	Message    string
	Line       string
}

func processLine(line string) {
	parsedLog := parse(line)
	fmt.Println(parsedLog)

	write(parsedLog)
}

func parse(line string) Log {
	slice1 := strings.Split(line, "::")
	pid := strings.Split(slice1[0], ":")[0]
	threadId := strings.Split(slice1[0], ":")[1]
	threadName := strings.Split(slice1[1], " ")[0]
	msg := strings.Split(slice1[1], " - ")[1]

	threadNameLen := len(threadName)
	ts := slice1[1][threadNameLen+1 : threadNameLen+1+23]

	return Log{
		Pid:        pid,
		ThreadId:   threadId,
		ThreadName: threadName,
		Timestamp:  ts,
		Message:    msg,
		Line:       line,
	}
}

func write(logLine Log) {
	key := logLine.Pid + ":" + logLine.ThreadId

	val, exists := pIdThreadIdStateMap[key]
	if !exists {
		pIdThreadIdStateMap[key] = Conf{
			Count: 1,
			State: startState,
		}
	} else if val.State == endState {
		// createFile and openFile
		cfg := pIdThreadIdStateMap[key]
		cfg.Count++
		cfg.State = startState
		pIdThreadIdStateMap[key] = cfg
	} else if logLine.Message == endStr {
		cfg := pIdThreadIdStateMap[key]
		cfg.State = endState
		pIdThreadIdStateMap[key] = cfg
		// closeFile
	}

	fileName := getFileName(logLine, pIdThreadIdStateMap[key].Count)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(logLine.Line + "\n"); err != nil {
		panic(err)
	}

}

func getFileName(logLine Log, cnt int) string {
	return logLine.Pid + "_" + logLine.ThreadId + "_" + strconv.Itoa(cnt)
}
