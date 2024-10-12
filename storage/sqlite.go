package storage

import "golangSampleCRUD/types"

type SqliteStorage struct{}

func (s *SqliteStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "John Doe",
	}
}
