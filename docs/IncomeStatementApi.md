# \IncomeStatementApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomeStatement**](IncomeStatementApi.md#GetIncomeStatement) | **Get** /companies({company_id})/incomeStatement({incomeStatement_lineNumber}) | getIncomeStatement
[**ListIncomeStatement**](IncomeStatementApi.md#ListIncomeStatement) | **Get** /companies({company_id})/incomeStatement | listIncomeStatement


# **GetIncomeStatement**
> IncomeStatement GetIncomeStatement(ctx, companyId, incomeStatementLineNumber, optional)
getIncomeStatement

Retrieve the properties and relationships of an object of type incomeStatement for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **incomeStatementLineNumber** | **int32**| lineNumber for incomeStatement | 
 **optional** | ***GetIncomeStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetIncomeStatementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**IncomeStatement**](incomeStatement.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIncomeStatement**
> CompaniesApiIncomeStatementResponse ListIncomeStatement(ctx, companyId, optional)
listIncomeStatement

Returns a list of incomeStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListIncomeStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListIncomeStatementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiIncomeStatementResponse**](CompaniesAPIIncomeStatementResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

