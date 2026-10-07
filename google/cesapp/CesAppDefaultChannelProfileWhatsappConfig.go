// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppDefaultChannelProfileWhatsappConfig struct {
	// The Meta phone number ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#phone_number_id CesApp#phone_number_id}
	PhoneNumberId *string `field:"required" json:"phoneNumberId" yaml:"phoneNumberId"`
	// The WhatsApp Business Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#waba_id CesApp#waba_id}
	WabaId *string `field:"required" json:"wabaId" yaml:"wabaId"`
	// The phone number in E.164 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#phone_number CesApp#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

