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
	"time"
)

// VpnConnectionsOperations contains the methods for the VpnConnections group.
type VpnConnectionsOperations interface {
	// BeginCreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsCreateOrUpdateOptions) (*VpnConnectionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VpnConnectionPoller, error)
	// BeginDelete - Deletes a vpn connection.
	BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a vpn connection.
	Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsGetOptions) (*VpnConnectionResponse, error)
	// ListByVpnGateway - Retrieves all vpn connections for a particular virtual wan vpn gateway.
	ListByVpnGateway(resourceGroupName string, gatewayName string, options *VpnConnectionsListByVpnGatewayOptions) ListVpnConnectionsResultPager
	// BeginStartPacketCapture - Starts packet capture on Vpn connection in the specified resource group.
	BeginStartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error)
	// ResumeStartPacketCapture - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStartPacketCapture(token string) (StringPoller, error)
	// BeginStopPacketCapture - Stops packet capture on Vpn connection in the specified resource group.
	BeginStopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStopPacketCaptureOptions) (*StringPollerResponse, error)
	// ResumeStopPacketCapture - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStopPacketCapture(token string) (StringPoller, error)
}

// VpnConnectionsClient implements the VpnConnectionsOperations interface.
// Don't use this type directly, use NewVpnConnectionsClient() instead.
type VpnConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVpnConnectionsClient creates a new instance of VpnConnectionsClient with the specified values.
func NewVpnConnectionsClient(con *armcore.Connection, subscriptionID string) VpnConnectionsOperations {
	return &VpnConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VpnConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VpnConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsCreateOrUpdateOptions) (*VpnConnectionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VpnConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnConnectionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnConnectionsClient) ResumeCreateOrUpdate(token string) (VpnConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnConnectionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
func (client *VpnConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VpnConnectionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnConnectionParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VpnConnectionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VpnConnectionResponse, error) {
	result := VpnConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnConnection)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VpnConnectionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VpnConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnConnectionsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnConnectionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnConnectionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a vpn connection.
func (client *VpnConnectionsClient) Delete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *VpnConnectionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// DeleteHandleError handles the Delete error response.
func (client *VpnConnectionsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a vpn connection.
func (client *VpnConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsGetOptions) (*VpnConnectionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
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
func (client *VpnConnectionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *VpnConnectionsClient) GetHandleResponse(resp *azcore.Response) (*VpnConnectionResponse, error) {
	result := VpnConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnConnection)
}

// GetHandleError handles the Get error response.
func (client *VpnConnectionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByVpnGateway - Retrieves all vpn connections for a particular virtual wan vpn gateway.
func (client *VpnConnectionsClient) ListByVpnGateway(resourceGroupName string, gatewayName string, options *VpnConnectionsListByVpnGatewayOptions) ListVpnConnectionsResultPager {
	return &listVpnConnectionsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByVpnGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
		},
		responder: client.ListByVpnGatewayHandleResponse,
		errorer:   client.ListByVpnGatewayHandleError,
		advancer: func(ctx context.Context, resp *ListVpnConnectionsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnConnectionsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByVpnGatewayCreateRequest creates the ListByVpnGateway request.
func (client *VpnConnectionsClient) ListByVpnGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnConnectionsListByVpnGatewayOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// ListByVpnGatewayHandleResponse handles the ListByVpnGateway response.
func (client *VpnConnectionsClient) ListByVpnGatewayHandleResponse(resp *azcore.Response) (*ListVpnConnectionsResultResponse, error) {
	result := ListVpnConnectionsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnConnectionsResult)
}

// ListByVpnGatewayHandleError handles the ListByVpnGateway error response.
func (client *VpnConnectionsClient) ListByVpnGatewayHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VpnConnectionsClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error) {
	resp, err := client.StartPacketCapture(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
	if err != nil {
		return nil, err
	}
	result := &StringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnConnectionsClient.StartPacketCapture", "location", resp, client.StartPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnConnectionsClient) ResumeStartPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnConnectionsClient.StartPacketCapture", token, client.StartPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// StartPacketCapture - Starts packet capture on Vpn connection in the specified resource group.
func (client *VpnConnectionsClient) StartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStartPacketCaptureOptions) (*azcore.Response, error) {
	req, err := client.StartPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.StartPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// StartPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VpnConnectionsClient) StartPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStartPacketCaptureOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/startpacketcapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnConnectionName}", url.PathEscape(vpnConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Parameters)
	}
	return req, nil
}

// StartPacketCaptureHandleResponse handles the StartPacketCapture response.
func (client *VpnConnectionsClient) StartPacketCaptureHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// StartPacketCaptureHandleError handles the StartPacketCapture error response.
func (client *VpnConnectionsClient) StartPacketCaptureHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VpnConnectionsClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStopPacketCaptureOptions) (*StringPollerResponse, error) {
	resp, err := client.StopPacketCapture(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
	if err != nil {
		return nil, err
	}
	result := &StringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnConnectionsClient.StopPacketCapture", "location", resp, client.StopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnConnectionsClient) ResumeStopPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnConnectionsClient.StopPacketCapture", token, client.StopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// StopPacketCapture - Stops packet capture on Vpn connection in the specified resource group.
func (client *VpnConnectionsClient) StopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStopPacketCaptureOptions) (*azcore.Response, error) {
	req, err := client.StopPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.StopPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// StopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VpnConnectionsClient) StopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VpnConnectionsStopPacketCaptureOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/stoppacketcapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnConnectionName}", url.PathEscape(vpnConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Parameters)
	}
	return req, nil
}

// StopPacketCaptureHandleResponse handles the StopPacketCapture response.
func (client *VpnConnectionsClient) StopPacketCaptureHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// StopPacketCaptureHandleError handles the StopPacketCapture error response.
func (client *VpnConnectionsClient) StopPacketCaptureHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
