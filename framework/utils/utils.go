package utils

import "github.com/google/uuid"

func GetUuid() string{
	uuidObj := uuid.New()
	return uuidObj.String()
}
