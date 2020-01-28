# \BankAccountApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBankAccount**](BankAccountApi.md#DeleteBankAccount) | **Delete** /companies({company_id})/bankAccounts({bankAccount_id}) | deleteBankAccount
[**GetBankAccount**](BankAccountApi.md#GetBankAccount) | **Get** /companies({company_id})/bankAccounts({bankAccount_id}) | getBankAccount
[**ListBankAccounts**](BankAccountApi.md#ListBankAccounts) | **Get** /companies({company_id})/bankAccounts | listBankAccounts
[**PatchBankAccount**](BankAccountApi.md#PatchBankAccount) | **Patch** /companies({company_id})/bankAccounts({bankAccount_id}) | patchBankAccount
[**PostBankAccount**](BankAccountApi.md#PostBankAccount) | **Post** /companies({company_id})/bankAccounts | postBankAccount


# **DeleteBankAccount**
> DeleteBankAccount(ctx, companyId, bankAccountId)
deleteBankAccount

Deletes an object of type bankAccount in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **bankAccountId** | [**string**](.md)| id for bankAccount | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAccount**
> BankAccount GetBankAccount(ctx, companyId, bankAccountId, optional)
getBankAccount

Retrieve the properties and relationships of an object of type bankAccount for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **bankAccountId** | [**string**](.md)| id for bankAccount | 
 **optional** | ***GetBankAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBankAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**BankAccount**](bankAccount.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBankAccounts**
> CompaniesApiBankAccountsResponse ListBankAccounts(ctx, companyId, optional)
listBankAccounts

Returns a list of bankAccounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListBankAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListBankAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiBankAccountsResponse**](CompaniesAPIBankAccountsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchBankAccount**
> BankAccount PatchBankAccount(ctx, companyId, bankAccountId, contentType, ifMatch, body)
patchBankAccount

Updates an object of type bankAccount in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **bankAccountId** | [**string**](.md)| id for bankAccount | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**BankAccount**](BankAccount.md)|  | 

### Return type

[**BankAccount**](bankAccount.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBankAccount**
> BankAccount PostBankAccount(ctx, companyId, contentType, body)
postBankAccount

Creates an object of type bankAccount in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**BankAccount**](BankAccount.md)|  | 

### Return type

[**BankAccount**](bankAccount.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

