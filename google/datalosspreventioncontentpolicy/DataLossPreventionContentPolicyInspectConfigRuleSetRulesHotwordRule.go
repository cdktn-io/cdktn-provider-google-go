// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#hotword_regex DataLossPreventionContentPolicy#hotword_regex}
	HotwordRegex *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRuleHotwordRegex `field:"required" json:"hotwordRegex" yaml:"hotwordRegex"`
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#likelihood_adjustment DataLossPreventionContentPolicy#likelihood_adjustment}
	LikelihoodAdjustment *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `field:"required" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#proximity DataLossPreventionContentPolicy#proximity}
	Proximity *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRuleProximity `field:"required" json:"proximity" yaml:"proximity"`
}

