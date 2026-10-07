// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig


type BigqueryDataTransferConfigSensitiveParams struct {
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/bigquery_data_transfer_config#secret_access_key BigqueryDataTransferConfig#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/bigquery_data_transfer_config#secret_access_key_wo BigqueryDataTransferConfig#secret_access_key_wo}
	SecretAccessKeyWo *string `field:"optional" json:"secretAccessKeyWo" yaml:"secretAccessKeyWo"`
	// Triggers update of 'secret_access_key_wo' write-only.
	//
	// Increment this value when an update to 'secret_access_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/bigquery_data_transfer_config#secret_access_key_wo_version BigqueryDataTransferConfig#secret_access_key_wo_version}
	SecretAccessKeyWoVersion *string `field:"optional" json:"secretAccessKeyWoVersion" yaml:"secretAccessKeyWoVersion"`
}

