// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection struct {
	// include_regexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/data_loss_prevention_discovery_config#include_regexes DataLossPreventionDiscoveryConfig#include_regexes}
	IncludeRegexes *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexes `field:"optional" json:"includeRegexes" yaml:"includeRegexes"`
}

