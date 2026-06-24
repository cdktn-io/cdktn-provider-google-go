// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolGoogleSearchToolPromptConfig struct {
	// Optional.
	//
	// Defines the prompt used for the system instructions when interacting with the
	// agent in chat conversations. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/ces_tool#text_prompt CesTool#text_prompt}
	TextPrompt *string `field:"optional" json:"textPrompt" yaml:"textPrompt"`
	// Optional.
	//
	// Defines the prompt used for the system instructions when interacting with the
	// agent in voice conversations. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/ces_tool#voice_prompt CesTool#voice_prompt}
	VoicePrompt *string `field:"optional" json:"voicePrompt" yaml:"voicePrompt"`
}

