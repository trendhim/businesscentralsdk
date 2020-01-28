# \SalesQuoteLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesQuoteLine**](SalesQuoteLineApi.md#DeleteSalesQuoteLine) | **Delete** /companies({company_id})/salesQuoteLines(&#39;{salesQuoteLine_id}&#39;) | deleteSalesQuoteLine
[**DeleteSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#DeleteSalesQuoteLineForSalesQuote) | **Delete** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines(&#39;{salesQuoteLine_id}&#39;) | deleteSalesQuoteLineForSalesQuote
[**GetSalesQuoteLine**](SalesQuoteLineApi.md#GetSalesQuoteLine) | **Get** /companies({company_id})/salesQuoteLines(&#39;{salesQuoteLine_id}&#39;) | getSalesQuoteLine
[**GetSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#GetSalesQuoteLineForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines(&#39;{salesQuoteLine_id}&#39;) | getSalesQuoteLineForSalesQuote
[**ListSalesQuoteLines**](SalesQuoteLineApi.md#ListSalesQuoteLines) | **Get** /companies({company_id})/salesQuoteLines | listSalesQuoteLines
[**ListSalesQuoteLinesForSalesQuote**](SalesQuoteLineApi.md#ListSalesQuoteLinesForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines | listSalesQuoteLinesForSalesQuote
[**PatchSalesQuoteLine**](SalesQuoteLineApi.md#PatchSalesQuoteLine) | **Patch** /companies({company_id})/salesQuoteLines(&#39;{salesQuoteLine_id}&#39;) | patchSalesQuoteLine
[**PatchSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#PatchSalesQuoteLineForSalesQuote) | **Patch** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines(&#39;{salesQuoteLine_id}&#39;) | patchSalesQuoteLineForSalesQuote
[**PostSalesQuoteLine**](SalesQuoteLineApi.md#PostSalesQuoteLine) | **Post** /companies({company_id})/salesQuoteLines | postSalesQuoteLine
[**PostSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#PostSalesQuoteLineForSalesQuote) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines | postSalesQuoteLineForSalesQuote


# **DeleteSalesQuoteLine**
> DeleteSalesQuoteLine(ctx, companyId, salesQuoteLineId)
deleteSalesQuoteLine

Deletes an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesQuoteLineForSalesQuote**
> DeleteSalesQuoteLineForSalesQuote(ctx, companyId, salesQuoteId, salesQuoteLineId)
deleteSalesQuoteLineForSalesQuote

Deletes an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesQuoteLine**
> SalesQuoteLine GetSalesQuoteLine(ctx, companyId, salesQuoteLineId, optional)
getSalesQuoteLine

Retrieve the properties and relationships of an object of type salesQuoteLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
 **optional** | ***GetSalesQuoteLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesQuoteLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesQuoteLineForSalesQuote**
> SalesQuoteLine GetSalesQuoteLineForSalesQuote(ctx, companyId, salesQuoteId, salesQuoteLineId, optional)
getSalesQuoteLineForSalesQuote

Retrieve the properties and relationships of an object of type salesQuoteLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
 **optional** | ***GetSalesQuoteLineForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesQuoteLineForSalesQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesQuoteLines**
> CompaniesApiSalesQuoteLinesResponse ListSalesQuoteLines(ctx, companyId, optional)
listSalesQuoteLines

Returns a list of salesQuoteLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesQuoteLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesQuoteLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesQuoteLinesResponse**](CompaniesAPISalesQuoteLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesQuoteLinesForSalesQuote**
> CompaniesApiSalesQuotesApiSalesQuoteLinesResponse ListSalesQuoteLinesForSalesQuote(ctx, companyId, salesQuoteId, optional)
listSalesQuoteLinesForSalesQuote

Returns a list of salesQuoteLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***ListSalesQuoteLinesForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesQuoteLinesForSalesQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesQuotesApiSalesQuoteLinesResponse**](CompaniesAPISalesQuotesAPISalesQuoteLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesQuoteLine**
> SalesQuoteLine PatchSalesQuoteLine(ctx, companyId, salesQuoteLineId, contentType, ifMatch, body)
patchSalesQuoteLine

Updates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesQuoteLinesApiRequest**](CompaniesApiSalesQuoteLinesApiRequest.md)|  | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesQuoteLineForSalesQuote**
> SalesQuoteLine PatchSalesQuoteLineForSalesQuote(ctx, companyId, salesQuoteId, salesQuoteLineId, contentType, ifMatch, body)
patchSalesQuoteLineForSalesQuote

Updates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesQuotesApiSalesQuoteLinesSalesQuoteLineIdApiRequest**](CompaniesApiSalesQuotesApiSalesQuoteLinesSalesQuoteLineIdApiRequest.md)|  | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesQuoteLine**
> SalesQuoteLine PostSalesQuoteLine(ctx, companyId, contentType, body)
postSalesQuoteLine

Creates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesQuoteLinesRequest**](CompaniesApiSalesQuoteLinesRequest.md)|  | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesQuoteLineForSalesQuote**
> SalesQuoteLine PostSalesQuoteLineForSalesQuote(ctx, companyId, salesQuoteId, contentType, body)
postSalesQuoteLineForSalesQuote

Creates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesQuotesApiSalesQuoteLinesRequest**](CompaniesApiSalesQuotesApiSalesQuoteLinesRequest.md)|  | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

