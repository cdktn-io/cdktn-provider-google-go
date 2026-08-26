// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusVpcAccessibleServices struct {
	// allowed_service_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/access_context_manager_service_perimeter#allowed_service_patterns AccessContextManagerServicePerimeter#allowed_service_patterns}
	AllowedServicePatterns interface{} `field:"optional" json:"allowedServicePatterns" yaml:"allowedServicePatterns"`
	// The list of APIs usable within the Service Perimeter. Must be empty unless 'enableRestriction' is True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/access_context_manager_service_perimeter#allowed_services AccessContextManagerServicePerimeter#allowed_services}
	AllowedServices *[]*string `field:"optional" json:"allowedServices" yaml:"allowedServices"`
	// Whether to restrict API calls within the Service Perimeter to the list of APIs specified in 'allowedServices'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/access_context_manager_service_perimeter#enable_restriction AccessContextManagerServicePerimeter#enable_restriction}
	EnableRestriction interface{} `field:"optional" json:"enableRestriction" yaml:"enableRestriction"`
	// Defines the enforcement scopes of service patterns. Possible values: ["GOOGLE_APIS_VIA_PRIVATE_PATH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/access_context_manager_service_perimeter#service_patterns_enforcement_scopes AccessContextManagerServicePerimeter#service_patterns_enforcement_scopes}
	ServicePatternsEnforcementScopes *[]*string `field:"optional" json:"servicePatternsEnforcementScopes" yaml:"servicePatternsEnforcementScopes"`
}

