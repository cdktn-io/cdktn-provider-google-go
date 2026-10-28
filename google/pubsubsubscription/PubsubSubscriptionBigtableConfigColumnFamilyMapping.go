// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription


type PubsubSubscriptionBigtableConfigColumnFamilyMapping struct {
	// delimited_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/pubsub_subscription#delimited_key PubsubSubscription#delimited_key}
	DelimitedKey *PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey `field:"optional" json:"delimitedKey" yaml:"delimitedKey"`
	// row_key_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/pubsub_subscription#row_key_schema PubsubSubscription#row_key_schema}
	RowKeySchema *PubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema `field:"optional" json:"rowKeySchema" yaml:"rowKeySchema"`
}

