// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsImageContainmentType struct {
	// encloses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#encloses DataLossPreventionContentPolicy#encloses}
	Encloses *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsImageContainmentTypeEncloses `field:"optional" json:"encloses" yaml:"encloses"`
	// fully_inside block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#fully_inside DataLossPreventionContentPolicy#fully_inside}
	FullyInside *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsImageContainmentTypeFullyInside `field:"optional" json:"fullyInside" yaml:"fullyInside"`
	// overlaps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#overlaps DataLossPreventionContentPolicy#overlaps}
	Overlaps *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsImageContainmentTypeOverlaps `field:"optional" json:"overlaps" yaml:"overlaps"`
}

