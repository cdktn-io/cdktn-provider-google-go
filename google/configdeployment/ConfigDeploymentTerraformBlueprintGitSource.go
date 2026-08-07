// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configdeployment


type ConfigDeploymentTerraformBlueprintGitSource struct {
	// Repository URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/config_deployment#repo ConfigDeployment#repo}
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// Subdirectory within the repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/config_deployment#directory ConfigDeployment#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Git reference (branch or tag).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/config_deployment#ref ConfigDeployment#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

