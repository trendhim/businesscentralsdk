# \TaxGroupApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaxGroup**](TaxGroupApi.md#DeleteTaxGroup) | **Delete** /companies({company_id})/taxGroups({taxGroup_id}) | deleteTaxGroup
[**GetTaxGroup**](TaxGroupApi.md#GetTaxGroup) | **Get** /companies({company_id})/taxGroups({taxGroup_id}) | getTaxGroup
[**ListTaxGroups**](TaxGroupApi.md#ListTaxGroups) | **Get** /companies({company_id})/taxGroups | listTaxGroups
[**PatchTaxGroup**](TaxGroupApi.md#PatchTaxGroup) | **Patch** /companies({company_id})/taxGroups({taxGroup_id}) | patchTaxGroup
[**PostTaxGroup**](TaxGroupApi.md#PostTaxGroup) | **Post** /companies({company_id})/taxGroups | postTaxGroup


# **DeleteTaxGroup**
> DeleteTaxGroup(ctx, companyId, taxGroupId)
deleteTaxGroup

Deletes an object of type taxGroup in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxGroupId** | [**string**](.md)| id for taxGroup | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxGroup**
> TaxGroup GetTaxGroup(ctx, companyId, taxGroupId, optional)
getTaxGroup

Retrieve the properties and relationships of an object of type taxGroup for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxGroupId** | [**string**](.md)| id for taxGroup | 
 **optional** | ***GetTaxGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTaxGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TaxGroup**](taxGroup.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaxGroups**
> CompaniesApiTaxGroupsResponse ListTaxGroups(ctx, companyId, optional)
listTaxGroups

Returns a list of taxGroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListTaxGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTaxGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiTaxGroupsResponse**](CompaniesAPITaxGroupsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTaxGroup**
> TaxGroup PatchTaxGroup(ctx, companyId, taxGroupId, contentType, ifMatch, body)
patchTaxGroup

Updates an object of type taxGroup in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxGroupId** | [**string**](.md)| id for taxGroup | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**TaxGroup**](TaxGroup.md)|  | 

### Return type

[**TaxGroup**](taxGroup.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTaxGroup**
> TaxGroup PostTaxGroup(ctx, companyId, contentType, body)
postTaxGroup

Creates an object of type taxGroup in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**TaxGroup**](TaxGroup.md)|  | 

### Return type

[**TaxGroup**](taxGroup.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

