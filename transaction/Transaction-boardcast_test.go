package transaction

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BroadCastTC struct {
	payload        BroadCastPayload
	expectedErr    error
	expectedStatus []string
}

func TestTransactionBroadcast(t *testing.T) {
	testCases := []BroadCastTC{
		{
			payload: BroadCastPayload{
				Symbol:    "ETH",
				Price:     4500,
				Timestamp: 1678912345,
			},
			expectedErr:    nil,
			expectedStatus: []string{CONFIRMED, PENDING},
		},
		{
			payload: BroadCastPayload{
				Symbol:    "BTC",
				Price:     4,
				Timestamp: 1678912345,
			},
			expectedErr:    nil,
			expectedStatus: []string{CONFIRMED, PENDING},
		},
		{
			payload: BroadCastPayload{
				Symbol:    "BNB",
				Price:     200,
				Timestamp: 1678912345,
			},
			expectedErr:    nil,
			expectedStatus: []string{CONFIRMED, PENDING},
		},
		{
			payload: BroadCastPayload{ // in reality can't be empty,but in assignment is working for empty symbol
				Symbol:    "",
				Price:     200,
				Timestamp: 1678912345,
			},
			expectedErr:    fmt.Errorf("invalid Symbol, Symbol is empty"),
			expectedStatus: []string{CONFIRMED, PENDING},
		},
		{
			payload: BroadCastPayload{
				Symbol:    "BNB",
				Price:     0,
				Timestamp: 1678912345,
			},
			expectedErr:    fmt.Errorf("invalid Price, Price can't less than or equal 0"),
			expectedStatus: []string{CONFIRMED, PENDING},
		},
		{
			payload: BroadCastPayload{
				Symbol:    "BNB",
				Price:     200,
				Timestamp: 0,
			},
			expectedErr:    fmt.Errorf("invalid timestamp, timestamp can't less than or equal 0"),
			expectedStatus: []string{CONFIRMED, PENDING},
		}}

	var wg sync.WaitGroup
	for idx, tt := range testCases {
		wg.Add(1)
		go func (idx int,tt BroadCastTC)  {
			defer wg.Done()
			t.Run(fmt.Sprintf("Test #%d", idx+1), func(t *testing.T) {
				tnxHash, err := TransactionBroadcastResty(&tt.payload)
				if tt.expectedErr != nil {
					assert.EqualError(t, err, tt.expectedErr.Error(), "error in TransactionBroadcast due to: %v", err)
				} else {
					assert.NoError(t, err, "error in TransactionBroadcast due to: %v", err)
	
					status, err := CheckTransactionStatusResty(tnxHash)
					assert.NoError(t, err, "error in CheckTransactionStatus due to: %v", err)
					assert.Contains(t, tt.expectedStatus, status)
					fmt.Println("view status", status)
				}
			})
		}(idx,tt)
		
	}

	wg.Wait()
}
