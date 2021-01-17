# kube-split-yaml
Simple tool to split one file with kubernetest manifests into smaller ones.

# Installation and usage
1. Install
	```
	$ got get -v  github.com/borisputerka/kube-split-yaml
	```
2. Use plugin
	```
	$ kube-split-yaml --input <input_file> --outpud-dir <output_dir>
	```

# Flags reference
Flag         | Description                                        | Default
-------------|----------------------------------------------------|-------------
`input`      | Input file to split. Can be replaced with `stdin`  | `""`
`output-dir` | Output directory where split yamls will be written | `split_yaml`
