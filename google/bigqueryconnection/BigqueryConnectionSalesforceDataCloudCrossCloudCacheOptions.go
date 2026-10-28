// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionSalesforceDataCloudCrossCloudCacheOptions struct {
	// Whether cross-cloud caching is enabled for queries through BigQuery.
	//
	// When enabled, read data and metadata are cached, which can improve performance
	// and reduce network egress costs for cross-cloud queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/bigquery_connection#enabled BigqueryConnection#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

