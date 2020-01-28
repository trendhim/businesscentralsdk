# \PdfDocumentApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPdfDocument**](PdfDocumentApi.md#GetPdfDocument) | **Get** /companies({company_id})/pdfDocument({pdfDocument_id}) | getPdfDocument
[**GetPdfDocumentForPurchaseInvoice**](PdfDocumentApi.md#GetPdfDocumentForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/pdfDocument({pdfDocument_id}) | getPdfDocumentForPurchaseInvoice
[**GetPdfDocumentForSalesCreditMemo**](PdfDocumentApi.md#GetPdfDocumentForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/pdfDocument({pdfDocument_id}) | getPdfDocumentForSalesCreditMemo
[**GetPdfDocumentForSalesInvoice**](PdfDocumentApi.md#GetPdfDocumentForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/pdfDocument({pdfDocument_id}) | getPdfDocumentForSalesInvoice
[**GetPdfDocumentForSalesQuote**](PdfDocumentApi.md#GetPdfDocumentForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/pdfDocument({pdfDocument_id}) | getPdfDocumentForSalesQuote
[**ListPdfDocument**](PdfDocumentApi.md#ListPdfDocument) | **Get** /companies({company_id})/pdfDocument | listPdfDocument
[**ListPdfDocumentForPurchaseInvoice**](PdfDocumentApi.md#ListPdfDocumentForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/pdfDocument | listPdfDocumentForPurchaseInvoice
[**ListPdfDocumentForSalesCreditMemo**](PdfDocumentApi.md#ListPdfDocumentForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/pdfDocument | listPdfDocumentForSalesCreditMemo
[**ListPdfDocumentForSalesInvoice**](PdfDocumentApi.md#ListPdfDocumentForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/pdfDocument | listPdfDocumentForSalesInvoice
[**ListPdfDocumentForSalesQuote**](PdfDocumentApi.md#ListPdfDocumentForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/pdfDocument | listPdfDocumentForSalesQuote


# **GetPdfDocument**
> PdfDocument GetPdfDocument(ctx, companyId, pdfDocumentId, optional)
getPdfDocument

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPdfDocumentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForPurchaseInvoice**
> PdfDocument GetPdfDocumentForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, pdfDocumentId, optional)
getPdfDocumentForPurchaseInvoice

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPdfDocumentForPurchaseInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForSalesCreditMemo**
> PdfDocument GetPdfDocumentForSalesCreditMemo(ctx, companyId, salesCreditMemoId, pdfDocumentId, optional)
getPdfDocumentForSalesCreditMemo

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPdfDocumentForSalesCreditMemoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForSalesInvoice**
> PdfDocument GetPdfDocumentForSalesInvoice(ctx, companyId, salesInvoiceId, pdfDocumentId, optional)
getPdfDocumentForSalesInvoice

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPdfDocumentForSalesInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForSalesQuote**
> PdfDocument GetPdfDocumentForSalesQuote(ctx, companyId, salesQuoteId, pdfDocumentId, optional)
getPdfDocumentForSalesQuote

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPdfDocumentForSalesQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocument**
> CompaniesApiPdfDocumentResponse ListPdfDocument(ctx, companyId, optional)
listPdfDocument

Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPdfDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPdfDocumentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPdfDocumentResponse**](CompaniesAPIPdfDocumentResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForPurchaseInvoice**
> CompaniesApiPurchaseInvoicesApiPdfDocumentResponse ListPdfDocumentForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)
listPdfDocumentForPurchaseInvoice

Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***ListPdfDocumentForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPdfDocumentForPurchaseInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiPurchaseInvoicesApiPdfDocumentResponse**](CompaniesAPIPurchaseInvoicesAPIPdfDocumentResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForSalesCreditMemo**
> CompaniesApiSalesCreditMemosApiPdfDocumentResponse ListPdfDocumentForSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)
listPdfDocumentForSalesCreditMemo

Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***ListPdfDocumentForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPdfDocumentForSalesCreditMemoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesCreditMemosApiPdfDocumentResponse**](CompaniesAPISalesCreditMemosAPIPdfDocumentResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForSalesInvoice**
> CompaniesApiSalesInvoicesApiPdfDocumentResponse ListPdfDocumentForSalesInvoice(ctx, companyId, salesInvoiceId, optional)
listPdfDocumentForSalesInvoice

Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***ListPdfDocumentForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPdfDocumentForSalesInvoiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesInvoicesApiPdfDocumentResponse**](CompaniesAPISalesInvoicesAPIPdfDocumentResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForSalesQuote**
> CompaniesApiSalesQuotesApiPdfDocumentResponse ListPdfDocumentForSalesQuote(ctx, companyId, salesQuoteId, optional)
listPdfDocumentForSalesQuote

Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***ListPdfDocumentForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPdfDocumentForSalesQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CompaniesApiSalesQuotesApiPdfDocumentResponse**](CompaniesAPISalesQuotesAPIPdfDocumentResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

