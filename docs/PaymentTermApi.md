# \PaymentTermApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePaymentTerm**](PaymentTermApi.md#DeletePaymentTerm) | **Delete** /companies({company_id})/paymentTerms({paymentTerm_id}) | deletePaymentTerm
[**GetPaymentTerm**](PaymentTermApi.md#GetPaymentTerm) | **Get** /companies({company_id})/paymentTerms({paymentTerm_id}) | getPaymentTerm
[**ListPaymentTerms**](PaymentTermApi.md#ListPaymentTerms) | **Get** /companies({company_id})/paymentTerms | listPaymentTerms
[**PatchPaymentTerm**](PaymentTermApi.md#PatchPaymentTerm) | **Patch** /companies({company_id})/paymentTerms({paymentTerm_id}) | patchPaymentTerm
[**PostPaymentTerm**](PaymentTermApi.md#PostPaymentTerm) | **Post** /companies({company_id})/paymentTerms | postPaymentTerm


# **DeletePaymentTerm**
> DeletePaymentTerm(ctx, companyId, paymentTermId)
deletePaymentTerm

Deletes an object of type paymentTerm in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentTermId** | [**string**](.md)| id for paymentTerm | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentTerm**
> PaymentTerm GetPaymentTerm(ctx, companyId, paymentTermId, optional)
getPaymentTerm

Retrieve the properties and relationships of an object of type paymentTerm for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentTermId** | [**string**](.md)| id for paymentTerm | 
 **optional** | ***GetPaymentTermOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPaymentTermOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PaymentTerm**](paymentTerm.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentTerms**
> CompaniesApiPaymentTermsResponse ListPaymentTerms(ctx, companyId, optional)
listPaymentTerms

Returns a list of paymentTerms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPaymentTermsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPaymentTermsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPaymentTermsResponse**](CompaniesAPIPaymentTermsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPaymentTerm**
> PaymentTerm PatchPaymentTerm(ctx, companyId, paymentTermId, contentType, ifMatch, body)
patchPaymentTerm

Updates an object of type paymentTerm in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentTermId** | [**string**](.md)| id for paymentTerm | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**PaymentTerm**](PaymentTerm.md)|  | 

### Return type

[**PaymentTerm**](paymentTerm.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPaymentTerm**
> PaymentTerm PostPaymentTerm(ctx, companyId, contentType, body)
postPaymentTerm

Creates an object of type paymentTerm in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**PaymentTerm**](PaymentTerm.md)|  | 

### Return type

[**PaymentTerm**](paymentTerm.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

