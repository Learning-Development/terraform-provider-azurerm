package configurationstores

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type UpdateResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Update ...
func (c ConfigurationStoresClient) Update(ctx context.Context, id ConfigurationStoreId, input ConfigurationStoreUpdateParameters) (result UpdateResponse, err error) {
	req, err := c.UpdatePreparer(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationstores.ConfigurationStoresClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = c.UpdateSender(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationstores.ConfigurationStoresClient", "Update", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateThenPoll performs Update then polls until it's completed
func (c ConfigurationStoresClient) UpdateThenPoll(ctx context.Context, id ConfigurationStoreId, input ConfigurationStoreUpdateParameters) error {
	result, err := c.Update(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Update: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Update: %+v", err)
	}

	return nil
}

// UpdatePreparer prepares the Update request.
func (c ConfigurationStoresClient) UpdatePreparer(ctx context.Context, id ConfigurationStoreId, input ConfigurationStoreUpdateParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (c ConfigurationStoresClient) UpdateSender(ctx context.Context, req *http.Request) (future UpdateResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
