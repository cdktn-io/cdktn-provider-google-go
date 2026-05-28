// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionConfigurationEndpoint struct {
	// Host and port in the format of 'host:port' for the connector endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/bigquery_connection#host_port BigqueryConnection#host_port}
	HostPort *string `field:"optional" json:"hostPort" yaml:"hostPort"`
}

