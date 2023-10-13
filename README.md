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
