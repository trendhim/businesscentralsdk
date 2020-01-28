# \UnitOfMeasureApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUnitOfMeasure**](UnitOfMeasureApi.md#DeleteUnitOfMeasure) | **Delete** /companies({company_id})/unitsOfMeasure({unitOfMeasure_id}) | deleteUnitOfMeasure
[**GetUnitOfMeasure**](UnitOfMeasureApi.md#GetUnitOfMeasure) | **Get** /companies({company_id})/unitsOfMeasure({unitOfMeasure_id}) | getUnitOfMeasure
[**ListUnitsOfMeasure**](UnitOfMeasureApi.md#ListUnitsOfMeasure) | **Get** /companies({company_id})/unitsOfMeasure | listUnitsOfMeasure
[**PatchUnitOfMeasure**](UnitOfMeasureApi.md#PatchUnitOfMeasure) | **Patch** /companies({company_id})/unitsOfMeasure({unitOfMeasure_id}) | patchUnitOfMeasure
[**PostUnitOfMeasure**](UnitOfMeasureApi.md#PostUnitOfMeasure) | **Post** /companies({company_id})/unitsOfMeasure | postUnitOfMeasure


# **DeleteUnitOfMeasure**
> DeleteUnitOfMeasure(ctx, companyId, unitOfMeasureId)
deleteUnitOfMeasure

Deletes an object of type unitOfMeasure in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **unitOfMeasureId** | [**string**](.md)| id for unitOfMeasure | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnitOfMeasure**
> UnitOfMeasure GetUnitOfMeasure(ctx, companyId, unitOfMeasureId, optional)
getUnitOfMeasure

Retrieve the properties and relationships of an object of type unitOfMeasure for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **unitOfMeasureId** | [**string**](.md)| id for unitOfMeasure | 
 **optional** | ***GetUnitOfMeasureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUnitOfMeasureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**UnitOfMeasure**](unitOfMeasure.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUnitsOfMeasure**
> CompaniesApiUnitsOfMeasureResponse ListUnitsOfMeasure(ctx, companyId, optional)
listUnitsOfMeasure

Returns a list of unitsOfMeasure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListUnitsOfMeasureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListUnitsOfMeasureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiUnitsOfMeasureResponse**](CompaniesAPIUnitsOfMeasureResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUnitOfMeasure**
> UnitOfMeasure PatchUnitOfMeasure(ctx, companyId, unitOfMeasureId, contentType, ifMatch, body)
patchUnitOfMeasure

Updates an object of type unitOfMeasure in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **unitOfMeasureId** | [**string**](.md)| id for unitOfMeasure | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**UnitOfMeasure**](UnitOfMeasure.md)|  | 

### Return type

[**UnitOfMeasure**](unitOfMeasure.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUnitOfMeasure**
> UnitOfMeasure PostUnitOfMeasure(ctx, companyId, contentType, body)
postUnitOfMeasure

Creates an object of type unitOfMeasure in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**UnitOfMeasure**](UnitOfMeasure.md)|  | 

### Return type

[**UnitOfMeasure**](unitOfMeasure.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

