package scanner

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type ScanOptions struct {
	Format   string
	SkipPull bool
}

type CommandRunner interface {
	RunCommand(name string, arg ...string) (output []byte, err error)
}

// ScanImage scans a Docker container image for vulnerabilities using Trivy
func ScanImage(imageName string, options ScanOptions, runner CommandRunner) (int, error) {
	log.Printf("Pulling image: %s", imageName)

	if !options.SkipPull {
		log.Printf("Pulling image: %s", imageName)
		pullCmd := exec.Command("docker", "pull", imageName)
		pullOutput, pullErr := pullCmd.CombinedOutput()
		if pullErr != nil {
			log.Fatalf("Error pulling image: %v\n%s", pullErr, pullOutput)
			return 0, fmt.Errorf("error pulling image: %v\n%s", pullErr, pullOutput)
		}
	}

	// First, pull the Docker image

	log.Printf("Scanning image: %s", imageName)

	// Run Trivy command to scan the image
	trivyCmd := exec.Command("trivy", "image", "--format", "json", imageName)
	scanOutput, scanErr := trivyCmd.CombinedOutput()
	if scanErr != nil {
		log.Fatalf("Error scanning image: %v\n%s", scanErr, scanOutput)
		return 0, fmt.Errorf("error scanning image: %v\n%s", scanErr, scanOutput)
	}

	// Parse Trivy output for vulnerabilities (you can extend this part)
	vulnerabilities := strings.Count(string(scanOutput), "VulnerabilityID")
	log.Printf("Vulnerabilities found: %d", vulnerabilities)

	return vulnerabilities, nil
}
