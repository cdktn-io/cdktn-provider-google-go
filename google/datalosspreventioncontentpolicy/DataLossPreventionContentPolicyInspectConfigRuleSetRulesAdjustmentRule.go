// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule struct {
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#likelihood_adjustment DataLossPreventionContentPolicy#likelihood_adjustment}
	LikelihoodAdjustment *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment `field:"required" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// adjust_by_image_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#adjust_by_image_findings DataLossPreventionContentPolicy#adjust_by_image_findings}
	AdjustByImageFindings *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings `field:"optional" json:"adjustByImageFindings" yaml:"adjustByImageFindings"`
	// adjust_by_matching_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#adjust_by_matching_info_types DataLossPreventionContentPolicy#adjust_by_matching_info_types}
	AdjustByMatchingInfoTypes *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes `field:"optional" json:"adjustByMatchingInfoTypes" yaml:"adjustByMatchingInfoTypes"`
}

