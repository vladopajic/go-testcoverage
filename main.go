package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vladopajic/testcoverage/pkg/testcoverage"
)

//nolint:forbidigo // relax
func main() {
	cfg, err := readConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	stats, err := testcoverage.GenerateCoverageStats(cfg.Profile)
	if err != nil {
		fmt.Printf("failed to generate coverage statistics: %v\n", err)
		os.Exit(1)
	}

	if ok := testcoverage.Analyze(*cfg, stats); !ok {
		os.Exit(1)
	}
}

var errConfigNotSpecified = fmt.Errorf("-config argument not specified")

func readConfig() (*testcoverage.Config, error) {
	configPath := ""
	flag.StringVar(
		&configPath,
		"config",
		"",
		"testcoverage config file",
	)
	flag.Parse()

	if configPath == "" {
		return nil, errConfigNotSpecified
	}

	cfg, err := testcoverage.ConfigFromFile(configPath)
	if err != nil {
		return nil, err //nolint:wrapcheck // relax
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config file is not valid: %w", err)
	}

	return cfg, nil
}
