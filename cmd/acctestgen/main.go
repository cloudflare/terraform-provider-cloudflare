package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	service := flag.Arg(0)
	err := Validate(service)
	if err != nil {
		fmt.Printf("\nERROR:\n%s\n\n", err)
		os.Exit(1)
	}
	PrintPreview(service)
}

// PrintPreview prints a list of paths that represent the files and directories
// that would be created.
func PrintPreview(service string) {
	fmt.Println("To be created:")
	fmt.Printf("- file: internal/services/%s/data_source_test.go\n", service)
	fmt.Printf("- file: internal/services/%s/resource_test.go\n", service)
	fmt.Printf("- dir: internal/services/%s/testdata/\n", service)
	fmt.Printf("- file: internal/services/%s/testdata/config.tf\n", service)
}

// Validate checks that the service provided exists, and to prevent
// overwriting existing test configurations and test cases, that
// none of the files and directories that would be generated already exist.
func Validate(service string) error {
	if service == "" {
		return fmt.Errorf("a service name must be provided")
	}

	basePath := fmt.Sprintf("./internal/services/%s/", service)
	// check if service does not exist
	if _, err := os.Stat(basePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("service '%s' not found", service)
	}
	// check if any of the test files/configs already exist
	var errs []error
	if _, err := os.Stat(basePath + "resource_test.go"); err == nil {
		errs = append(errs, fmt.Errorf("the service '%s' already has a 'resource_test.go' file", service))
	}
	if _, err := os.Stat(basePath + "data_source_test.go"); err == nil {
		errs = append(errs, fmt.Errorf("the service '%s' already has a 'data_source_test.go' file", service))
	}
	if _, err := os.Stat(basePath + "testdata"); err == nil {
		errs = append(errs, fmt.Errorf("the service '%s' already has a 'testdata' directory", service))
	}

	return errors.Join(errs...)
}
