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
	newLine   = []byte("\n")

	outputDir = kingpin.Flag(
		"output-dir",
		"Directory to write split files.",
	).Default("split_yaml").String()

	input = kingpin.Flag(
		"input",
		"Input file with kube manifests.",
	).String()

	splitCommand = kingpin.Command(
		"split",
		"Splits one kubernetes yaml file with multiple resources into separated files",
	)

	versionCommand = kingpin.Command(
		"version",
		"Display version and git commit of this binary",
	)

	version   = ""
	gitCommit = ""
)

// BuildInfo is a structure containing information about the build
type BuildInfo struct {
	// Version is the current semver.
	Version string `json:"version,omitempty"`
	// GitCommit is the git sha1.
	GitCommit string `json:"git_commit,omitempty"`
}

func main() {
	switch kingpin.Parse() {
	case versionCommand.FullCommand():
		buildInfo := BuildInfo{
			Version:   version,
			GitCommit: gitCommit,
		}
		log.Printf("%#v", buildInfo)
		return
	default:
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
}
