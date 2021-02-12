package pkg

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

var (
	groupByKind = kingpin.Flag(
		"group-by-kind",
		"Group resources with same Kind into one file",
	).Bool()
)

const (
	yamlSplitter = "---\n"
)

type manifestCore struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   metadata `yaml:"metadata"`
}

type metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

func SplitYaml(inputData string, dirName string) error {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0751)
		if err != nil {
			return err
		}
	}

	for _, document := range strings.Split(inputData, yamlSplitter) {
		var manifest manifestCore
		err = yaml.Unmarshal([]byte(document), &manifest)
		if err != nil {
			return err
		}

		var fileName string
		if *groupByKind {
			fileName = strcase.ToKebab(manifest.Kind) + ".yaml"
			f, err := os.OpenFile(dirName+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			_, err = f.Write([]byte(yamlSplitter + document))
			if err != nil {
				return err
			}
			err = f.Close()
			if err != nil {
				return err
			}
		} else {
			apiVersionDotted := strings.ReplaceAll(manifest.ApiVersion, "/", ".")
			if manifest.Metadata.Namespace == "" {
				fileName = strings.Join(
					[]string{manifest.Kind, apiVersionDotted, manifest.Metadata.Name + ".yaml"},
					"-")
			} else {
				fileName = strings.Join(
					[]string{manifest.Kind, apiVersionDotted, manifest.Metadata.Namespace, manifest.Metadata.Name + ".yaml"},
					"-")
			}

			err = ioutil.WriteFile(dirName+"/"+fileName, []byte(document), 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
