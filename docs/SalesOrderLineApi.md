# \SalesOrderLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrderLine**](SalesOrderLineApi.md#DeleteSalesOrderLine) | **Delete** /companies({company_id})/salesOrderLines(&#39;{salesOrderLine_id}&#39;) | deleteSalesOrderLine
[**DeleteSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#DeleteSalesOrderLineForSalesOrder) | **Delete** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines(&#39;{salesOrderLine_id}&#39;) | deleteSalesOrderLineForSalesOrder
[**GetSalesOrderLine**](SalesOrderLineApi.md#GetSalesOrderLine) | **Get** /companies({company_id})/salesOrderLines(&#39;{salesOrderLine_id}&#39;) | getSalesOrderLine
[**GetSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#GetSalesOrderLineForSalesOrder) | **Get** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines(&#39;{salesOrderLine_id}&#39;) | getSalesOrderLineForSalesOrder
[**ListSalesOrderLines**](SalesOrderLineApi.md#ListSalesOrderLines) | **Get** /companies({company_id})/salesOrderLines | listSalesOrderLines
[**ListSalesOrderLinesForSalesOrder**](SalesOrderLineApi.md#ListSalesOrderLinesForSalesOrder) | **Get** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines | listSalesOrderLinesForSalesOrder
[**PatchSalesOrderLine**](SalesOrderLineApi.md#PatchSalesOrderLine) | **Patch** /companies({company_id})/salesOrderLines(&#39;{salesOrderLine_id}&#39;) | patchSalesOrderLine
[**PatchSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#PatchSalesOrderLineForSalesOrder) | **Patch** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines(&#39;{salesOrderLine_id}&#39;) | patchSalesOrderLineForSalesOrder
[**PostSalesOrderLine**](SalesOrderLineApi.md#PostSalesOrderLine) | **Post** /companies({company_id})/salesOrderLines | postSalesOrderLine
[**PostSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#PostSalesOrderLineForSalesOrder) | **Post** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines | postSalesOrderLineForSalesOrder


# **DeleteSalesOrderLine**
> DeleteSalesOrderLine(ctx, companyId, salesOrderLineId)
deleteSalesOrderLine

Deletes an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesOrderLineForSalesOrder**
> DeleteSalesOrderLineForSalesOrder(ctx, companyId, salesOrderId, salesOrderLineId)
deleteSalesOrderLineForSalesOrder

Deletes an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesOrderLine**
> SalesOrderLine GetSalesOrderLine(ctx, companyId, salesOrderLineId, optional)
getSalesOrderLine

Retrieve the properties and relationships of an object of type salesOrderLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
 **optional** | ***GetSalesOrderLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesOrderLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesOrderLineForSalesOrder**
> SalesOrderLine GetSalesOrderLineForSalesOrder(ctx, companyId, salesOrderId, salesOrderLineId, optional)
getSalesOrderLineForSalesOrder

Retrieve the properties and relationships of an object of type salesOrderLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
 **optional** | ***GetSalesOrderLineForSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSalesOrderLineForSalesOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesOrderLines**
> CompaniesApiSalesOrderLinesResponse ListSalesOrderLines(ctx, companyId, optional)
listSalesOrderLines

Returns a list of salesOrderLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesOrderLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesOrderLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesOrderLinesResponse**](CompaniesAPISalesOrderLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesOrderLinesForSalesOrder**
> CompaniesApiSalesOrdersApiSalesOrderLinesResponse ListSalesOrderLinesForSalesOrder(ctx, companyId, salesOrderId, optional)
listSalesOrderLinesForSalesOrder

Returns a list of salesOrderLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
 **optional** | ***ListSalesOrderLinesForSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSalesOrderLinesForSalesOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesOrdersApiSalesOrderLinesResponse**](CompaniesAPISalesOrdersAPISalesOrderLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesOrderLine**
> SalesOrderLine PatchSalesOrderLine(ctx, companyId, salesOrderLineId, contentType, ifMatch, body)
patchSalesOrderLine

Updates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesOrderLinesApiRequest**](CompaniesApiSalesOrderLinesApiRequest.md)|  | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesOrderLineForSalesOrder**
> SalesOrderLine PatchSalesOrderLineForSalesOrder(ctx, companyId, salesOrderId, salesOrderLineId, contentType, ifMatch, body)
patchSalesOrderLineForSalesOrder

Updates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiSalesOrdersApiSalesOrderLinesSalesOrderLineIdApiRequest**](CompaniesApiSalesOrdersApiSalesOrderLinesSalesOrderLineIdApiRequest.md)|  | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesOrderLine**
> SalesOrderLine PostSalesOrderLine(ctx, companyId, contentType, body)
postSalesOrderLine

Creates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesOrderLinesRequest**](CompaniesApiSalesOrderLinesRequest.md)|  | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesOrderLineForSalesOrder**
> SalesOrderLine PostSalesOrderLineForSalesOrder(ctx, companyId, salesOrderId, contentType, body)
postSalesOrderLineForSalesOrder

Creates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiSalesOrdersApiSalesOrderLinesRequest**](CompaniesApiSalesOrdersApiSalesOrderLinesRequest.md)|  | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

