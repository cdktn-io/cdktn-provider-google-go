// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource struct {
	// amazon_s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/data_loss_prevention_discovery_config#amazon_s3_bucket DataLossPreventionDiscoveryConfig#amazon_s3_bucket}
	AmazonS3Bucket *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket `field:"optional" json:"amazonS3Bucket" yaml:"amazonS3Bucket"`
}

