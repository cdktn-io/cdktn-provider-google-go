// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance


type FilestoreInstanceDirectoryServices struct {
	// ldap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.33.0/docs/resources/filestore_instance#ldap FilestoreInstance#ldap}
	Ldap *FilestoreInstanceDirectoryServicesLdap `field:"optional" json:"ldap" yaml:"ldap"`
}

