// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshublistingsubscription


type BigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscription struct {
	// pubsub_subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/bigquery_analytics_hub_listing_subscription#pubsub_subscription BigqueryAnalyticsHubListingSubscription#pubsub_subscription}
	PubsubSubscription *BigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscription `field:"required" json:"pubsubSubscription" yaml:"pubsubSubscription"`
}

