// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetConnectorToolsetConnectorActionsEntityOperation struct {
	// ID of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_toolset#entity_id CesToolset#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Operation to perform on the entity. Possible values: LIST GET CREATE UPDATE DELETE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_toolset#operation CesToolset#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
}

