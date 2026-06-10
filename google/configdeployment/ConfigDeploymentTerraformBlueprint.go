// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configdeployment


type ConfigDeploymentTerraformBlueprint struct {
	// URI of a GCS object containing the zipped Terraform blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/config_deployment#gcs_source ConfigDeployment#gcs_source}
	GcsSource *string `field:"optional" json:"gcsSource" yaml:"gcsSource"`
	// git_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/config_deployment#git_source ConfigDeployment#git_source}
	GitSource *ConfigDeploymentTerraformBlueprintGitSource `field:"optional" json:"gitSource" yaml:"gitSource"`
	// input_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/config_deployment#input_values ConfigDeployment#input_values}
	InputValues interface{} `field:"optional" json:"inputValues" yaml:"inputValues"`
}

