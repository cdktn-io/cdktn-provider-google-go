// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventarctrigger


type EventarcTriggerRetryPolicy struct {
	// The maximum number of delivery attempts for any message. The only valid value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/eventarc_trigger#max_attempts EventarcTrigger#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
}

