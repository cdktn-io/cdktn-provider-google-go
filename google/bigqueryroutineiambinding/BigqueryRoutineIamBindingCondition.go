// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryroutineiambinding


type BigqueryRoutineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/bigquery_routine_iam_binding#expression BigqueryRoutineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/bigquery_routine_iam_binding#title BigqueryRoutineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/bigquery_routine_iam_binding#description BigqueryRoutineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

