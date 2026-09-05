// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamconnectionprofile


type DatastreamConnectionProfilePostgresqlProfileSslConfigServerVerification struct {
	// PEM-encoded server root CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/datastream_connection_profile#ca_certificate DatastreamConnectionProfile#ca_certificate}
	CaCertificate *string `field:"required" json:"caCertificate" yaml:"caCertificate"`
}

