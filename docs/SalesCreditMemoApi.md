# \SalesCreditMemoApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelActionSalesCreditMemos**](SalesCreditMemoApi.md#CancelActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.cancel | cancelActionSalesCreditMemos
[**CancelAndSendActionSalesCreditMemos**](SalesCreditMemoApi.md#CancelAndSendActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.cancelAndSend | cancelAndSendActionSalesCreditMemos
[**DeleteSalesCreditMemo**](SalesCreditMemoApi.md#DeleteSalesCreditMemo) | **Delete** /companies({company_id})/salesCreditMemos({salesCreditMemo_id}) | deleteSalesCreditMemo
[**GetSalesCreditMemo**](SalesCreditMemoApi.md#GetSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id}) | getSalesCreditMemo
[**ListSalesCreditMemos**](SalesCreditMemoApi.md#ListSalesCreditMemos) | **Get** /companies({company_id})/salesCreditMemos | listSalesCreditMemos
[**PatchSalesCreditMemo**](SalesCreditMemoApi.md#PatchSalesCreditMemo) | **Patch** /companies({company_id})/salesCreditMemos({salesCreditMemo_id}) | patchSalesCreditMemo
[**PostActionSalesCreditMemos**](SalesCreditMemoApi.md#PostActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.post | postActionSalesCreditMemos
[**PostAndSendActionSalesCreditMemos**](SalesCreditMemoApi.md#PostAndSendActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.postAndSend | postAndSendActionSalesCreditMemos
[**PostSalesCreditMemo**](SalesCreditMemoApi.md#PostSalesCreditMemo) | **Post** /companies({company_id})/salesCreditMemos | postSalesCreditMemo
[**SendActionSalesCreditMemos**](SalesCreditMemoApi.md#SendActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.send | sendActionSalesCreditMemos


# **CancelActionSalesCreditMemos**
> CancelActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)
cancelActionSalesCreditMemos

Performs the cancel action for salesCreditMemos entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelAndSendActionSalesCreditMemos**
> CancelAndSendActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)
cancelAndSendActionSalesCreditMemos

Performs the cancelAndSend action for salesCreditMemos entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesCreditMemo**
> DeleteSalesCreditMemo(ctx, companyId, salesCreditMemoId)
deleteSalesCreditMemo

Deletes an object of type salesCreditMemo in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesCreditMemo**
> SalesCreditMemo GetSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)
getSalesCreditMemo

Retrieve the properties and relationships of an object of type salesCreditMemo for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***GetSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesCreditMemoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesCreditMemo**](salesCreditMemo.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesCreditMemos**
> CompaniesApiSalesCreditMemosResponse ListSalesCreditMemos(ctx, companyId, optional)
listSalesCreditMemos

Returns a list of salesCreditMemos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesCreditMemosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesCreditMemosOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesCreditMemosResponse**](CompaniesAPISalesCreditMemosResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesCreditMemo**
> SalesCreditMemo PatchSalesCreditMemo(ctx, companyId, salesCreditMemoId, contentType, ifMatch, body)
patchSalesCreditMemo

Updates an object of type salesCreditMemo in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesCreditMemosApiRequest**](CompaniesApiSalesCreditMemosApiRequest.md)|  | 

### Return type

[**SalesCreditMemo**](salesCreditMemo.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostActionSalesCreditMemos**
> PostActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)
postActionSalesCreditMemos

Performs the post action for salesCreditMemos entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAndSendActionSalesCreditMemos**
> PostAndSendActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)
postAndSendActionSalesCreditMemos

Performs the postAndSend action for salesCreditMemos entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesCreditMemo**
> SalesCreditMemo PostSalesCreditMemo(ctx, companyId, contentType, body)
postSalesCreditMemo

Creates an object of type salesCreditMemo in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesCreditMemosRequest**](CompaniesApiSalesCreditMemosRequest.md)|  | 

### Return type

[**SalesCreditMemo**](salesCreditMemo.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendActionSalesCreditMemos**
> SendActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)
sendActionSalesCreditMemos

Performs the send action for salesCreditMemos entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

