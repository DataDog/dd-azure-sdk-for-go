//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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

// OperationStatusesClient contains the methods for the OperationStatuses group.
// Don't use this type directly, use NewOperationStatusesClient() instead.
type OperationStatusesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationStatusesClient creates a new instance of OperationStatusesClient with the specified values.
// subscriptionID - Unique identifier of the Azure subscription. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationStatusesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationStatusesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &OperationStatusesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the current status of an async operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
// location - The Azure region
// operationID - The ID of an ongoing async operation
// options - OperationStatusesClientGetOptions contains the optional parameters for the OperationStatusesClient.Get method.
func (client *OperationStatusesClient) Get(ctx context.Context, location string, operationID string, options *OperationStatusesClientGetOptions) (OperationStatusesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, operationID, options)
	if err != nil {
		return OperationStatusesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationStatusesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return OperationStatusesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationStatusesClient) getCreateRequest(ctx context.Context, location string, operationID string, options *OperationStatusesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevCenter/locations/{location}/operationStatuses/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationStatusesClient) getHandleResponse(resp *http.Response) (OperationStatusesClientGetResponse, error) {
	result := OperationStatusesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return OperationStatusesClientGetResponse{}, err
	}
	return result, nil
}
