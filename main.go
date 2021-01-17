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
	if *input != "" {
		inputData, err = ioutil.ReadFile(*input)
		if err != nil {
			fmt.Errorf("failed to read input file: %v", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Errorf("failed to read stdin: %v", scanner.Err())
		}
		inputData = scanner.Bytes()
	}

	err = pkg.SplitYaml(string(inputData), *outputDir)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

}
