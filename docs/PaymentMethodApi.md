# \PaymentMethodApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePaymentMethod**](PaymentMethodApi.md#DeletePaymentMethod) | **Delete** /companies({company_id})/paymentMethods({paymentMethod_id}) | deletePaymentMethod
[**GetPaymentMethod**](PaymentMethodApi.md#GetPaymentMethod) | **Get** /companies({company_id})/paymentMethods({paymentMethod_id}) | getPaymentMethod
[**ListPaymentMethods**](PaymentMethodApi.md#ListPaymentMethods) | **Get** /companies({company_id})/paymentMethods | listPaymentMethods
[**PatchPaymentMethod**](PaymentMethodApi.md#PatchPaymentMethod) | **Patch** /companies({company_id})/paymentMethods({paymentMethod_id}) | patchPaymentMethod
[**PostPaymentMethod**](PaymentMethodApi.md#PostPaymentMethod) | **Post** /companies({company_id})/paymentMethods | postPaymentMethod


# **DeletePaymentMethod**
> DeletePaymentMethod(ctx, companyId, paymentMethodId)
deletePaymentMethod

Deletes an object of type paymentMethod in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentMethodId** | [**string**](.md)| id for paymentMethod | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethod**
> PaymentMethod GetPaymentMethod(ctx, companyId, paymentMethodId, optional)
getPaymentMethod

Retrieve the properties and relationships of an object of type paymentMethod for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentMethodId** | [**string**](.md)| id for paymentMethod | 
 **optional** | ***GetPaymentMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPaymentMethodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PaymentMethod**](paymentMethod.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentMethods**
> CompaniesApiPaymentMethodsResponse ListPaymentMethods(ctx, companyId, optional)
listPaymentMethods

Returns a list of paymentMethods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPaymentMethodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPaymentMethodsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPaymentMethodsResponse**](CompaniesAPIPaymentMethodsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPaymentMethod**
> PaymentMethod PatchPaymentMethod(ctx, companyId, paymentMethodId, contentType, ifMatch, body)
patchPaymentMethod

Updates an object of type paymentMethod in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentMethodId** | [**string**](.md)| id for paymentMethod | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**PaymentMethod**](PaymentMethod.md)|  | 

### Return type

[**PaymentMethod**](paymentMethod.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPaymentMethod**
> PaymentMethod PostPaymentMethod(ctx, companyId, contentType, body)
postPaymentMethod

Creates an object of type paymentMethod in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**PaymentMethod**](PaymentMethod.md)|  | 

### Return type

[**PaymentMethod**](paymentMethod.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

