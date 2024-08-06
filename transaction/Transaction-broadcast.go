package transaction

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type BroadCastPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type BroadCaseResponse struct {
	TxHash string `json:"tx_hash"`
}
type TransactionResponse struct {
	TxStatus string `json:"tx_status"`
}

const URL_BROADCAST = "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"

const (
	URL_BROADCAST_CHECK = "https://mock-node-wgqbnxruha-as.a.run.app/check/%s"
	CONFIRMED           = "CONFIRMED"
	FAILED              = "FAILED"
	PENDING             = "PENDING"
	DNE                 = "DNE"
)
const (
	RETRY_COUNT             = 3
	SET_RETRY_WAIT_TIME     = 12 * time.Second //SetRetryWaitTime method sets default wait time to sleep before retrying request.
)

func TransactionBroadcastResty(payload *BroadCastPayload) (string, error) {
	err := validatePayload(payload)
	if err != nil {
		return "", err
	}
	client := resty.New()
	var response BroadCaseResponse
	_, err = client.R().
		SetBody(payload).
		SetResult(&response).
		Post(URL_BROADCAST)
	fmt.Println("requested with payload :", payload)
	if err != nil {
		fmt.Println("Error while resty client call :: ", err)
		return "", err
	}

	return response.TxHash, nil
}
func CheckTransactionStatusResty(tnxHash string) (string, error) {
	if len(strings.TrimSpace(tnxHash)) == 0 {
		return "", fmt.Errorf("tnxHash is empty")
	}
	url := fmt.Sprintf(URL_BROADCAST_CHECK, tnxHash)
	var response TransactionResponse
	client := resty.New()
	client.
		SetRetryCount(RETRY_COUNT).
		SetRetryWaitTime(SET_RETRY_WAIT_TIME).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				if err != nil {
					return false
				}
				var res TransactionResponse
				if err := json.Unmarshal(r.Body(), &res); err != nil {
					return false
				}
				return res.TxStatus == PENDING || res.TxStatus == FAILED
			},
		)
	_, err := client.R().
		SetResult(&response).
		Get(url)
	if err != nil {
		fmt.Println("Error while resty client call :: ", err)
		return "", err
	}
	return response.TxStatus, nil
}

func validatePayload(payload *BroadCastPayload) error {
	if len(strings.TrimSpace(payload.Symbol)) == 0 {
		return fmt.Errorf("invalid Symbol, Symbol is empty")
	}
	if payload.Price <= 0 {
		return fmt.Errorf("invalid Price, Price can't less than or equal 0")
	}
	if payload.Timestamp <= 0 {
		return fmt.Errorf("invalid timestamp, timestamp can't less than or equal 0")
	}
	return nil
}

// func TransactionBroadcast(payload BroadCastPayload) (string, error) {
// 	jsonPayload, err := json.Marshal(payload)
// 	if err != nil {
// 		fmt.Println("Error while unmarshal :: ", err)
// 		return "", err
// 	}
// 	resp, err := http.Post(URL_BROADCAST, CONTENT_TYPE, bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		fmt.Println("Error while broadcasting :: ", err)
// 		return "", err
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error while reading server response :: ", err)
// 		return "", err
// 	}
// 	fmt.Println("POST", string(body))
// 	var response BroadCaseResponse
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		fmt.Println("Error while unmarshal  :: ", err)
// 		return "", err
// 	}
// 	return response.TxHash, nil
// }

// func CheckTransactionStatus(tnxHash string) (string, error) {
// 	url := fmt.Sprintf(URL_BROADCAST_CHECK, tnxHash)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error while broadcasting :: ", err)
// 		return "", err
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error while reading server response :: ", err)
// 		return "", err
// 	}
// 	fmt.Println("GET", string(body))
// 	var response TransactionResponse
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		fmt.Println("Error while unmarshal  :: ", err)
// 		return "", err
// 	}
// 	return response.TxStatus, nil
// }
