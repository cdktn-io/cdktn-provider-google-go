// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjectsDatabasesCollections struct {
	// Collection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/datastream_stream#collection DatastreamStream#collection}
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/datastream_stream#fields DatastreamStream#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

