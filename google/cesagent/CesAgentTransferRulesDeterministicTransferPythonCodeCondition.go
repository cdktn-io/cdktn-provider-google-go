// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentTransferRulesDeterministicTransferPythonCodeCondition struct {
	// The python code to execute. The function must be named 'should_trigger_transfer_callback'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#python_code CesAgent#python_code}
	PythonCode *string `field:"required" json:"pythonCode" yaml:"pythonCode"`
}

