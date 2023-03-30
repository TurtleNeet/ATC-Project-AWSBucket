package findir

import (
	"fmt"
	"os"
)

// GetDir := gets user input to check file exists in user-specified directory
func GetDir() {
	var sysDir, fileExt string

	fmt.Print("\n\t Enter a file directory: ")
	fmt.Scanln(&sysDir)

	fmt.Print("\n\t Enter a file extension: ")
	fmt.Scanln(&fileExt)

	checkDir(sysDir, fileExt)
}

func checkDir(sysDir string, fileExt string) {

	path := sysDir + fileExt

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("\n\t Error! Unable to locate file.")

	} else {
		fmt.Printf("\n\t The id for your file (%v) exists!\n", file)
	}

}
