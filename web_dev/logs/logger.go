package logs

import (
	"fmt"
	"os"
	"time"
)

const LOG_DIR = "logs"
const ERR_LOG_FILE = "error.log"
const ACCESS_LOG_FILE = "access.log"
const DATE_FORMAT_LAYOUT = "2006-01-02" // this date is used in Golang for formatting because its Golang's birthday

func GetCurrentDate() string {
	// Fetches the current date and returns it into YYYY_MM_DD format for maintaining current day's log directory
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(DATE_FORMAT_LAYOUT))
}

func GetCurrentLogDir() string {
	cwd, _ := os.Getwd()
	log_dir := fmt.Sprintf("%v/%v/%v", cwd, LOG_DIR, GetCurrentDate())
	return log_dir
}

func CloseLogFiles() {
	log_dir := GetCurrentLogDir()
	log_files := []string{ERR_LOG_FILE, ACCESS_LOG_FILE}
	for _, log_file := range log_files {
		_, err := os.Stat(fmt.Sprintf("%v/%v", log_dir, log_file))
		if err != nil {
			continue
		} else {
			// Close the file
		}
	}
}

func Logger(log_string string, isErr bool) {

	// Takes the log string and write it into error.logs/access.logs based on isErr
	log_dir := GetCurrentLogDir()

	var log_file string
	if isErr {
		log_file = ERR_LOG_FILE
	} else {
		log_file = ACCESS_LOG_FILE
	}

	_, err := os.Stat(log_dir)
	if err != nil {
		os.Mkdir(log_dir, 0777)
	}

	log_file_path := fmt.Sprintf("%v/%v", log_dir, log_file)
	if _, err := os.Stat(log_file_path); err != nil {
		// Log file does not exit, creating
		os.Create(log_file_path)
	}

	// Write the log into the file in append-mode
	// getting file-object by opening the file
	f, err := os.OpenFile(log_file_path, os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(f, log_string) // Note: Use Fprintln for appending in a newline.

}
