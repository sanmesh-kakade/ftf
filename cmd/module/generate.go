package module

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	ModuleTemplatePath = "../../templates/module"
)

var (
	intent      string
	flavor      string
	cloud       string
	title       string
	description string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new terraform module template",
	Long:  `This command is used to generate a skeleton terraform module template by templating predefined jinja2 files which follows the Facets.Cloud module structure.`,
	Run:   generateModule,
}

func init() {

	generateCmd.Flags().StringVarP(&intent, "intent", "i", "", "Intent of the module")
	generateCmd.Flags().StringVarP(&flavor, "flavor", "f", "", "Flavor of the module")
	generateCmd.Flags().StringVarP(&cloud, "cloud", "c", "", "Cloud provider for the module")
	generateCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the module")
	generateCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the module")

	generateCmd.MarkFlagRequired("intent")
	generateCmd.MarkFlagRequired("flavor")
	generateCmd.MarkFlagRequired("cloud")
	generateCmd.MarkFlagRequired("title")
	generateCmd.MarkFlagRequired("description")

	ModuleCmd.AddCommand(generateCmd)
}

// Function to execute the command
func generateModule(cmd *cobra.Command, args []string) {

	// Create the directory structure
	path := filepath.Join(intent, flavor, "1.0")
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("❌ Failed to create directory: %v\n", err)
		return
	}

	// Get the directory of the current file (generate.go)
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("❌ Failed to get the current file path")
		return
	}
	currentDir := filepath.Dir(currentFile)

	// Resolve the templates path relative to the current file
	templatesPath := filepath.Join(currentDir, ModuleTemplatePath)

	// List of template files to process
	templateFiles := []string{"main.tf.j2", "variables.tf.j2", "output.tf.j2", "facets.yaml.j2"}

	// Render and write templatess
	for _, templateFile := range templateFiles {
		templatePath := filepath.Join(templatesPath, templateFile)

		// Parse the template file
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			fmt.Printf("❌ Failed to parse template %s: %v\n", templateFile, err)
			return
		}

		// Render the template with data
		outputFileName := templateFile[:len(templateFile)-3] // Remove ".j2" from the file name
		outputFilePath := filepath.Join(path, outputFileName)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			fmt.Printf("❌ Failed to create file %s: %v\n", outputFilePath, err)
			return
		}
		defer outputFile.Close()

		// Execute the template with the provided data
		err = tmpl.Execute(outputFile, map[string]string{
			"intent":      intent,
			"flavor":      flavor,
			"cloud":       cloud,
			"title":       title,
			"description": description,
		})
		if err != nil {
			fmt.Printf("❌ Failed to render template %s: %v\n", templateFile, err)
			return
		}
	}

	fmt.Printf("✅ Module generated at: %s\n", path)
}
