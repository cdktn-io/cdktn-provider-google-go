// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package clouddeploycustomtargettype


type ClouddeployCustomTargetTypeTasks struct {
	// deploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/clouddeploy_custom_target_type#deploy ClouddeployCustomTargetType#deploy}
	Deploy *ClouddeployCustomTargetTypeTasksDeploy `field:"required" json:"deploy" yaml:"deploy"`
	// render block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/clouddeploy_custom_target_type#render ClouddeployCustomTargetType#render}
	Render *ClouddeployCustomTargetTypeTasksRender `field:"optional" json:"render" yaml:"render"`
}

