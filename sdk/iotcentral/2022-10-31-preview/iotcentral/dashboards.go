package iotcentral

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// DashboardsClient is the azure IoT Central is a service that makes it easy to connect, monitor, and manage your IoT
// devices at scale.
type DashboardsClient struct {
	BaseClient
}

// NewDashboardsClient creates an instance of the DashboardsClient client.
func NewDashboardsClient(subdomain string) DashboardsClient {
	return DashboardsClient{New(subdomain)}
}

// Create sends the create request.
// Parameters:
// dashboardID - unique ID for the dashboard.
// body - dashboard definition.
func (client DashboardsClient) Create(ctx context.Context, dashboardID string, body Dashboard) (result Dashboard, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dashboardID,
			Constraints: []validation.Constraint{{Target: "dashboardID", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "dashboardID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9:;]*$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.DisplayName", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.DisplayName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
				{Target: "body.Tiles", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Tiles", Name: validation.MaxItems, Rule: 100, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("iotcentral.DashboardsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, dashboardID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DashboardsClient) CreatePreparer(ctx context.Context, dashboardID string, body Dashboard) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"dashboardId": autorest.Encode("path", dashboardID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Personal = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/dashboards/{dashboardId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DashboardsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DashboardsClient) CreateResponder(resp *http.Response) (result Dashboard, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// dashboardID - unique ID for the dashboard.
func (client DashboardsClient) Get(ctx context.Context, dashboardID string) (result Dashboard, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dashboardID,
			Constraints: []validation.Constraint{{Target: "dashboardID", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "dashboardID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9:;]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DashboardsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, dashboardID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DashboardsClient) GetPreparer(ctx context.Context, dashboardID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"dashboardId": autorest.Encode("path", dashboardID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/dashboards/{dashboardId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DashboardsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DashboardsClient) GetResponder(resp *http.Response) (result Dashboard, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// filter - an expression on the resource type that selects the resources to be returned.
// maxpagesize - the maximum number of resources to return from one response.
// orderby - an expression that specify the order of the returned resources.
func (client DashboardsClient) List(ctx context.Context, filter string, maxpagesize *int32, orderby string) (result DashboardCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardsClient.List")
		defer func() {
			sc := -1
			if result.dc.Response.Response != nil {
				sc = result.dc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxpagesize,
			Constraints: []validation.Constraint{{Target: "maxpagesize", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxpagesize", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
					{Target: "maxpagesize", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("iotcentral.DashboardsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, maxpagesize, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "List", resp, "Failure sending request")
		return
	}

	result.dc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dc.hasNextLink() && result.dc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DashboardsClient) ListPreparer(ctx context.Context, filter string, maxpagesize *int32, orderby string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["filter"] = autorest.Encode("query", filter)
	}
	if maxpagesize != nil {
		queryParameters["maxpagesize"] = autorest.Encode("query", *maxpagesize)
	}
	if len(orderby) > 0 {
		queryParameters["orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPath("/dashboards"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DashboardsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DashboardsClient) ListResponder(resp *http.Response) (result DashboardCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DashboardsClient) listNextResults(ctx context.Context, lastResults DashboardCollection) (result DashboardCollection, err error) {
	req, err := lastResults.dashboardCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DashboardsClient) ListComplete(ctx context.Context, filter string, maxpagesize *int32, orderby string) (result DashboardCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, maxpagesize, orderby)
	return
}

// Remove sends the remove request.
// Parameters:
// dashboardID - unique ID for the dashboard.
func (client DashboardsClient) Remove(ctx context.Context, dashboardID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardsClient.Remove")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dashboardID,
			Constraints: []validation.Constraint{{Target: "dashboardID", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "dashboardID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9:;]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DashboardsClient", "Remove", err.Error())
	}

	req, err := client.RemovePreparer(ctx, dashboardID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Remove", resp, "Failure responding to request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client DashboardsClient) RemovePreparer(ctx context.Context, dashboardID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"dashboardId": autorest.Encode("path", dashboardID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/dashboards/{dashboardId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client DashboardsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client DashboardsClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update sends the update request.
// Parameters:
// dashboardID - unique ID for the dashboard.
// body - dashboard definition.
// ifMatch - only perform the operation if the entity's etag matches one of the etags provided or * is
// provided.
func (client DashboardsClient) Update(ctx context.Context, dashboardID string, body interface{}, ifMatch string) (result Dashboard, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dashboardID,
			Constraints: []validation.Constraint{{Target: "dashboardID", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "dashboardID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9:;]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DashboardsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, dashboardID, body, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DashboardsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DashboardsClient) UpdatePreparer(ctx context.Context, dashboardID string, body interface{}, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"dashboardId": autorest.Encode("path", dashboardID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/merge-patch+json"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/dashboards/{dashboardId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DashboardsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DashboardsClient) UpdateResponder(resp *http.Response) (result Dashboard, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}