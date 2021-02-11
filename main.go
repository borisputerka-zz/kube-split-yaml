package main

import (
	"bufio"
	"io/ioutil"
	"log"
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
		inputData, err = ioutil.ReadFile(*input)
		if err != nil {
			log.Fatal(err)
    	}

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			log.Fatal("err: couldn't scan stdin")
		}
		for scanner.Scan() {
			inputData = append(inputData, scanner.Text()...)
			inputData = append(inputData, newLine...)
		}
	}

	err = pkg.SplitYaml(string(inputData), *outputDir)
	if err != nil {
		log.Fatal(err)
	}

}
