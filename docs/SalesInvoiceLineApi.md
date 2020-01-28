# \SalesInvoiceLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesInvoiceLine**](SalesInvoiceLineApi.md#DeleteSalesInvoiceLine) | **Delete** /companies({company_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | deleteSalesInvoiceLine
[**DeleteSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#DeleteSalesInvoiceLineForSalesInvoice) | **Delete** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | deleteSalesInvoiceLineForSalesInvoice
[**GetSalesInvoiceLine**](SalesInvoiceLineApi.md#GetSalesInvoiceLine) | **Get** /companies({company_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | getSalesInvoiceLine
[**GetSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#GetSalesInvoiceLineForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | getSalesInvoiceLineForSalesInvoice
[**ListSalesInvoiceLines**](SalesInvoiceLineApi.md#ListSalesInvoiceLines) | **Get** /companies({company_id})/salesInvoiceLines | listSalesInvoiceLines
[**ListSalesInvoiceLinesForSalesInvoice**](SalesInvoiceLineApi.md#ListSalesInvoiceLinesForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines | listSalesInvoiceLinesForSalesInvoice
[**PatchSalesInvoiceLine**](SalesInvoiceLineApi.md#PatchSalesInvoiceLine) | **Patch** /companies({company_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | patchSalesInvoiceLine
[**PatchSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#PatchSalesInvoiceLineForSalesInvoice) | **Patch** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | patchSalesInvoiceLineForSalesInvoice
[**PostSalesInvoiceLine**](SalesInvoiceLineApi.md#PostSalesInvoiceLine) | **Post** /companies({company_id})/salesInvoiceLines | postSalesInvoiceLine
[**PostSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#PostSalesInvoiceLineForSalesInvoice) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines | postSalesInvoiceLineForSalesInvoice


# **DeleteSalesInvoiceLine**
> DeleteSalesInvoiceLine(ctx, companyId, salesInvoiceLineId)
deleteSalesInvoiceLine

Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceLineId** | **string**| id for salesInvoiceLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesInvoiceLineForSalesInvoice**
> DeleteSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, salesInvoiceLineId)
deleteSalesInvoiceLineForSalesInvoice

Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **salesInvoiceLineId** | **string**| id for salesInvoiceLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesInvoiceLine**
> SalesInvoiceLine GetSalesInvoiceLine(ctx, companyId, salesInvoiceLineId, optional)
getSalesInvoiceLine

Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
 **optional** | ***GetSalesInvoiceLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesInvoiceLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesInvoiceLineForSalesInvoice**
> SalesInvoiceLine GetSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, salesInvoiceLineId, optional)
getSalesInvoiceLineForSalesInvoice

Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
 **optional** | ***GetSalesInvoiceLineForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesInvoiceLineForSalesInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesInvoiceLines**
> CompaniesApiSalesInvoiceLinesResponse ListSalesInvoiceLines(ctx, companyId, optional)
listSalesInvoiceLines

Returns a list of salesInvoiceLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesInvoiceLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesInvoiceLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesInvoiceLinesResponse**](CompaniesAPISalesInvoiceLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesInvoiceLinesForSalesInvoice**
> CompaniesApiSalesInvoicesApiSalesInvoiceLinesResponse ListSalesInvoiceLinesForSalesInvoice(ctx, companyId, salesInvoiceId, optional)
listSalesInvoiceLinesForSalesInvoice

Returns a list of salesInvoiceLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***ListSalesInvoiceLinesForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesInvoiceLinesForSalesInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesInvoicesApiSalesInvoiceLinesResponse**](CompaniesAPISalesInvoicesAPISalesInvoiceLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesInvoiceLine**
> SalesInvoiceLine PatchSalesInvoiceLine(ctx, companyId, salesInvoiceLineId, contentType, ifMatch, body)
patchSalesInvoiceLine

Updates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesInvoiceLinesApiRequest**](CompaniesApiSalesInvoiceLinesApiRequest.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesInvoiceLineForSalesInvoice**
> SalesInvoiceLine PatchSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, salesInvoiceLineId, contentType, ifMatch, body)
patchSalesInvoiceLineForSalesInvoice

Updates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesInvoicesApiSalesInvoiceLinesSalesInvoiceLineIdApiRequest**](CompaniesApiSalesInvoicesApiSalesInvoiceLinesSalesInvoiceLineIdApiRequest.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesInvoiceLine**
> SalesInvoiceLine PostSalesInvoiceLine(ctx, companyId, contentType, body)
postSalesInvoiceLine

Creates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesInvoiceLinesRequest**](CompaniesApiSalesInvoiceLinesRequest.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesInvoiceLineForSalesInvoice**
> SalesInvoiceLine PostSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, contentType, body)
postSalesInvoiceLineForSalesInvoice

Creates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesInvoicesApiSalesInvoiceLinesRequest**](CompaniesApiSalesInvoicesApiSalesInvoiceLinesRequest.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

