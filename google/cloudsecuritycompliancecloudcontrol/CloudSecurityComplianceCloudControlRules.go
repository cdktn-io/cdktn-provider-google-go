// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlRules struct {
	// The functionality enabled by the Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/cloud_security_compliance_cloud_control#rule_action_types CloudSecurityComplianceCloudControl#rule_action_types}
	RuleActionTypes *[]*string `field:"required" json:"ruleActionTypes" yaml:"ruleActionTypes"`
	// cel_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/cloud_security_compliance_cloud_control#cel_expression CloudSecurityComplianceCloudControl#cel_expression}
	CelExpression *CloudSecurityComplianceCloudControlRulesCelExpression `field:"optional" json:"celExpression" yaml:"celExpression"`
	// Description of the Rule. The maximum length is 2000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/cloud_security_compliance_cloud_control#description CloudSecurityComplianceCloudControl#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

