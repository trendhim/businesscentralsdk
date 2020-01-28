# \TaxAreaApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaxArea**](TaxAreaApi.md#DeleteTaxArea) | **Delete** /companies({company_id})/taxAreas({taxArea_id}) | deleteTaxArea
[**GetTaxArea**](TaxAreaApi.md#GetTaxArea) | **Get** /companies({company_id})/taxAreas({taxArea_id}) | getTaxArea
[**ListTaxAreas**](TaxAreaApi.md#ListTaxAreas) | **Get** /companies({company_id})/taxAreas | listTaxAreas
[**PatchTaxArea**](TaxAreaApi.md#PatchTaxArea) | **Patch** /companies({company_id})/taxAreas({taxArea_id}) | patchTaxArea
[**PostTaxArea**](TaxAreaApi.md#PostTaxArea) | **Post** /companies({company_id})/taxAreas | postTaxArea


# **DeleteTaxArea**
> DeleteTaxArea(ctx, companyId, taxAreaId)
deleteTaxArea

Deletes an object of type taxArea in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxAreaId** | [**string**](.md)| id for taxArea | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxArea**
> TaxArea GetTaxArea(ctx, companyId, taxAreaId, optional)
getTaxArea

Retrieve the properties and relationships of an object of type taxArea for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxAreaId** | [**string**](.md)| id for taxArea | 
 **optional** | ***GetTaxAreaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTaxAreaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TaxArea**](taxArea.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaxAreas**
> CompaniesApiTaxAreasResponse ListTaxAreas(ctx, companyId, optional)
listTaxAreas

Returns a list of taxAreas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListTaxAreasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTaxAreasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiTaxAreasResponse**](CompaniesAPITaxAreasResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTaxArea**
> TaxArea PatchTaxArea(ctx, companyId, taxAreaId, contentType, ifMatch, body)
patchTaxArea

Updates an object of type taxArea in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxAreaId** | [**string**](.md)| id for taxArea | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**TaxArea**](TaxArea.md)|  | 

### Return type

[**TaxArea**](taxArea.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTaxArea**
> TaxArea PostTaxArea(ctx, companyId, contentType, body)
postTaxArea

Creates an object of type taxArea in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**TaxArea**](TaxArea.md)|  | 

### Return type

[**TaxArea**](taxArea.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

