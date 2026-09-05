// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects struct {
	// databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/datastream_stream#databases DatastreamStream#databases}
	Databases interface{} `field:"optional" json:"databases" yaml:"databases"`
}

