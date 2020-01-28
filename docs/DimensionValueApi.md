# \DimensionValueApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDimensionValue**](DimensionValueApi.md#GetDimensionValue) | **Get** /companies({company_id})/dimensionValues({dimensionValue_id}) | getDimensionValue
[**GetDimensionValueForDimension**](DimensionValueApi.md#GetDimensionValueForDimension) | **Get** /companies({company_id})/dimensions({dimension_id})/dimensionValues({dimensionValue_id}) | getDimensionValueForDimension
[**ListDimensionValues**](DimensionValueApi.md#ListDimensionValues) | **Get** /companies({company_id})/dimensionValues | listDimensionValues
[**ListDimensionValuesForDimension**](DimensionValueApi.md#ListDimensionValuesForDimension) | **Get** /companies({company_id})/dimensions({dimension_id})/dimensionValues | listDimensionValuesForDimension


# **GetDimensionValue**
> DimensionValue GetDimensionValue(ctx, companyId, dimensionValueId, optional)
getDimensionValue

Retrieve the properties and relationships of an object of type dimensionValue for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionValueId** | [**string**](.md)| id for dimensionValue | 
 **optional** | ***GetDimensionValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDimensionValueOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DimensionValue**](dimensionValue.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDimensionValueForDimension**
> DimensionValue GetDimensionValueForDimension(ctx, companyId, dimensionId, dimensionValueId, optional)
getDimensionValueForDimension

Retrieve the properties and relationships of an object of type dimensionValue for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionId** | [**string**](.md)| id for dimension | 
  **dimensionValueId** | [**string**](.md)| id for dimensionValue | 
 **optional** | ***GetDimensionValueForDimensionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDimensionValueForDimensionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DimensionValue**](dimensionValue.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDimensionValues**
> CompaniesApiDimensionValuesResponse ListDimensionValues(ctx, companyId, optional)
listDimensionValues

Returns a list of dimensionValues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListDimensionValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDimensionValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiDimensionValuesResponse**](CompaniesAPIDimensionValuesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDimensionValuesForDimension**
> CompaniesApiDimensionsApiDimensionValuesResponse ListDimensionValuesForDimension(ctx, companyId, dimensionId, optional)
listDimensionValuesForDimension

Returns a list of dimensionValues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionId** | [**string**](.md)| id for dimension | 
 **optional** | ***ListDimensionValuesForDimensionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDimensionValuesForDimensionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiDimensionsApiDimensionValuesResponse**](CompaniesAPIDimensionsAPIDimensionValuesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

