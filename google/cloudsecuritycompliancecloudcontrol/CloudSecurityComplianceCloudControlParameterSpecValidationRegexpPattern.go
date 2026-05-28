// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern struct {
	// Regex Pattern to match the value(s) of parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/cloud_security_compliance_cloud_control#pattern CloudSecurityComplianceCloudControl#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
}

