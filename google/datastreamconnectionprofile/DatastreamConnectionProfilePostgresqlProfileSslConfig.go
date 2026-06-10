// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamconnectionprofile


type DatastreamConnectionProfilePostgresqlProfileSslConfig struct {
	// server_and_client_verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/datastream_connection_profile#server_and_client_verification DatastreamConnectionProfile#server_and_client_verification}
	ServerAndClientVerification *DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification `field:"optional" json:"serverAndClientVerification" yaml:"serverAndClientVerification"`
	// server_verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/datastream_connection_profile#server_verification DatastreamConnectionProfile#server_verification}
	ServerVerification *DatastreamConnectionProfilePostgresqlProfileSslConfigServerVerification `field:"optional" json:"serverVerification" yaml:"serverVerification"`
}

