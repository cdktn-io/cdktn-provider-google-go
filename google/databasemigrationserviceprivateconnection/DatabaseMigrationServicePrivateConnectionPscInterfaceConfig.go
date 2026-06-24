// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceprivateconnection


type DatabaseMigrationServicePrivateConnectionPscInterfaceConfig struct {
	// Fully qualified name of the Network Attachment that DMS will connect to. Format: projects/{project}/regions/{region}/networkAttachments/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/database_migration_service_private_connection#network_attachment DatabaseMigrationServicePrivateConnection#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

