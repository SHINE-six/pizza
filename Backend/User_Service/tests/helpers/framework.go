package helpers

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var step_counter int = 0

func Action(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	shortFile := file[strings.LastIndex(file, "/")+1:]
	step_counter++
	
	// remove \t from the message
	message = strings.Replace(message, "\t", "", -1)

	// report output to the console
	customLog(shortFile + ":" + fmt.Sprint(line))
	customLog("Step " + fmt.Sprint(step_counter) + ": " + message)
}

func Result(message string) {
	// remove \t from the message
	message = strings.Replace(message, "\t", "", -1)

	// report output to the console
	customLog("Passed: " + message)
	customLog("Step " + fmt.Sprint(step_counter) + " passed!" + "\n")
}

func Title(message string) {
	// report output to the console
	customLog("---------------------------------")
	customLog("\t" + message)
	customLog("---------------------------------")
}

func customLog(message string) {
	fmt.Fprintln(os.Stdout, message)
}