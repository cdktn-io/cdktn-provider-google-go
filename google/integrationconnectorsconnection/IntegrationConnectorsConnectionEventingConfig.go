// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfig struct {
	// registration_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#registration_destination_config IntegrationConnectorsConnection#registration_destination_config}
	RegistrationDestinationConfig *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig `field:"required" json:"registrationDestinationConfig" yaml:"registrationDestinationConfig"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#additional_variable IntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// List of allowed event types for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#allowed_event_types IntegrationConnectorsConnection#allowed_event_types}
	AllowedEventTypes *[]*string `field:"optional" json:"allowedEventTypes" yaml:"allowedEventTypes"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#auth_config IntegrationConnectorsConnection#auth_config}
	AuthConfig *IntegrationConnectorsConnectionEventingConfigAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#dead_letter_config IntegrationConnectorsConnection#dead_letter_config}
	DeadLetterConfig *IntegrationConnectorsConnectionEventingConfigDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// enrichment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#enrichment_config IntegrationConnectorsConnection#enrichment_config}
	EnrichmentConfig *IntegrationConnectorsConnectionEventingConfigEnrichmentConfig `field:"optional" json:"enrichmentConfig" yaml:"enrichmentConfig"`
	// Enrichment Enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#enrichment_enabled IntegrationConnectorsConnection#enrichment_enabled}
	EnrichmentEnabled interface{} `field:"optional" json:"enrichmentEnabled" yaml:"enrichmentEnabled"`
	// Filter to be applied on the events to be received by the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#global_event_filter IntegrationConnectorsConnection#global_event_filter}
	GlobalEventFilter *string `field:"optional" json:"globalEventFilter" yaml:"globalEventFilter"`
}

