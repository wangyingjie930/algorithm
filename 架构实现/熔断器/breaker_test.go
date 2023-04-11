/**
  @author: wangyingjie
  @since: 2023/4/11
  @desc:
**/

package breaker

import (
	"fmt"
	"testing"
	"time"
)

func TestCircuitBreaker_Call(t *testing.T) {
	type fields struct {
		State                   State
		TripThreshold           int
		SuccessThreshold        int
		ConsecutiveSuccessCount int
		ConsecutiveErrorCount   int
		SleepWindow             time.Duration
		LastAttemptTime         time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "case1", fields: fields{
			State:                   Closed,
			TripThreshold:           3,
			SuccessThreshold:        2,
			ConsecutiveSuccessCount: 0,
			ConsecutiveErrorCount:   0,
			SleepWindow:             5 * time.Second,
			LastAttemptTime:         time.Now(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cb := &CircuitBreaker{
				State:                   tt.fields.State,
				TripThreshold:           tt.fields.TripThreshold,
				SuccessThreshold:        tt.fields.SuccessThreshold,
				ConsecutiveSuccessCount: tt.fields.ConsecutiveSuccessCount,
				ConsecutiveErrorCount:   tt.fields.ConsecutiveErrorCount,
				SleepWindow:             tt.fields.SleepWindow,
				LastAttemptTime:         tt.fields.LastAttemptTime,
			}

			// 进行 10 次服务调用
			for i := 0; i < 10; i++ {
				fmt.Printf("Call %d: ", i+1)
				err := cb.Call()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				fmt.Println()
				time.Sleep(500 * time.Millisecond)
			}
		})
	}
}
