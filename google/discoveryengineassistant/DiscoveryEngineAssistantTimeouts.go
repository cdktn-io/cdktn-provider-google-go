// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package discoveryengineassistant


type DiscoveryEngineAssistantTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/discovery_engine_assistant#create DiscoveryEngineAssistant#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/discovery_engine_assistant#delete DiscoveryEngineAssistant#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/discovery_engine_assistant#update DiscoveryEngineAssistant#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

