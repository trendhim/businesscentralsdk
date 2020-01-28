# \SalesOrderApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrder**](SalesOrderApi.md#DeleteSalesOrder) | **Delete** /companies({company_id})/salesOrders({salesOrder_id}) | deleteSalesOrder
[**GetSalesOrder**](SalesOrderApi.md#GetSalesOrder) | **Get** /companies({company_id})/salesOrders({salesOrder_id}) | getSalesOrder
[**ListSalesOrders**](SalesOrderApi.md#ListSalesOrders) | **Get** /companies({company_id})/salesOrders | listSalesOrders
[**PatchSalesOrder**](SalesOrderApi.md#PatchSalesOrder) | **Patch** /companies({company_id})/salesOrders({salesOrder_id}) | patchSalesOrder
[**PostSalesOrder**](SalesOrderApi.md#PostSalesOrder) | **Post** /companies({company_id})/salesOrders | postSalesOrder
[**ShipAndInvoiceActionSalesOrders**](SalesOrderApi.md#ShipAndInvoiceActionSalesOrders) | **Post** /companies({company_id})/salesOrders({salesOrder_id})/Microsoft.NAV.shipAndInvoice | shipAndInvoiceActionSalesOrders


# **DeleteSalesOrder**
> DeleteSalesOrder(ctx, companyId, salesOrderId)
deleteSalesOrder

Deletes an object of type salesOrder in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesOrder**
> SalesOrder GetSalesOrder(ctx, companyId, salesOrderId, optional)
getSalesOrder

Retrieve the properties and relationships of an object of type salesOrder for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
 **optional** | ***GetSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesOrder**](salesOrder.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesOrders**
> CompaniesApiSalesOrdersResponse ListSalesOrders(ctx, companyId, optional)
listSalesOrders

Returns a list of salesOrders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesOrdersResponse**](CompaniesAPISalesOrdersResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesOrder**
> SalesOrder PatchSalesOrder(ctx, companyId, salesOrderId, contentType, ifMatch, body)
patchSalesOrder

Updates an object of type salesOrder in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesOrdersApiRequest**](CompaniesApiSalesOrdersApiRequest.md)|  | 

### Return type

[**SalesOrder**](salesOrder.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesOrder**
> SalesOrder PostSalesOrder(ctx, companyId, contentType, body)
postSalesOrder

Creates an object of type salesOrder in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesOrdersRequest**](CompaniesApiSalesOrdersRequest.md)|  | 

### Return type

[**SalesOrder**](salesOrder.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShipAndInvoiceActionSalesOrders**
> ShipAndInvoiceActionSalesOrders(ctx, companyId, salesOrderId)
shipAndInvoiceActionSalesOrders

Performs the shipAndInvoice action for salesOrders entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

