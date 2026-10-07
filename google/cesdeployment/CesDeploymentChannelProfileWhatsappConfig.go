// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesdeployment


type CesDeploymentChannelProfileWhatsappConfig struct {
	// Required. The Meta phone number ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_deployment#phone_number_id CesDeployment#phone_number_id}
	PhoneNumberId *string `field:"required" json:"phoneNumberId" yaml:"phoneNumberId"`
	// Required. The WhatsApp Business Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_deployment#waba_id CesDeployment#waba_id}
	WabaId *string `field:"required" json:"wabaId" yaml:"wabaId"`
	// Optional. The phone number in E.164 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_deployment#phone_number CesDeployment#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

