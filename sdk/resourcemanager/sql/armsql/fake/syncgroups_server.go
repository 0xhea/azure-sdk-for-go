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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// SyncGroupsServer is a fake server for instances of the armsql.SyncGroupsClient type.
type SyncGroupsServer struct {
	// CancelSync is the fake for method SyncGroupsClient.CancelSync
	// HTTP status codes to indicate success: http.StatusOK
	CancelSync func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *armsql.SyncGroupsClientCancelSyncOptions) (resp azfake.Responder[armsql.SyncGroupsClientCancelSyncResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method SyncGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, parameters armsql.SyncGroup, options *armsql.SyncGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.SyncGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SyncGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *armsql.SyncGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.SyncGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SyncGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *armsql.SyncGroupsClientGetOptions) (resp azfake.Responder[armsql.SyncGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method SyncGroupsClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.SyncGroupsClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.SyncGroupsClientListByDatabaseResponse])

	// NewListHubSchemasPager is the fake for method SyncGroupsClient.NewListHubSchemasPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListHubSchemasPager func(resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *armsql.SyncGroupsClientListHubSchemasOptions) (resp azfake.PagerResponder[armsql.SyncGroupsClientListHubSchemasResponse])

	// NewListLogsPager is the fake for method SyncGroupsClient.NewListLogsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListLogsPager func(resourceGroupName string, serverName string, databaseName string, syncGroupName string, startTime string, endTime string, typeParam armsql.SyncGroupsType, options *armsql.SyncGroupsClientListLogsOptions) (resp azfake.PagerResponder[armsql.SyncGroupsClientListLogsResponse])

	// NewListSyncDatabaseIDsPager is the fake for method SyncGroupsClient.NewListSyncDatabaseIDsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSyncDatabaseIDsPager func(locationName string, options *armsql.SyncGroupsClientListSyncDatabaseIDsOptions) (resp azfake.PagerResponder[armsql.SyncGroupsClientListSyncDatabaseIDsResponse])

	// BeginRefreshHubSchema is the fake for method SyncGroupsClient.BeginRefreshHubSchema
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefreshHubSchema func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *armsql.SyncGroupsClientBeginRefreshHubSchemaOptions) (resp azfake.PollerResponder[armsql.SyncGroupsClientRefreshHubSchemaResponse], errResp azfake.ErrorResponder)

	// TriggerSync is the fake for method SyncGroupsClient.TriggerSync
	// HTTP status codes to indicate success: http.StatusOK
	TriggerSync func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, options *armsql.SyncGroupsClientTriggerSyncOptions) (resp azfake.Responder[armsql.SyncGroupsClientTriggerSyncResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method SyncGroupsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, parameters armsql.SyncGroup, options *armsql.SyncGroupsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsql.SyncGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSyncGroupsServerTransport creates a new instance of SyncGroupsServerTransport with the provided implementation.
// The returned SyncGroupsServerTransport instance is connected to an instance of armsql.SyncGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSyncGroupsServerTransport(srv *SyncGroupsServer) *SyncGroupsServerTransport {
	return &SyncGroupsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armsql.SyncGroupsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armsql.SyncGroupsClientDeleteResponse]](),
		newListByDatabasePager:      newTracker[azfake.PagerResponder[armsql.SyncGroupsClientListByDatabaseResponse]](),
		newListHubSchemasPager:      newTracker[azfake.PagerResponder[armsql.SyncGroupsClientListHubSchemasResponse]](),
		newListLogsPager:            newTracker[azfake.PagerResponder[armsql.SyncGroupsClientListLogsResponse]](),
		newListSyncDatabaseIDsPager: newTracker[azfake.PagerResponder[armsql.SyncGroupsClientListSyncDatabaseIDsResponse]](),
		beginRefreshHubSchema:       newTracker[azfake.PollerResponder[armsql.SyncGroupsClientRefreshHubSchemaResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armsql.SyncGroupsClientUpdateResponse]](),
	}
}

