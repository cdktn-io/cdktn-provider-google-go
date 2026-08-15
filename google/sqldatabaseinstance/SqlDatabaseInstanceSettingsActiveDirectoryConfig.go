// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceSettingsActiveDirectoryConfig struct {
	// Domain name of the Active Directory for SQL Server (e.g., mydomain.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/sql_database_instance#domain SqlDatabaseInstance#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The secret manager key storing the administrator credential. (e.g., projects/{project}/secrets/{secret}).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/sql_database_instance#admin_credential_secret_name SqlDatabaseInstance#admin_credential_secret_name}
	AdminCredentialSecretName *string `field:"optional" json:"adminCredentialSecretName" yaml:"adminCredentialSecretName"`
	// Domain controller IPv4 addresses used to bootstrap Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/sql_database_instance#dns_servers SqlDatabaseInstance#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// The mode of the Active Directory configuration. Can be MANAGED_ACTIVE_DIRECTORY or CUSTOMER_MANAGED_ACTIVE_DIRECTORY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/sql_database_instance#mode SqlDatabaseInstance#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The organizational unit distinguished name. This is the full hierarchical path to the organizational unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/sql_database_instance#organizational_unit SqlDatabaseInstance#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
}

