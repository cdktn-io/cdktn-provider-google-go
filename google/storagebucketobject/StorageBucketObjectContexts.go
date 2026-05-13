// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagebucketobject


type StorageBucketObjectContexts struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/storage_bucket_object#custom StorageBucketObject#custom}
	Custom interface{} `field:"required" json:"custom" yaml:"custom"`
}

