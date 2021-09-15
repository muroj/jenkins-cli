# ApplicationAlertConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
**ApplicationId** | **string** |  | [optional] 
**Applications** | [**map[string]ApplicationNode**](ApplicationNode.md) |  | 
**BoundaryScope** | **string** |  | 
**CustomPayloadFields** | [**[]StaticStringField**](StaticStringField.md) |  | 
**Description** | **string** |  | 
**EvaluationType** | **string** |  | 
**Granularity** | **int32** |  | [optional] [default to GRANULARITY__600000]
**IncludeInternal** | **bool** |  | [optional] 
**IncludeSynthetic** | **bool** |  | [optional] 
**Name** | **string** |  | 
**Rule** | [**ApplicationAlertRule**](ApplicationAlertRule.md) |  | 
**Severity** | **int32** |  | [optional] 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TagFilters** | [**[]TagFilter**](TagFilter.md) |  | [optional] 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**ApplicationTimeThreshold**](ApplicationTimeThreshold.md) |  | 
**Triggering** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


