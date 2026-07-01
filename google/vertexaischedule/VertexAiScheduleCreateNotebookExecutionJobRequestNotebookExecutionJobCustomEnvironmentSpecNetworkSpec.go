// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule


type VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec struct {
	// Whether to enable public internet access. Default false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/vertex_ai_schedule#enable_internet_access VertexAiSchedule#enable_internet_access}
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/vertex_ai_schedule#network VertexAiSchedule#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnet that this instance is in. Format: 'projects/{project_id_or_number}/regions/{region}/subnetworks/{subnetwork_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/vertex_ai_schedule#subnetwork VertexAiSchedule#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

