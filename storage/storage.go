package storage

import "golangSampleCRUD/types"

type Storage interface {
	Get(int) *types.User
}
