package validation

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// type TerraformValidator interface {
// 	ValidateTerraformModule() error
// }

// type TerraformValidatorImpl struct {
// 	path string
// }

// // ValidateTerraformModule implements TerraformValidator.
// func (t *TerraformValidatorImpl) ValidateTerraformModule() error {
// 	panic("unimplemented")
// }

// func NewTerraformValidator(path string) TerraformValidator {
// 	return &TerraformValidatorImpl{}
// }

// // Helper function to run a command
// func runCommand(command []string, workingDir string) error {
// 	cmd := exec.Command(command[0], command[1:]...)
// 	if workingDir != "" {
// 		cmd.Dir = workingDir
// 	}
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }

// func runTfsec(path string) error {
// 	// Ensure the path exists
// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		return fmt.Errorf("❌ specified path does not exist: %s", path)
// 	}

// 	// Check if the tfsec binary is installed
// 	if _, err := exec.LookPath("tfsec"); err != nil {
// 		fmt.Println("❌ tfsec is not installed. Please install it using the following command:")
// 		fmt.Println("   brew install tfsec   # For macOS")
// 		fmt.Println("   scoop install tfsec  # For Windows")
// 		fmt.Println("   apt install tfsec    # For Linux (Debian-based)")
// 		return fmt.Errorf("tfsec binary not found")
// 	}

// 	// Run the tfsec command
// 	cmd := exec.Command("tfsec", path, "-m", "CRITICAL")

// 	if output, err := cmd.CombinedOutput(); err != nil {
// 		return fmt.Errorf("❌ tfsec validation failed: %v\n%v", err, string(output))
// 	}

// 	return nil
// }
