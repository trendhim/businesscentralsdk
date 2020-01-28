# \DefaultDimensionsApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDefaultDimensions**](DefaultDimensionsApi.md#DeleteDefaultDimensions) | **Delete** /companies({company_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | deleteDefaultDimensions
[**DeleteDefaultDimensionsForCustomer**](DefaultDimensionsApi.md#DeleteDefaultDimensionsForCustomer) | **Delete** /companies({company_id})/customers({customer_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | deleteDefaultDimensionsForCustomer
[**DeleteDefaultDimensionsForEmployee**](DefaultDimensionsApi.md#DeleteDefaultDimensionsForEmployee) | **Delete** /companies({company_id})/employees({employee_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | deleteDefaultDimensionsForEmployee
[**DeleteDefaultDimensionsForItem**](DefaultDimensionsApi.md#DeleteDefaultDimensionsForItem) | **Delete** /companies({company_id})/items({item_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | deleteDefaultDimensionsForItem
[**DeleteDefaultDimensionsForVendor**](DefaultDimensionsApi.md#DeleteDefaultDimensionsForVendor) | **Delete** /companies({company_id})/vendors({vendor_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | deleteDefaultDimensionsForVendor
[**GetDefaultDimensions**](DefaultDimensionsApi.md#GetDefaultDimensions) | **Get** /companies({company_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | getDefaultDimensions
[**GetDefaultDimensionsForCustomer**](DefaultDimensionsApi.md#GetDefaultDimensionsForCustomer) | **Get** /companies({company_id})/customers({customer_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | getDefaultDimensionsForCustomer
[**GetDefaultDimensionsForEmployee**](DefaultDimensionsApi.md#GetDefaultDimensionsForEmployee) | **Get** /companies({company_id})/employees({employee_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | getDefaultDimensionsForEmployee
[**GetDefaultDimensionsForItem**](DefaultDimensionsApi.md#GetDefaultDimensionsForItem) | **Get** /companies({company_id})/items({item_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | getDefaultDimensionsForItem
[**GetDefaultDimensionsForVendor**](DefaultDimensionsApi.md#GetDefaultDimensionsForVendor) | **Get** /companies({company_id})/vendors({vendor_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | getDefaultDimensionsForVendor
[**ListDefaultDimensions**](DefaultDimensionsApi.md#ListDefaultDimensions) | **Get** /companies({company_id})/defaultDimensions | listDefaultDimensions
[**ListDefaultDimensionsForCustomer**](DefaultDimensionsApi.md#ListDefaultDimensionsForCustomer) | **Get** /companies({company_id})/customers({customer_id})/defaultDimensions | listDefaultDimensionsForCustomer
[**ListDefaultDimensionsForEmployee**](DefaultDimensionsApi.md#ListDefaultDimensionsForEmployee) | **Get** /companies({company_id})/employees({employee_id})/defaultDimensions | listDefaultDimensionsForEmployee
[**ListDefaultDimensionsForItem**](DefaultDimensionsApi.md#ListDefaultDimensionsForItem) | **Get** /companies({company_id})/items({item_id})/defaultDimensions | listDefaultDimensionsForItem
[**ListDefaultDimensionsForVendor**](DefaultDimensionsApi.md#ListDefaultDimensionsForVendor) | **Get** /companies({company_id})/vendors({vendor_id})/defaultDimensions | listDefaultDimensionsForVendor
[**PatchDefaultDimensions**](DefaultDimensionsApi.md#PatchDefaultDimensions) | **Patch** /companies({company_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | patchDefaultDimensions
[**PatchDefaultDimensionsForCustomer**](DefaultDimensionsApi.md#PatchDefaultDimensionsForCustomer) | **Patch** /companies({company_id})/customers({customer_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | patchDefaultDimensionsForCustomer
[**PatchDefaultDimensionsForEmployee**](DefaultDimensionsApi.md#PatchDefaultDimensionsForEmployee) | **Patch** /companies({company_id})/employees({employee_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | patchDefaultDimensionsForEmployee
[**PatchDefaultDimensionsForItem**](DefaultDimensionsApi.md#PatchDefaultDimensionsForItem) | **Patch** /companies({company_id})/items({item_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | patchDefaultDimensionsForItem
[**PatchDefaultDimensionsForVendor**](DefaultDimensionsApi.md#PatchDefaultDimensionsForVendor) | **Patch** /companies({company_id})/vendors({vendor_id})/defaultDimensions({defaultDimensions_parentId},{defaultDimensions_dimensionId}) | patchDefaultDimensionsForVendor
[**PostDefaultDimensions**](DefaultDimensionsApi.md#PostDefaultDimensions) | **Post** /companies({company_id})/defaultDimensions | postDefaultDimensions
[**PostDefaultDimensionsForCustomer**](DefaultDimensionsApi.md#PostDefaultDimensionsForCustomer) | **Post** /companies({company_id})/customers({customer_id})/defaultDimensions | postDefaultDimensionsForCustomer
[**PostDefaultDimensionsForEmployee**](DefaultDimensionsApi.md#PostDefaultDimensionsForEmployee) | **Post** /companies({company_id})/employees({employee_id})/defaultDimensions | postDefaultDimensionsForEmployee
[**PostDefaultDimensionsForItem**](DefaultDimensionsApi.md#PostDefaultDimensionsForItem) | **Post** /companies({company_id})/items({item_id})/defaultDimensions | postDefaultDimensionsForItem
[**PostDefaultDimensionsForVendor**](DefaultDimensionsApi.md#PostDefaultDimensionsForVendor) | **Post** /companies({company_id})/vendors({vendor_id})/defaultDimensions | postDefaultDimensionsForVendor


# **DeleteDefaultDimensions**
> DeleteDefaultDimensions(ctx, companyId, defaultDimensionsParentId, defaultDimensionsDimensionId)
deleteDefaultDimensions

Deletes an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefaultDimensionsForCustomer**
> DeleteDefaultDimensionsForCustomer(ctx, companyId, customerId, defaultDimensionsParentId, defaultDimensionsDimensionId)
deleteDefaultDimensionsForCustomer

Deletes an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefaultDimensionsForEmployee**
> DeleteDefaultDimensionsForEmployee(ctx, companyId, employeeId, defaultDimensionsParentId, defaultDimensionsDimensionId)
deleteDefaultDimensionsForEmployee

Deletes an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefaultDimensionsForItem**
> DeleteDefaultDimensionsForItem(ctx, companyId, itemId, defaultDimensionsParentId, defaultDimensionsDimensionId)
deleteDefaultDimensionsForItem

Deletes an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefaultDimensionsForVendor**
> DeleteDefaultDimensionsForVendor(ctx, companyId, vendorId, defaultDimensionsParentId, defaultDimensionsDimensionId)
deleteDefaultDimensionsForVendor

Deletes an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultDimensions**
> DefaultDimensions GetDefaultDimensions(ctx, companyId, defaultDimensionsParentId, defaultDimensionsDimensionId, optional)
getDefaultDimensions

Retrieve the properties and relationships of an object of type defaultDimensions for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
 **optional** | ***GetDefaultDimensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDefaultDimensionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultDimensionsForCustomer**
> DefaultDimensions GetDefaultDimensionsForCustomer(ctx, companyId, customerId, defaultDimensionsParentId, defaultDimensionsDimensionId, optional)
getDefaultDimensionsForCustomer

Retrieve the properties and relationships of an object of type defaultDimensions for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
 **optional** | ***GetDefaultDimensionsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDefaultDimensionsForCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultDimensionsForEmployee**
> DefaultDimensions GetDefaultDimensionsForEmployee(ctx, companyId, employeeId, defaultDimensionsParentId, defaultDimensionsDimensionId, optional)
getDefaultDimensionsForEmployee

Retrieve the properties and relationships of an object of type defaultDimensions for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
 **optional** | ***GetDefaultDimensionsForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDefaultDimensionsForEmployeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultDimensionsForItem**
> DefaultDimensions GetDefaultDimensionsForItem(ctx, companyId, itemId, defaultDimensionsParentId, defaultDimensionsDimensionId, optional)
getDefaultDimensionsForItem

Retrieve the properties and relationships of an object of type defaultDimensions for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
 **optional** | ***GetDefaultDimensionsForItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDefaultDimensionsForItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultDimensionsForVendor**
> DefaultDimensions GetDefaultDimensionsForVendor(ctx, companyId, vendorId, defaultDimensionsParentId, defaultDimensionsDimensionId, optional)
getDefaultDimensionsForVendor

Retrieve the properties and relationships of an object of type defaultDimensions for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
 **optional** | ***GetDefaultDimensionsForVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDefaultDimensionsForVendorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDefaultDimensions**
> CompaniesApiDefaultDimensionsResponse ListDefaultDimensions(ctx, companyId, optional)
listDefaultDimensions

Returns a list of defaultDimensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListDefaultDimensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDefaultDimensionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiDefaultDimensionsResponse**](CompaniesAPIDefaultDimensionsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDefaultDimensionsForCustomer**
> CompaniesApiCustomersApiDefaultDimensionsResponse ListDefaultDimensionsForCustomer(ctx, companyId, customerId, optional)
listDefaultDimensionsForCustomer

Returns a list of defaultDimensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
 **optional** | ***ListDefaultDimensionsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDefaultDimensionsForCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCustomersApiDefaultDimensionsResponse**](CompaniesAPICustomersAPIDefaultDimensionsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDefaultDimensionsForEmployee**
> CompaniesApiEmployeesApiDefaultDimensionsResponse ListDefaultDimensionsForEmployee(ctx, companyId, employeeId, optional)
listDefaultDimensionsForEmployee

Returns a list of defaultDimensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
 **optional** | ***ListDefaultDimensionsForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDefaultDimensionsForEmployeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiEmployeesApiDefaultDimensionsResponse**](CompaniesAPIEmployeesAPIDefaultDimensionsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDefaultDimensionsForItem**
> CompaniesApiItemsApiDefaultDimensionsResponse ListDefaultDimensionsForItem(ctx, companyId, itemId, optional)
listDefaultDimensionsForItem

Returns a list of defaultDimensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
 **optional** | ***ListDefaultDimensionsForItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDefaultDimensionsForItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiItemsApiDefaultDimensionsResponse**](CompaniesAPIItemsAPIDefaultDimensionsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDefaultDimensionsForVendor**
> CompaniesApiVendorsApiDefaultDimensionsResponse ListDefaultDimensionsForVendor(ctx, companyId, vendorId, optional)
listDefaultDimensionsForVendor

Returns a list of defaultDimensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
 **optional** | ***ListDefaultDimensionsForVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDefaultDimensionsForVendorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiVendorsApiDefaultDimensionsResponse**](CompaniesAPIVendorsAPIDefaultDimensionsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDefaultDimensions**
> DefaultDimensions PatchDefaultDimensions(ctx, companyId, defaultDimensionsParentId, defaultDimensionsDimensionId, contentType, ifMatch, body)
patchDefaultDimensions

Updates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiDefaultDimensionsDefaultDimensionsDimensionIdApiRequest**](CompaniesApiDefaultDimensionsDefaultDimensionsDimensionIdApiRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDefaultDimensionsForCustomer**
> DefaultDimensions PatchDefaultDimensionsForCustomer(ctx, companyId, customerId, defaultDimensionsParentId, defaultDimensionsDimensionId, contentType, ifMatch, body)
patchDefaultDimensionsForCustomer

Updates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiCustomersApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest**](CompaniesApiCustomersApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDefaultDimensionsForEmployee**
> DefaultDimensions PatchDefaultDimensionsForEmployee(ctx, companyId, employeeId, defaultDimensionsParentId, defaultDimensionsDimensionId, contentType, ifMatch, body)
patchDefaultDimensionsForEmployee

Updates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiEmployeesApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest**](CompaniesApiEmployeesApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDefaultDimensionsForItem**
> DefaultDimensions PatchDefaultDimensionsForItem(ctx, companyId, itemId, defaultDimensionsParentId, defaultDimensionsDimensionId, contentType, ifMatch, body)
patchDefaultDimensionsForItem

Updates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiItemsApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest**](CompaniesApiItemsApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDefaultDimensionsForVendor**
> DefaultDimensions PatchDefaultDimensionsForVendor(ctx, companyId, vendorId, defaultDimensionsParentId, defaultDimensionsDimensionId, contentType, ifMatch, body)
patchDefaultDimensionsForVendor

Updates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **defaultDimensionsParentId** | [**string**](.md)| parentId for defaultDimensions | 
  **defaultDimensionsDimensionId** | [**string**](.md)| dimensionId for defaultDimensions | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiVendorsApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest**](CompaniesApiVendorsApiDefaultDimensionsDefaultDimensionsParentIdDefaultDimensionsDimensionIdApiRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDefaultDimensions**
> DefaultDimensions PostDefaultDimensions(ctx, companyId, contentType, body)
postDefaultDimensions

Creates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiDefaultDimensionsRequest**](CompaniesApiDefaultDimensionsRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDefaultDimensionsForCustomer**
> DefaultDimensions PostDefaultDimensionsForCustomer(ctx, companyId, customerId, contentType, body)
postDefaultDimensionsForCustomer

Creates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiCustomersApiDefaultDimensionsRequest**](CompaniesApiCustomersApiDefaultDimensionsRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDefaultDimensionsForEmployee**
> DefaultDimensions PostDefaultDimensionsForEmployee(ctx, companyId, employeeId, contentType, body)
postDefaultDimensionsForEmployee

Creates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiEmployeesApiDefaultDimensionsRequest**](CompaniesApiEmployeesApiDefaultDimensionsRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDefaultDimensionsForItem**
> DefaultDimensions PostDefaultDimensionsForItem(ctx, companyId, itemId, contentType, body)
postDefaultDimensionsForItem

Creates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiItemsApiDefaultDimensionsRequest**](CompaniesApiItemsApiDefaultDimensionsRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDefaultDimensionsForVendor**
> DefaultDimensions PostDefaultDimensionsForVendor(ctx, companyId, vendorId, contentType, body)
postDefaultDimensionsForVendor

Creates an object of type defaultDimensions in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiVendorsApiDefaultDimensionsRequest**](CompaniesApiVendorsApiDefaultDimensionsRequest.md)|  | 

### Return type

[**DefaultDimensions**](defaultDimensions.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

