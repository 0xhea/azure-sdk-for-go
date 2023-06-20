//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// ActionGroupsServer is a fake server for instances of the armmonitor.ActionGroupsClient type.
type ActionGroupsServer struct {
	// BeginCreateNotificationsAtActionGroupResourceLevel is the fake for method ActionGroupsClient.BeginCreateNotificationsAtActionGroupResourceLevel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateNotificationsAtActionGroupResourceLevel func(ctx context.Context, resourceGroupName string, actionGroupName string, notificationRequest armmonitor.NotificationRequestBody, options *armmonitor.ActionGroupsClientBeginCreateNotificationsAtActionGroupResourceLevelOptions) (resp azfake.PollerResponder[armmonitor.ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method ActionGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup armmonitor.ActionGroupResource, options *armmonitor.ActionGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armmonitor.ActionGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ActionGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, actionGroupName string, options *armmonitor.ActionGroupsClientDeleteOptions) (resp azfake.Responder[armmonitor.ActionGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// EnableReceiver is the fake for method ActionGroupsClient.EnableReceiver
	// HTTP status codes to indicate success: http.StatusOK
	EnableReceiver func(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest armmonitor.EnableRequest, options *armmonitor.ActionGroupsClientEnableReceiverOptions) (resp azfake.Responder[armmonitor.ActionGroupsClientEnableReceiverResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ActionGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, actionGroupName string, options *armmonitor.ActionGroupsClientGetOptions) (resp azfake.Responder[armmonitor.ActionGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// GetTestNotificationsAtActionGroupResourceLevel is the fake for method ActionGroupsClient.GetTestNotificationsAtActionGroupResourceLevel
	// HTTP status codes to indicate success: http.StatusOK
	GetTestNotificationsAtActionGroupResourceLevel func(ctx context.Context, resourceGroupName string, actionGroupName string, notificationID string, options *armmonitor.ActionGroupsClientGetTestNotificationsAtActionGroupResourceLevelOptions) (resp azfake.Responder[armmonitor.ActionGroupsClientGetTestNotificationsAtActionGroupResourceLevelResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ActionGroupsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmonitor.ActionGroupsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmonitor.ActionGroupsClientListByResourceGroupResponse])

	// NewListBySubscriptionIDPager is the fake for method ActionGroupsClient.NewListBySubscriptionIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionIDPager func(options *armmonitor.ActionGroupsClientListBySubscriptionIDOptions) (resp azfake.PagerResponder[armmonitor.ActionGroupsClientListBySubscriptionIDResponse])

	// Update is the fake for method ActionGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch armmonitor.ActionGroupPatchBody, options *armmonitor.ActionGroupsClientUpdateOptions) (resp azfake.Responder[armmonitor.ActionGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewActionGroupsServerTransport creates a new instance of ActionGroupsServerTransport with the provided implementation.
// The returned ActionGroupsServerTransport instance is connected to an instance of armmonitor.ActionGroupsClient by way of the
// undefined.Transporter field.
func NewActionGroupsServerTransport(srv *ActionGroupsServer) *ActionGroupsServerTransport {
	return &ActionGroupsServerTransport{srv: srv}
}

// ActionGroupsServerTransport connects instances of armmonitor.ActionGroupsClient to instances of ActionGroupsServer.
// Don't use this type directly, use NewActionGroupsServerTransport instead.
type ActionGroupsServerTransport struct {
	srv                                                *ActionGroupsServer
	beginCreateNotificationsAtActionGroupResourceLevel *azfake.PollerResponder[armmonitor.ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse]
	newListByResourceGroupPager                        *azfake.PagerResponder[armmonitor.ActionGroupsClientListByResourceGroupResponse]
	newListBySubscriptionIDPager                       *azfake.PagerResponder[armmonitor.ActionGroupsClientListBySubscriptionIDResponse]
}

// Do implements the policy.Transporter interface for ActionGroupsServerTransport.
func (a *ActionGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ActionGroupsClient.BeginCreateNotificationsAtActionGroupResourceLevel":
		resp, err = a.dispatchBeginCreateNotificationsAtActionGroupResourceLevel(req)
	case "ActionGroupsClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "ActionGroupsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "ActionGroupsClient.EnableReceiver":
		resp, err = a.dispatchEnableReceiver(req)
	case "ActionGroupsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ActionGroupsClient.GetTestNotificationsAtActionGroupResourceLevel":
		resp, err = a.dispatchGetTestNotificationsAtActionGroupResourceLevel(req)
	case "ActionGroupsClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "ActionGroupsClient.NewListBySubscriptionIDPager":
		resp, err = a.dispatchNewListBySubscriptionIDPager(req)
	case "ActionGroupsClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchBeginCreateNotificationsAtActionGroupResourceLevel(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateNotificationsAtActionGroupResourceLevel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateNotificationsAtActionGroupResourceLevel not implemented")}
	}
	if a.beginCreateNotificationsAtActionGroupResourceLevel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createNotifications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmonitor.NotificationRequestBody](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateNotificationsAtActionGroupResourceLevel(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginCreateNotificationsAtActionGroupResourceLevel = &respr
	}

	resp, err := server.PollerResponderNext(a.beginCreateNotificationsAtActionGroupResourceLevel, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginCreateNotificationsAtActionGroupResourceLevel) {
		a.beginCreateNotificationsAtActionGroupResourceLevel = nil
	}

	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.ActionGroupResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActionGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchEnableReceiver(req *http.Request) (*http.Response, error) {
	if a.srv.EnableReceiver == nil {
		return nil, &nonRetriableError{errors.New("fake for method EnableReceiver not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscribe`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.EnableRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.EnableReceiver(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActionGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchGetTestNotificationsAtActionGroupResourceLevel(req *http.Request) (*http.Response, error) {
	if a.srv.GetTestNotificationsAtActionGroupResourceLevel == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTestNotificationsAtActionGroupResourceLevel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationStatus/(?P<notificationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
	if err != nil {
		return nil, err
	}
	notificationIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("notificationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetTestNotificationsAtActionGroupResourceLevel(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, notificationIDUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TestNotificationDetailsResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if a.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		a.newListByResourceGroupPager = &resp
	}
	resp, err := server.PagerResponderNext(a.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListByResourceGroupPager) {
		a.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchNewListBySubscriptionIDPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionIDPager not implemented")}
	}
	if a.newListBySubscriptionIDPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionIDPager(nil)
		a.newListBySubscriptionIDPager = &resp
	}
	resp, err := server.PagerResponderNext(a.newListBySubscriptionIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListBySubscriptionIDPager) {
		a.newListBySubscriptionIDPager = nil
	}
	return resp, nil
}

func (a *ActionGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/actionGroups/(?P<actionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.ActionGroupPatchBody](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	actionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("actionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameUnescaped, actionGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActionGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}