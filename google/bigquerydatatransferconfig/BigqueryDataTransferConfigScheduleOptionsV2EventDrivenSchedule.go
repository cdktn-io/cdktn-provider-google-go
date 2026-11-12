// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig


type BigqueryDataTransferConfigScheduleOptionsV2EventDrivenSchedule struct {
	// Pub/Sub subscription name used to receive events. Only the Google Cloud Storage data source supports this option. Format: projects/{project}/subscriptions/{subscription}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/bigquery_data_transfer_config#pubsub_subscription BigqueryDataTransferConfig#pubsub_subscription}
	PubsubSubscription *string `field:"required" json:"pubsubSubscription" yaml:"pubsubSubscription"`
}

