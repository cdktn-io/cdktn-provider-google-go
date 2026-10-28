// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataflowflextemplatejob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataflowFlexTemplateJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#container_spec_gcs_path DataflowFlexTemplateJob#container_spec_gcs_path}.
	ContainerSpecGcsPath *string `field:"required" json:"containerSpecGcsPath" yaml:"containerSpecGcsPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#name DataflowFlexTemplateJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#additional_experiments DataflowFlexTemplateJob#additional_experiments}
	AdditionalExperiments *[]*string `field:"optional" json:"additionalExperiments" yaml:"additionalExperiments"`
	// List of pipeline options that should be used by the job. An example value is ["numberOfWorkerHarnessThreads=20"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#additional_pipeline_options DataflowFlexTemplateJob#additional_pipeline_options}
	AdditionalPipelineOptions *[]*string `field:"optional" json:"additionalPipelineOptions" yaml:"additionalPipelineOptions"`
	// The algorithm to use for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#autoscaling_algorithm DataflowFlexTemplateJob#autoscaling_algorithm}
	AutoscalingAlgorithm *string `field:"optional" json:"autoscalingAlgorithm" yaml:"autoscalingAlgorithm"`
	// If true, if a 409 AlreadyExists error is returned on create, the provider will ignore it and adopt the existing resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#create_ignore_already_exists DataflowFlexTemplateJob#create_ignore_already_exists}
	CreateIgnoreAlreadyExists interface{} `field:"optional" json:"createIgnoreAlreadyExists" yaml:"createIgnoreAlreadyExists"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#deletion_policy DataflowFlexTemplateJob#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Indicates if the job should use the streaming engine feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#enable_streaming_engine DataflowFlexTemplateJob#enable_streaming_engine}
	EnableStreamingEngine interface{} `field:"optional" json:"enableStreamingEngine" yaml:"enableStreamingEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#id DataflowFlexTemplateJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#ip_configuration DataflowFlexTemplateJob#ip_configuration}
	IpConfiguration *string `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#kms_key_name DataflowFlexTemplateJob#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// User labels to be specified for the job.
	//
	// Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#labels DataflowFlexTemplateJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The machine type to use for launching the job. The default is n1-standard-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#launcher_machine_type DataflowFlexTemplateJob#launcher_machine_type}
	LauncherMachineType *string `field:"optional" json:"launcherMachineType" yaml:"launcherMachineType"`
	// The machine type to use for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#machine_type DataflowFlexTemplateJob#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The maximum number of Google Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#max_workers DataflowFlexTemplateJob#max_workers}
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#network DataflowFlexTemplateJob#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The initial number of Google Compute Engine instances for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#num_workers DataflowFlexTemplateJob#num_workers}
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#on_delete DataflowFlexTemplateJob#on_delete}.
	OnDelete *string `field:"optional" json:"onDelete" yaml:"onDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#parameters DataflowFlexTemplateJob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#project DataflowFlexTemplateJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the created job should run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#region DataflowFlexTemplateJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docker registry location of container image to use for the 'worker harness.
	//
	// Default is the container for the version of the SDK. Note this field is only valid for portable pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#sdk_container_image DataflowFlexTemplateJob#sdk_container_image}
	SdkContainerImage *string `field:"optional" json:"sdkContainerImage" yaml:"sdkContainerImage"`
	// The Service Account email used to create the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#service_account_email DataflowFlexTemplateJob#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from terraform state and moving on.
	//
	// WARNING: this will lead to job name conflicts if you do not ensure that the job names are different, e.g. by embedding a release ID or by using a random_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#skip_wait_on_job_termination DataflowFlexTemplateJob#skip_wait_on_job_termination}
	SkipWaitOnJobTermination interface{} `field:"optional" json:"skipWaitOnJobTermination" yaml:"skipWaitOnJobTermination"`
	// The Cloud Storage path to use for staging files. Must be a valid Cloud Storage URL, beginning with gs://.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#staging_location DataflowFlexTemplateJob#staging_location}
	StagingLocation *string `field:"optional" json:"stagingLocation" yaml:"stagingLocation"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#subnetwork DataflowFlexTemplateJob#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#temp_location DataflowFlexTemplateJob#temp_location}
	TempLocation *string `field:"optional" json:"tempLocation" yaml:"tempLocation"`
	// Only applicable when updating a pipeline.
	//
	// Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job#transform_name_mapping DataflowFlexTemplateJob#transform_name_mapping}
	TransformNameMapping *map[string]*string `field:"optional" json:"transformNameMapping" yaml:"transformNameMapping"`
}

