// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#action DataLossPreventionContentPolicy#action}
	Action *DataLossPreventionContentPolicyRulesAction `field:"required" json:"action" yaml:"action"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#conditions DataLossPreventionContentPolicy#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

