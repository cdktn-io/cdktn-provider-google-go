// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentTransferRules struct {
	// The resource name of the child agent the rule applies to. Format: 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#child_agent CesAgent#child_agent}
	ChildAgent *string `field:"required" json:"childAgent" yaml:"childAgent"`
	// The direction of the transfer. Possible values: ["PARENT_TO_CHILD", "CHILD_TO_PARENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#direction CesAgent#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// deterministic_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#deterministic_transfer CesAgent#deterministic_transfer}
	DeterministicTransfer *CesAgentTransferRulesDeterministicTransfer `field:"optional" json:"deterministicTransfer" yaml:"deterministicTransfer"`
	// disable_planner_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#disable_planner_transfer CesAgent#disable_planner_transfer}
	DisablePlannerTransfer *CesAgentTransferRulesDisablePlannerTransfer `field:"optional" json:"disablePlannerTransfer" yaml:"disablePlannerTransfer"`
}

