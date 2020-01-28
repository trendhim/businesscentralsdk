# \SalesInvoiceApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelActionSalesInvoices**](SalesInvoiceApi.md#CancelActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.cancel | cancelActionSalesInvoices
[**CancelAndSendActionSalesInvoices**](SalesInvoiceApi.md#CancelAndSendActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.cancelAndSend | cancelAndSendActionSalesInvoices
[**DeleteSalesInvoice**](SalesInvoiceApi.md#DeleteSalesInvoice) | **Delete** /companies({company_id})/salesInvoices({salesInvoice_id}) | deleteSalesInvoice
[**GetSalesInvoice**](SalesInvoiceApi.md#GetSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id}) | getSalesInvoice
[**ListSalesInvoices**](SalesInvoiceApi.md#ListSalesInvoices) | **Get** /companies({company_id})/salesInvoices | listSalesInvoices
[**MakeCorrectiveCreditMemoActionSalesInvoices**](SalesInvoiceApi.md#MakeCorrectiveCreditMemoActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.makeCorrectiveCreditMemo | makeCorrectiveCreditMemoActionSalesInvoices
[**PatchSalesInvoice**](SalesInvoiceApi.md#PatchSalesInvoice) | **Patch** /companies({company_id})/salesInvoices({salesInvoice_id}) | patchSalesInvoice
[**PostActionSalesInvoices**](SalesInvoiceApi.md#PostActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.post | postActionSalesInvoices
[**PostAndSendActionSalesInvoices**](SalesInvoiceApi.md#PostAndSendActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.postAndSend | postAndSendActionSalesInvoices
[**PostSalesInvoice**](SalesInvoiceApi.md#PostSalesInvoice) | **Post** /companies({company_id})/salesInvoices | postSalesInvoice
[**SendActionSalesInvoices**](SalesInvoiceApi.md#SendActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.send | sendActionSalesInvoices


# **CancelActionSalesInvoices**
> CancelActionSalesInvoices(ctx, companyId, salesInvoiceId)
cancelActionSalesInvoices

Performs the cancel action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelAndSendActionSalesInvoices**
> CancelAndSendActionSalesInvoices(ctx, companyId, salesInvoiceId)
cancelAndSendActionSalesInvoices

Performs the cancelAndSend action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesInvoice**
> DeleteSalesInvoice(ctx, companyId, salesInvoiceId)
deleteSalesInvoice

Deletes an object of type salesInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesInvoice**
> SalesInvoice GetSalesInvoice(ctx, companyId, salesInvoiceId, optional)
getSalesInvoice

Retrieve the properties and relationships of an object of type salesInvoice for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***GetSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesInvoice**](salesInvoice.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesInvoices**
> CompaniesApiSalesInvoicesResponse ListSalesInvoices(ctx, companyId, optional)
listSalesInvoices

Returns a list of salesInvoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesInvoicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesInvoicesResponse**](CompaniesAPISalesInvoicesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeCorrectiveCreditMemoActionSalesInvoices**
> MakeCorrectiveCreditMemoActionSalesInvoices(ctx, companyId, salesInvoiceId)
makeCorrectiveCreditMemoActionSalesInvoices

Performs the makeCorrectiveCreditMemo action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesInvoice**
> SalesInvoice PatchSalesInvoice(ctx, companyId, salesInvoiceId, contentType, ifMatch, body)
patchSalesInvoice

Updates an object of type salesInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesInvoicesApiRequest**](CompaniesApiSalesInvoicesApiRequest.md)|  | 

### Return type

[**SalesInvoice**](salesInvoice.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostActionSalesInvoices**
> PostActionSalesInvoices(ctx, companyId, salesInvoiceId)
postActionSalesInvoices

Performs the post action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAndSendActionSalesInvoices**
> PostAndSendActionSalesInvoices(ctx, companyId, salesInvoiceId)
postAndSendActionSalesInvoices

Performs the postAndSend action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesInvoice**
> SalesInvoice PostSalesInvoice(ctx, companyId, contentType, body)
postSalesInvoice

Creates an object of type salesInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesInvoicesRequest**](CompaniesApiSalesInvoicesRequest.md)|  | 

### Return type

[**SalesInvoice**](salesInvoice.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendActionSalesInvoices**
> SendActionSalesInvoices(ctx, companyId, salesInvoiceId)
sendActionSalesInvoices

Performs the send action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

