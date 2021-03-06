# \CustomerSaleApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerSale**](CustomerSaleApi.md#GetCustomerSale) | **Get** /companies({company_id})/customerSales({customerSale_customerId},&#39;{customerSale_customerNumber}&#39;,&#39;{customerSale_name}&#39;) | getCustomerSale
[**ListCustomerSales**](CustomerSaleApi.md#ListCustomerSales) | **Get** /companies({company_id})/customerSales | listCustomerSales


# **GetCustomerSale**
> CustomerSale GetCustomerSale(ctx, companyId, customerSaleCustomerId, customerSaleCustomerNumber, customerSaleName, optional)
getCustomerSale

Retrieve the properties and relationships of an object of type customerSale for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerSaleCustomerId** | [**string**](.md)| customerId for customerSale | 
  **customerSaleCustomerNumber** | **string**| customerNumber for customerSale | 
  **customerSaleName** | **string**| name for customerSale | 
 **optional** | ***GetCustomerSaleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerSaleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerSale**](customerSale.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerSales**
> CompaniesApiCustomerSalesResponse ListCustomerSales(ctx, companyId, optional)
listCustomerSales

Returns a list of customerSales

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListCustomerSalesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCustomerSalesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCustomerSalesResponse**](CompaniesAPICustomerSalesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

