# testcoverage

[![lint](https://github.com/vladopajic/testcoverage/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/vladopajic/testcoverage/actions/workflows/lint.yml)
[![test](https://github.com/vladopajic/testcoverage/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vladopajic/testcoverage/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/vladopajic/testcoverage?cache=v1)](https://goreportcard.com/report/github.com/vladopajic/testcoverage)
[![codecov](https://codecov.io/gh/vladopajic/testcoverage/branch/main/graph/badge.svg?token=WYCKb1MLgl)](https://codecov.io/gh/vladopajic/testcoverage)
[![GoDoc](https://godoc.org/github.com/vladopajic/testcoverage?status.svg)](https://godoc.org/github.com/vladopajic/testcoverage)
[![Release](https://img.shields.io/github/release/vladopajic/testcoverage.svg?style=flat-square)](https://github.com/vladopajic/testcoverage/releases/latest)


`testcoverage` is tool which reports issues when test coverage of a file or package is below set threshold.

### Usage

```yml
name: Go test coverage check
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
      
      - name: test (generate coverage)
        run: go test ./... -coverprofile=./cover.out

      - name: check test coverage
        uses: vladopajic/testcoverage@v1
        with:
          config: ./.testcoverage.yml  # (config file is mandatory)
```

### Config
Example of [.testcoverage.yml](./.testcoverage.example.yml) config file.

```yml
# (mandatory) Path to coverprofile file (output of `go test -coverprofile` command)
profile: cover.out

# (optional) When specified reported file paths will not contain local prefix in the output
localPrefix: "github.com/user/project"

# Holds coverage thresholds percentages, values should be in range [0-100]
threshold:
  # (optional; default 0) The minimum coverage that each file should have
  file: 80

  # (optional; default 0) The minimum coverage that each package should have
  package: 80

  # (optional; default 50) The minimum total coverage project should have
  total: 95
```

## Contribution

All contributions are useful, whether it is a simple typo, a more complex change, or just pointing out an issue. We welcome any contribution so feel free to open PR or issue. 
