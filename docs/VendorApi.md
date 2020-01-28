# \VendorApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVendor**](VendorApi.md#DeleteVendor) | **Delete** /companies({company_id})/vendors({vendor_id}) | deleteVendor
[**GetVendor**](VendorApi.md#GetVendor) | **Get** /companies({company_id})/vendors({vendor_id}) | getVendor
[**ListVendors**](VendorApi.md#ListVendors) | **Get** /companies({company_id})/vendors | listVendors
[**PatchVendor**](VendorApi.md#PatchVendor) | **Patch** /companies({company_id})/vendors({vendor_id}) | patchVendor
[**PostVendor**](VendorApi.md#PostVendor) | **Post** /companies({company_id})/vendors | postVendor


# **DeleteVendor**
> DeleteVendor(ctx, companyId, vendorId)
deleteVendor

Deletes an object of type vendor in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVendor**
> Vendor GetVendor(ctx, companyId, vendorId, optional)
getVendor

Retrieve the properties and relationships of an object of type vendor for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
 **optional** | ***GetVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVendorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Vendor**](vendor.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVendors**
> CompaniesApiVendorsResponse ListVendors(ctx, companyId, optional)
listVendors

Returns a list of vendors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListVendorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVendorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiVendorsResponse**](CompaniesAPIVendorsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchVendor**
> Vendor PatchVendor(ctx, companyId, vendorId, contentType, ifMatch, body)
patchVendor

Updates an object of type vendor in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiVendorsApiRequest**](CompaniesApiVendorsApiRequest.md)|  | 

### Return type

[**Vendor**](vendor.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVendor**
> Vendor PostVendor(ctx, companyId, contentType, body)
postVendor

Creates an object of type vendor in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiVendorsRequest**](CompaniesApiVendorsRequest.md)|  | 

### Return type

[**Vendor**](vendor.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

