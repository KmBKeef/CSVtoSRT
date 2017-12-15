package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func returnError() (x []string, err error) {
	return
}

// func doesExist(fileToCheck string) (x []string, err error) {
// 	if _, err := os.Stat(fileToCheck); err == nil {
// 	}
// 	fmt.Printf("\n File %s not found.", fileToCheck)
// 	return

// }

func ReadCsvFile(filePath string) {
	f, _ := os.Open(filePath)

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage : " + os.Args[0] + " <file name>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	if _, err := os.Stat(fileName); err == nil {
		ReadCsvFile(fileName)
	} else {
		fmt.Printf("File %s not found.\n", fileName)
		returnError()
	}

}
