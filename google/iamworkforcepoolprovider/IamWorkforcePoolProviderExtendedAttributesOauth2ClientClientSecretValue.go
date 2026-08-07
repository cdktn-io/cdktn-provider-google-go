// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider


type IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretValue struct {
	// The plain text of the client secret value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/iam_workforce_pool_provider#plain_text IamWorkforcePoolProvider#plain_text}
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// The plain text of the client secret value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/iam_workforce_pool_provider#plain_text_wo IamWorkforcePoolProvider#plain_text_wo}
	PlainTextWo *string `field:"optional" json:"plainTextWo" yaml:"plainTextWo"`
	// Triggers update of 'plain_text_wo' write-only.
	//
	// Increment this value when an update to 'plain_text_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/iam_workforce_pool_provider#plain_text_wo_version IamWorkforcePoolProvider#plain_text_wo_version}
	PlainTextWoVersion *string `field:"optional" json:"plainTextWoVersion" yaml:"plainTextWoVersion"`
}

