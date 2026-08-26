// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclebigqueryexport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleBigQueryExportConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#instance ChronicleBigQueryExport#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#location ChronicleBigQueryExport#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The BigQueryExportPackage entitled for the Chronicle instance. Possible values: ["BIG_QUERY_EXPORT_PACKAGE_BYOBQ", "BIG_QUERY_EXPORT_PACKAGE_ADVANCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#big_query_export_package ChronicleBigQueryExport#big_query_export_package}
	BigQueryExportPackage *string `field:"optional" json:"bigQueryExportPackage" yaml:"bigQueryExportPackage"`
	// entity_graph_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#entity_graph_settings ChronicleBigQueryExport#entity_graph_settings}
	EntityGraphSettings *ChronicleBigQueryExportEntityGraphSettings `field:"optional" json:"entityGraphSettings" yaml:"entityGraphSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#id ChronicleBigQueryExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ioc_matches_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#ioc_matches_settings ChronicleBigQueryExport#ioc_matches_settings}
	IocMatchesSettings *ChronicleBigQueryExportIocMatchesSettings `field:"optional" json:"iocMatchesSettings" yaml:"iocMatchesSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#project ChronicleBigQueryExport#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rule_detections_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#rule_detections_settings ChronicleBigQueryExport#rule_detections_settings}
	RuleDetectionsSettings *ChronicleBigQueryExportRuleDetectionsSettings `field:"optional" json:"ruleDetectionsSettings" yaml:"ruleDetectionsSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#timeouts ChronicleBigQueryExport#timeouts}
	Timeouts *ChronicleBigQueryExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// udm_events_aggregates_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#udm_events_aggregates_settings ChronicleBigQueryExport#udm_events_aggregates_settings}
	UdmEventsAggregatesSettings *ChronicleBigQueryExportUdmEventsAggregatesSettings `field:"optional" json:"udmEventsAggregatesSettings" yaml:"udmEventsAggregatesSettings"`
	// udm_events_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_big_query_export#udm_events_settings ChronicleBigQueryExport#udm_events_settings}
	UdmEventsSettings *ChronicleBigQueryExportUdmEventsSettings `field:"optional" json:"udmEventsSettings" yaml:"udmEventsSettings"`
}

