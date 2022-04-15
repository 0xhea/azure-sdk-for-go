//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerinstance

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationClient contains the methods for the Location group.
// Don't use this type directly, use NewLocationClient() instead.
type LocationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationClient creates a new instance of LocationClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LocationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListCachedImagesPager - Get the list of cached images on specific OS type for a subscription in a region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The identifier for the physical azure location.
// options - LocationClientListCachedImagesOptions contains the optional parameters for the LocationClient.ListCachedImages
// method.
func (client *LocationClient) NewListCachedImagesPager(location string, options *LocationClientListCachedImagesOptions) *runtime.Pager[LocationClientListCachedImagesResponse] {
	return runtime.NewPager(runtime.PageProcessor[LocationClientListCachedImagesResponse]{
		More: func(page LocationClientListCachedImagesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationClientListCachedImagesResponse) (LocationClientListCachedImagesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCachedImagesCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationClientListCachedImagesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LocationClientListCachedImagesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationClientListCachedImagesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listCachedImagesHandleResponse(resp)
		},
	})
}

// listCachedImagesCreateRequest creates the ListCachedImages request.
func (client *LocationClient) listCachedImagesCreateRequest(ctx context.Context, location string, options *LocationClientListCachedImagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/cachedImages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listCachedImagesHandleResponse handles the ListCachedImages response.
func (client *LocationClient) listCachedImagesHandleResponse(resp *http.Response) (LocationClientListCachedImagesResponse, error) {
	result := LocationClientListCachedImagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CachedImagesListResult); err != nil {
		return LocationClientListCachedImagesResponse{}, err
	}
	return result, nil
}

// NewListCapabilitiesPager - Get the list of CPU/memory/GPU capabilities of a region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The identifier for the physical azure location.
// options - LocationClientListCapabilitiesOptions contains the optional parameters for the LocationClient.ListCapabilities
// method.
func (client *LocationClient) NewListCapabilitiesPager(location string, options *LocationClientListCapabilitiesOptions) *runtime.Pager[LocationClientListCapabilitiesResponse] {
	return runtime.NewPager(runtime.PageProcessor[LocationClientListCapabilitiesResponse]{
		More: func(page LocationClientListCapabilitiesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationClientListCapabilitiesResponse) (LocationClientListCapabilitiesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCapabilitiesCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationClientListCapabilitiesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LocationClientListCapabilitiesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationClientListCapabilitiesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listCapabilitiesHandleResponse(resp)
		},
	})
}

// listCapabilitiesCreateRequest creates the ListCapabilities request.
func (client *LocationClient) listCapabilitiesCreateRequest(ctx context.Context, location string, options *LocationClientListCapabilitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/capabilities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listCapabilitiesHandleResponse handles the ListCapabilities response.
func (client *LocationClient) listCapabilitiesHandleResponse(resp *http.Response) (LocationClientListCapabilitiesResponse, error) {
	result := LocationClientListCapabilitiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilitiesListResult); err != nil {
		return LocationClientListCapabilitiesResponse{}, err
	}
	return result, nil
}

// NewListUsagePager - Get the usage for a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// location - The identifier for the physical azure location.
// options - LocationClientListUsageOptions contains the optional parameters for the LocationClient.ListUsage method.
func (client *LocationClient) NewListUsagePager(location string, options *LocationClientListUsageOptions) *runtime.Pager[LocationClientListUsageResponse] {
	return runtime.NewPager(runtime.PageProcessor[LocationClientListUsageResponse]{
		More: func(page LocationClientListUsageResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LocationClientListUsageResponse) (LocationClientListUsageResponse, error) {
			req, err := client.listUsageCreateRequest(ctx, location, options)
			if err != nil {
				return LocationClientListUsageResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LocationClientListUsageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationClientListUsageResponse{}, runtime.NewResponseError(resp)
			}
			return client.listUsageHandleResponse(resp)
		},
	})
}

// listUsageCreateRequest creates the ListUsage request.
func (client *LocationClient) listUsageCreateRequest(ctx context.Context, location string, options *LocationClientListUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listUsageHandleResponse handles the ListUsage response.
func (client *LocationClient) listUsageHandleResponse(resp *http.Response) (LocationClientListUsageResponse, error) {
	result := LocationClientListUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return LocationClientListUsageResponse{}, err
	}
	return result, nil
}
