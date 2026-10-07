// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentTransferRulesDisablePlannerTransferExpressionCondition struct {
	// The string representation of cloud.api.Expression condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#expression CesAgent#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

