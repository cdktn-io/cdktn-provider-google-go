// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription


type PubsubSubscriptionMessageTransformsCompression struct {
	// Specifies the compression algorithm to use. Possible values include 'ZLIB'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/pubsub_subscription#compression_algorithm PubsubSubscription#compression_algorithm}
	CompressionAlgorithm *string `field:"required" json:"compressionAlgorithm" yaml:"compressionAlgorithm"`
	// Specifies whether to compress or decompress the message. Possible values: ["COMPRESS", "DECOMPRESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/pubsub_subscription#compression_mode PubsubSubscription#compression_mode}
	CompressionMode *string `field:"required" json:"compressionMode" yaml:"compressionMode"`
}

