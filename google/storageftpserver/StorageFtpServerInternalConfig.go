// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageftpserver


type StorageFtpServerInternalConfig struct {
	// consumer_accept_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/storage_ftp_server#consumer_accept_list StorageFtpServer#consumer_accept_list}
	ConsumerAcceptList interface{} `field:"optional" json:"consumerAcceptList" yaml:"consumerAcceptList"`
	// consumer_reject_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/storage_ftp_server#consumer_reject_list StorageFtpServer#consumer_reject_list}
	ConsumerRejectList interface{} `field:"optional" json:"consumerRejectList" yaml:"consumerRejectList"`
}

