// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package discoveryenginecontrol


type DiscoveryEngineControlPromoteAction struct {
	// The data store to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/discovery_engine_control#data_store DiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// search_link_promotion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/discovery_engine_control#search_link_promotion DiscoveryEngineControl#search_link_promotion}
	SearchLinkPromotion *DiscoveryEngineControlPromoteActionSearchLinkPromotion `field:"required" json:"searchLinkPromotion" yaml:"searchLinkPromotion"`
}

