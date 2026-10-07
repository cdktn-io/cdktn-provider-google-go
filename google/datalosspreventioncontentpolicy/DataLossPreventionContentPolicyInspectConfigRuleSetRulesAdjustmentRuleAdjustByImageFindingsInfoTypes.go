// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsInfoTypes struct {
	// Name of the information type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#name DataLossPreventionContentPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#sensitivity_score DataLossPreventionContentPolicy#sensitivity_score}
	SensitivityScore *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#version DataLossPreventionContentPolicy#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

