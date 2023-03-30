package fthash

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// S3IOParam := gets user input for the bucket, directory, and region
func S3IOParam() (string, string, string) {

	var bucket, region, direct string

	fmt.Print("\n\t Enter a file directory path: ")
	fmt.Scanln(&direct)

	fmt.Print("\n\t Enter an S3 bucket name: ")
	fmt.Scanln(&bucket)

	fmt.Print("\n\t Enter the S3 bucket region: ")
	fmt.Scanln(&region)

	if bucket == "" {
		fmt.Print("\n\t Bucket value missing. Please try again.\n")
	}

	if direct == "" {
		fmt.Print("\n\t Directory value missing. Please try again.\n")
	}

	if region == "" {
		fmt.Print("\n\t Region value missing. Please try again.\n")
	}

	return direct, bucket, region
}

// RefreshTerminal := clears terminal screen
func RefreshTerminal() {
	fmt.Println("\n\t Please wait a few moments. The application is being refreshed.")
	time.Sleep(2 * time.Second)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
