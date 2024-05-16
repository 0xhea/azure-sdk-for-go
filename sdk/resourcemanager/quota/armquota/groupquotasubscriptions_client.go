//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GroupQuotaSubscriptionsClient contains the methods for the GroupQuotaSubscriptions group.
// Don't use this type directly, use NewGroupQuotaSubscriptionsClient() instead.
type GroupQuotaSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGroupQuotaSubscriptionsClient creates a new instance of GroupQuotaSubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGroupQuotaSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GroupQuotaSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GroupQuotaSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Adds a subscription to GroupQuotas. The subscriptions will be validated based on the additionalAttributes
// defined in the GroupQuota. The additionalAttributes works as filter for the subscriptions,
// which can be included in the GroupQuotas. The request's TenantId is validated against the subscription's TenantId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - managementGroupID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotaSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupQuotaSubscriptionsClient.BeginCreateOrUpdate
//     method.
func (client *GroupQuotaSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotaSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, managementGroupID, groupQuotaName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupQuotaSubscriptionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupQuotaSubscriptionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Adds a subscription to GroupQuotas. The subscriptions will be validated based on the additionalAttributes
// defined in the GroupQuota. The additionalAttributes works as filter for the subscriptions,
// which can be included in the GroupQuotas. The request's TenantId is validated against the subscription's TenantId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *GroupQuotaSubscriptionsClient) createOrUpdate(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupQuotaSubscriptionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, managementGroupID, groupQuotaName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GroupQuotaSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptions/{subscriptionId}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Removes the subscription from GroupQuotas. The request's TenantId is validated against the subscription's
// TenantId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - managementGroupID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotaSubscriptionsClientBeginDeleteOptions contains the optional parameters for the GroupQuotaSubscriptionsClient.BeginDelete
//     method.
func (client *GroupQuotaSubscriptionsClient) BeginDelete(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[GroupQuotaSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, managementGroupID, groupQuotaName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupQuotaSubscriptionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupQuotaSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Removes the subscription from GroupQuotas. The request's TenantId is validated against the subscription's TenantId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *GroupQuotaSubscriptionsClient) deleteOperation(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupQuotaSubscriptionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, managementGroupID, groupQuotaName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GroupQuotaSubscriptionsClient) deleteCreateRequest(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptions/{subscriptionId}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the subscriptionIds along with its provisioning state for being associated with the GroupQuota. If the subscription
// is not a member of GroupQuota, it will return 404, else 200.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - managementGroupID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotaSubscriptionsClientGetOptions contains the optional parameters for the GroupQuotaSubscriptionsClient.Get
//     method.
func (client *GroupQuotaSubscriptionsClient) Get(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientGetOptions) (GroupQuotaSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "GroupQuotaSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, managementGroupID, groupQuotaName, options)
	if err != nil {
		return GroupQuotaSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupQuotaSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GroupQuotaSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GroupQuotaSubscriptionsClient) getCreateRequest(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptions/{subscriptionId}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GroupQuotaSubscriptionsClient) getHandleResponse(resp *http.Response) (GroupQuotaSubscriptionsClientGetResponse, error) {
	result := GroupQuotaSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupQuotaSubscriptionID); err != nil {
		return GroupQuotaSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of the subscriptionIds associated with the GroupQuotas.
//
// Generated from API version 2023-06-01-preview
//   - managementGroupID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotaSubscriptionsClientListOptions contains the optional parameters for the GroupQuotaSubscriptionsClient.NewListPager
//     method.
func (client *GroupQuotaSubscriptionsClient) NewListPager(managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientListOptions) *runtime.Pager[GroupQuotaSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GroupQuotaSubscriptionsClientListResponse]{
		More: func(page GroupQuotaSubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GroupQuotaSubscriptionsClientListResponse) (GroupQuotaSubscriptionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GroupQuotaSubscriptionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, managementGroupID, groupQuotaName, options)
			}, nil)
			if err != nil {
				return GroupQuotaSubscriptionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GroupQuotaSubscriptionsClient) listCreateRequest(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GroupQuotaSubscriptionsClient) listHandleResponse(resp *http.Response) (GroupQuotaSubscriptionsClientListResponse, error) {
	result := GroupQuotaSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupQuotaSubscriptionIDList); err != nil {
		return GroupQuotaSubscriptionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the GroupQuotas with the subscription to add to the subscriptions list. The subscriptions will be
// validated if additionalAttributes are defined in the GroupQuota. The request's TenantId is
// validated against the subscription's TenantId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - managementGroupID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotaSubscriptionsClientBeginUpdateOptions contains the optional parameters for the GroupQuotaSubscriptionsClient.BeginUpdate
//     method.
func (client *GroupQuotaSubscriptionsClient) BeginUpdate(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[GroupQuotaSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, managementGroupID, groupQuotaName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupQuotaSubscriptionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupQuotaSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the GroupQuotas with the subscription to add to the subscriptions list. The subscriptions will be validated
// if additionalAttributes are defined in the GroupQuota. The request's TenantId is
// validated against the subscription's TenantId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *GroupQuotaSubscriptionsClient) update(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupQuotaSubscriptionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, managementGroupID, groupQuotaName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *GroupQuotaSubscriptionsClient) updateCreateRequest(ctx context.Context, managementGroupID string, groupQuotaName string, options *GroupQuotaSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptions/{subscriptionId}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}