// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gkehubrolloutsequence

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GkeHubRolloutSequenceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The user-provided identifier of the RolloutSequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#rollout_sequence_id GkeHubRolloutSequence#rollout_sequence_id}
	RolloutSequenceId *string `field:"required" json:"rolloutSequenceId" yaml:"rolloutSequenceId"`
	// stages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#stages GkeHubRolloutSequence#stages}
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// auto_upgrade_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#auto_upgrade_config GkeHubRolloutSequence#auto_upgrade_config}
	AutoUpgradeConfig *GkeHubRolloutSequenceAutoUpgradeConfig `field:"optional" json:"autoUpgradeConfig" yaml:"autoUpgradeConfig"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#deletion_policy GkeHubRolloutSequence#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Human readable display name of the Rollout Sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#display_name GkeHubRolloutSequence#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#id GkeHubRolloutSequence#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ignored_clusters_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#ignored_clusters_selector GkeHubRolloutSequence#ignored_clusters_selector}
	IgnoredClustersSelector *GkeHubRolloutSequenceIgnoredClustersSelector `field:"optional" json:"ignoredClustersSelector" yaml:"ignoredClustersSelector"`
	// Labels for this Rollout Sequence.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#labels GkeHubRolloutSequence#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#project GkeHubRolloutSequence#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gke_hub_rollout_sequence#timeouts GkeHubRolloutSequence#timeouts}
	Timeouts *GkeHubRolloutSequenceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

