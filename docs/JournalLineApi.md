# \JournalLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJournalLine**](JournalLineApi.md#DeleteJournalLine) | **Delete** /companies({company_id})/journalLines({journalLine_id}) | deleteJournalLine
[**DeleteJournalLineForJournal**](JournalLineApi.md#DeleteJournalLineForJournal) | **Delete** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id}) | deleteJournalLineForJournal
[**GetJournalLine**](JournalLineApi.md#GetJournalLine) | **Get** /companies({company_id})/journalLines({journalLine_id}) | getJournalLine
[**GetJournalLineForJournal**](JournalLineApi.md#GetJournalLineForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id}) | getJournalLineForJournal
[**ListJournalLines**](JournalLineApi.md#ListJournalLines) | **Get** /companies({company_id})/journalLines | listJournalLines
[**ListJournalLinesForJournal**](JournalLineApi.md#ListJournalLinesForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines | listJournalLinesForJournal
[**PatchJournalLine**](JournalLineApi.md#PatchJournalLine) | **Patch** /companies({company_id})/journalLines({journalLine_id}) | patchJournalLine
[**PatchJournalLineForJournal**](JournalLineApi.md#PatchJournalLineForJournal) | **Patch** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id}) | patchJournalLineForJournal
[**PostJournalLine**](JournalLineApi.md#PostJournalLine) | **Post** /companies({company_id})/journalLines | postJournalLine
[**PostJournalLineForJournal**](JournalLineApi.md#PostJournalLineForJournal) | **Post** /companies({company_id})/journals({journal_id})/journalLines | postJournalLineForJournal


# **DeleteJournalLine**
> DeleteJournalLine(ctx, companyId, journalLineId)
deleteJournalLine

Deletes an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJournalLineForJournal**
> DeleteJournalLineForJournal(ctx, companyId, journalId, journalLineId)
deleteJournalLineForJournal

Deletes an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJournalLine**
> JournalLine GetJournalLine(ctx, companyId, journalLineId, optional)
getJournalLine

Retrieve the properties and relationships of an object of type journalLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***GetJournalLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetJournalLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJournalLineForJournal**
> JournalLine GetJournalLineForJournal(ctx, companyId, journalId, journalLineId, optional)
getJournalLineForJournal

Retrieve the properties and relationships of an object of type journalLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***GetJournalLineForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetJournalLineForJournalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJournalLines**
> CompaniesApiJournalLinesResponse ListJournalLines(ctx, companyId, optional)
listJournalLines

Returns a list of journalLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListJournalLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListJournalLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiJournalLinesResponse**](CompaniesAPIJournalLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJournalLinesForJournal**
> CompaniesApiJournalsApiJournalLinesResponse ListJournalLinesForJournal(ctx, companyId, journalId, optional)
listJournalLinesForJournal

Returns a list of journalLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
 **optional** | ***ListJournalLinesForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListJournalLinesForJournalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiJournalsApiJournalLinesResponse**](CompaniesAPIJournalsAPIJournalLinesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchJournalLine**
> JournalLine PatchJournalLine(ctx, companyId, journalLineId, contentType, ifMatch, body)
patchJournalLine

Updates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiJournalLinesApiRequest**](CompaniesApiJournalLinesApiRequest.md)|  | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchJournalLineForJournal**
> JournalLine PatchJournalLineForJournal(ctx, companyId, journalId, journalLineId, contentType, ifMatch, body)
patchJournalLineForJournal

Updates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**CompaniesApiJournalsApiJournalLinesJournalLineIdApiRequest**](CompaniesApiJournalsApiJournalLinesJournalLineIdApiRequest.md)|  | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostJournalLine**
> JournalLine PostJournalLine(ctx, companyId, contentType, body)
postJournalLine

Creates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiJournalLinesRequest**](CompaniesApiJournalLinesRequest.md)|  | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostJournalLineForJournal**
> JournalLine PostJournalLineForJournal(ctx, companyId, journalId, contentType, body)
postJournalLineForJournal

Creates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **contentType** | **string**| application/json | 
  **body** | [**CompaniesApiJournalsApiJournalLinesRequest**](CompaniesApiJournalsApiJournalLinesRequest.md)|  | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

