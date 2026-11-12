// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentType struct {
	// encloses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#encloses DataLossPreventionInspectTemplate#encloses}
	Encloses *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentTypeEncloses `field:"optional" json:"encloses" yaml:"encloses"`
	// fully_inside block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#fully_inside DataLossPreventionInspectTemplate#fully_inside}
	FullyInside *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentTypeFullyInside `field:"optional" json:"fullyInside" yaml:"fullyInside"`
	// overlaps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#overlaps DataLossPreventionInspectTemplate#overlaps}
	Overlaps *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentTypeOverlaps `field:"optional" json:"overlaps" yaml:"overlaps"`
}

