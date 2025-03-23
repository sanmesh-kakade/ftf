package validation

// type ModuleValidator interface {
// 	ValidateTerraformModule()
// }

// type ModuleValidatorImpl struct {
// 	yamlValidator YAMLValidator
// 	options       *ModuleValidatorOptions
// }

// type ModuleValidatorOptions struct {
// 	Check                   bool
// 	SkipTerraformValidation bool
// 	Path                    string
// }

// func (m *ModuleValidatorImpl) ValidateTerraformModule() {

// 	err := m.yamlValidator.ValidateFacetsYAML()
// 	if err != nil {
// 		// return err
// 	}

// 	// // m.options.checkif Terraform is installed
// 	// if _, err := exec.LookPath("terraform"); err != nil {
// 	// 	return fmt.Errorf("❌ Terraform is not installed. Please install Terraform to continue.")
// 	// }

// 	// // Run terraform fmt
// 	// fmtCommand := []string{"terraform", "fmt"}
// 	// if m.options.Check {
// 	// 	fmtCommand = append(fmtCommand, "-check")
// 	// }
// 	// if err := runCommand(fmtCommand, m.options.Path); err != nil {
// 	// 	if m.options.Check {
// 	// 		return fmt.Errorf("❌ Error: Terraform files are not correctly formatted. Please run `terraform fmt` locally to format the files.")
// 	// 	} else {
// 	// 		return fmt.Errorf("❌ Error formatting Terraform files: %v\n", err)
// 	// 	}
// 	// }
// 	// if m.options.Check {
// 	// 	fmt.Println("✅ Terraform files are correctly formatted.")
// 	// } else {
// 	// 	fmt.Println("🎨 Terraform files formatted.")
// 	// }

// 	// // Skip Terraform validation if the flag is set
// 	// if m.options.SkipTerraformValidation {
// 	// 	fmt.Println("⏭ Skipping Terraform validation as per flag.")
// 	// } else {
// 	// 	// Run terraform init
// 	// 	if err := runCommand([]string{"terraform", "-chdir=" + m.options.Path, "init", "-backend=false"}, ""); err != nil {
// 	// 		return fmt.Errorf("❌ Error initializing Terraform: %v\n", err)
// 	// 	}
// 	// 	fmt.Println("🚀 Terraform initialized.")

// 	// 	// Run terraform validate
// 	// 	if err := runCommand([]string{"terraform", "-chdir=" + m.options.Path, "validate"}, ""); err != nil {
// 	// 		return fmt.Errorf("❌ Terraform validation failed: %v\n", err)
// 	// 	}
// 	// 	fmt.Println("🔍 Terraform validation successful.")
// 	// }

// 	// // Run tfsec
// 	// if err := runTfsec(m.options.Path); err != nil {
// 	// 	return fmt.Errorf("❌ Error encountered while running tfsec validation.\n%s", err)
// 	// }
// 	// fmt.Println("✅ tfsec validation passed.")

// 	// return nil
// }

// func NewModuleValidator(options *ModuleValidatorOptions) ModuleValidator {
// 	YamlValidator := NewYAMLValidator(options.Path)
// 	return &ModuleValidatorImpl{
// 		yamlValidator: YamlValidator,
// 		options:       options,
// 	}
// }
