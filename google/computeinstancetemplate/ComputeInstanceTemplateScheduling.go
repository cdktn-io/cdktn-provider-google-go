// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateScheduling struct {
	// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user).
	//
	// This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#automatic_restart ComputeInstanceTemplate#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the availability domain, which this instance should be scheduled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#availability_domain ComputeInstanceTemplate#availability_domain}
	AvailabilityDomain *float64 `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Specify the time in seconds for host error detection, the value must be within the range of [90, 330] with the increment of 30, if unset, the default behavior of host error recovery will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#host_error_timeout_seconds ComputeInstanceTemplate#host_error_timeout_seconds}
	HostErrorTimeoutSeconds *float64 `field:"optional" json:"hostErrorTimeoutSeconds" yaml:"hostErrorTimeoutSeconds"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#instance_termination_action ComputeInstanceTemplate#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#local_ssd_recovery_timeout ComputeInstanceTemplate#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout interface{} `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#max_run_duration ComputeInstanceTemplate#max_run_duration}
	MaxRunDuration *ComputeInstanceTemplateSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Minimum number of cpus for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#min_node_cpus ComputeInstanceTemplate#min_node_cpus}
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#node_affinities ComputeInstanceTemplate#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Defines the maintenance behavior for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#on_host_maintenance ComputeInstanceTemplate#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#on_instance_stop_action ComputeInstanceTemplate#on_instance_stop_action}
	OnInstanceStopAction *ComputeInstanceTemplateSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Allows instance to be preempted. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#preemptible ComputeInstanceTemplate#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Describes the desired provisioning model for the instance.
	//
	// Possible values are STANDARD, SPOT, FLEX_START, and RESERVATION_BOUND. For STANDARD, resources are provisioned immediately. For SPOT, resources are offered at a discount compared to standard pricing but may be preempted. For FLEX_START, resources are offered at a discount with flexible start times. For RESERVATION_BOUND, the instance is bound to a specific reservation and will only consume capacity from that reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#provisioning_model ComputeInstanceTemplate#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format.
	//
	// If specified, the instance termination action
	// will be performed at the termination time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_instance_template#termination_time ComputeInstanceTemplate#termination_time}
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

