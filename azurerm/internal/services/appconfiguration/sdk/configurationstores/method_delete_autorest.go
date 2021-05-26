package configurationstores

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type DeleteResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Delete ...
func (c ConfigurationStoresClient) Delete(ctx context.Context, id ConfigurationStoreId) (result DeleteResponse, err error) {
	req, err := c.DeletePreparer(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationstores.ConfigurationStoresClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = c.DeleteSender(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationstores.ConfigurationStoresClient", "Delete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteThenPoll performs Delete then polls until it's completed
func (c ConfigurationStoresClient) DeleteThenPoll(ctx context.Context, id ConfigurationStoreId) error {
	result, err := c.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Delete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Delete: %+v", err)
	}

	return nil
}

// DeletePreparer prepares the Delete request.
func (c ConfigurationStoresClient) DeletePreparer(ctx context.Context, id ConfigurationStoreId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (c ConfigurationStoresClient) DeleteSender(ctx context.Context, req *http.Request) (future DeleteResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
