// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigDeadLetterConfig struct {
	// Project which has the topic given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#project_id IntegrationConnectorsConnection#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Topic to push events which couldn't be processed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#topic IntegrationConnectorsConnection#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

