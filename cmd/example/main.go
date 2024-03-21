package main

import (
	"flag"
	"fmt"
	"github.com/KPI-3-Architecture-Labs/lab2"
	"io"
	"log"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to store computed expression")
)

func main() {
	flag.Parse()

	if *inputExpression == "" && *inputFile == "" {
		log.Fatal("no expression provided.")
	}

	if *inputExpression != "" && *inputFile != "" {
		log.Fatal("flags -e and -f can't both be used")
	}

	var reader io.Reader

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			log.Fatal("no file found")
		}
		reader = file
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal("Error closing file")
			}
		}(file)
	}

	var writer io.Writer

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal("something went wrong while creating file")
		}
		writer = file
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal("Error closing file")
			}
		}(file)
	} else {
		writer = &Writer{}
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		log.Fatal("Error, ", err)
	}
}

type Writer struct{}

func (w *Writer) Write(data []byte) (n int, err error) {
	fmt.Println(string(data))
	return len(data), nil
}
