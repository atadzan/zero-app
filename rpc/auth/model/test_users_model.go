package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TestUsersModel = (*customTestUsersModel)(nil)

type (
	// TestUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTestUsersModel.
	TestUsersModel interface {
		testUsersModel
	}

	customTestUsersModel struct {
		*defaultTestUsersModel
	}
)

// NewTestUsersModel returns a model for the database table.
func NewTestUsersModel(conn sqlx.SqlConn) TestUsersModel {
	return &customTestUsersModel{
		defaultTestUsersModel: newTestUsersModel(conn),
	}
}
