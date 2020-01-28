# \SalesCreditMemoLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesCreditMemoLine**](SalesCreditMemoLineApi.md#DeleteSalesCreditMemoLine) | **Delete** /companies({company_id})/salesCreditMemoLines(&#39;{salesCreditMemoLine_id}&#39;) | deleteSalesCreditMemoLine
[**DeleteSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#DeleteSalesCreditMemoLineForSalesCreditMemo) | **Delete** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines(&#39;{salesCreditMemoLine_id}&#39;) | deleteSalesCreditMemoLineForSalesCreditMemo
[**GetSalesCreditMemoLine**](SalesCreditMemoLineApi.md#GetSalesCreditMemoLine) | **Get** /companies({company_id})/salesCreditMemoLines(&#39;{salesCreditMemoLine_id}&#39;) | getSalesCreditMemoLine
[**GetSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#GetSalesCreditMemoLineForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines(&#39;{salesCreditMemoLine_id}&#39;) | getSalesCreditMemoLineForSalesCreditMemo
[**ListSalesCreditMemoLines**](SalesCreditMemoLineApi.md#ListSalesCreditMemoLines) | **Get** /companies({company_id})/salesCreditMemoLines | listSalesCreditMemoLines
[**ListSalesCreditMemoLinesForSalesCreditMemo**](SalesCreditMemoLineApi.md#ListSalesCreditMemoLinesForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines | listSalesCreditMemoLinesForSalesCreditMemo
[**PatchSalesCreditMemoLine**](SalesCreditMemoLineApi.md#PatchSalesCreditMemoLine) | **Patch** /companies({company_id})/salesCreditMemoLines(&#39;{salesCreditMemoLine_id}&#39;) | patchSalesCreditMemoLine
[**PatchSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#PatchSalesCreditMemoLineForSalesCreditMemo) | **Patch** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines(&#39;{salesCreditMemoLine_id}&#39;) | patchSalesCreditMemoLineForSalesCreditMemo
[**PostSalesCreditMemoLine**](SalesCreditMemoLineApi.md#PostSalesCreditMemoLine) | **Post** /companies({company_id})/salesCreditMemoLines | postSalesCreditMemoLine
[**PostSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#PostSalesCreditMemoLineForSalesCreditMemo) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines | postSalesCreditMemoLineForSalesCreditMemo


# **DeleteSalesCreditMemoLine**
> DeleteSalesCreditMemoLine(ctx, companyId, salesCreditMemoLineId)
deleteSalesCreditMemoLine

Deletes an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesCreditMemoLineForSalesCreditMemo**
> DeleteSalesCreditMemoLineForSalesCreditMemo(ctx, companyId, salesCreditMemoId, salesCreditMemoLineId)
deleteSalesCreditMemoLineForSalesCreditMemo

Deletes an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesCreditMemoLine**
> SalesCreditMemoLine GetSalesCreditMemoLine(ctx, companyId, salesCreditMemoLineId, optional)
getSalesCreditMemoLine

Retrieve the properties and relationships of an object of type salesCreditMemoLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
 **optional** | ***GetSalesCreditMemoLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesCreditMemoLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesCreditMemoLineForSalesCreditMemo**
> SalesCreditMemoLine GetSalesCreditMemoLineForSalesCreditMemo(ctx, companyId, salesCreditMemoId, salesCreditMemoLineId, optional)
getSalesCreditMemoLineForSalesCreditMemo

Retrieve the properties and relationships of an object of type salesCreditMemoLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
 **optional** | ***GetSalesCreditMemoLineForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesCreditMemoLineForSalesCreditMemoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesCreditMemoLines**
> CompaniesApiSalesCreditMemoLinesResponse ListSalesCreditMemoLines(ctx, companyId, optional)
listSalesCreditMemoLines

Returns a list of salesCreditMemoLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesCreditMemoLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesCreditMemoLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesCreditMemoLinesResponse**](CompaniesAPISalesCreditMemoLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesCreditMemoLinesForSalesCreditMemo**
> CompaniesApiSalesCreditMemosApiSalesCreditMemoLinesResponse ListSalesCreditMemoLinesForSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)
listSalesCreditMemoLinesForSalesCreditMemo

Returns a list of salesCreditMemoLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***ListSalesCreditMemoLinesForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesCreditMemoLinesForSalesCreditMemoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesCreditMemosApiSalesCreditMemoLinesResponse**](CompaniesAPISalesCreditMemosAPISalesCreditMemoLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesCreditMemoLine**
> SalesCreditMemoLine PatchSalesCreditMemoLine(ctx, companyId, salesCreditMemoLineId, contentType, ifMatch, body)
patchSalesCreditMemoLine

Updates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesCreditMemoLinesApiRequest**](CompaniesApiSalesCreditMemoLinesApiRequest.md)|  | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesCreditMemoLineForSalesCreditMemo**
> SalesCreditMemoLine PatchSalesCreditMemoLineForSalesCreditMemo(ctx, companyId, salesCreditMemoId, salesCreditMemoLineId, contentType, ifMatch, body)
patchSalesCreditMemoLineForSalesCreditMemo

Updates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesCreditMemosApiSalesCreditMemoLinesSalesCreditMemoLineIdApiRequest**](CompaniesApiSalesCreditMemosApiSalesCreditMemoLinesSalesCreditMemoLineIdApiRequest.md)|  | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesCreditMemoLine**
> SalesCreditMemoLine PostSalesCreditMemoLine(ctx, companyId, contentType, body)
postSalesCreditMemoLine

Creates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesCreditMemoLinesRequest**](CompaniesApiSalesCreditMemoLinesRequest.md)|  | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesCreditMemoLineForSalesCreditMemo**
> SalesCreditMemoLine PostSalesCreditMemoLineForSalesCreditMemo(ctx, companyId, salesCreditMemoId, contentType, body)
postSalesCreditMemoLineForSalesCreditMemo

Creates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesCreditMemosApiSalesCreditMemoLinesRequest**](CompaniesApiSalesCreditMemosApiSalesCreditMemoLinesRequest.md)|  | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

