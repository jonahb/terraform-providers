package main

import "github.com/hashicorp/terraform/configs"
import "os"
import "path/filepath"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %v <path-to-config>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	configPath := os.Args[1]
	parser := configs.NewParser(nil)
	mod, diags := parser.LoadConfigDir(configPath)

	if len(diags) != 0 {
		for _, diag := range diags {
			fmt.Fprintln(os.Stderr, diag)
			os.Exit(1)
		}
	}

	for _, provider := range mod.ProviderConfigs {
		fmt.Println(provider.Name)
	}
}
