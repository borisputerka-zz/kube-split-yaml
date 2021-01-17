package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/borisputerka/kube-split-yaml/pkg"
)

var (
	suffix string
)

func main() {
	if strings.HasSuffix(os.Args[1], ".yaml") {
		suffix = ".yaml"
	} else if strings.HasSuffix(os.Args[1], ".yaml") {
		suffix = ".yml"
	} else {
		fmt.Errorf("missing a yaml/yml suffix")
	}

	inputData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Errorf("failed to read input file: %v", err)
	}

	filePathSplit := strings.Split(os.Args[1], "/")
	dirName := strings.TrimSuffix(filePathSplit[len(filePathSplit)-1], suffix)
	err = os.Mkdir(dirName, 0751)
	if err != nil {
		fmt.Errorf("create directory: %v", err)
	}

	err = pkg.SplitYaml(string(inputData), dirName)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

}
