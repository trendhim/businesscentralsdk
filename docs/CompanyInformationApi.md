# \CompanyInformationApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyInformation**](CompanyInformationApi.md#GetCompanyInformation) | **Get** /companies({company_id})/companyInformation({companyInformation_id}) | getCompanyInformation
[**ListCompanyInformation**](CompanyInformationApi.md#ListCompanyInformation) | **Get** /companies({company_id})/companyInformation | listCompanyInformation
[**PatchCompanyInformation**](CompanyInformationApi.md#PatchCompanyInformation) | **Patch** /companies({company_id})/companyInformation({companyInformation_id}) | patchCompanyInformation


# **GetCompanyInformation**
> CompanyInformation GetCompanyInformation(ctx, companyId, companyInformationId, optional)
getCompanyInformation

Retrieve the properties and relationships of an object of type companyInformation for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **companyInformationId** | [**string**](.md)| id for companyInformation | 
 **optional** | ***GetCompanyInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCompanyInformationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompanyInformation**](companyInformation.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCompanyInformation**
> CompaniesApiCompanyInformationResponse ListCompanyInformation(ctx, companyId, optional)
listCompanyInformation

Returns a list of companyInformation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListCompanyInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCompanyInformationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCompanyInformationResponse**](CompaniesAPICompanyInformationResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCompanyInformation**
> CompanyInformation PatchCompanyInformation(ctx, companyId, companyInformationId, contentType, ifMatch, body)
patchCompanyInformation

Updates an object of type companyInformation in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **companyInformationId** | [**string**](.md)| id for companyInformation | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompanyInformation**](CompanyInformation.md)|  | 

### Return type

[**CompanyInformation**](companyInformation.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

