package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/borisputerka/kube-split-yaml/pkg"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	suffix    string
	inputData []byte
	err       error
	newLine = []byte("\n")

	outputDir = kingpin.Flag(
		"output-dir",
		"Directory to write split files.",
	).Default("split_yaml").String()

	input = kingpin.Flag(
		"input",
		"Input file with kube manifests.",
	).String()
)

func main() {
	kingpin.Parse()
	if *input != "" {
		fmt.Println("som tu")
		inputData, err = ioutil.ReadFile(*input)
		if err != nil {
			fmt.Errorf("failed to read input file: %v", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Errorf("failed to read stdin: %v", scanner.Err())
		}
		for scanner.Scan() {
			inputData = append(inputData, scanner.Text()...)
			inputData = append(inputData, newLine...)
		}
	}

	err = pkg.SplitYaml(string(inputData), *outputDir)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

}
