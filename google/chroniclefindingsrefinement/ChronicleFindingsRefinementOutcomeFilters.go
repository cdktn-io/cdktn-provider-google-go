// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefindingsrefinement


type ChronicleFindingsRefinementOutcomeFilters struct {
	// The operator to be applied to the outcome variable. Possible values: EQUAL CONTAINS MATCHES_REGEX MATCHES_CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_findings_refinement#outcome_filter_operator ChronicleFindingsRefinement#outcome_filter_operator}
	OutcomeFilterOperator *string `field:"required" json:"outcomeFilterOperator" yaml:"outcomeFilterOperator"`
	// The value of the outcome variable to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_findings_refinement#outcome_value ChronicleFindingsRefinement#outcome_value}
	OutcomeValue *string `field:"required" json:"outcomeValue" yaml:"outcomeValue"`
	// The outcome variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_findings_refinement#outcome_variable ChronicleFindingsRefinement#outcome_variable}
	OutcomeVariable *string `field:"required" json:"outcomeVariable" yaml:"outcomeVariable"`
}

