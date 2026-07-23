// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerinstance


type SecureSourceManagerInstancePrivateConfig struct {
	// 'Indicate if it's private instance.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/secure_source_manager_instance#is_private SecureSourceManagerInstance#is_private}
	IsPrivate interface{} `field:"required" json:"isPrivate" yaml:"isPrivate"`
	// CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/secure_source_manager_instance#ca_pool SecureSourceManagerInstance#ca_pool}
	CaPool *string `field:"optional" json:"caPool" yaml:"caPool"`
	// custom_host_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/secure_source_manager_instance#custom_host_config SecureSourceManagerInstance#custom_host_config}
	CustomHostConfig *SecureSourceManagerInstancePrivateConfigCustomHostConfig `field:"optional" json:"customHostConfig" yaml:"customHostConfig"`
	// Optional.
	//
	// Additional allowed projects for setting up PSC connections.
	// Instance host project is automatically allowed and does not need to be included in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/secure_source_manager_instance#psc_allowed_projects SecureSourceManagerInstance#psc_allowed_projects}
	PscAllowedProjects *[]*string `field:"optional" json:"pscAllowedProjects" yaml:"pscAllowedProjects"`
}

