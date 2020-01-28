# \GeneralLedgerEntryApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeneralLedgerEntry**](GeneralLedgerEntryApi.md#GetGeneralLedgerEntry) | **Get** /companies({company_id})/generalLedgerEntries({generalLedgerEntry_id}) | getGeneralLedgerEntry
[**ListGeneralLedgerEntries**](GeneralLedgerEntryApi.md#ListGeneralLedgerEntries) | **Get** /companies({company_id})/generalLedgerEntries | listGeneralLedgerEntries


# **GetGeneralLedgerEntry**
> GeneralLedgerEntry GetGeneralLedgerEntry(ctx, companyId, generalLedgerEntryId, optional)
getGeneralLedgerEntry

Retrieve the properties and relationships of an object of type generalLedgerEntry for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **generalLedgerEntryId** | **int32**| id for generalLedgerEntry | 
 **optional** | ***GetGeneralLedgerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGeneralLedgerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**GeneralLedgerEntry**](generalLedgerEntry.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGeneralLedgerEntries**
> CompaniesApiGeneralLedgerEntriesResponse ListGeneralLedgerEntries(ctx, companyId, optional)
listGeneralLedgerEntries

Returns a list of generalLedgerEntries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListGeneralLedgerEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListGeneralLedgerEntriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiGeneralLedgerEntriesResponse**](CompaniesAPIGeneralLedgerEntriesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

