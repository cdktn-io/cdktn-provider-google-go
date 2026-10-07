// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshublistingsubscription


type BigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionRetryPolicy struct {
	// The maximum delay between consecutive deliveries of a given message.
	//
	// Value should be between 0 and 600 seconds. Defaults to 600 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/bigquery_analytics_hub_listing_subscription#maximum_backoff BigqueryAnalyticsHubListingSubscription#maximum_backoff}
	MaximumBackoff *string `field:"optional" json:"maximumBackoff" yaml:"maximumBackoff"`
	// The minimum delay between consecutive deliveries of a given message.
	//
	// Value should be between 0 and 600 seconds. Defaults to 10 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/bigquery_analytics_hub_listing_subscription#minimum_backoff BigqueryAnalyticsHubListingSubscription#minimum_backoff}
	MinimumBackoff *string `field:"optional" json:"minimumBackoff" yaml:"minimumBackoff"`
}

