package jsontime

import (
	"encoding/json"
	"testing"
	"time"

	tassert "github.com/stretchr/testify/assert"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	type testStruct struct {
		Date Time `json:"date"`
	}

	tests := []struct {
		name    string
		Time    string
		data    []byte
		wantErr bool
	}{
		{
			name: "Successfully unmarshalls a formatted time",
			Time: "2021-07-16T12:01:02-11:00",
			data: []byte(`{"date": "2021-07-16T12:01:02 -11:00"}`),
		},
		{
			name: "Fails to unmarshall a bad time",
			//Time: "2021-07-16T12:01:02-11:00",
			data:    []byte(`{"date": "2021-07-16T12:01:02-11:00"}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := tassert.New(t)
			ts := testStruct{}
			err := json.Unmarshal(tt.data, &ts)
			if tt.wantErr {
				assert.Error(err)
			}
			if tt.Time != "" {
				expectedTime, err := time.Parse(time.RFC3339, tt.Time)
				assert.NoError(err)
				assert.Equal(ts.Date.Time, expectedTime)
			}
		})
	}
}
