# \CountryRegionApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCountryRegion**](CountryRegionApi.md#DeleteCountryRegion) | **Delete** /companies({company_id})/countriesRegions({countryRegion_id}) | deleteCountryRegion
[**GetCountryRegion**](CountryRegionApi.md#GetCountryRegion) | **Get** /companies({company_id})/countriesRegions({countryRegion_id}) | getCountryRegion
[**ListCountriesRegions**](CountryRegionApi.md#ListCountriesRegions) | **Get** /companies({company_id})/countriesRegions | listCountriesRegions
[**PatchCountryRegion**](CountryRegionApi.md#PatchCountryRegion) | **Patch** /companies({company_id})/countriesRegions({countryRegion_id}) | patchCountryRegion
[**PostCountryRegion**](CountryRegionApi.md#PostCountryRegion) | **Post** /companies({company_id})/countriesRegions | postCountryRegion


# **DeleteCountryRegion**
> DeleteCountryRegion(ctx, companyId, countryRegionId)
deleteCountryRegion

Deletes an object of type countryRegion in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **countryRegionId** | [**string**](.md)| id for countryRegion | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryRegion**
> CountryRegion GetCountryRegion(ctx, companyId, countryRegionId, optional)
getCountryRegion

Retrieve the properties and relationships of an object of type countryRegion for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **countryRegionId** | [**string**](.md)| id for countryRegion | 
 **optional** | ***GetCountryRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCountryRegionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CountryRegion**](countryRegion.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCountriesRegions**
> CompaniesApiCountriesRegionsResponse ListCountriesRegions(ctx, companyId, optional)
listCountriesRegions

Returns a list of countriesRegions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListCountriesRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCountriesRegionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCountriesRegionsResponse**](CompaniesAPICountriesRegionsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCountryRegion**
> CountryRegion PatchCountryRegion(ctx, companyId, countryRegionId, contentType, ifMatch, body)
patchCountryRegion

Updates an object of type countryRegion in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **countryRegionId** | [**string**](.md)| id for countryRegion | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CountryRegion**](CountryRegion.md)|  | 

### Return type

[**CountryRegion**](countryRegion.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCountryRegion**
> CountryRegion PostCountryRegion(ctx, companyId, contentType, body)
postCountryRegion

Creates an object of type countryRegion in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CountryRegion**](CountryRegion.md)|  | 

### Return type

[**CountryRegion**](countryRegion.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

