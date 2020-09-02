package utils

import "github.com/google/uuid"

// 唯一Id
func GetUuid() string{
	uuidObj := uuid.New()
	return uuidObj.String()
}
