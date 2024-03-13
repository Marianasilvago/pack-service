package functionaltest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

type PackResult struct {
	Data map[string]interface{} `json:"data"`
}

func (p PackResult) GetResult() map[int]int {
	currentData := p.Data
	changed := make(map[int]int, 0)
	for k, v := range currentData {
		intK, _ := strconv.ParseInt(fmt.Sprintf("%s", k), 10, 64)
		intV, _ := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		changed[int(intK)] = int(intV)
	}
	return changed
}

func TestLivePackHandler(t *testing.T) {
	baseURL := "http://pack-svc-service:8080"

	prepopulatePacksFor(t, baseURL, []int{250, 500, 1000, 2000, 5000})

	tests := []struct {
		description    string
		orderSize      int
		expectedPacks  map[int]int
		expectedStatus int
	}{
		{
			description:    "Order 1 item",
			orderSize:      1,
			expectedPacks:  map[int]int{250: 1},
			expectedStatus: http.StatusOK,
		},
		{
			description:    "Order 250 items",
			orderSize:      250,
			expectedPacks:  map[int]int{250: 1},
			expectedStatus: http.StatusOK,
		},
		{
			description:    "Order 251 items",
			orderSize:      251,
			expectedPacks:  map[int]int{500: 1},
			expectedStatus: http.StatusOK,
		},
		{
			description:    "Order 501 items",
			orderSize:      501,
			expectedPacks:  map[int]int{500: 1, 250: 1},
			expectedStatus: http.StatusOK,
		},
		{
			description:    "Order 12001 items",
			orderSize:      12001,
			expectedPacks:  map[int]int{5000: 2, 2000: 1, 250: 1},
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			url := fmt.Sprintf("%s/calculate-packs", baseURL)
			payloadBytes, _ := json.Marshal(struct{ OrderSize int }{OrderSize: test.orderSize})
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
			if err != nil {
				t.Fatalf("Could not create HTTP request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Could not execute HTTP request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != test.expectedStatus {
				t.Errorf("Expected status code %d, got %d for order size %d", test.expectedStatus, resp.StatusCode, test.orderSize)
				return
			}

			var packResult PackResult
			if err := json.NewDecoder(resp.Body).Decode(&packResult); err != nil {
				t.Fatalf("Could not decode response: %v", err)
			}
			actualResult := packResult.GetResult()

			if !isMapEqual(actualResult, test.expectedPacks) {
				t.Errorf("For order size %d, expected packs %v, got %v", test.orderSize, test.expectedPacks, actualResult)
			}
		})
	}
}

func prepopulatePacksFor(t *testing.T, baseURL string, sizes []int) {
	for _, size := range sizes {
		url := fmt.Sprintf("%s/pack-sizes", baseURL)
		payloadBytes, _ := json.Marshal(struct{ Size int }{Size: size})
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
		if err != nil {
			t.Fatalf("Could not create HTTP request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Could not execute HTTP request: %v", err)
		}
		defer resp.Body.Close()
	}
}

func isMapEqual(map1, map2 map[int]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, valueMap1 := range map1 {
		if valueMap2, ok := map2[key]; !ok || valueMap1 != valueMap2 {

			return false
		}
	}

	return true
}