// SyncGroupsServerTransport connects instances of armsql.SyncGroupsClient to instances of SyncGroupsServer.
// Don't use this type directly, use NewSyncGroupsServerTransport instead.
type SyncGroupsServerTransport struct {
	srv                         *SyncGroupsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armsql.SyncGroupsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armsql.SyncGroupsClientDeleteResponse]]
	newListByDatabasePager      *tracker[azfake.PagerResponder[armsql.SyncGroupsClientListByDatabaseResponse]]
	newListHubSchemasPager      *tracker[azfake.PagerResponder[armsql.SyncGroupsClientListHubSchemasResponse]]
	newListLogsPager            *tracker[azfake.PagerResponder[armsql.SyncGroupsClientListLogsResponse]]
	newListSyncDatabaseIDsPager *tracker[azfake.PagerResponder[armsql.SyncGroupsClientListSyncDatabaseIDsResponse]]
	beginRefreshHubSchema       *tracker[azfake.PollerResponder[armsql.SyncGroupsClientRefreshHubSchemaResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armsql.SyncGroupsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for SyncGroupsServerTransport.
func (s *SyncGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SyncGroupsClient.CancelSync":
		resp, err = s.dispatchCancelSync(req)
	case "SyncGroupsClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SyncGroupsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SyncGroupsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SyncGroupsClient.NewListByDatabasePager":
		resp, err = s.dispatchNewListByDatabasePager(req)
	case "SyncGroupsClient.NewListHubSchemasPager":
		resp, err = s.dispatchNewListHubSchemasPager(req)
	case "SyncGroupsClient.NewListLogsPager":
		resp, err = s.dispatchNewListLogsPager(req)
	case "SyncGroupsClient.NewListSyncDatabaseIDsPager":
		resp, err = s.dispatchNewListSyncDatabaseIDsPager(req)
	case "SyncGroupsClient.BeginRefreshHubSchema":
		resp, err = s.dispatchBeginRefreshHubSchema(req)
	case "SyncGroupsClient.TriggerSync":
		resp, err = s.dispatchTriggerSync(req)
	case "SyncGroupsClient.BeginUpdate":
		resp, err = s.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchCancelSync(req *http.Request) (*http.Response, error) {
	if s.srv.CancelSync == nil {
		return nil, &nonRetriableError{errors.New("fake for method CancelSync not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancelSync`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CancelSync(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, nil)
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

func (s *SyncGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.SyncGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SyncGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := s.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		s.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.SyncGroupsClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		s.newListByDatabasePager.remove(req)
	}
	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchNewListHubSchemasPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListHubSchemasPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListHubSchemasPager not implemented")}
	}
	newListHubSchemasPager := s.newListHubSchemasPager.get(req)
	if newListHubSchemasPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hubSchemas`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListHubSchemasPager(resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, nil)
		newListHubSchemasPager = &resp
		s.newListHubSchemasPager.add(req, newListHubSchemasPager)
		server.PagerResponderInjectNextLinks(newListHubSchemasPager, req, func(page *armsql.SyncGroupsClientListHubSchemasResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListHubSchemasPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListHubSchemasPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListHubSchemasPager) {
		s.newListHubSchemasPager.remove(req)
	}
	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchNewListLogsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListLogsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListLogsPager not implemented")}
	}
	newListLogsPager := s.newListLogsPager.get(req)
	if newListLogsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/logs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
		if err != nil {
			return nil, err
		}
		startTimeParam, err := url.QueryUnescape(qp.Get("startTime"))
		if err != nil {
			return nil, err
		}
		endTimeParam, err := url.QueryUnescape(qp.Get("endTime"))
		if err != nil {
			return nil, err
		}
		typeParamParam, err := parseWithCast(qp.Get("type"), func(v string) (armsql.SyncGroupsType, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.SyncGroupsType(p), nil
		})
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armsql.SyncGroupsClientListLogsOptions
		if continuationTokenParam != nil {
			options = &armsql.SyncGroupsClientListLogsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := s.srv.NewListLogsPager(resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, startTimeParam, endTimeParam, typeParamParam, options)
		newListLogsPager = &resp
		s.newListLogsPager.add(req, newListLogsPager)
		server.PagerResponderInjectNextLinks(newListLogsPager, req, func(page *armsql.SyncGroupsClientListLogsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListLogsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListLogsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListLogsPager) {
		s.newListLogsPager.remove(req)
	}
	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchNewListSyncDatabaseIDsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListSyncDatabaseIDsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSyncDatabaseIDsPager not implemented")}
	}
	newListSyncDatabaseIDsPager := s.newListSyncDatabaseIDsPager.get(req)
	if newListSyncDatabaseIDsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncDatabaseIds`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListSyncDatabaseIDsPager(locationNameParam, nil)
		newListSyncDatabaseIDsPager = &resp
		s.newListSyncDatabaseIDsPager.add(req, newListSyncDatabaseIDsPager)
		server.PagerResponderInjectNextLinks(newListSyncDatabaseIDsPager, req, func(page *armsql.SyncGroupsClientListSyncDatabaseIDsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSyncDatabaseIDsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListSyncDatabaseIDsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSyncDatabaseIDsPager) {
		s.newListSyncDatabaseIDsPager.remove(req)
	}
	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchBeginRefreshHubSchema(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRefreshHubSchema == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefreshHubSchema not implemented")}
	}
	beginRefreshHubSchema := s.beginRefreshHubSchema.get(req)
	if beginRefreshHubSchema == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshHubSchema`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginRefreshHubSchema(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefreshHubSchema = &respr
		s.beginRefreshHubSchema.add(req, beginRefreshHubSchema)
	}

	resp, err := server.PollerResponderNext(beginRefreshHubSchema, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginRefreshHubSchema.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefreshHubSchema) {
		s.beginRefreshHubSchema.remove(req)
	}

	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchTriggerSync(req *http.Request) (*http.Response, error) {
	if s.srv.TriggerSync == nil {
		return nil, &nonRetriableError{errors.New("fake for method TriggerSync not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerSync`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.TriggerSync(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, nil)
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

func (s *SyncGroupsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.SyncGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, syncGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}