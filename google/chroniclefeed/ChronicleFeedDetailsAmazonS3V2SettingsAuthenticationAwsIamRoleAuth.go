// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth struct {
	// AWS IAM Role for Identity Federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#aws_iam_role_arn ChronicleFeed#aws_iam_role_arn}
	AwsIamRoleArn *string `field:"optional" json:"awsIamRoleArn" yaml:"awsIamRoleArn"`
	// Subject ID to use for S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#subject_id ChronicleFeed#subject_id}
	SubjectId *string `field:"optional" json:"subjectId" yaml:"subjectId"`
}

