# \AttachmentsApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttachments**](AttachmentsApi.md#DeleteAttachments) | **Delete** /companies({company_id})/attachments({attachments_parentId},{attachments_id}) | deleteAttachments
[**DeleteAttachmentsForJournalLine**](AttachmentsApi.md#DeleteAttachmentsForJournalLine) | **Delete** /companies({company_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | deleteAttachmentsForJournalLine
[**DeleteAttachmentsForJournalLineForJournal**](AttachmentsApi.md#DeleteAttachmentsForJournalLineForJournal) | **Delete** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | deleteAttachmentsForJournalLineForJournal
[**GetAttachments**](AttachmentsApi.md#GetAttachments) | **Get** /companies({company_id})/attachments({attachments_parentId},{attachments_id}) | getAttachments
[**GetAttachmentsForJournalLine**](AttachmentsApi.md#GetAttachmentsForJournalLine) | **Get** /companies({company_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | getAttachmentsForJournalLine
[**GetAttachmentsForJournalLineForJournal**](AttachmentsApi.md#GetAttachmentsForJournalLineForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | getAttachmentsForJournalLineForJournal
[**ListAttachments**](AttachmentsApi.md#ListAttachments) | **Get** /companies({company_id})/attachments | listAttachments
[**ListAttachmentsForJournalLine**](AttachmentsApi.md#ListAttachmentsForJournalLine) | **Get** /companies({company_id})/journalLines({journalLine_id})/attachments | listAttachmentsForJournalLine
[**ListAttachmentsForJournalLineForJournal**](AttachmentsApi.md#ListAttachmentsForJournalLineForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments | listAttachmentsForJournalLineForJournal
[**PatchAttachments**](AttachmentsApi.md#PatchAttachments) | **Patch** /companies({company_id})/attachments({attachments_parentId},{attachments_id}) | patchAttachments
[**PatchAttachmentsForJournalLine**](AttachmentsApi.md#PatchAttachmentsForJournalLine) | **Patch** /companies({company_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | patchAttachmentsForJournalLine
[**PatchAttachmentsForJournalLineForJournal**](AttachmentsApi.md#PatchAttachmentsForJournalLineForJournal) | **Patch** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | patchAttachmentsForJournalLineForJournal
[**PostAttachments**](AttachmentsApi.md#PostAttachments) | **Post** /companies({company_id})/attachments | postAttachments
[**PostAttachmentsForJournalLine**](AttachmentsApi.md#PostAttachmentsForJournalLine) | **Post** /companies({company_id})/journalLines({journalLine_id})/attachments | postAttachmentsForJournalLine
[**PostAttachmentsForJournalLineForJournal**](AttachmentsApi.md#PostAttachmentsForJournalLineForJournal) | **Post** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments | postAttachmentsForJournalLineForJournal


# **DeleteAttachments**
> DeleteAttachments(ctx, companyId, attachmentsParentId, attachmentsId)
deleteAttachments

Deletes an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAttachmentsForJournalLine**
> DeleteAttachmentsForJournalLine(ctx, companyId, journalLineId, attachmentsParentId, attachmentsId)
deleteAttachmentsForJournalLine

Deletes an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAttachmentsForJournalLineForJournal**
> DeleteAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, attachmentsParentId, attachmentsId)
deleteAttachmentsForJournalLineForJournal

Deletes an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 

### Return type

 (empty response body)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttachments**
> Attachments GetAttachments(ctx, companyId, attachmentsParentId, attachmentsId, optional)
getAttachments

Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 
 **optional** | ***GetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAttachmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttachmentsForJournalLine**
> Attachments GetAttachmentsForJournalLine(ctx, companyId, journalLineId, attachmentsParentId, attachmentsId, optional)
getAttachmentsForJournalLine

Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 
 **optional** | ***GetAttachmentsForJournalLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAttachmentsForJournalLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttachmentsForJournalLineForJournal**
> Attachments GetAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, attachmentsParentId, attachmentsId, optional)
getAttachmentsForJournalLineForJournal

Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 
 **optional** | ***GetAttachmentsForJournalLineForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAttachmentsForJournalLineForJournalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachments**
> CompaniesApiAttachmentsResponse ListAttachments(ctx, companyId, optional)
listAttachments

Returns a list of attachments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAttachmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiAttachmentsResponse**](CompaniesAPIAttachmentsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentsForJournalLine**
> CompaniesApiJournalLinesApiAttachmentsResponse ListAttachmentsForJournalLine(ctx, companyId, journalLineId, optional)
listAttachmentsForJournalLine

Returns a list of attachments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***ListAttachmentsForJournalLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAttachmentsForJournalLineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiJournalLinesApiAttachmentsResponse**](CompaniesAPIJournalLinesAPIAttachmentsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentsForJournalLineForJournal**
> CompaniesApiJournalsApiJournalLinesJournalLineIdApiAttachmentsResponse ListAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, optional)
listAttachmentsForJournalLineForJournal

Returns a list of attachments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***ListAttachmentsForJournalLineForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAttachmentsForJournalLineForJournalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiJournalsApiJournalLinesJournalLineIdApiAttachmentsResponse**](CompaniesAPIJournalsAPIJournalLinesJournalLineIdAPIAttachmentsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAttachments**
> Attachments PatchAttachments(ctx, companyId, attachmentsParentId, attachmentsId, contentType, ifMatch, body)
patchAttachments

Updates an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Attachments**](Attachments.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAttachmentsForJournalLine**
> Attachments PatchAttachmentsForJournalLine(ctx, companyId, journalLineId, attachmentsParentId, attachmentsId, contentType, ifMatch, body)
patchAttachmentsForJournalLine

Updates an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Attachments**](Attachments.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAttachmentsForJournalLineForJournal**
> Attachments PatchAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, attachmentsParentId, attachmentsId, contentType, ifMatch, body)
patchAttachmentsForJournalLineForJournal

Updates an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **attachmentsParentId** | [**string**](.md)| parentId for attachments | 
  **attachmentsId** | [**string**](.md)| id for attachments | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
  **body** | [**Attachments**](Attachments.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAttachments**
> Attachments PostAttachments(ctx, companyId, contentType, body)
postAttachments

Creates an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 
  **body** | [**Attachments**](Attachments.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAttachmentsForJournalLine**
> Attachments PostAttachmentsForJournalLine(ctx, companyId, journalLineId, contentType, body)
postAttachmentsForJournalLine

Creates an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **contentType** | **string**| application/json | 
  **body** | [**Attachments**](Attachments.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAttachmentsForJournalLineForJournal**
> Attachments PostAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, contentType, body)
postAttachmentsForJournalLineForJournal

Creates an object of type attachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **contentType** | **string**| application/json | 
  **body** | [**Attachments**](Attachments.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

