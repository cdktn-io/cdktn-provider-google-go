// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageftpserver


type StorageFtpServerExternalConfig struct {
	// A list of allowed IPv4 or IPv6 CIDR block ranges that can connect to this server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/storage_ftp_server#allowed_cidr_blocks StorageFtpServer#allowed_cidr_blocks}
	AllowedCidrBlocks *[]*string `field:"optional" json:"allowedCidrBlocks" yaml:"allowedCidrBlocks"`
}

