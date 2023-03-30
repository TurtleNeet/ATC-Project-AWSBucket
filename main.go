package main

import (
	"aWS/findir"
	"aWS/fthash"
	"aWS/getDate"
	"fmt"
	"os"
	"time"
)

func main() {

	option := menu()

	switch option {

	case 1:
		fthash.RefreshTerminal()
		findir.GetDir()
		var newMenu = menu()
		resetMenu(newMenu)
		break

	case 2:
		fthash.RefreshTerminal()
		getDate.GetFileDate()
		var newMenu = menu()
		resetMenu(newMenu)
		break

	case 3:
		fthash.RefreshTerminal()
		fmt.Println("\n\t The program has ended.")
		os.Exit(0)
		break

	default:
		fmt.Println("\n\t Error detected! Please try again.")
	}
}

func menu() int {
	var option int
	fmt.Print("\n\t                   Menu               ")
	fmt.Print("\n\t ----------------------------------------")
	fmt.Print("\n\t 1. Check a directory for a specific file")
	fmt.Print("\n\t 2. Check a file's creation date & time)")
	fmt.Print("\n\t 3. Exit the program\n")
	fmt.Print("\n\t Option: ")
	fmt.Scanln(&option)
	return option
}

func resetMenu(newMenu int) {

	time.Sleep(1 * time.Second)

	switch newMenu {

	case 1:
		fthash.RefreshTerminal()
		findir.GetDir()
		break

	case 2:
		fthash.RefreshTerminal()
		getDate.GetFileDate()
		break

	case 3:
		fthash.RefreshTerminal()
		fmt.Println("\n\t The program has ended.")
		os.Exit(0)
		break

	default:
		fmt.Print("\n\t Error detected! Please try again.")
	}

}
