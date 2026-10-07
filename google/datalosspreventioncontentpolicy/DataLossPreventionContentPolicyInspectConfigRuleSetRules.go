// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRules struct {
	// adjustment_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#adjustment_rule DataLossPreventionContentPolicy#adjustment_rule}
	AdjustmentRule *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule `field:"optional" json:"adjustmentRule" yaml:"adjustmentRule"`
	// exclusion_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#exclusion_rule DataLossPreventionContentPolicy#exclusion_rule}
	ExclusionRule *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule `field:"optional" json:"exclusionRule" yaml:"exclusionRule"`
	// hotword_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#hotword_rule DataLossPreventionContentPolicy#hotword_rule}
	HotwordRule *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule `field:"optional" json:"hotwordRule" yaml:"hotwordRule"`
}

