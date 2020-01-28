# \AgedAccountsReceivableApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgedAccountsReceivable**](AgedAccountsReceivableApi.md#GetAgedAccountsReceivable) | **Get** /companies({company_id})/agedAccountsReceivable({agedAccountsReceivable_customerId}) | getAgedAccountsReceivable
[**ListAgedAccountsReceivable**](AgedAccountsReceivableApi.md#ListAgedAccountsReceivable) | **Get** /companies({company_id})/agedAccountsReceivable | listAgedAccountsReceivable


# **GetAgedAccountsReceivable**
> AgedAccountsReceivable GetAgedAccountsReceivable(ctx, companyId, agedAccountsReceivableCustomerId, optional)
getAgedAccountsReceivable

Retrieve the properties and relationships of an object of type agedAccountsReceivable for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **agedAccountsReceivableCustomerId** | [**string**](.md)| customerId for agedAccountsReceivable | 
 **optional** | ***GetAgedAccountsReceivableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgedAccountsReceivableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**AgedAccountsReceivable**](agedAccountsReceivable.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAgedAccountsReceivable**
> CompaniesApiAgedAccountsReceivableResponse ListAgedAccountsReceivable(ctx, companyId, optional)
listAgedAccountsReceivable

Returns a list of agedAccountsReceivable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListAgedAccountsReceivableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAgedAccountsReceivableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiAgedAccountsReceivableResponse**](CompaniesAPIAgedAccountsReceivableResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

