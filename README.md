# Kube Split Yaml
Simple tool to split one file with kubernetes manifests into separated ones.

# Use cases
Useful when you are using [kustomize](https://kustomize.io) and you want to evaluate generated manifests after `kustomize build .` with [OPA](https://www.openpolicyagent.org).

	$ kustomize build <path_to_kustomization> >> build.yaml
	$ kube-split-yaml split --input build.yaml --output-dir manifests
	$ for i in $(ls manifests); do opa eval -d <validation_rego> -i manifests/$i data.kubernetes.admission.deny --format=pretty; done

# Installation and usage
1. Install
	```
	$ go get -v github.com/borisputerka/kube-split-yaml
	```
2. Use plugin
	```
	$ kube-split-yaml split --input <input_file> --output-dir <output_dir>
	```

# Flags reference
By default `kube-split-yaml` splits each kubernetes resource from input file into separated file in format:

```go
{{.kind}}-{{.apiVersion}}-{{.metadata.namespace}}-{{.metadata.name}}.yaml
```

Here are flags that can be used with `split` command:

Flag            | Description                                        | Default
----------------|----------------------------------------------------|-------------
`input`         | Input file to split. Can be replaced with `stdin`                | `""`
`output-dir`    | Output directory where divided yamls will be written             | `split_yaml`
`group-by-kind` | Group manifests based on their `kind` (in format `<.kind>.yaml`) | `false`
