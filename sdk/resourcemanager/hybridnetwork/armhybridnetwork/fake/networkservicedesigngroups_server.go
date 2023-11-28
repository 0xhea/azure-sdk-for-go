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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkServiceDesignGroupsServer is a fake server for instances of the armhybridnetwork.NetworkServiceDesignGroupsClient type.
type NetworkServiceDesignGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method NetworkServiceDesignGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters armhybridnetwork.NetworkServiceDesignGroup, options *armhybridnetwork.NetworkServiceDesignGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkServiceDesignGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *armhybridnetwork.NetworkServiceDesignGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkServiceDesignGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *armhybridnetwork.NetworkServiceDesignGroupsClientGetOptions) (resp azfake.Responder[armhybridnetwork.NetworkServiceDesignGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByPublisherPager is the fake for method NetworkServiceDesignGroupsClient.NewListByPublisherPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByPublisherPager func(resourceGroupName string, publisherName string, options *armhybridnetwork.NetworkServiceDesignGroupsClientListByPublisherOptions) (resp azfake.PagerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientListByPublisherResponse])

	// Update is the fake for method NetworkServiceDesignGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters armhybridnetwork.TagsObject, options *armhybridnetwork.NetworkServiceDesignGroupsClientUpdateOptions) (resp azfake.Responder[armhybridnetwork.NetworkServiceDesignGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewNetworkServiceDesignGroupsServerTransport creates a new instance of NetworkServiceDesignGroupsServerTransport with the provided implementation.
// The returned NetworkServiceDesignGroupsServerTransport instance is connected to an instance of armhybridnetwork.NetworkServiceDesignGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkServiceDesignGroupsServerTransport(srv *NetworkServiceDesignGroupsServer) *NetworkServiceDesignGroupsServerTransport {
	return &NetworkServiceDesignGroupsServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientDeleteResponse]](),
		newListByPublisherPager: newTracker[azfake.PagerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientListByPublisherResponse]](),
	}
}

// NetworkServiceDesignGroupsServerTransport connects instances of armhybridnetwork.NetworkServiceDesignGroupsClient to instances of NetworkServiceDesignGroupsServer.
// Don't use this type directly, use NewNetworkServiceDesignGroupsServerTransport instead.
type NetworkServiceDesignGroupsServerTransport struct {
	srv                     *NetworkServiceDesignGroupsServer
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientDeleteResponse]]
	newListByPublisherPager *tracker[azfake.PagerResponder[armhybridnetwork.NetworkServiceDesignGroupsClientListByPublisherResponse]]
}

// Do implements the policy.Transporter interface for NetworkServiceDesignGroupsServerTransport.
func (n *NetworkServiceDesignGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkServiceDesignGroupsClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NetworkServiceDesignGroupsClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NetworkServiceDesignGroupsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkServiceDesignGroupsClient.NewListByPublisherPager":
		resp, err = n.dispatchNewListByPublisherPager(req)
	case "NetworkServiceDesignGroupsClient.Update":
		resp, err = n.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkServiceDesignGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.NetworkServiceDesignGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NetworkServiceDesignGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkServiceDesignGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkServiceDesignGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkServiceDesignGroupsServerTransport) dispatchNewListByPublisherPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByPublisherPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByPublisherPager not implemented")}
	}
	newListByPublisherPager := n.newListByPublisherPager.get(req)
	if newListByPublisherPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByPublisherPager(resourceGroupNameParam, publisherNameParam, nil)
		newListByPublisherPager = &resp
		n.newListByPublisherPager.add(req, newListByPublisherPager)
		server.PagerResponderInjectNextLinks(newListByPublisherPager, req, func(page *armhybridnetwork.NetworkServiceDesignGroupsClientListByPublisherResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByPublisherPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByPublisherPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByPublisherPager) {
		n.newListByPublisherPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkServiceDesignGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkServiceDesignGroups/(?P<networkServiceDesignGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	networkServiceDesignGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkServiceDesignGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Update(req.Context(), resourceGroupNameParam, publisherNameParam, networkServiceDesignGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkServiceDesignGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}