// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex struct {
	// aws_account_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/data_loss_prevention_discovery_config#aws_account_regex DataLossPreventionDiscoveryConfig#aws_account_regex}
	AwsAccountRegex *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex `field:"optional" json:"awsAccountRegex" yaml:"awsAccountRegex"`
	// Regex to test the bucket name against. If empty, all buckets match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/data_loss_prevention_discovery_config#bucket_name_regex DataLossPreventionDiscoveryConfig#bucket_name_regex}
	BucketNameRegex *string `field:"optional" json:"bucketNameRegex" yaml:"bucketNameRegex"`
}

