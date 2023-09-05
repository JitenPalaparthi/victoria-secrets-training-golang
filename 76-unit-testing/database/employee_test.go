package database_test

import (
	"demo/database"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestInsertEmployee(t *testing.T) {

	ctrl := gomock.NewController(t)

	mockemp := database.NewMockIEmployee(ctrl)

	emp := &database.Employee{}
	expectedEmp := &database.Employee{ID: "101", Name: "Jiten"}

	mockemp.EXPECT().InsertEmployee(emp).Return(expectedEmp, nil).MaxTimes(1)

}
