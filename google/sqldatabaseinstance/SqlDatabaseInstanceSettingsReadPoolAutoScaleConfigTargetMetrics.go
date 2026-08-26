// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceSettingsReadPoolAutoScaleConfigTargetMetrics struct {
	// Metric name for Read Pool Auto Scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/sql_database_instance#metric SqlDatabaseInstance#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Target value for Read Pool Auto Scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/sql_database_instance#target_value SqlDatabaseInstance#target_value}
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

