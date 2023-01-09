package controllers

import (
	"net/http"

	"github.com/leonardogazio/golang-account-crud/db"
	"github.com/leonardogazio/golang-account-crud/db/repositories/operationtype"
	"github.com/leonardogazio/golang-account-crud/models"
)

const (
	PersistKindInsert = iota
	PersistKindUpdate
	PersistKindDelete
)

func persistOperationType(kind int, m *models.OperationType, strErrMsg, strSuccessMsg string) *models.Response {
	conn, err := db.GetConnection()
	if err != nil {
		return models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	rep, err := operationtype.New(conn)
	if err != nil {
		return models.NewResponse(http.StatusInternalServerError, err.Error(), nil)

	}

	httpResponseCode := http.StatusOK

	switch kind {
	case PersistKindInsert:
		err = rep.Insert(m)
		httpResponseCode = http.StatusCreated
	case PersistKindUpdate:
		err = rep.Update(m)
	default:
		err = rep.Delete(m)
	}

	if err != nil {
		return models.ParseSQLError(err, strErrMsg)
	}

	return models.NewResponse(httpResponseCode, strSuccessMsg, nil)
}

// CreateOperationType ...
func CreateOperationType(m *models.OperationType) *models.Response {
	return persistOperationType(PersistKindInsert, m, strInsertFail, strCreated)
}

// UpdateOperationType ...
func UpdateOperationType(m *models.OperationType) *models.Response {
	return persistOperationType(PersistKindUpdate, m, strUpdateFail, strUpdated)
}

// DeleteOperationType ...
func DeleteOperationType(m *models.OperationType) *models.Response {
	return persistOperationType(PersistKindDelete, m, strDeleteFail, strDeleted)
}

// GetOperationTypes ...
func GetOperationTypes(id int) *models.Response {
	conn, err := db.GetConnection()
	if err != nil {
		return models.NewResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	rep, err := operationtype.New(conn)
	if err != nil {
		return models.NewResponse(http.StatusInternalServerError, err.Error(), nil)

	}

	var res interface{}

	if id > 0 {
		res, err = rep.GetByID(id)
	} else {
		res, err = rep.Get(id)
	}

	if err != nil {
		return models.ParseSQLError(err, strGetFail)
	}

	return models.NewResponse(http.StatusOK, "", res)
}
