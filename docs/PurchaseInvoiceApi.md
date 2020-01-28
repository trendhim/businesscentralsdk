# \PurchaseInvoiceApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchaseInvoice**](PurchaseInvoiceApi.md#DeletePurchaseInvoice) | **Delete** /companies({company_id})/purchaseInvoices({purchaseInvoice_id}) | deletePurchaseInvoice
[**GetPurchaseInvoice**](PurchaseInvoiceApi.md#GetPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id}) | getPurchaseInvoice
[**ListPurchaseInvoices**](PurchaseInvoiceApi.md#ListPurchaseInvoices) | **Get** /companies({company_id})/purchaseInvoices | listPurchaseInvoices
[**PatchPurchaseInvoice**](PurchaseInvoiceApi.md#PatchPurchaseInvoice) | **Patch** /companies({company_id})/purchaseInvoices({purchaseInvoice_id}) | patchPurchaseInvoice
[**PostActionPurchaseInvoices**](PurchaseInvoiceApi.md#PostActionPurchaseInvoices) | **Post** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/Microsoft.NAV.post | postActionPurchaseInvoices
[**PostPurchaseInvoice**](PurchaseInvoiceApi.md#PostPurchaseInvoice) | **Post** /companies({company_id})/purchaseInvoices | postPurchaseInvoice


# **DeletePurchaseInvoice**
> DeletePurchaseInvoice(ctx, companyId, purchaseInvoiceId)
deletePurchaseInvoice

Deletes an object of type purchaseInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchaseInvoice**
> PurchaseInvoice GetPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)
getPurchaseInvoice

Retrieve the properties and relationships of an object of type purchaseInvoice for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***GetPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPurchaseInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PurchaseInvoice**](purchaseInvoice.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPurchaseInvoices**
> CompaniesApiPurchaseInvoicesResponse ListPurchaseInvoices(ctx, companyId, optional)
listPurchaseInvoices

Returns a list of purchaseInvoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPurchaseInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPurchaseInvoicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPurchaseInvoicesResponse**](CompaniesAPIPurchaseInvoicesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPurchaseInvoice**
> PurchaseInvoice PatchPurchaseInvoice(ctx, companyId, purchaseInvoiceId, contentType, ifMatch, body)
patchPurchaseInvoice

Updates an object of type purchaseInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiPurchaseInvoicesApiRequest**](CompaniesApiPurchaseInvoicesApiRequest.md)|  | 

### Return type

[**PurchaseInvoice**](purchaseInvoice.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostActionPurchaseInvoices**
> PostActionPurchaseInvoices(ctx, companyId, purchaseInvoiceId)
postActionPurchaseInvoices

Performs the post action for purchaseInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchaseInvoice**
> PurchaseInvoice PostPurchaseInvoice(ctx, companyId, contentType, body)
postPurchaseInvoice

Creates an object of type purchaseInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiPurchaseInvoicesRequest**](CompaniesApiPurchaseInvoicesRequest.md)|  | 

### Return type

[**PurchaseInvoice**](purchaseInvoice.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

