package operationtype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

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
			data, rsc := handlers.MockedReq(http.MethodPost, "/v1/operation_types", bytes.NewBuffer(jsn))
			assert.Equal(test.wantStatusCode, rsc.Code)

			body := models.Response{}
			err := json.Unmarshal(data, &body)
			assert.Nil(err)

			if err == nil {
				assert.Equal(test.wantRes, body)
			}
		})
	}

}

func TestRead(t *testing.T) {
	assert := assert.New(t)

	pur := "Routes: HTTP 1.1 GET /operation_types"
	qry := "SELECT (.+) FROM operation_types"
	columns := []string{"id", "description", "created_at", "updated_at"}

	tt := []struct {
		purpose        string
		req            string
		mock           func(string)
		wantRes        models.Response
		wantStatusCode int
	}{
		{
			purpose: fmt.Sprintf("%s - Success Case /Should return list.", pur),
			mock: func(req string) {
				q := db.Mock.ExpectQuery(qry)
				if req != "" {
					q.WithArgs(req)
				}
				q.WillReturnRows(
					sqlxmock.NewRows(columns).
						AddRow(1, "COMPRA A VISTA", time.Unix(1673281418, 0), time.Unix(1673281418, 0)).
						AddRow(2, "COMPRA PARCELADA", time.Unix(1673281461, 0), time.Unix(1673281461, 0)).
						AddRow(3, "SAQUE", time.Unix(1673281499, 0), time.Unix(1673281499, 0)).
						AddRow(4, "PAGAMENTO", time.Unix(1673281543, 0), time.Unix(1673281543, 0)),
				)
			},
			wantRes: models.Response{
				Data: []interface{}{
					map[string]interface{}{
						"id":          1.,
						"description": "COMPRA A VISTA",
						"created_at":  "2023-01-09T13:23:38-03:00",
						"updated_at":  "2023-01-09T13:23:38-03:00",
					},
					map[string]interface{}{
						"id":          2.,
						"description": "COMPRA PARCELADA",
						"created_at":  "2023-01-09T13:24:21-03:00",
						"updated_at":  "2023-01-09T13:24:21-03:00",
					},
					map[string]interface{}{
						"id":          3.,
						"description": "SAQUE",
						"created_at":  "2023-01-09T13:24:59-03:00",
						"updated_at":  "2023-01-09T13:24:59-03:00",
					},
					map[string]interface{}{
						"id":          4.,
						"description": "PAGAMENTO",
						"created_at":  "2023-01-09T13:25:43-03:00",
						"updated_at":  "2023-01-09T13:25:43-03:00",
					},
				},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			purpose: fmt.Sprintf("%s - Success Case /Should return COMPRA A VISTA by its ID 1.", pur),
			req:     "1",
			mock: func(r string) {
				q := db.Mock.ExpectQuery(qry)
				if r != "" {
					r, _ := strconv.Atoi(r)
					q.WithArgs(r)
				}
				q.WillReturnRows(
					sqlxmock.NewRows(columns).
						AddRow(1, "COMPRA A VISTA", time.Unix(1673281418, 0), time.Unix(1673281418, 0)).
						AddRow(2, "COMPRA PARCELADA", time.Unix(1673281461, 0), time.Unix(1673281461, 0)).
						AddRow(3, "SAQUE", time.Unix(1673281499, 0), time.Unix(1673281499, 0)).
						AddRow(4, "PAGAMENTO", time.Unix(1673281543, 0), time.Unix(1673281543, 0)),
				)
			},
			wantRes: models.Response{
				Data: map[string]interface{}{
					"id":          1.,
					"description": "COMPRA A VISTA",
					"created_at":  "2023-01-09T13:23:38-03:00",
					"updated_at":  "2023-01-09T13:23:38-03:00",
				},
			},
			wantStatusCode: http.StatusOK,
		},
	}

	for _, test := range tt {
		t.Run(test.purpose, func(_ *testing.T) {
			if test.mock != nil {
				test.mock(test.req)
			}

			uri := "/v1/operation_types"

			if test.req != "" {
				uri = fmt.Sprintf("%s/%s", uri, test.req)
			}

			data, rsc := handlers.MockedReq(http.MethodGet, uri, nil)
			assert.Equal(test.wantStatusCode, rsc.Code)

			var body models.Response
			err := json.Unmarshal(data, &body)
			assert.Nil(err)

			if err == nil {
				assert.Equal(test.wantRes, body)
			}
		})
	}

}
