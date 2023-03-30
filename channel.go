package main

import (
	"aWS/fthash"
	"aWS/ftios3"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// FileParam := stores filename & path for user
type FileParam struct {
	fileName string
	path     string
}

const buffLen = 1000000
const numGoRoutines = 2

var c1 = make(chan FileParam, buffLen)
var c2 = make(chan string)
var wg = sync.WaitGroup{}
var numFiles = 0
var direct, bucket, region = fthash.S3IOParam()

func main() {

	start := time.Now()

	for i := 1; i <= numGoRoutines; i++ {
		wg.Add(1)
		go sendFile(i)
	}

	run()

	for len(c1) > 0 {
		fmt.Println("\n\t Channel Not Empty. Msgs:", len(c1))
		time.Sleep(time.Minute * 1)
	}

	elapsed := time.Since(start)
	fmt.Println("\n\t Total Elasped Time: ", elapsed.Minutes())

	for i := 1; i <= numGoRoutines; i++ {
		c2 <- "Exit"
	}

	close(c1)
	fthash.RefreshTerminal()
	fmt.Println("\n\t All Files Found! \n\t Total Files Uploaded: ", numFiles)
	fmt.Println("\n\t All Directory Files have been sent to S3. Program exited: 0")
	os.Exit(0)
}

func sendFile(i int) {
	numMessages := 0
	for {
		select {

		case msg1 := <-c1:
			numMessages++
			ftios3.PutFileS3(direct, msg1.fileName, bucket, region)
			runtime.Gosched()
			time.Sleep(500 * time.Millisecond)

		case msg2 := <-c2:
			fmt.Println("\n\t P: ", i, "\n\t Total Messages Processed: ", numMessages)
			fmt.Println("\n\t Break: ", msg2)
			return

		default:
			runtime.Gosched()
			fmt.Println("\n\t Files Received: ", i, "\n\t Files Processed: ", numMessages)
			fthash.RefreshTerminal()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func run() error {

	searchDir := direct

	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if f.IsDir() {
			return err
		}

		numFiles++
		fName := f.Name()
		indx := strings.LastIndex(path, fName)
		newPath := truncate(path, int(indx))
		c1 <- FileParam{fName, newPath}
		return err
	})

	if e != nil {
		panic(e)
	}

	return nil
}

func truncate(s string, to int) string {
	return s[:to]
}
