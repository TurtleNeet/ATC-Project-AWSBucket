package getDate

import (
	"fmt"
	"os"
)

// GetFileDate := gets user input for file & directory to share date
func GetFileDate() {
	var myDir, myFile string

	fmt.Print("\n\t Enter a file directory: ")
	fmt.Scanln(&myDir)

	fmt.Print("\n\t Enter a file name with the extension: ")
	fmt.Scanln(&myFile)

	showFileDate(myDir, myFile)
}

// showFileDate := shows latest file date for user
func showFileDate(myDir string, myFile string) {

	filename := myDir + myFile

	// get last modified time
	file, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err)
	}

	modifiedtime := file.ModTime()

	fmt.Print("\n\t Last modified time : ", modifiedtime)
}
