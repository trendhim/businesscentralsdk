/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package businesscentralsdk

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type GeneralLedgerEntryAttachmentsApiService service

/* 
GeneralLedgerEntryAttachmentsApiService deleteGeneralLedgerEntryAttachments
Deletes an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param generalLedgerEntryAttachmentsGeneralLedgerEntryNumber generalLedgerEntryNumber for generalLedgerEntryAttachments
 * @param generalLedgerEntryAttachmentsId id for generalLedgerEntryAttachments


*/
func (a *GeneralLedgerEntryAttachmentsApiService) DeleteGeneralLedgerEntryAttachments(ctx context.Context, companyId string, generalLedgerEntryAttachmentsGeneralLedgerEntryNumber int32, generalLedgerEntryAttachmentsId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/generalLedgerEntryAttachments({generalLedgerEntryAttachments_generalLedgerEntryNumber},{generalLedgerEntryAttachments_id})"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", fmt.Sprintf("%v", companyId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"generalLedgerEntryAttachments_generalLedgerEntryNumber"+"}", fmt.Sprintf("%v", generalLedgerEntryAttachmentsGeneralLedgerEntryNumber), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"generalLedgerEntryAttachments_id"+"}", fmt.Sprintf("%v", generalLedgerEntryAttachmentsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
GeneralLedgerEntryAttachmentsApiService getGeneralLedgerEntryAttachments
Retrieve the properties and relationships of an object of type generalLedgerEntryAttachments for Dynamics 365 Business Central.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param generalLedgerEntryAttachmentsGeneralLedgerEntryNumber generalLedgerEntryNumber for generalLedgerEntryAttachments
 * @param generalLedgerEntryAttachmentsId id for generalLedgerEntryAttachments
 * @param optional nil or *GetGeneralLedgerEntryAttachmentsOpts - Optional Parameters:
     * @param "Expand" (optional.Interface of []string) -  Entities to expand
     * @param "Select_" (optional.Interface of []string) -  Selected properties to be retrieved

@return GeneralLedgerEntryAttachments
*/

type GetGeneralLedgerEntryAttachmentsOpts struct { 
	Expand optional.Interface
	Select_ optional.Interface
}

func (a *GeneralLedgerEntryAttachmentsApiService) GetGeneralLedgerEntryAttachments(ctx context.Context, companyId string, generalLedgerEntryAttachmentsGeneralLedgerEntryNumber int32, generalLedgerEntryAttachmentsId string, localVarOptionals *GetGeneralLedgerEntryAttachmentsOpts) (GeneralLedgerEntryAttachments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue GeneralLedgerEntryAttachments
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/generalLedgerEntryAttachments({generalLedgerEntryAttachments_generalLedgerEntryNumber},{generalLedgerEntryAttachments_id})"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", fmt.Sprintf("%v", companyId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"generalLedgerEntryAttachments_generalLedgerEntryNumber"+"}", fmt.Sprintf("%v", generalLedgerEntryAttachmentsGeneralLedgerEntryNumber), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"generalLedgerEntryAttachments_id"+"}", fmt.Sprintf("%v", generalLedgerEntryAttachmentsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v GeneralLedgerEntryAttachments
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
GeneralLedgerEntryAttachmentsApiService listGeneralLedgerEntryAttachments
Returns a list of generalLedgerEntryAttachments
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param optional nil or *ListGeneralLedgerEntryAttachmentsOpts - Optional Parameters:
     * @param "Top" (optional.Int32) -  Number of items to return from the top of the list
     * @param "Skip" (optional.Int32) -  Number of items to skip from the list
     * @param "Limit" (optional.Int32) -  Number of items to return from the list
     * @param "Filter" (optional.String) -  Filtering expression
     * @param "Expand" (optional.Interface of []string) -  Entities to expand
     * @param "Select_" (optional.Interface of []string) -  Selected properties to be retrieved

@return CompaniesApiGeneralLedgerEntryAttachmentsResponse
*/

type ListGeneralLedgerEntryAttachmentsOpts struct { 
	Top optional.Int32
	Skip optional.Int32
	Limit optional.Int32
	Filter optional.String
	Expand optional.Interface
	Select_ optional.Interface
}

func (a *GeneralLedgerEntryAttachmentsApiService) ListGeneralLedgerEntryAttachments(ctx context.Context, companyId string, localVarOptionals *ListGeneralLedgerEntryAttachmentsOpts) (CompaniesApiGeneralLedgerEntryAttachmentsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CompaniesApiGeneralLedgerEntryAttachmentsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/generalLedgerEntryAttachments"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", fmt.Sprintf("%v", companyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("$limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("$filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v CompaniesApiGeneralLedgerEntryAttachmentsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
GeneralLedgerEntryAttachmentsApiService patchGeneralLedgerEntryAttachments
Updates an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param generalLedgerEntryAttachmentsGeneralLedgerEntryNumber generalLedgerEntryNumber for generalLedgerEntryAttachments
 * @param generalLedgerEntryAttachmentsId id for generalLedgerEntryAttachments
 * @param contentType application/json
 * @param ifMatch Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated.
 * @param body 

@return GeneralLedgerEntryAttachments
*/
func (a *GeneralLedgerEntryAttachmentsApiService) PatchGeneralLedgerEntryAttachments(ctx context.Context, companyId string, generalLedgerEntryAttachmentsGeneralLedgerEntryNumber int32, generalLedgerEntryAttachmentsId string, contentType string, ifMatch string, body CompaniesApiGeneralLedgerEntryAttachmentsGeneralLedgerEntryAttachmentsIdApiRequest) (GeneralLedgerEntryAttachments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue GeneralLedgerEntryAttachments
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/generalLedgerEntryAttachments({generalLedgerEntryAttachments_generalLedgerEntryNumber},{generalLedgerEntryAttachments_id})"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", fmt.Sprintf("%v", companyId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"generalLedgerEntryAttachments_generalLedgerEntryNumber"+"}", fmt.Sprintf("%v", generalLedgerEntryAttachmentsGeneralLedgerEntryNumber), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"generalLedgerEntryAttachments_id"+"}", fmt.Sprintf("%v", generalLedgerEntryAttachmentsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["If-Match"] = parameterToString(ifMatch, "")
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v GeneralLedgerEntryAttachments
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
GeneralLedgerEntryAttachmentsApiService postGeneralLedgerEntryAttachments
Creates an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param contentType application/json
 * @param body 

@return GeneralLedgerEntryAttachments
*/
func (a *GeneralLedgerEntryAttachmentsApiService) PostGeneralLedgerEntryAttachments(ctx context.Context, companyId string, contentType string, body CompaniesApiGeneralLedgerEntryAttachmentsRequest) (GeneralLedgerEntryAttachments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue GeneralLedgerEntryAttachments
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/generalLedgerEntryAttachments"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", fmt.Sprintf("%v", companyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 201 {
			var v GeneralLedgerEntryAttachments
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
