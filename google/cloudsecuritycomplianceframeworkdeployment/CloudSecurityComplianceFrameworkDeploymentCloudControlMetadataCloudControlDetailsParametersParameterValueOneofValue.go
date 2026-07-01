// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParametersParameterValueOneofValue struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/cloud_security_compliance_framework_deployment#name CloudSecurityComplianceFrameworkDeployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/cloud_security_compliance_framework_deployment#parameter_value CloudSecurityComplianceFrameworkDeployment#parameter_value}
	ParameterValue *CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParametersParameterValueOneofValueParameterValue `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

