package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// init invoke stty to modify terminal configuration
// enabling cbreak mode and disabling disabling cursor echo
func init() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

// cleanup restores cooked mode (default mode)
// Same as enabling cbreak, but with flags reversed
func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

// read input from stdin
func readInput() (string, error) {
	buffer := make([]byte, 100)

	// returns the number of bytes read and an error value
	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	// if we read just one byte and if that byte is the escape key.
	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil
}

var maze []string

func loadMaze() error {
	f, err := os.Open("maze01.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

func printScreen() {
	clearScreen()
	for _, line := range maze {
		fmt.Println(line)
	}
}

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

func main() {
	// initialize game

	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}
	// load resources

	// game loop
	for {
		// update screen
		printScreen()

		// process input
		input, err := readInput()
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop
		if input == "ESC" {
			break
		}

		// repeat
	}
}
