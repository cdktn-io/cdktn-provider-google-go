// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionConfigurationNetwork struct {
	// private_service_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/bigquery_connection#private_service_connect BigqueryConnection#private_service_connect}
	PrivateServiceConnect *BigqueryConnectionConfigurationNetworkPrivateServiceConnect `field:"optional" json:"privateServiceConnect" yaml:"privateServiceConnect"`
}

