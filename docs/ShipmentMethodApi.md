# \ShipmentMethodApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteShipmentMethod**](ShipmentMethodApi.md#DeleteShipmentMethod) | **Delete** /companies({company_id})/shipmentMethods({shipmentMethod_id}) | deleteShipmentMethod
[**GetShipmentMethod**](ShipmentMethodApi.md#GetShipmentMethod) | **Get** /companies({company_id})/shipmentMethods({shipmentMethod_id}) | getShipmentMethod
[**ListShipmentMethods**](ShipmentMethodApi.md#ListShipmentMethods) | **Get** /companies({company_id})/shipmentMethods | listShipmentMethods
[**PatchShipmentMethod**](ShipmentMethodApi.md#PatchShipmentMethod) | **Patch** /companies({company_id})/shipmentMethods({shipmentMethod_id}) | patchShipmentMethod
[**PostShipmentMethod**](ShipmentMethodApi.md#PostShipmentMethod) | **Post** /companies({company_id})/shipmentMethods | postShipmentMethod


# **DeleteShipmentMethod**
> DeleteShipmentMethod(ctx, companyId, shipmentMethodId)
deleteShipmentMethod

Deletes an object of type shipmentMethod in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **shipmentMethodId** | [**string**](.md)| id for shipmentMethod | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipmentMethod**
> ShipmentMethod GetShipmentMethod(ctx, companyId, shipmentMethodId, optional)
getShipmentMethod

Retrieve the properties and relationships of an object of type shipmentMethod for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **shipmentMethodId** | [**string**](.md)| id for shipmentMethod | 
 **optional** | ***GetShipmentMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetShipmentMethodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**ShipmentMethod**](shipmentMethod.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListShipmentMethods**
> CompaniesApiShipmentMethodsResponse ListShipmentMethods(ctx, companyId, optional)
listShipmentMethods

Returns a list of shipmentMethods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListShipmentMethodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListShipmentMethodsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiShipmentMethodsResponse**](CompaniesAPIShipmentMethodsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchShipmentMethod**
> ShipmentMethod PatchShipmentMethod(ctx, companyId, shipmentMethodId, contentType, ifMatch, body)
patchShipmentMethod

Updates an object of type shipmentMethod in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **shipmentMethodId** | [**string**](.md)| id for shipmentMethod | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**ShipmentMethod**](ShipmentMethod.md)|  | 

### Return type

[**ShipmentMethod**](shipmentMethod.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostShipmentMethod**
> ShipmentMethod PostShipmentMethod(ctx, companyId, contentType, body)
postShipmentMethod

Creates an object of type shipmentMethod in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**ShipmentMethod**](ShipmentMethod.md)|  | 

### Return type

[**ShipmentMethod**](shipmentMethod.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

