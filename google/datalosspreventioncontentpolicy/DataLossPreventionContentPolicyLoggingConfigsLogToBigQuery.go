// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyLoggingConfigsLogToBigQuery struct {
	// The dataset ID of the BigQuery table to log to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#dataset_id DataLossPreventionContentPolicy#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The project ID of the BigQuery table to log to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#project_id DataLossPreventionContentPolicy#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The table ID of the BigQuery table to log to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#table_id DataLossPreventionContentPolicy#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

