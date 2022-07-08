package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// create a filename for the migration
// filename is the timestamp of the migration plus the name of the migration
// filename is in the format YYYYMMDDHHMMSS_name.sql
// name comes from the terminal
// output goes to the terminal
func CreateMigration(name string) {
	// get the current time
	t := time.Now()
	// format the time
	timestamp := t.Format("20060102150405")
	// create the filename
	filename := "V" + timestamp + "_" + name + ".sql"
	// output the filename
	println(filename)
	// copy fileename to clipboard
	clipboard(filename)
	// notify user that the file has been copied to the clipboard
	println("Copied to clipboard")
}

func copyUnix(s string) {
	// copy the string to the clipboard
	clipboard := `echo '` + s + `' | pbcopy`
	// run the command
	cmd := exec.Command("sh", "-c", clipboard)
	// run the command
	cmd.Run()
}

func copyWindows(s string) {
	// copy the string to the clipboard
	clipboard := `echo '` + s + `' | clip`
	// run the command
	cmd := exec.Command("cmd", "/C", clipboard)
	// run the command
	cmd.Run()
}

func clipboard(s string) {
	os := runtime.GOOS
	switch os {
	case "windows":
		copyWindows(s)
	default:
		copyUnix(s)
	}
}


func readInput() string {
	// read the input
	var input string
	fmt.Println("Enter the name of the migration:")
	// read the input
	fmt.Scanln(&input)
	// return the input
	return input
}

func main() {
	CreateMigration(readInput())
}
