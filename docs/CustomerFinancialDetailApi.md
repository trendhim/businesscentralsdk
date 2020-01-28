# \CustomerFinancialDetailApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerFinancialDetail**](CustomerFinancialDetailApi.md#GetCustomerFinancialDetail) | **Get** /companies({company_id})/customerFinancialDetails({customerFinancialDetail_id}) | getCustomerFinancialDetail
[**GetCustomerFinancialDetailForCustomer**](CustomerFinancialDetailApi.md#GetCustomerFinancialDetailForCustomer) | **Get** /companies({company_id})/customers({customer_id})/customerFinancialDetails({customerFinancialDetail_id}) | getCustomerFinancialDetailForCustomer
[**ListCustomerFinancialDetails**](CustomerFinancialDetailApi.md#ListCustomerFinancialDetails) | **Get** /companies({company_id})/customerFinancialDetails | listCustomerFinancialDetails
[**ListCustomerFinancialDetailsForCustomer**](CustomerFinancialDetailApi.md#ListCustomerFinancialDetailsForCustomer) | **Get** /companies({company_id})/customers({customer_id})/customerFinancialDetails | listCustomerFinancialDetailsForCustomer


# **GetCustomerFinancialDetail**
> CustomerFinancialDetail GetCustomerFinancialDetail(ctx, companyId, customerFinancialDetailId, optional)
getCustomerFinancialDetail

Retrieve the properties and relationships of an object of type customerFinancialDetail for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerFinancialDetailId** | [**string**](.md)| id for customerFinancialDetail | 
 **optional** | ***GetCustomerFinancialDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerFinancialDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerFinancialDetail**](customerFinancialDetail.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerFinancialDetailForCustomer**
> CustomerFinancialDetail GetCustomerFinancialDetailForCustomer(ctx, companyId, customerId, customerFinancialDetailId, optional)
getCustomerFinancialDetailForCustomer

Retrieve the properties and relationships of an object of type customerFinancialDetail for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **customerFinancialDetailId** | [**string**](.md)| id for customerFinancialDetail | 
 **optional** | ***GetCustomerFinancialDetailForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerFinancialDetailForCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerFinancialDetail**](customerFinancialDetail.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerFinancialDetails**
> CompaniesApiCustomerFinancialDetailsResponse ListCustomerFinancialDetails(ctx, companyId, optional)
listCustomerFinancialDetails

Returns a list of customerFinancialDetails

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListCustomerFinancialDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCustomerFinancialDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCustomerFinancialDetailsResponse**](CompaniesAPICustomerFinancialDetailsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerFinancialDetailsForCustomer**
> CompaniesApiCustomersApiCustomerFinancialDetailsResponse ListCustomerFinancialDetailsForCustomer(ctx, companyId, customerId, optional)
listCustomerFinancialDetailsForCustomer

Returns a list of customerFinancialDetails

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
 **optional** | ***ListCustomerFinancialDetailsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCustomerFinancialDetailsForCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCustomersApiCustomerFinancialDetailsResponse**](CompaniesAPICustomersAPICustomerFinancialDetailsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

