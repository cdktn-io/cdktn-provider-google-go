// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsInfoTypes struct {
	// Name of the information type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#name DataLossPreventionInspectTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#sensitivity_score DataLossPreventionInspectTemplate#sensitivity_score}
	SensitivityScore *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#version DataLossPreventionInspectTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

