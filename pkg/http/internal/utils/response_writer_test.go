package utils_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pack-svc/pkg/http/internal/resperr"
	"pack-svc/pkg/http/internal/utils"
	"testing"
)

func TestWriteSuccessResponse(t *testing.T) {
	testCases := map[string]struct {
		inputData    interface{}
		expectedCode int
		expectedResp string
	}{
		"write success response success": {
			inputData: []struct {
				K string
				V string
			}{{"resp_id", "resp-id"}, {"resp_data", "resp data"}},
			expectedCode: http.StatusCreated,
			expectedResp: `{"data":[{"K":"resp_id","V":"resp-id"},{"K":"resp_data","V":"resp data"}],"success":true}`,
		},
		"write success response failure": {
			inputData:    make(chan int),
			expectedCode: http.StatusInternalServerError,
			expectedResp: "server error",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testWriteSuccessResponse(t, testCase.expectedResp, testCase.expectedCode, testCase.inputData)
		})
	}
}

func TestWriteFailureResponse(t *testing.T) {
	err := resperr.NewResponseError(http.StatusBadRequest, "failed to parse")

	w := httptest.NewRecorder()

	utils.WriteFailureResponse(w, err)

	expectedCode := http.StatusBadRequest
	expectedResp := "{\"error\":{\"description\":\"failed to parse\"},\"success\":false}"

	assert.Equal(t, expectedCode, w.Code)
	assert.Equal(t, expectedResp, w.Body.String())
}

func testWriteSuccessResponse(t *testing.T, expectedResp string, expectedCode int, inputData interface{}) {
	w := httptest.NewRecorder()

	utils.WriteSuccessResponse(w, http.StatusCreated, inputData)

	assert.Equal(t, expectedCode, w.Code)
	assert.Equal(t, expectedResp, w.Body.String())
}
