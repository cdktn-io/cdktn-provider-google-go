// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AmazonKinesisConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference
	AmazonKinesisConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties
	AmazonRedshiftConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionPropertiesOutputReference
	AmazonRedshiftConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties
	AmazonS3ConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionPropertiesOutputReference
	AmazonS3ConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties
	AzureDataLakeStorageConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionPropertiesOutputReference
	AzureDataLakeStorageConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties
	AzureSynapseAnalyticsConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionPropertiesOutputReference
	AzureSynapseAnalyticsConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabricksConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionPropertiesOutputReference
	DatabricksConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties
	Db2ConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference
	Db2ConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ElasticsearchConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionPropertiesOutputReference
	ElasticsearchConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties
	// Experimental.
	Fqn() *string
	GenericConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGenericConnectionPropertiesOutputReference
	GenericConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties
	GoldengateConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionPropertiesOutputReference
	GoldengateConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties
	GoogleBigQueryConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionPropertiesOutputReference
	GoogleBigQueryConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties
	GoogleCloudStorageConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionPropertiesOutputReference
	GoogleCloudStorageConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties
	GooglePubsubConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionPropertiesOutputReference
	GooglePubsubConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties
	HdfsConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionPropertiesOutputReference
	HdfsConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties
	IcebergConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesOutputReference
	IcebergConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties
	IngressIpAddresses() *[]*string
	InternalValue() *OracleDatabaseGoldengateConnectionProperties
	SetInternalValue(val *OracleDatabaseGoldengateConnectionProperties)
	JavaMessageServiceConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference
	JavaMessageServiceConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties
	KafkaConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference
	KafkaConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties
	KafkaSchemaRegistryConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference
	KafkaSchemaRegistryConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties
	LifecycleDetails() *string
	LifecycleState() *string
	MicrosoftFabricConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionPropertiesOutputReference
	MicrosoftFabricConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties
	MicrosoftSqlserverConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference
	MicrosoftSqlserverConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties
	MongodbConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionPropertiesOutputReference
	MongodbConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties
	MysqlConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference
	MysqlConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties
	Ocid() *string
	OciObjectStorageConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionPropertiesOutputReference
	OciObjectStorageConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties
	OracleAiDataPlatformConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference
	OracleAiDataPlatformConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties
	OracleConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOracleConnectionPropertiesOutputReference
	OracleConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties
	OracleNosqlConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference
	OracleNosqlConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	PostgresqlConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionPropertiesOutputReference
	PostgresqlConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties
	RedisConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesRedisConnectionPropertiesOutputReference
	RedisConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties
	RoutingMethod() *string
	SetRoutingMethod(val *string)
	RoutingMethodInput() *string
	SnowflakeConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionPropertiesOutputReference
	SnowflakeConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAmazonKinesisConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties)
	PutAmazonRedshiftConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties)
	PutAmazonS3ConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties)
	PutAzureDataLakeStorageConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties)
	PutAzureSynapseAnalyticsConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties)
	PutDatabricksConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties)
	PutDb2ConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties)
	PutElasticsearchConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties)
	PutGenericConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties)
	PutGoldengateConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties)
	PutGoogleBigQueryConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties)
	PutGoogleCloudStorageConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties)
	PutGooglePubsubConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties)
	PutHdfsConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties)
	PutIcebergConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties)
	PutJavaMessageServiceConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties)
	PutKafkaConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties)
	PutKafkaSchemaRegistryConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties)
	PutMicrosoftFabricConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties)
	PutMicrosoftSqlserverConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties)
	PutMongodbConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties)
	PutMysqlConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties)
	PutOciObjectStorageConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties)
	PutOracleAiDataPlatformConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties)
	PutOracleConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties)
	PutOracleNosqlConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties)
	PutPostgresqlConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties)
	PutRedisConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties)
	PutSnowflakeConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties)
	ResetAmazonKinesisConnectionProperties()
	ResetAmazonRedshiftConnectionProperties()
	ResetAmazonS3ConnectionProperties()
	ResetAzureDataLakeStorageConnectionProperties()
	ResetAzureSynapseAnalyticsConnectionProperties()
	ResetDatabricksConnectionProperties()
	ResetDb2ConnectionProperties()
	ResetDescription()
	ResetElasticsearchConnectionProperties()
	ResetGenericConnectionProperties()
	ResetGoldengateConnectionProperties()
	ResetGoogleBigQueryConnectionProperties()
	ResetGoogleCloudStorageConnectionProperties()
	ResetGooglePubsubConnectionProperties()
	ResetHdfsConnectionProperties()
	ResetIcebergConnectionProperties()
	ResetJavaMessageServiceConnectionProperties()
	ResetKafkaConnectionProperties()
	ResetKafkaSchemaRegistryConnectionProperties()
	ResetMicrosoftFabricConnectionProperties()
	ResetMicrosoftSqlserverConnectionProperties()
	ResetMongodbConnectionProperties()
	ResetMysqlConnectionProperties()
	ResetOciObjectStorageConnectionProperties()
	ResetOracleAiDataPlatformConnectionProperties()
	ResetOracleConnectionProperties()
	ResetOracleNosqlConnectionProperties()
	ResetPostgresqlConnectionProperties()
	ResetRedisConnectionProperties()
	ResetRoutingMethod()
	ResetSnowflakeConnectionProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateConnectionPropertiesOutputReference
type jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonKinesisConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonKinesisConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonKinesisConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties
	_jsii_.Get(
		j,
		"amazonKinesisConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonRedshiftConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonRedshiftConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonRedshiftConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties
	_jsii_.Get(
		j,
		"amazonRedshiftConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonS3ConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonS3ConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonS3ConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties
	_jsii_.Get(
		j,
		"amazonS3ConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureDataLakeStorageConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"azureDataLakeStorageConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureDataLakeStorageConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties
	_jsii_.Get(
		j,
		"azureDataLakeStorageConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureSynapseAnalyticsConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"azureSynapseAnalyticsConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureSynapseAnalyticsConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties
	_jsii_.Get(
		j,
		"azureSynapseAnalyticsConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) DatabricksConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"databricksConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) DatabricksConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties
	_jsii_.Get(
		j,
		"databricksConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) Db2ConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"db2ConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) Db2ConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties
	_jsii_.Get(
		j,
		"db2ConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ElasticsearchConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ElasticsearchConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties
	_jsii_.Get(
		j,
		"elasticsearchConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GenericConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGenericConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesGenericConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"genericConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GenericConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties
	_jsii_.Get(
		j,
		"genericConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GoldengateConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"goldengateConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GoldengateConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties
	_jsii_.Get(
		j,
		"goldengateConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleBigQueryConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"googleBigQueryConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleBigQueryConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties
	_jsii_.Get(
		j,
		"googleBigQueryConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleCloudStorageConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleCloudStorageConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties
	_jsii_.Get(
		j,
		"googleCloudStorageConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GooglePubsubConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"googlePubsubConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GooglePubsubConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties
	_jsii_.Get(
		j,
		"googlePubsubConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) HdfsConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"hdfsConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) HdfsConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties
	_jsii_.Get(
		j,
		"hdfsConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) IcebergConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"icebergConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) IcebergConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties
	_jsii_.Get(
		j,
		"icebergConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) IngressIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) InternalValue() *OracleDatabaseGoldengateConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) JavaMessageServiceConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"javaMessageServiceConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) JavaMessageServiceConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties
	_jsii_.Get(
		j,
		"javaMessageServiceConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"kafkaConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties
	_jsii_.Get(
		j,
		"kafkaConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaSchemaRegistryConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"kafkaSchemaRegistryConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaSchemaRegistryConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties
	_jsii_.Get(
		j,
		"kafkaSchemaRegistryConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftFabricConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"microsoftFabricConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftFabricConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties
	_jsii_.Get(
		j,
		"microsoftFabricConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftSqlserverConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"microsoftSqlserverConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftSqlserverConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties
	_jsii_.Get(
		j,
		"microsoftSqlserverConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MongodbConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"mongodbConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MongodbConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties
	_jsii_.Get(
		j,
		"mongodbConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MysqlConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"mysqlConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) MysqlConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties
	_jsii_.Get(
		j,
		"mysqlConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OciObjectStorageConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"ociObjectStorageConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OciObjectStorageConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties
	_jsii_.Get(
		j,
		"ociObjectStorageConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleAiDataPlatformConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"oracleAiDataPlatformConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleAiDataPlatformConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties
	_jsii_.Get(
		j,
		"oracleAiDataPlatformConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOracleConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesOracleConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"oracleConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties
	_jsii_.Get(
		j,
		"oracleConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleNosqlConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"oracleNosqlConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleNosqlConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	_jsii_.Get(
		j,
		"oracleNosqlConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PostgresqlConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"postgresqlConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PostgresqlConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties
	_jsii_.Get(
		j,
		"postgresqlConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) RedisConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesRedisConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesRedisConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"redisConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) RedisConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties
	_jsii_.Get(
		j,
		"redisConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) RoutingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) RoutingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) SnowflakeConnectionProperties() OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionPropertiesOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"snowflakeConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) SnowflakeConnectionPropertiesInput() *OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties
	_jsii_.Get(
		j,
		"snowflakeConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateConnectionPropertiesOutputReference_Override(o OracleDatabaseGoldengateConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetInternalValue(val *OracleDatabaseGoldengateConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetRoutingMethod(val *string) {
	if err := j.validateSetRoutingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingMethod",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAmazonKinesisConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties) {
	if err := o.validatePutAmazonKinesisConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonKinesisConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAmazonRedshiftConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties) {
	if err := o.validatePutAmazonRedshiftConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonRedshiftConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAmazonS3ConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties) {
	if err := o.validatePutAmazonS3ConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonS3ConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAzureDataLakeStorageConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties) {
	if err := o.validatePutAzureDataLakeStorageConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAzureDataLakeStorageConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAzureSynapseAnalyticsConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties) {
	if err := o.validatePutAzureSynapseAnalyticsConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAzureSynapseAnalyticsConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutDatabricksConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties) {
	if err := o.validatePutDatabricksConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatabricksConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutDb2ConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties) {
	if err := o.validatePutDb2ConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDb2ConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutElasticsearchConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties) {
	if err := o.validatePutElasticsearchConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putElasticsearchConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGenericConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties) {
	if err := o.validatePutGenericConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGenericConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGoldengateConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties) {
	if err := o.validatePutGoldengateConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoldengateConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGoogleBigQueryConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties) {
	if err := o.validatePutGoogleBigQueryConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoogleBigQueryConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGoogleCloudStorageConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties) {
	if err := o.validatePutGoogleCloudStorageConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoogleCloudStorageConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGooglePubsubConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties) {
	if err := o.validatePutGooglePubsubConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGooglePubsubConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutHdfsConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties) {
	if err := o.validatePutHdfsConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHdfsConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutIcebergConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties) {
	if err := o.validatePutIcebergConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIcebergConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutJavaMessageServiceConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties) {
	if err := o.validatePutJavaMessageServiceConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJavaMessageServiceConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutKafkaConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties) {
	if err := o.validatePutKafkaConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putKafkaConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutKafkaSchemaRegistryConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties) {
	if err := o.validatePutKafkaSchemaRegistryConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putKafkaSchemaRegistryConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMicrosoftFabricConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties) {
	if err := o.validatePutMicrosoftFabricConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMicrosoftFabricConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMicrosoftSqlserverConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties) {
	if err := o.validatePutMicrosoftSqlserverConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMicrosoftSqlserverConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMongodbConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties) {
	if err := o.validatePutMongodbConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMongodbConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMysqlConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties) {
	if err := o.validatePutMysqlConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMysqlConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOciObjectStorageConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties) {
	if err := o.validatePutOciObjectStorageConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOciObjectStorageConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOracleAiDataPlatformConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties) {
	if err := o.validatePutOracleAiDataPlatformConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOracleAiDataPlatformConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOracleConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties) {
	if err := o.validatePutOracleConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOracleConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOracleNosqlConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties) {
	if err := o.validatePutOracleNosqlConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOracleNosqlConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutPostgresqlConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties) {
	if err := o.validatePutPostgresqlConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPostgresqlConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutRedisConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties) {
	if err := o.validatePutRedisConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRedisConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) PutSnowflakeConnectionProperties(value *OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties) {
	if err := o.validatePutSnowflakeConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSnowflakeConnectionProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAmazonKinesisConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonKinesisConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAmazonRedshiftConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonRedshiftConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAmazonS3ConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonS3ConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAzureDataLakeStorageConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetAzureDataLakeStorageConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAzureSynapseAnalyticsConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetAzureSynapseAnalyticsConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetDatabricksConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabricksConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetDb2ConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetDb2ConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetElasticsearchConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticsearchConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGenericConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetGenericConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGoldengateConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetGoldengateConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGoogleBigQueryConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleBigQueryConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGoogleCloudStorageConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleCloudStorageConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGooglePubsubConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetGooglePubsubConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetHdfsConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetHdfsConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetIcebergConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetIcebergConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetJavaMessageServiceConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetJavaMessageServiceConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetKafkaConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetKafkaConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetKafkaSchemaRegistryConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetKafkaSchemaRegistryConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMicrosoftFabricConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetMicrosoftFabricConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMicrosoftSqlserverConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetMicrosoftSqlserverConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMongodbConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetMongodbConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMysqlConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetMysqlConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOciObjectStorageConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetOciObjectStorageConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOracleAiDataPlatformConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetOracleAiDataPlatformConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOracleConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetOracleConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOracleNosqlConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetOracleNosqlConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetPostgresqlConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetPostgresqlConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetRedisConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetRedisConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetRoutingMethod() {
	_jsii_.InvokeVoid(
		o,
		"resetRoutingMethod",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetSnowflakeConnectionProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetSnowflakeConnectionProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

