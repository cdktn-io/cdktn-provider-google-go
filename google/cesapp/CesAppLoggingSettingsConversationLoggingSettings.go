// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppLoggingSettingsConversationLoggingSettings struct {
	// Whether to disable conversation logging for the sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#disable_conversation_logging CesApp#disable_conversation_logging}
	DisableConversationLogging interface{} `field:"optional" json:"disableConversationLogging" yaml:"disableConversationLogging"`
	// Controls the retention window for the conversation. If not set, the conversation will be retained for 365 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#retention_window CesApp#retention_window}
	RetentionWindow *string `field:"optional" json:"retentionWindow" yaml:"retentionWindow"`
}

