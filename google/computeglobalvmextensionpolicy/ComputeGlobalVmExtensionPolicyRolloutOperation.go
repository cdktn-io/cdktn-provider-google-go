// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeglobalvmextensionpolicy


type ComputeGlobalVmExtensionPolicyRolloutOperation struct {
	// rollout_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_global_vm_extension_policy#rollout_input ComputeGlobalVmExtensionPolicy#rollout_input}
	RolloutInput *ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput `field:"required" json:"rolloutInput" yaml:"rolloutInput"`
}

