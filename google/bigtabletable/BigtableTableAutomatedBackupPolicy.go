// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigtabletable


type BigtableTableAutomatedBackupPolicy struct {
	// How frequently automated backups should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/bigtable_table#frequency BigtableTable#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// A list of Cloud Bigtable zones where automated backups are allowed to be created.
	//
	// If empty, automated backups will be created in all zones of the instance. Locations are in the format projects/{project}/locations/{zone}. This field can only be set for tables in Enterprise Plus instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/bigtable_table#locations BigtableTable#locations}
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// How long the automated backups should be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/bigtable_table#retention_period BigtableTable#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

