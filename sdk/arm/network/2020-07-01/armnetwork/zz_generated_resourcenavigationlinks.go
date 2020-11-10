// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ResourceNavigationLinksOperations contains the methods for the ResourceNavigationLinks group.
type ResourceNavigationLinksOperations interface {
	// List - Gets a list of resource navigation links for a subnet.
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksListOptions) (*ResourceNavigationLinksListResultResponse, error)
}

// ResourceNavigationLinksClient implements the ResourceNavigationLinksOperations interface.
// Don't use this type directly, use NewResourceNavigationLinksClient() instead.
type ResourceNavigationLinksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewResourceNavigationLinksClient creates a new instance of ResourceNavigationLinksClient with the specified values.
func NewResourceNavigationLinksClient(con *armcore.Connection, subscriptionID string) ResourceNavigationLinksOperations {
	return &ResourceNavigationLinksClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *ResourceNavigationLinksClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// List - Gets a list of resource navigation links for a subnet.
func (client *ResourceNavigationLinksClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksListOptions) (*ResourceNavigationLinksListResultResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *ResourceNavigationLinksClient) ListCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *ResourceNavigationLinksClient) ListHandleResponse(resp *azcore.Response) (*ResourceNavigationLinksListResultResponse, error) {
	result := ResourceNavigationLinksListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ResourceNavigationLinksListResult)
}

// ListHandleError handles the List error response.
func (client *ResourceNavigationLinksClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
