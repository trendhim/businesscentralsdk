# \TimeRegistrationEntryApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeRegistrationEntry**](TimeRegistrationEntryApi.md#DeleteTimeRegistrationEntry) | **Delete** /companies({company_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | deleteTimeRegistrationEntry
[**DeleteTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#DeleteTimeRegistrationEntryForEmployee) | **Delete** /companies({company_id})/employees({employee_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | deleteTimeRegistrationEntryForEmployee
[**GetTimeRegistrationEntry**](TimeRegistrationEntryApi.md#GetTimeRegistrationEntry) | **Get** /companies({company_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | getTimeRegistrationEntry
[**GetTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#GetTimeRegistrationEntryForEmployee) | **Get** /companies({company_id})/employees({employee_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | getTimeRegistrationEntryForEmployee
[**ListTimeRegistrationEntries**](TimeRegistrationEntryApi.md#ListTimeRegistrationEntries) | **Get** /companies({company_id})/timeRegistrationEntries | listTimeRegistrationEntries
[**ListTimeRegistrationEntriesForEmployee**](TimeRegistrationEntryApi.md#ListTimeRegistrationEntriesForEmployee) | **Get** /companies({company_id})/employees({employee_id})/timeRegistrationEntries | listTimeRegistrationEntriesForEmployee
[**PatchTimeRegistrationEntry**](TimeRegistrationEntryApi.md#PatchTimeRegistrationEntry) | **Patch** /companies({company_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | patchTimeRegistrationEntry
[**PatchTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#PatchTimeRegistrationEntryForEmployee) | **Patch** /companies({company_id})/employees({employee_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | patchTimeRegistrationEntryForEmployee
[**PostTimeRegistrationEntry**](TimeRegistrationEntryApi.md#PostTimeRegistrationEntry) | **Post** /companies({company_id})/timeRegistrationEntries | postTimeRegistrationEntry
[**PostTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#PostTimeRegistrationEntryForEmployee) | **Post** /companies({company_id})/employees({employee_id})/timeRegistrationEntries | postTimeRegistrationEntryForEmployee


# **DeleteTimeRegistrationEntry**
> DeleteTimeRegistrationEntry(ctx, companyId, timeRegistrationEntryId)
deleteTimeRegistrationEntry

Deletes an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTimeRegistrationEntryForEmployee**
> DeleteTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, timeRegistrationEntryId)
deleteTimeRegistrationEntryForEmployee

Deletes an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeRegistrationEntry**
> TimeRegistrationEntry GetTimeRegistrationEntry(ctx, companyId, timeRegistrationEntryId, optional)
getTimeRegistrationEntry

Retrieve the properties and relationships of an object of type timeRegistrationEntry for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
 **optional** | ***GetTimeRegistrationEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTimeRegistrationEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeRegistrationEntryForEmployee**
> TimeRegistrationEntry GetTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, timeRegistrationEntryId, optional)
getTimeRegistrationEntryForEmployee

Retrieve the properties and relationships of an object of type timeRegistrationEntry for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
 **optional** | ***GetTimeRegistrationEntryForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTimeRegistrationEntryForEmployeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimeRegistrationEntries**
> CompaniesApiTimeRegistrationEntriesResponse ListTimeRegistrationEntries(ctx, companyId, optional)
listTimeRegistrationEntries

Returns a list of timeRegistrationEntries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListTimeRegistrationEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTimeRegistrationEntriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiTimeRegistrationEntriesResponse**](CompaniesAPITimeRegistrationEntriesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimeRegistrationEntriesForEmployee**
> CompaniesApiEmployeesApiTimeRegistrationEntriesResponse ListTimeRegistrationEntriesForEmployee(ctx, companyId, employeeId, optional)
listTimeRegistrationEntriesForEmployee

Returns a list of timeRegistrationEntries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
 **optional** | ***ListTimeRegistrationEntriesForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTimeRegistrationEntriesForEmployeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiEmployeesApiTimeRegistrationEntriesResponse**](CompaniesAPIEmployeesAPITimeRegistrationEntriesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTimeRegistrationEntry**
> TimeRegistrationEntry PatchTimeRegistrationEntry(ctx, companyId, timeRegistrationEntryId, contentType, ifMatch, body)
patchTimeRegistrationEntry

Updates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiTimeRegistrationEntriesApiRequest**](CompaniesApiTimeRegistrationEntriesApiRequest.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTimeRegistrationEntryForEmployee**
> TimeRegistrationEntry PatchTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, timeRegistrationEntryId, contentType, ifMatch, body)
patchTimeRegistrationEntryForEmployee

Updates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiEmployeesApiTimeRegistrationEntriesTimeRegistrationEntryIdApiRequest**](CompaniesApiEmployeesApiTimeRegistrationEntriesTimeRegistrationEntryIdApiRequest.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTimeRegistrationEntry**
> TimeRegistrationEntry PostTimeRegistrationEntry(ctx, companyId, contentType, body)
postTimeRegistrationEntry

Creates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiTimeRegistrationEntriesRequest**](CompaniesApiTimeRegistrationEntriesRequest.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTimeRegistrationEntryForEmployee**
> TimeRegistrationEntry PostTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, contentType, body)
postTimeRegistrationEntryForEmployee

Creates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiEmployeesApiTimeRegistrationEntriesRequest**](CompaniesApiEmployeesApiTimeRegistrationEntriesRequest.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

