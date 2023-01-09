package operationtype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leonardogazio/golang-account-crud/db"
	"github.com/leonardogazio/golang-account-crud/models"
	"github.com/leonardogazio/golang-account-crud/server/routes/handlers"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestMain(m *testing.M) {
	handlers.Init(m)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	qry := "INSERT INTO operation_types"
	pur := "Routes: HTTP 1.1 POST /operation_types"

	tt := []struct {
		purpose        string
		mock           func()
		req            models.OperationType
		wantRes        models.Response
		wantStatusCode int
	}{
		{
			purpose: fmt.Sprintf("%s - Success Case /Should create.", pur),
			req:     models.OperationType{Description: "COMPRA PARCELADA"},
			mock: func() {
				db.Mock.ExpectBegin()
				db.Mock.ExpectExec(qry).WithArgs("COMPRA PARCELADA").WillReturnResult(sqlxmock.NewResult(1, 1))
				db.Mock.ExpectCommit()
			},
			wantRes:        models.Response{Message: "Successfully Created!"},
			wantStatusCode: http.StatusCreated,
		},
		{
			purpose: fmt.Sprintf("%s - Error Case /Should fail validation.", pur),
			req:     models.OperationType{},
			wantRes: models.Response{
				Message: "Invalid Params",
				Data: map[string]interface{}{
					"validations": []interface{}{
						map[string]interface{}{
							"field":   "Description",
							"message": "Field 'Description' is 'required'.",
						},
					},
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			purpose: fmt.Sprintf("%s - Error Case /Should return unique constraint violate.", pur),
			req:     models.OperationType{Description: "COMPRA PARCELADA"},
			mock: func() {
				db.Mock.ExpectBegin()
				db.Mock.ExpectExec(qry).WithArgs("COMPRA PARCELADA").
					WillReturnError(&pq.Error{Code: "23505", Message: `duplicate key value violates unique constraint "operation_types_description_key"`})
				db.Mock.ExpectRollback()
			},
			wantRes:        models.Response{Message: `duplicate key value violates unique constraint "operation_types_description_key"`},
			wantStatusCode: http.StatusConflict,
		},
	}

	for _, test := range tt {
		t.Run(test.purpose, func(_ *testing.T) {
			if test.mock != nil {
				test.mock()
			}

			jsn, _ := json.Marshal(test.req)
			req, _ := http.NewRequest(http.MethodPost, "/v1/operation_types", bytes.NewBuffer(jsn))
			rsc := httptest.NewRecorder()

			handlers.TestEngine.ServeHTTP(rsc, req)
			assert.Equal(test.wantStatusCode, rsc.Code)

			data, _ := ioutil.ReadAll(rsc.Result().Body)
			var body models.Response

			err := json.Unmarshal(data, &body)
			assert.Nil(err)

			if err == nil {
				assert.Equal(test.wantRes, body)
			}
		})
	}

}
