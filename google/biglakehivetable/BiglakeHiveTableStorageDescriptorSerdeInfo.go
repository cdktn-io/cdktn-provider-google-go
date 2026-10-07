// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakehivetable


type BiglakeHiveTableStorageDescriptorSerdeInfo struct {
	// Name of the SerDe, table name by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#name BiglakeHiveTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The fully qualified Java class name of the serialization library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#serialization_lib BiglakeHiveTable#serialization_lib}
	SerializationLib *string `field:"required" json:"serializationLib" yaml:"serializationLib"`
	// Description of the SerDe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#description BiglakeHiveTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The fully qualified Java class name of the deserializer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#deserializer_class BiglakeHiveTable#deserializer_class}
	DeserializerClass *string `field:"optional" json:"deserializerClass" yaml:"deserializerClass"`
	// Parameters of the SerDe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#parameters BiglakeHiveTable#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The SerDe type. Possible values: ["SERDE_TYPE_UNSPECIFIED", "HIVE", "SCHEMA_REGISTRY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#serde_type BiglakeHiveTable#serde_type}
	SerdeType *string `field:"optional" json:"serdeType" yaml:"serdeType"`
	// The fully qualified Java class name of the serializer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/biglake_hive_table#serializer_class BiglakeHiveTable#serializer_class}
	SerializerClass *string `field:"optional" json:"serializerClass" yaml:"serializerClass"`
}

