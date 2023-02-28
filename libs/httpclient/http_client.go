package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func Send[Body any, Resp any](ctx context.Context, url string, body Body) (Resp, error) {
	var response Resp

	rawJSON, err := json.Marshal(body)
	if err != nil {
		return response, errors.Wrap(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(rawJSON))
	if err != nil {
		return response, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return response, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return response, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return response, errors.Wrap(err, "decoding json")
	}

	return response, nil
}
