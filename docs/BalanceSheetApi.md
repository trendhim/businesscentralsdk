# \BalanceSheetApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalanceSheet**](BalanceSheetApi.md#GetBalanceSheet) | **Get** /companies({company_id})/balanceSheet({balanceSheet_lineNumber}) | getBalanceSheet
[**ListBalanceSheet**](BalanceSheetApi.md#ListBalanceSheet) | **Get** /companies({company_id})/balanceSheet | listBalanceSheet


# **GetBalanceSheet**
> BalanceSheet GetBalanceSheet(ctx, companyId, balanceSheetLineNumber, optional)
getBalanceSheet

Retrieve the properties and relationships of an object of type balanceSheet for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **balanceSheetLineNumber** | **int32**| lineNumber for balanceSheet | 
 **optional** | ***GetBalanceSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBalanceSheetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**BalanceSheet**](balanceSheet.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBalanceSheet**
> CompaniesApiBalanceSheetResponse ListBalanceSheet(ctx, companyId, optional)
listBalanceSheet

Returns a list of balanceSheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListBalanceSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListBalanceSheetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiBalanceSheetResponse**](CompaniesAPIBalanceSheetResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

