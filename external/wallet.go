package external

import (
	"bytes"
	"context"
	"encoding/json"
	"ewallet-framework/helpers"
	"net/http"

	"github.com/pkg/errors"
)

type Wallet struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}

type ExWallet struct {
}

func (e *ExWallet) CreateWallet(ctx context.Context, UserID int) (*Wallet, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal json")
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_CREATE", "")

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create http request")
	}

	client := &http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send http request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("got error response from wallet service")
	}

	result := &Wallet{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.Errorf("failed to decode response body: %d", err)
	}
	defer resp.Body.Close()

	return result, nil
}
