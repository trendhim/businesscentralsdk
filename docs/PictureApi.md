# \PictureApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePicture**](PictureApi.md#DeletePicture) | **Delete** /companies({company_id})/picture({picture_id}) | deletePicture
[**DeletePictureForCustomer**](PictureApi.md#DeletePictureForCustomer) | **Delete** /companies({company_id})/customers({customer_id})/picture({picture_id}) | deletePictureForCustomer
[**DeletePictureForEmployee**](PictureApi.md#DeletePictureForEmployee) | **Delete** /companies({company_id})/employees({employee_id})/picture({picture_id}) | deletePictureForEmployee
[**DeletePictureForItem**](PictureApi.md#DeletePictureForItem) | **Delete** /companies({company_id})/items({item_id})/picture({picture_id}) | deletePictureForItem
[**DeletePictureForVendor**](PictureApi.md#DeletePictureForVendor) | **Delete** /companies({company_id})/vendors({vendor_id})/picture({picture_id}) | deletePictureForVendor
[**GetPicture**](PictureApi.md#GetPicture) | **Get** /companies({company_id})/picture({picture_id}) | getPicture
[**GetPictureForCustomer**](PictureApi.md#GetPictureForCustomer) | **Get** /companies({company_id})/customers({customer_id})/picture({picture_id}) | getPictureForCustomer
[**GetPictureForEmployee**](PictureApi.md#GetPictureForEmployee) | **Get** /companies({company_id})/employees({employee_id})/picture({picture_id}) | getPictureForEmployee
[**GetPictureForItem**](PictureApi.md#GetPictureForItem) | **Get** /companies({company_id})/items({item_id})/picture({picture_id}) | getPictureForItem
[**GetPictureForVendor**](PictureApi.md#GetPictureForVendor) | **Get** /companies({company_id})/vendors({vendor_id})/picture({picture_id}) | getPictureForVendor
[**ListPicture**](PictureApi.md#ListPicture) | **Get** /companies({company_id})/picture | listPicture
[**ListPictureForCustomer**](PictureApi.md#ListPictureForCustomer) | **Get** /companies({company_id})/customers({customer_id})/picture | listPictureForCustomer
[**ListPictureForEmployee**](PictureApi.md#ListPictureForEmployee) | **Get** /companies({company_id})/employees({employee_id})/picture | listPictureForEmployee
[**ListPictureForItem**](PictureApi.md#ListPictureForItem) | **Get** /companies({company_id})/items({item_id})/picture | listPictureForItem
[**ListPictureForVendor**](PictureApi.md#ListPictureForVendor) | **Get** /companies({company_id})/vendors({vendor_id})/picture | listPictureForVendor
[**PatchPicture**](PictureApi.md#PatchPicture) | **Patch** /companies({company_id})/picture({picture_id}) | patchPicture
[**PatchPictureForCustomer**](PictureApi.md#PatchPictureForCustomer) | **Patch** /companies({company_id})/customers({customer_id})/picture({picture_id}) | patchPictureForCustomer
[**PatchPictureForEmployee**](PictureApi.md#PatchPictureForEmployee) | **Patch** /companies({company_id})/employees({employee_id})/picture({picture_id}) | patchPictureForEmployee
[**PatchPictureForItem**](PictureApi.md#PatchPictureForItem) | **Patch** /companies({company_id})/items({item_id})/picture({picture_id}) | patchPictureForItem
[**PatchPictureForVendor**](PictureApi.md#PatchPictureForVendor) | **Patch** /companies({company_id})/vendors({vendor_id})/picture({picture_id}) | patchPictureForVendor


# **DeletePicture**
> DeletePicture(ctx, companyId, pictureId)
deletePicture

Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForCustomer**
> DeletePictureForCustomer(ctx, companyId, customerId, pictureId)
deletePictureForCustomer

Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForEmployee**
> DeletePictureForEmployee(ctx, companyId, employeeId, pictureId)
deletePictureForEmployee

Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForItem**
> DeletePictureForItem(ctx, companyId, itemId, pictureId)
deletePictureForItem

Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForVendor**
> DeletePictureForVendor(ctx, companyId, vendorId, pictureId)
deletePictureForVendor

Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPicture**
> Picture GetPicture(ctx, companyId, pictureId, optional)
getPicture

Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***GetPictureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPictureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForCustomer**
> Picture GetPictureForCustomer(ctx, companyId, customerId, pictureId, optional)
getPictureForCustomer

Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***GetPictureForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPictureForCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForEmployee**
> Picture GetPictureForEmployee(ctx, companyId, employeeId, pictureId, optional)
getPictureForEmployee

Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***GetPictureForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPictureForEmployeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForItem**
> Picture GetPictureForItem(ctx, companyId, itemId, pictureId, optional)
getPictureForItem

Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***GetPictureForItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPictureForItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForVendor**
> Picture GetPictureForVendor(ctx, companyId, vendorId, pictureId, optional)
getPictureForVendor

Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***GetPictureForVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPictureForVendorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPicture**
> CompaniesApiPictureResponse ListPicture(ctx, companyId, optional)
listPicture

Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPictureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPictureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPictureResponse**](CompaniesAPIPictureResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForCustomer**
> CompaniesApiCustomersApiPictureResponse ListPictureForCustomer(ctx, companyId, customerId, optional)
listPictureForCustomer

Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
 **optional** | ***ListPictureForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPictureForCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiCustomersApiPictureResponse**](CompaniesAPICustomersAPIPictureResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForEmployee**
> CompaniesApiEmployeesApiPictureResponse ListPictureForEmployee(ctx, companyId, employeeId, optional)
listPictureForEmployee

Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
 **optional** | ***ListPictureForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPictureForEmployeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiEmployeesApiPictureResponse**](CompaniesAPIEmployeesAPIPictureResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForItem**
> CompaniesApiItemsApiPictureResponse ListPictureForItem(ctx, companyId, itemId, optional)
listPictureForItem

Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
 **optional** | ***ListPictureForItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPictureForItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiItemsApiPictureResponse**](CompaniesAPIItemsAPIPictureResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForVendor**
> CompaniesApiVendorsApiPictureResponse ListPictureForVendor(ctx, companyId, vendorId, optional)
listPictureForVendor

Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
 **optional** | ***ListPictureForVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPictureForVendorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiVendorsApiPictureResponse**](CompaniesAPIVendorsAPIPictureResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPicture**
> Picture PatchPicture(ctx, companyId, pictureId, contentType, ifMatch, body)
patchPicture

Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Picture**](Picture.md)|  | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForCustomer**
> Picture PatchPictureForCustomer(ctx, companyId, customerId, pictureId, contentType, ifMatch, body)
patchPictureForCustomer

Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Picture**](Picture.md)|  | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForEmployee**
> Picture PatchPictureForEmployee(ctx, companyId, employeeId, pictureId, contentType, ifMatch, body)
patchPictureForEmployee

Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Picture**](Picture.md)|  | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForItem**
> Picture PatchPictureForItem(ctx, companyId, itemId, pictureId, contentType, ifMatch, body)
patchPictureForItem

Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Picture**](Picture.md)|  | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForVendor**
> Picture PatchPictureForVendor(ctx, companyId, vendorId, pictureId, contentType, ifMatch, body)
patchPictureForVendor

Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Picture**](Picture.md)|  | 

### Return type

[**Picture**](picture.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

