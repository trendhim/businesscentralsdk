# \VendorPurchaseApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendorPurchase**](VendorPurchaseApi.md#GetVendorPurchase) | **Get** /companies({company_id})/vendorPurchases({vendorPurchase_vendorId},&#39;{vendorPurchase_vendorNumber}&#39;,&#39;{vendorPurchase_name}&#39;) | getVendorPurchase
[**ListVendorPurchases**](VendorPurchaseApi.md#ListVendorPurchases) | **Get** /companies({company_id})/vendorPurchases | listVendorPurchases


# **GetVendorPurchase**
> VendorPurchase GetVendorPurchase(ctx, companyId, vendorPurchaseVendorId, vendorPurchaseVendorNumber, vendorPurchaseName, optional)
getVendorPurchase

Retrieve the properties and relationships of an object of type vendorPurchase for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorPurchaseVendorId** | [**string**](.md)| vendorId for vendorPurchase | 
  **vendorPurchaseVendorNumber** | **string**| vendorNumber for vendorPurchase | 
  **vendorPurchaseName** | **string**| name for vendorPurchase | 
 **optional** | ***GetVendorPurchaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVendorPurchaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**VendorPurchase**](vendorPurchase.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVendorPurchases**
> CompaniesApiVendorPurchasesResponse ListVendorPurchases(ctx, companyId, optional)
listVendorPurchases

Returns a list of vendorPurchases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListVendorPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVendorPurchasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiVendorPurchasesResponse**](CompaniesAPIVendorPurchasesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

