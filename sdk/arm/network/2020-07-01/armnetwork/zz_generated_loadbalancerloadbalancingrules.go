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

// LoadBalancerLoadBalancingRulesOperations contains the methods for the LoadBalancerLoadBalancingRules group.
type LoadBalancerLoadBalancingRulesOperations interface {
	// Get - Gets the specified load balancer load balancing rule.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesGetOptions) (*LoadBalancingRuleResponse, error)
	// List - Gets all the load balancing rules in a load balancer.
	List(resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesListOptions) LoadBalancerLoadBalancingRuleListResultPager
}

// LoadBalancerLoadBalancingRulesClient implements the LoadBalancerLoadBalancingRulesOperations interface.
// Don't use this type directly, use NewLoadBalancerLoadBalancingRulesClient() instead.
type LoadBalancerLoadBalancingRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerLoadBalancingRulesClient creates a new instance of LoadBalancerLoadBalancingRulesClient with the specified values.
func NewLoadBalancerLoadBalancingRulesClient(con *armcore.Connection, subscriptionID string) LoadBalancerLoadBalancingRulesOperations {
	return &LoadBalancerLoadBalancingRulesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *LoadBalancerLoadBalancingRulesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets the specified load balancer load balancing rule.
func (client *LoadBalancerLoadBalancingRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesGetOptions) (*LoadBalancingRuleResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, loadBalancerName, loadBalancingRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *LoadBalancerLoadBalancingRulesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancingRuleName}", url.PathEscape(loadBalancingRuleName))
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

// GetHandleResponse handles the Get response.
func (client *LoadBalancerLoadBalancingRulesClient) GetHandleResponse(resp *azcore.Response) (*LoadBalancingRuleResponse, error) {
	result := LoadBalancingRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancingRule)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerLoadBalancingRulesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the load balancing rules in a load balancer.
func (client *LoadBalancerLoadBalancingRulesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesListOptions) LoadBalancerLoadBalancingRuleListResultPager {
	return &loadBalancerLoadBalancingRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *LoadBalancerLoadBalancingRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerLoadBalancingRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerLoadBalancingRulesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancerLoadBalancingRulesClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerLoadBalancingRuleListResultResponse, error) {
	result := LoadBalancerLoadBalancingRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerLoadBalancingRuleListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerLoadBalancingRulesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
