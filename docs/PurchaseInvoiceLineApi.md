# \PurchaseInvoiceLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#DeletePurchaseInvoiceLine) | **Delete** /companies({company_id})/purchaseInvoiceLines(&#39;{purchaseInvoiceLine_id}&#39;) | deletePurchaseInvoiceLine
[**DeletePurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#DeletePurchaseInvoiceLineForPurchaseInvoice) | **Delete** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines(&#39;{purchaseInvoiceLine_id}&#39;) | deletePurchaseInvoiceLineForPurchaseInvoice
[**GetPurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#GetPurchaseInvoiceLine) | **Get** /companies({company_id})/purchaseInvoiceLines(&#39;{purchaseInvoiceLine_id}&#39;) | getPurchaseInvoiceLine
[**GetPurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#GetPurchaseInvoiceLineForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines(&#39;{purchaseInvoiceLine_id}&#39;) | getPurchaseInvoiceLineForPurchaseInvoice
[**ListPurchaseInvoiceLines**](PurchaseInvoiceLineApi.md#ListPurchaseInvoiceLines) | **Get** /companies({company_id})/purchaseInvoiceLines | listPurchaseInvoiceLines
[**ListPurchaseInvoiceLinesForPurchaseInvoice**](PurchaseInvoiceLineApi.md#ListPurchaseInvoiceLinesForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines | listPurchaseInvoiceLinesForPurchaseInvoice
[**PatchPurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#PatchPurchaseInvoiceLine) | **Patch** /companies({company_id})/purchaseInvoiceLines(&#39;{purchaseInvoiceLine_id}&#39;) | patchPurchaseInvoiceLine
[**PatchPurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#PatchPurchaseInvoiceLineForPurchaseInvoice) | **Patch** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines(&#39;{purchaseInvoiceLine_id}&#39;) | patchPurchaseInvoiceLineForPurchaseInvoice
[**PostPurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#PostPurchaseInvoiceLine) | **Post** /companies({company_id})/purchaseInvoiceLines | postPurchaseInvoiceLine
[**PostPurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#PostPurchaseInvoiceLineForPurchaseInvoice) | **Post** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines | postPurchaseInvoiceLineForPurchaseInvoice


# **DeletePurchaseInvoiceLine**
> DeletePurchaseInvoiceLine(ctx, companyId, purchaseInvoiceLineId)
deletePurchaseInvoiceLine

Deletes an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePurchaseInvoiceLineForPurchaseInvoice**
> DeletePurchaseInvoiceLineForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, purchaseInvoiceLineId)
deletePurchaseInvoiceLineForPurchaseInvoice

Deletes an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchaseInvoiceLine**
> PurchaseInvoiceLine GetPurchaseInvoiceLine(ctx, companyId, purchaseInvoiceLineId, optional)
getPurchaseInvoiceLine

Retrieve the properties and relationships of an object of type purchaseInvoiceLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
 **optional** | ***GetPurchaseInvoiceLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPurchaseInvoiceLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchaseInvoiceLineForPurchaseInvoice**
> PurchaseInvoiceLine GetPurchaseInvoiceLineForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, purchaseInvoiceLineId, optional)
getPurchaseInvoiceLineForPurchaseInvoice

Retrieve the properties and relationships of an object of type purchaseInvoiceLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
 **optional** | ***GetPurchaseInvoiceLineForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPurchaseInvoiceLineForPurchaseInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPurchaseInvoiceLines**
> CompaniesApiPurchaseInvoiceLinesResponse ListPurchaseInvoiceLines(ctx, companyId, optional)
listPurchaseInvoiceLines

Returns a list of purchaseInvoiceLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPurchaseInvoiceLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPurchaseInvoiceLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPurchaseInvoiceLinesResponse**](CompaniesAPIPurchaseInvoiceLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPurchaseInvoiceLinesForPurchaseInvoice**
> CompaniesApiPurchaseInvoicesApiPurchaseInvoiceLinesResponse ListPurchaseInvoiceLinesForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)
listPurchaseInvoiceLinesForPurchaseInvoice

Returns a list of purchaseInvoiceLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***ListPurchaseInvoiceLinesForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPurchaseInvoiceLinesForPurchaseInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPurchaseInvoicesApiPurchaseInvoiceLinesResponse**](CompaniesAPIPurchaseInvoicesAPIPurchaseInvoiceLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPurchaseInvoiceLine**
> PurchaseInvoiceLine PatchPurchaseInvoiceLine(ctx, companyId, purchaseInvoiceLineId, contentType, ifMatch, body)
patchPurchaseInvoiceLine

Updates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiPurchaseInvoiceLinesApiRequest**](CompaniesApiPurchaseInvoiceLinesApiRequest.md)|  | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPurchaseInvoiceLineForPurchaseInvoice**
> PurchaseInvoiceLine PatchPurchaseInvoiceLineForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, purchaseInvoiceLineId, contentType, ifMatch, body)
patchPurchaseInvoiceLineForPurchaseInvoice

Updates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiPurchaseInvoicesApiPurchaseInvoiceLinesPurchaseInvoiceLineIdApiRequest**](CompaniesApiPurchaseInvoicesApiPurchaseInvoiceLinesPurchaseInvoiceLineIdApiRequest.md)|  | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchaseInvoiceLine**
> PurchaseInvoiceLine PostPurchaseInvoiceLine(ctx, companyId, contentType, body)
postPurchaseInvoiceLine

Creates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiPurchaseInvoiceLinesRequest**](CompaniesApiPurchaseInvoiceLinesRequest.md)|  | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchaseInvoiceLineForPurchaseInvoice**
> PurchaseInvoiceLine PostPurchaseInvoiceLineForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, contentType, body)
postPurchaseInvoiceLineForPurchaseInvoice

Creates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiPurchaseInvoicesApiPurchaseInvoiceLinesRequest**](CompaniesApiPurchaseInvoicesApiPurchaseInvoiceLinesRequest.md)|  | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

