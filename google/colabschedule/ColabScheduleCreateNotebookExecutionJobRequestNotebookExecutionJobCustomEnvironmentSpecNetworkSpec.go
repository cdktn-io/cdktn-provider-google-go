// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec struct {
	// Whether to enable public internet access. Default false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/colab_schedule#enable_internet_access ColabSchedule#enable_internet_access}
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/colab_schedule#network ColabSchedule#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnet that this instance is in. Format: 'projects/{project_id_or_number}/regions/{region}/subnetworks/{subnetwork_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/colab_schedule#subnetwork ColabSchedule#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

