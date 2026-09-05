// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdataproduct


type DataplexDataProductAccessApprovalConfig struct {
	// Specifies the email addresses of users who are potential approvers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/dataplex_data_product#approver_emails DataplexDataProduct#approver_emails}
	ApproverEmails *[]*string `field:"optional" json:"approverEmails" yaml:"approverEmails"`
}

