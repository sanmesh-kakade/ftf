package module

// import (
// 	"fmt"
// 	"os"

// 	"github.com/sanmesh-kakade/ftf/internal/validation"
// 	"github.com/spf13/cobra"
// )

// // Flags for the command
// var (
// 	check                   bool
// 	skipTerraformValidation bool
// 	path                    string
// )

// // validateCmd represents the validate command
// var validateCmd = &cobra.Command{
// 	Use:   "validate",
// 	Short: "Validate the Terraform module and its security aspects",
// 	Long:  ``,
// 	Run:   validate,
// }

// func init() {

// 	validateCmd.Flags().BoolVar(&check, "check", false, "Check if Terraform files are correctly formatted without modifying them.")
// 	validateCmd.Flags().BoolVar(&skipTerraformValidation, "skip-terraform-validation", false, "Skip Terraform validation steps if set to true.")
// 	validateCmd.Flags().StringVarP(&path, "path", "p", "", "Path to the Terraform module directory.")

// 	validateCmd.MarkFlagRequired("path")

// 	ModuleCmd.AddCommand(validateCmd)
// }

// // Function to execute the command
// func validate(cmd *cobra.Command, args []string) {

// 	// Check if the path exists
// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		fmt.Printf("❌ Path does not exist: %s\n", path)
// 		return
// 	}

// 	options := &validation.ModuleValidatorOptions{
// 		Path:                    path,
// 		Check:                   check,
// 		SkipTerraformValidation: skipTerraformValidation,
// 	}

// 	// Create a new YAML moduleValidator
// 	moduleValidator := validation.NewModuleValidator(options)

// 	// Validate the facets.yaml file
// 	moduleValidator.ValidateTerraformModule()
// }
