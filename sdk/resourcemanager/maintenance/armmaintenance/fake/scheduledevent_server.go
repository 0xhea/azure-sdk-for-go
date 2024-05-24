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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"net/http"
	"net/url"
	"regexp"
)

// ScheduledEventServer is a fake server for instances of the armmaintenance.ScheduledEventClient type.
type ScheduledEventServer struct {
	// Acknowledge is the fake for method ScheduledEventClient.Acknowledge
	// HTTP status codes to indicate success: http.StatusOK
	Acknowledge func(ctx context.Context, resourceGroupName string, resourceType string, resourceName string, scheduledEventID string, options *armmaintenance.ScheduledEventClientAcknowledgeOptions) (resp azfake.Responder[armmaintenance.ScheduledEventClientAcknowledgeResponse], errResp azfake.ErrorResponder)
}

// NewScheduledEventServerTransport creates a new instance of ScheduledEventServerTransport with the provided implementation.
// The returned ScheduledEventServerTransport instance is connected to an instance of armmaintenance.ScheduledEventClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScheduledEventServerTransport(srv *ScheduledEventServer) *ScheduledEventServerTransport {
	return &ScheduledEventServerTransport{srv: srv}
}

// ScheduledEventServerTransport connects instances of armmaintenance.ScheduledEventClient to instances of ScheduledEventServer.
// Don't use this type directly, use NewScheduledEventServerTransport instead.
type ScheduledEventServerTransport struct {
	srv *ScheduledEventServer
}

// Do implements the policy.Transporter interface for ScheduledEventServerTransport.
func (s *ScheduledEventServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ScheduledEventClient.Acknowledge":
		resp, err = s.dispatchAcknowledge(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ScheduledEventServerTransport) dispatchAcknowledge(req *http.Request) (*http.Response, error) {
	if s.srv.Acknowledge == nil {
		return nil, &nonRetriableError{errors.New("fake for method Acknowledge not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/scheduledevents/(?P<scheduledEventId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/acknowledge`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	scheduledEventIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduledEventId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Acknowledge(req.Context(), resourceGroupNameParam, resourceTypeParam, resourceNameParam, scheduledEventIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScheduledEventApproveResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}