// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig struct {
	// CRM node in format organizations/{organization}, folders/{folder}, or projects/{project}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/cloud_security_compliance_framework_deployment#existing_target_resource CloudSecurityComplianceFrameworkDeployment#existing_target_resource}
	ExistingTargetResource *string `field:"optional" json:"existingTargetResource" yaml:"existingTargetResource"`
	// target_resource_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/cloud_security_compliance_framework_deployment#target_resource_creation_config CloudSecurityComplianceFrameworkDeployment#target_resource_creation_config}
	TargetResourceCreationConfig *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig `field:"optional" json:"targetResourceCreationConfig" yaml:"targetResourceCreationConfig"`
}

