// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsModifiersAddRequestHeader struct {
	// HTTP header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/access_context_manager_service_perimeters#key AccessContextManagerServicePerimeters#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// HTTP header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/access_context_manager_service_perimeters#value AccessContextManagerServicePerimeters#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

