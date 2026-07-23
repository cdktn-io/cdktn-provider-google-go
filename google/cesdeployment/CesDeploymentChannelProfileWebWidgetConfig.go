// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesdeployment


type CesDeploymentChannelProfileWebWidgetConfig struct {
	// The modality of the web widget. Possible values: MODALITY_UNSPECIFIED CHAT_AND_VOICE VOICE_ONLY CHAT_ONLY CHAT_VOICE_AND_VIDEO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/ces_deployment#modality CesDeployment#modality}
	Modality *string `field:"optional" json:"modality" yaml:"modality"`
	// security_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/ces_deployment#security_settings CesDeployment#security_settings}
	SecuritySettings *CesDeploymentChannelProfileWebWidgetConfigSecuritySettings `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// The theme of the web widget. Possible values: THEME_UNSPECIFIED LIGHT DARK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/ces_deployment#theme CesDeployment#theme}
	Theme *string `field:"optional" json:"theme" yaml:"theme"`
	// The title of the web widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/ces_deployment#web_widget_title CesDeployment#web_widget_title}
	WebWidgetTitle *string `field:"optional" json:"webWidgetTitle" yaml:"webWidgetTitle"`
}

