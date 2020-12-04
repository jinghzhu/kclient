package utils

import (
	"fmt"
	"runtime"
	"time"
)

// PanicHandler catches a panic and logs an error. Suppose to be called via defer.
func PanicHandler() (caller string, fileName string, lineNum int, stackTrace string, rec interface{}) {
	buf := make([]byte, stackBuffer)
	runtime.Stack(buf, false)
	name, file, line := GetCallerInfo(2)
	if r := recover(); r != nil {
		caller, fileName, stackTrace = name, file, string(buf)
		lineNum = line
		rec = r
		fmt.Printf("%s %s ln%d: PANIC Deferred : %v\n", name, file, line, r)
		fmt.Printf("%s %s ln%d: Stack Trace : %s", name, file, line, string(buf))
	}

	return caller, fileName, lineNum, stackTrace, rec
}

// GetCallerInfo returns the name of method caller and file name. It also returns the line number.
func GetCallerInfo(level int) (caller string, fileName string, lineNum int) {
	if level < 1 || level > maxCallerLevel {
		level = defaultCallerLevel
	}

	pc, file, line, ok := runtime.Caller(level)
	if ok {
		fileName = file
		lineNum = line
	}
	details := runtime.FuncForPC(pc)
	if details != nil {
		caller = details.Name()
	}

	return caller, fileName, lineNum
}

// Retry will retry the given condition function with specific time interval and retry round.
// It will return true if the condition is met. If it is timeout, it will return false. Otherwise,
// it will return the error encountered in the retry round.
func Retry(interval time.Duration, round int, retry func() (bool, error)) (bool, error) {
	if round < 1 {
		round = 1
	} else if round > 5 {
		round = 5
	}
	var err error
	done := false
	for i := 0; i < round; i++ {
		done, err = retry()
		if done {
			break
		}
		time.Sleep(interval)
	}

	if done {
		return true, nil
	}

	return false, err
}

// ContainsStrSli returns true if the a exists in the slice A.
func ContainsStrSli(A []string, a string) bool {
	if A == nil {
		return false
	}
	for _, val := range A {
		if val == a {
			return true
		}
	}

	return false
}
