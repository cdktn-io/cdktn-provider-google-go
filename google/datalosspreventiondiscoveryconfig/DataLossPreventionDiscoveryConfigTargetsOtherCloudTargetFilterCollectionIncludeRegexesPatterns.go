// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatterns struct {
	// amazon_s3_bucket_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/data_loss_prevention_discovery_config#amazon_s3_bucket_regex DataLossPreventionDiscoveryConfig#amazon_s3_bucket_regex}
	AmazonS3BucketRegex *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex `field:"optional" json:"amazonS3BucketRegex" yaml:"amazonS3BucketRegex"`
}

