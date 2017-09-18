package log

import (
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
	logger.Print(args...)
}

// Printf is the analog of the standard log.Printf function.
func Printf(s string, args ...interface{}) {
	logger.Printf(s, args...)
}

// Fatal is the analog of the standard log.Fatal function.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf is the analog of the standard log.Fatalf function.
func Fatalf(s string, args ...interface{}) {
	logger.Fatalf(s, args...)
}

// Close cleans up the log package open files.
func Close() {
	logfile.Close()
}
