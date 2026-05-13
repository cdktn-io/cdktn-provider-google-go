// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamRuleSetsCustomizationRulesBigqueryClustering struct {
	// Column names to set as clustering columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_stream#columns DatastreamStream#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
}

