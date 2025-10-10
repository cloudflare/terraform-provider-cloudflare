package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/resources"
)

func main() {
	// Command line flags
	var resourceList string
	flag.StringVar(&resourceList, "resources", "", "Comma-separated list of resources to enable (empty = all)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [flags] <terraform_file.tf>")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		fmt.Println("\nAvailable resources:", strings.Join(resources.GetAvailableResources(), ", "))
		os.Exit(1)
	}

	filename := flag.Arg(0)
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Create registry
	reg := registry.NewStrategyRegistry()

	// Register resources based on flags
	if resourceList == "" {
		// Register all available resources
		resources.RegisterAll(reg)
		log.Printf("Registered all %d available resources\n", reg.Count())
	} else {
		// Register specific resources
		requestedResources := strings.Split(resourceList, ",")
		resources.RegisterFromFactories(reg, requestedResources...)
		log.Printf("Registered %d resources: %s\n", reg.Count(), resourceList)
	}

	// Build the transformation pipeline
	pipeline := poc.BuildPipeline(reg)

	// TransformConfig the content
	result, err := pipeline.Transform(content, filename)
	if err != nil {
		log.Fatalf("Transformation failed: %v", err)
	}

	// Output the result
	fmt.Print(string(result))
}
