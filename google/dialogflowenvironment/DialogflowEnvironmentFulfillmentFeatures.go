// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowenvironment


type DialogflowEnvironmentFulfillmentFeatures struct {
	// The type of the feature that enabled for fulfillment. Possible values: ["TYPE_UNSPECIFIED", "SMALLTALK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/dialogflow_environment#type DialogflowEnvironment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

