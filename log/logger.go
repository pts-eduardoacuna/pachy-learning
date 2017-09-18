package log

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var logfile *os.File

// ToFile installs a file into the default logger unless a DEBUG environment variable is defined.
func ToFile(path string) {
	if os.Getenv("DEBUG") != "" {
		logfile = os.Stdout
	} else {
		var err error
		logfile, err = os.Create(path)
		if err != nil {
			log.Fatal("couldn't create logfile")
		}
	}

	logger = log.New(logfile, "❯ ", log.LstdFlags|log.Lshortfile)
}

// Print is the analog of the standard log.Print function.
func Print(args ...interface{}) {
	str := fmt.Sprint(args...)
	logger.Output(2, str)
}

// Printf is the analog of the standard log.Printf function.
func Printf(s string, args ...interface{}) {
	str := fmt.Sprintf(s, args...)
	logger.Output(2, str)
}

// Fatal is the analog of the standard log.Fatal function.
func Fatal(args ...interface{}) {
	str := fmt.Sprint(args...)
	logger.Output(2, str)
	os.Exit(1)
}

// Fatalf is the analog of the standard log.Fatalf function.
func Fatalf(s string, args ...interface{}) {
	str := fmt.Sprintf(s, args...)
	logger.Output(2, str)
	os.Exit(1)
}

// Close cleans up the log package open files.
func Close() {
	logfile.Close()
}
