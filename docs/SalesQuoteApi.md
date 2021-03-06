# \SalesQuoteApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesQuote**](SalesQuoteApi.md#DeleteSalesQuote) | **Delete** /companies({company_id})/salesQuotes({salesQuote_id}) | deleteSalesQuote
[**GetSalesQuote**](SalesQuoteApi.md#GetSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id}) | getSalesQuote
[**ListSalesQuotes**](SalesQuoteApi.md#ListSalesQuotes) | **Get** /companies({company_id})/salesQuotes | listSalesQuotes
[**MakeInvoiceActionSalesQuotes**](SalesQuoteApi.md#MakeInvoiceActionSalesQuotes) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/Microsoft.NAV.makeInvoice | makeInvoiceActionSalesQuotes
[**MakeOrderActionSalesQuotes**](SalesQuoteApi.md#MakeOrderActionSalesQuotes) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/Microsoft.NAV.makeOrder | makeOrderActionSalesQuotes
[**PatchSalesQuote**](SalesQuoteApi.md#PatchSalesQuote) | **Patch** /companies({company_id})/salesQuotes({salesQuote_id}) | patchSalesQuote
[**PostSalesQuote**](SalesQuoteApi.md#PostSalesQuote) | **Post** /companies({company_id})/salesQuotes | postSalesQuote
[**SendActionSalesQuotes**](SalesQuoteApi.md#SendActionSalesQuotes) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/Microsoft.NAV.send | sendActionSalesQuotes


# **DeleteSalesQuote**
> DeleteSalesQuote(ctx, companyId, salesQuoteId)
deleteSalesQuote

Deletes an object of type salesQuote in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesQuote**
> SalesQuote GetSalesQuote(ctx, companyId, salesQuoteId, optional)
getSalesQuote

Retrieve the properties and relationships of an object of type salesQuote for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***GetSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesQuote**](salesQuote.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesQuotes**
> CompaniesApiSalesQuotesResponse ListSalesQuotes(ctx, companyId, optional)
listSalesQuotes

Returns a list of salesQuotes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesQuotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesQuotesResponse**](CompaniesAPISalesQuotesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeInvoiceActionSalesQuotes**
> MakeInvoiceActionSalesQuotes(ctx, companyId, salesQuoteId)
makeInvoiceActionSalesQuotes

Performs the makeInvoice action for salesQuotes entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeOrderActionSalesQuotes**
> MakeOrderActionSalesQuotes(ctx, companyId, salesQuoteId)
makeOrderActionSalesQuotes

Performs the makeOrder action for salesQuotes entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesQuote**
> SalesQuote PatchSalesQuote(ctx, companyId, salesQuoteId, contentType, ifMatch, body)
patchSalesQuote

Updates an object of type salesQuote in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesQuotesApiRequest**](CompaniesApiSalesQuotesApiRequest.md)|  | 

### Return type

[**SalesQuote**](salesQuote.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesQuote**
> SalesQuote PostSalesQuote(ctx, companyId, contentType, body)
postSalesQuote

Creates an object of type salesQuote in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesQuotesRequest**](CompaniesApiSalesQuotesRequest.md)|  | 

### Return type

[**SalesQuote**](salesQuote.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendActionSalesQuotes**
> SendActionSalesQuotes(ctx, companyId, salesQuoteId)
sendActionSalesQuotes

Performs the send action for salesQuotes entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

