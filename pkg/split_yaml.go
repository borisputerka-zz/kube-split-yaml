package pkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

const (
yamlSplitter = "---"
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
	for _, document := range strings.Split(string(inputData), yamlSplitter) {
		var manifest manifestCore
		err := yaml.Unmarshal([]byte(document), &manifest)
		if err != nil {
			return err
		}

		manifest.ApiVersion = strings.ReplaceAll(manifest.ApiVersion, "/", ".")
		var fileName string
		if manifest.Metadata.Namespace == "" {
			fileName = strings.Join(
				[]string{manifest.Kind, manifest.ApiVersion, manifest.Metadata.Name + ".yaml"},
				"-")
		} else {
			fileName = strings.Join(
				[]string{manifest.Kind, manifest.ApiVersion, manifest.Metadata.Namespace, manifest.Metadata.Name + ".yaml"},
				"-")
		}

		err = ioutil.WriteFile(dirName+"/"+fileName, []byte(document), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}