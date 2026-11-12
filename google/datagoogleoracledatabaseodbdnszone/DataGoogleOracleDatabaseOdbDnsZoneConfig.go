// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabaseodbdnszone

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleOracleDatabaseOdbDnsZoneConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the OdbDnsZone to create.
	//
	// This value is restricted
	// to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63
	// characters in length. The value must start with a letter and end with
	// a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/data-sources/oracle_database_odb_dns_zone#odb_dns_zone_id DataGoogleOracleDatabaseOdbDnsZone#odb_dns_zone_id}
	OdbDnsZoneId *string `field:"required" json:"odbDnsZoneId" yaml:"odbDnsZoneId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/data-sources/oracle_database_odb_dns_zone#odbnetwork DataGoogleOracleDatabaseOdbDnsZone#odbnetwork}
	Odbnetwork *string `field:"required" json:"odbnetwork" yaml:"odbnetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/data-sources/oracle_database_odb_dns_zone#id DataGoogleOracleDatabaseOdbDnsZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/data-sources/oracle_database_odb_dns_zone#location DataGoogleOracleDatabaseOdbDnsZone#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/data-sources/oracle_database_odb_dns_zone#project DataGoogleOracleDatabaseOdbDnsZone#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

