package module

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

const (
	FACETS_YAML_CONTENT = "intent: test-intent\nflavor: test-flavor\nversion: \"1.0\"\nclouds: [aws]\ndescription: Test Description\nspec:\n  title: Test Title\n  description: Test Description\n  type: object\noutputs:\n  default:\n     type: \"@outputs/test-intent\"\nsample:\n  kind: test-intent\n  flavor: test-flavor\n  version: \"1.0\"\n  disabled: true\n  spec:\n"
	MAIN_TF_CONTENT     = "# main.tf - This file is for defining the main resources for the module\n"
	OUTPUT_TF_CONTENT   = "locals {\n  output_interfaces = {}\n  output_attributes = {}\n}\n"
	VARIABLE_TF_CONTENT = "variable \"instance\" {\n  description = \"Test Description\"\n  type = object({\n    kind    = string\n    flavor  = string\n    version = string\n    spec = object({\n    })\n  })\n}\nvariable \"instance_name\" {\n  description = \"The architectural name for the resource as added in the Facets blueprint designer.\"\n  type        = string\n}\nvariable \"environment\" {\n  description = \"An object containing details about the environment.\"\n  type = object({\n    name        = string\n    unique_name = string\n  })\n}\nvariable \"inputs\" {\n  description = \"A map of inputs requested by the module developer.\"\n  type        = object({})\n}\n"
)

func TestGenerateModuleWithRealTemplates(t *testing.T) {
	// Set global variables for the test
	intent = "test-intent"
	flavor = "test-flavor"
	cloud = "aws"
	title = "Test Title"
	description = "Test Description"

	// Mock a *cobra.Command
	mockCmd := &cobra.Command{}

	curDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	// Execute the generateModule function
	generateModule(mockCmd, []string{})

	// Clean up the generated files
	defer os.RemoveAll(filepath.Join(curDir, intent))

	// Verify the output directory structure
	outputDir := filepath.Join(curDir, intent, flavor, "1.0")
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Fatalf("Output directory %s was not created", outputDir)
	}

	// Verify the generated files
	expectedFiles := []string{"main.tf", "variables.tf", "output.tf", "facets.yaml"}
	for _, file := range expectedFiles {
		filePath := filepath.Join(outputDir, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected file %s was not created", filePath)
		}
	}

	// Verify the content of one of the generated files
	variablesFile := filepath.Join(outputDir, "variables.tf")
	content, err := os.ReadFile(variablesFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", variablesFile, err)
	}
	// Check if the content contains the expected description
	if !reflect.DeepEqual(string(content), VARIABLE_TF_CONTENT) {
		t.Errorf("Content of %s does not contain the expected description", variablesFile)
	}

	// Verify the content of one of the generated files
	facetsYAMLFile := filepath.Join(outputDir, "facets.yaml")
	content, err = os.ReadFile(facetsYAMLFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", facetsYAMLFile, err)
	}
	// Check if the content contains the expected description
	if !reflect.DeepEqual(string(content), FACETS_YAML_CONTENT) {
		t.Errorf("Content of %s does not contain the expected description", facetsYAMLFile)
	}

	// Verify the content of one of the generated files
	mainFile := filepath.Join(outputDir, "main.tf")
	content, err = os.ReadFile(mainFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", mainFile, err)
	}
	// Check if the content contains the expected description
	if !reflect.DeepEqual(string(content), MAIN_TF_CONTENT) {
		t.Errorf("Content of %s does not contain the expected description", mainFile)
	}

	// Verify the content of one of the generated files
	outputFile := filepath.Join(outputDir, "output.tf")
	content, err = os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", outputFile, err)
	}
	// Check if the content contains the expected description
	if !reflect.DeepEqual(string(content), OUTPUT_TF_CONTENT) {
		t.Errorf("Content of %s does not contain the expected description", outputFile)
	}
}
