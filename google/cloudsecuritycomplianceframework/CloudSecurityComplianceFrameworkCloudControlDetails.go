// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudsecuritycomplianceframework


type CloudSecurityComplianceFrameworkCloudControlDetails struct {
	// Major revision of cloudcontrol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_security_compliance_framework#major_revision_id CloudSecurityComplianceFramework#major_revision_id}
	MajorRevisionId *string `field:"required" json:"majorRevisionId" yaml:"majorRevisionId"`
	// The name of the CloudControl in the format: "{parent}/locations/{location}/cloudControls/{cloud-control}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_security_compliance_framework#name CloudSecurityComplianceFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_security_compliance_framework#parameters CloudSecurityComplianceFramework#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

