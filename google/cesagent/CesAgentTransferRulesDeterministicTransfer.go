// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentTransferRulesDeterministicTransfer struct {
	// expression_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#expression_condition CesAgent#expression_condition}
	ExpressionCondition *CesAgentTransferRulesDeterministicTransferExpressionCondition `field:"optional" json:"expressionCondition" yaml:"expressionCondition"`
	// python_code_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#python_code_condition CesAgent#python_code_condition}
	PythonCodeCondition *CesAgentTransferRulesDeterministicTransferPythonCodeCondition `field:"optional" json:"pythonCodeCondition" yaml:"pythonCodeCondition"`
}

