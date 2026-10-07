// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentTransferRulesDisablePlannerTransfer struct {
	// expression_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#expression_condition CesAgent#expression_condition}
	ExpressionCondition *CesAgentTransferRulesDisablePlannerTransferExpressionCondition `field:"required" json:"expressionCondition" yaml:"expressionCondition"`
}

