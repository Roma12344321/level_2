package main

import (
	"errors"
	"testing"
	"time"
)

type MockNTPClient struct {
	time time.Time
	err  error
}

func (m *MockNTPClient) GetTime() (time.Time, error) {
	return m.time, m.err
}

func TestTask1(t *testing.T) {
	tests := []struct {
		name       string
		mockClient NTPClient
		wantErr    bool
	}{
		{
			name: "successful time fetch",
			mockClient: &MockNTPClient{
				time: time.Date(2022, time.June, 12, 10, 0, 0, 0, time.UTC),
				err:  nil,
			},
			wantErr: false,
		},
		{
			name: "error fetching time",
			mockClient: &MockNTPClient{
				time: time.Time{},
				err:  errors.New("network error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		client := tt.mockClient
		currentTime, err := client.GetTime()
		if (err != nil) != tt.wantErr {
			t.Errorf("test %s failed: expected error: %v, got: %v", tt.name, tt.wantErr, err)
		}
		if !tt.wantErr && currentTime.IsZero() {
			t.Errorf("test %s failed: expected valid time, got zero time", tt.name)
		}
	}
}
