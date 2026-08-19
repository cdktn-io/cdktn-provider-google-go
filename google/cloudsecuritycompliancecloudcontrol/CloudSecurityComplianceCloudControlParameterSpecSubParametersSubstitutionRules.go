// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecSubParametersSubstitutionRules struct {
	// attribute_substitution_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/cloud_security_compliance_cloud_control#attribute_substitution_rule CloudSecurityComplianceCloudControl#attribute_substitution_rule}
	AttributeSubstitutionRule *CloudSecurityComplianceCloudControlParameterSpecSubParametersSubstitutionRulesAttributeSubstitutionRule `field:"optional" json:"attributeSubstitutionRule" yaml:"attributeSubstitutionRule"`
	// placeholder_substitution_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/cloud_security_compliance_cloud_control#placeholder_substitution_rule CloudSecurityComplianceCloudControl#placeholder_substitution_rule}
	PlaceholderSubstitutionRule *CloudSecurityComplianceCloudControlParameterSpecSubParametersSubstitutionRulesPlaceholderSubstitutionRule `field:"optional" json:"placeholderSubstitutionRule" yaml:"placeholderSubstitutionRule"`
}

