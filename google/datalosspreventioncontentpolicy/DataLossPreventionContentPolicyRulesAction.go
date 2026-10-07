// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyRulesAction struct {
	// If set, the verdict will be returned to the user. Possible values: ["ALLOW", "BLOCK"] Possible values: ["ALLOW", "BLOCK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#return_verdict DataLossPreventionContentPolicy#return_verdict}
	ReturnVerdict *string `field:"optional" json:"returnVerdict" yaml:"returnVerdict"`
}

