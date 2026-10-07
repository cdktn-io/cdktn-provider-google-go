// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageftpserver


type StorageFtpServerInternalConfigConsumerAcceptListStruct struct {
	// The maximum number of Private Service Connect endpoints that can be created in the consumer project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/storage_ftp_server#connection_limit StorageFtpServer#connection_limit}
	ConnectionLimit *float64 `field:"required" json:"connectionLimit" yaml:"connectionLimit"`
	// The project that is allowed to connect, in the format 'projects/{project}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/storage_ftp_server#project StorageFtpServer#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

