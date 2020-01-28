# \CashFlowStatementApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCashFlowStatement**](CashFlowStatementApi.md#GetCashFlowStatement) | **Get** /companies({company_id})/cashFlowStatement({cashFlowStatement_lineNumber}) | getCashFlowStatement
[**ListCashFlowStatement**](CashFlowStatementApi.md#ListCashFlowStatement) | **Get** /companies({company_id})/cashFlowStatement | listCashFlowStatement


# **GetCashFlowStatement**
> CashFlowStatement GetCashFlowStatement(ctx, companyId, cashFlowStatementLineNumber, optional)
getCashFlowStatement

Retrieve the properties and relationships of an object of type cashFlowStatement for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **cashFlowStatementLineNumber** | **int32**| lineNumber for cashFlowStatement | 
 **optional** | ***GetCashFlowStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCashFlowStatementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CashFlowStatement**](cashFlowStatement.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCashFlowStatement**
> CompaniesApiCashFlowStatementResponse ListCashFlowStatement(ctx, companyId, optional)
listCashFlowStatement

Returns a list of cashFlowStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListCashFlowStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCashFlowStatementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCashFlowStatementResponse**](CompaniesAPICashFlowStatementResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

