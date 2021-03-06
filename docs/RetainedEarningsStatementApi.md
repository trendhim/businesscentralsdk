# \RetainedEarningsStatementApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetainedEarningsStatement**](RetainedEarningsStatementApi.md#GetRetainedEarningsStatement) | **Get** /companies({company_id})/retainedEarningsStatement({retainedEarningsStatement_lineNumber}) | getRetainedEarningsStatement
[**ListRetainedEarningsStatement**](RetainedEarningsStatementApi.md#ListRetainedEarningsStatement) | **Get** /companies({company_id})/retainedEarningsStatement | listRetainedEarningsStatement


# **GetRetainedEarningsStatement**
> RetainedEarningsStatement GetRetainedEarningsStatement(ctx, companyId, retainedEarningsStatementLineNumber, optional)
getRetainedEarningsStatement

Retrieve the properties and relationships of an object of type retainedEarningsStatement for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **retainedEarningsStatementLineNumber** | **int32**| lineNumber for retainedEarningsStatement | 
 **optional** | ***GetRetainedEarningsStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRetainedEarningsStatementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**RetainedEarningsStatement**](retainedEarningsStatement.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRetainedEarningsStatement**
> CompaniesApiRetainedEarningsStatementResponse ListRetainedEarningsStatement(ctx, companyId, optional)
listRetainedEarningsStatement

Returns a list of retainedEarningsStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListRetainedEarningsStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListRetainedEarningsStatementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiRetainedEarningsStatementResponse**](CompaniesAPIRetainedEarningsStatementResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

