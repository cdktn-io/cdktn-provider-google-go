// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclenativedashboard


type ChronicleNativeDashboardFilters struct {
	// The IDs of charts that this filter applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#chart_ids ChronicleNativeDashboard#chart_ids}
	ChartIds *[]*string `field:"optional" json:"chartIds" yaml:"chartIds"`
	// The data source for the filter. Possible values: UDM, ENTITY, INGESTION_METRICS, RULE_DETECTIONS, RULESETS, GLOBAL, IOC_MATCHES, RULES, SOAR_CASES, SOAR_PLAYBOOKS, SOAR_CASE_HISTORY, DATA_TABLE, INVESTIGATION, INVESTIGATION_FEEDBACK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#data_source ChronicleNativeDashboard#data_source}
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The display name of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#display_name ChronicleNativeDashboard#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The UDM field path being filtered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#field_path ChronicleNativeDashboard#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// filter_operator_and_field_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#filter_operator_and_field_values ChronicleNativeDashboard#filter_operator_and_field_values}
	FilterOperatorAndFieldValues interface{} `field:"optional" json:"filterOperatorAndFieldValues" yaml:"filterOperatorAndFieldValues"`
	// The unique ID of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#id ChronicleNativeDashboard#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the filter is mandatory for the dashboard consumer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#is_mandatory ChronicleNativeDashboard#is_mandatory}
	IsMandatory interface{} `field:"optional" json:"isMandatory" yaml:"isMandatory"`
	// Whether the filter is a standard time range filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#is_standard_time_range_filter ChronicleNativeDashboard#is_standard_time_range_filter}
	IsStandardTimeRangeFilter interface{} `field:"optional" json:"isStandardTimeRangeFilter" yaml:"isStandardTimeRangeFilter"`
	// Whether the standard time range filter is currently enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/chronicle_native_dashboard#is_standard_time_range_filter_enabled ChronicleNativeDashboard#is_standard_time_range_filter_enabled}
	IsStandardTimeRangeFilterEnabled interface{} `field:"optional" json:"isStandardTimeRangeFilterEnabled" yaml:"isStandardTimeRangeFilterEnabled"`
}

