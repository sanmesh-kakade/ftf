package validation

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	"github.com/xeipuuv/gojsonschema"
// 	"gopkg.in/yaml.v3"
// )

// // YAML schema equivalent to the Python yaml_schema
// var yamlSchema = map[string]interface{}{
// 	"$schema": "https://json-schema.org/draft-07/schema#",
// 	"type":    "object",
// 	"properties": map[string]interface{}{
// 		"intent":      map[string]interface{}{"type": "string"},
// 		"flavor":      map[string]interface{}{"type": "string"},
// 		"version":     map[string]interface{}{"type": "string"},
// 		"description": map[string]interface{}{"type": "string"},
// 		"clouds": map[string]interface{}{
// 			"type": "array",
// 			"items": map[string]interface{}{
// 				"type": "string",
// 			},
// 		},
// 		"spec": map[string]interface{}{},
// 		"outputs": map[string]interface{}{
// 			"type": "object",
// 			"patternProperties": map[string]interface{}{
// 				".*": map[string]interface{}{
// 					"type": "object",
// 					"properties": map[string]interface{}{
// 						"type": map[string]interface{}{
// 							"type":    "string",
// 							"pattern": "^@outputs/.+",
// 						},
// 						"providers": map[string]interface{}{
// 							"type": "object",
// 							"patternProperties": map[string]interface{}{
// 								".*": map[string]interface{}{
// 									"type": "object",
// 									"properties": map[string]interface{}{
// 										"source":  map[string]interface{}{"type": "string"},
// 										"version": map[string]interface{}{"type": "string"},
// 										"attributes": map[string]interface{}{
// 											"type":                 "object",
// 											"additionalProperties": true,
// 										},
// 									},
// 									"required": []string{"source", "version", "attributes"},
// 								},
// 							},
// 						},
// 					},
// 					"required": []string{"type"},
// 				},
// 			},
// 		},
// 		"inputs": map[string]interface{}{
// 			"type": "object",
// 			"patternProperties": map[string]interface{}{
// 				".*": map[string]interface{}{
// 					"type": "object",
// 					"properties": map[string]interface{}{
// 						"type": map[string]interface{}{
// 							"type":    "string",
// 							"pattern": "^@outputs/.+",
// 						},
// 						"providers": map[string]interface{}{
// 							"type": "array",
// 							"items": map[string]interface{}{
// 								"type": "string",
// 							},
// 						},
// 					},
// 					"required": []string{"type"},
// 				},
// 			},
// 		},
// 		"sample": map[string]interface{}{"type": "object"},
// 		"artifact_inputs": map[string]interface{}{
// 			"type": "object",
// 			"properties": map[string]interface{}{
// 				"primary": map[string]interface{}{
// 					"type": "object",
// 					"properties": map[string]interface{}{
// 						"attribute_path": map[string]interface{}{"type": "string"},
// 						"artifact_type": map[string]interface{}{
// 							"type": "string",
// 							"enum": []string{"docker_image"},
// 						},
// 					},
// 					"required": []string{"attribute_path", "artifact_type"},
// 				},
// 			},
// 			"required": []string{"primary"},
// 		},
// 		"metadata": map[string]interface{}{},
// 	},
// 	"required": []string{"intent", "flavor", "version", "description", "spec"},
// }

// type YAMLValidator interface {
// 	ValidateFacetsYAML() error
// }

// type YAMLValidatorImpl struct {
// 	path string
// }

// // ValidateFacetsYAML implements YAMLValidator.
// func (y *YAMLValidatorImpl) ValidateFacetsYAML() error {

// 	yamlPath := filepath.Join(y.path, "facets.yaml")
// 	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
// 		return fmt.Errorf("❌ facets.yaml file does not exist at %s", yamlPath)
// 	}

// 	content, err := os.ReadFile(yamlPath)
// 	if err != nil {
// 		return fmt.Errorf("❌ failed to read facets.yaml: %v", err)
// 	}

// 	// Parse the YAML content into a map
// 	var data map[string]interface{}
// 	if err := yaml.Unmarshal(content, &data); err != nil {
// 		return fmt.Errorf("❌ facets.yaml is not a valid YAML file: %v", err)
// 	}

// 	// Convert the schema and data to JSON loaders
// 	schemaLoader := gojsonschema.NewGoLoader(yamlSchema)
// 	dataLoader := gojsonschema.NewGoLoader(data)

// 	// Perform the validation
// 	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
// 	if err != nil {
// 		return fmt.Errorf("error during validation: %w", err)
// 	}

// 	// Check if the validation was successful
// 	if !result.Valid() {
// 		// Collect all validation errors
// 		var validationErrors string
// 		for _, desc := range result.Errors() {
// 			validationErrors += fmt.Sprintf("- %s\n", desc)
// 		}
// 		return fmt.Errorf("validation failed:\n%s", validationErrors)
// 	}

// 	fmt.Println("Facets YAML validation successful!")
// 	return nil
// }

// func NewYAMLValidator(path string) YAMLValidator {

// 	// Read the YAML file
// 	return &YAMLValidatorImpl{
// 		path: path,
// 	}
// }
