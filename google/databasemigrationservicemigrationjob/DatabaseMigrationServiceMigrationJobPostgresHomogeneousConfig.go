// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob


type DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfig struct {
	// Whether the migration uses native logical replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/database_migration_service_migration_job#is_native_logical DatabaseMigrationServiceMigrationJob#is_native_logical}
	IsNativeLogical interface{} `field:"required" json:"isNativeLogical" yaml:"isNativeLogical"`
	// Maximum number of additional subscriptions to use for the migration job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/database_migration_service_migration_job#max_additional_subscriptions DatabaseMigrationServiceMigrationJob#max_additional_subscriptions}
	MaxAdditionalSubscriptions *float64 `field:"optional" json:"maxAdditionalSubscriptions" yaml:"maxAdditionalSubscriptions"`
}

