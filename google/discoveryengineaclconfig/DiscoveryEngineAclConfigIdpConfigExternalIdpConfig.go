// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package discoveryengineaclconfig


type DiscoveryEngineAclConfigIdpConfigExternalIdpConfig struct {
	// Workforce pool name: "locations/global/workforcePools/pool_id".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/discovery_engine_acl_config#workforce_pool_name DiscoveryEngineAclConfig#workforce_pool_name}
	WorkforcePoolName *string `field:"optional" json:"workforcePoolName" yaml:"workforcePoolName"`
}

