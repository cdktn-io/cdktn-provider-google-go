// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandling struct {
	// leave_untransformed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_deidentify_template#leave_untransformed DataLossPreventionDeidentifyTemplate#leave_untransformed}
	LeaveUntransformed *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed `field:"optional" json:"leaveUntransformed" yaml:"leaveUntransformed"`
	// throw_error block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_deidentify_template#throw_error DataLossPreventionDeidentifyTemplate#throw_error}
	ThrowError *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError `field:"optional" json:"throwError" yaml:"throwError"`
}

