// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configdeployment


type ConfigDeploymentTerraformBlueprintInputValues struct {
	// The value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/config_deployment#input_value ConfigDeployment#input_value}
	InputValue *string `field:"required" json:"inputValue" yaml:"inputValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/config_deployment#variable_name ConfigDeployment#variable_name}.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
}

