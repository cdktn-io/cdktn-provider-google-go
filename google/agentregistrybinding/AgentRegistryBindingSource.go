// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistrybinding


type AgentRegistryBindingSource struct {
	// The identifier of the source Agent. Format: 'urn:agent:{publisher}:{namespace}:{name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/agent_registry_binding#identifier AgentRegistryBinding#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

