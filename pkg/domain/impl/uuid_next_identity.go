package impl

import "github.com/google/uuid"

func UUIDNextIdentity() string {
	return uuid.NewString()
}
