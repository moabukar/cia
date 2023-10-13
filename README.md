# Container Image Analyser (CIA)

Container Image Analyser (CIA) is a lightweight CLI tool designed to help DevOps engineers and developers analyse Docker container images for vulnerabilities and potential issues. It simplifies the process of ensuring the security and quality of containerized applications.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Contributing](#contributing)
  - [Development Setup](#development-setup)
  - [Adding New Features](#adding-new-features)
- [License](#license)

## Features

- **Vulnerability Scanning**: CIA scans Docker container images for known vulnerabilities and rates their severity.
- **Detailed Reports**: Generate detailed reports about the vulnerabilities found, including their CVE IDs, severity levels, and package information.
- **Remediation Suggestions**: Get recommendations for remediation, such as updating specific packages or using alternative base images.
- **User-Friendly CLI**: An easy-to-use command-line interface that supports a variety of options for scanning and reporting.

## 📝 Usage

```bash
NAME:
   cia - Analyse container images

USAGE:
   cia [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value  Report format (json, xml) (default: "json")
   --skip-pull     Skip pulling image before scanning
   --help, -h      show help

```

## CLI usage

```bash
./cia --help
Container Image Analyser (CIA) - Deep dive into your container images!

Description:
  CIA is your go-to CLI tool for analyzing container images. It can pull images, scan for vulnerabilities, and output reports in multiple formats.

Usage:
  ./cia [command] [flags]

Commands:
  scan          Scans the given container image
  report        Generate a report of the last scan
  version       Show the version of CIA tool

Global Flags:
  -f, --format string      Choose report format. Options: json, xml (default "json")
  -h, --help               Show help information for CIA or a specific command
  -s, --skip-pull          Skip pulling the image from the registry, use local image
  -v, --verbose            Enable verbose output

Examples:
  ./cia scan ubuntu:latest             Scans the 'ubuntu:latest' image
  ./cia report -f xml                  Generate the last scan report in XML format
  ./cia scan alpine:latest --skip-pull  Scans the 'alpine:latest' image without pulling it

For more details and usage examples, visit https://github.com/moabukar/cia

```

## Getting Started

### Installation

You can install CIA via `go get` or by downloading the binary release for your platform from the [Releases](https://github.com/moabukar/cia/releases) page.

```bash
go get github.com/moabukar/cia

## Usage

```bash
cia scan my-container-image:latest

```

## Directory structure

```bash
cia/
├── cmd/
│   └── cia/
│       └── main.go               # Entry point of the CLI application
├── internal/
│   ├── cmd/                      # Command-related logic
│   │   └── ...
│   ├── cli/                      # CLI-specific code
│   │   └── ...
│   ├── scanner/
│   │   ├── scanner.go           # Image scanning logic
│   │   └── vulnerabilities.go   # Vulnerability database integration
│   └── report/
│       ├── report.go            # Report generation logic
│       └── remediation.go       # Remediation suggestions
├── data/
│   └── vulnerabilities.json     # Local copy of vulnerability database (optional)
├── tests/
│   ├── unit/      
│   └── integration/        
├── README.md                 
├── LICENSE                       
├── .gitignore                
├── go.mod                    
└── go.sum                       

```
