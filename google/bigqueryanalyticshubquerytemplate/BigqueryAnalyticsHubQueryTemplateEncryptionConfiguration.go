// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshubquerytemplate


type BigqueryAnalyticsHubQueryTemplateEncryptionConfiguration struct {
	// The KMS key used to encrypt the query template. Format: 'projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/bigquery_analytics_hub_query_template#kms_key_name BigqueryAnalyticsHubQueryTemplate#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

