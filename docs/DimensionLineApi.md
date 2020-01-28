# \DimensionLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDimensionLine**](DimensionLineApi.md#DeleteDimensionLine) | **Delete** /companies({company_id})/dimensionLines({dimensionLine_parentId},{dimensionLine_id}) | deleteDimensionLine
[**GetDimensionLine**](DimensionLineApi.md#GetDimensionLine) | **Get** /companies({company_id})/dimensionLines({dimensionLine_parentId},{dimensionLine_id}) | getDimensionLine
[**ListDimensionLines**](DimensionLineApi.md#ListDimensionLines) | **Get** /companies({company_id})/dimensionLines | listDimensionLines
[**PatchDimensionLine**](DimensionLineApi.md#PatchDimensionLine) | **Patch** /companies({company_id})/dimensionLines({dimensionLine_parentId},{dimensionLine_id}) | patchDimensionLine
[**PostDimensionLine**](DimensionLineApi.md#PostDimensionLine) | **Post** /companies({company_id})/dimensionLines | postDimensionLine


# **DeleteDimensionLine**
> DeleteDimensionLine(ctx, companyId, dimensionLineParentId, dimensionLineId)
deleteDimensionLine

Deletes an object of type dimensionLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionLineParentId** | [**string**](.md)| parentId for dimensionLine | 
  **dimensionLineId** | [**string**](.md)| id for dimensionLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDimensionLine**
> DimensionLine GetDimensionLine(ctx, companyId, dimensionLineParentId, dimensionLineId, optional)
getDimensionLine

Retrieve the properties and relationships of an object of type dimensionLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionLineParentId** | [**string**](.md)| parentId for dimensionLine | 
  **dimensionLineId** | [**string**](.md)| id for dimensionLine | 
 **optional** | ***GetDimensionLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDimensionLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DimensionLine**](dimensionLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDimensionLines**
> CompaniesApiDimensionLinesResponse ListDimensionLines(ctx, companyId, optional)
listDimensionLines

Returns a list of dimensionLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListDimensionLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDimensionLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiDimensionLinesResponse**](CompaniesAPIDimensionLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDimensionLine**
> DimensionLine PatchDimensionLine(ctx, companyId, dimensionLineParentId, dimensionLineId, contentType, ifMatch, body)
patchDimensionLine

Updates an object of type dimensionLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionLineParentId** | [**string**](.md)| parentId for dimensionLine | 
  **dimensionLineId** | [**string**](.md)| id for dimensionLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiDimensionLinesDimensionLineIdApiRequest**](CompaniesApiDimensionLinesDimensionLineIdApiRequest.md)|  | 

### Return type

[**DimensionLine**](dimensionLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDimensionLine**
> DimensionLine PostDimensionLine(ctx, companyId, contentType, body)
postDimensionLine

Creates an object of type dimensionLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiDimensionLinesRequest**](CompaniesApiDimensionLinesRequest.md)|  | 

### Return type

[**DimensionLine**](dimensionLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

