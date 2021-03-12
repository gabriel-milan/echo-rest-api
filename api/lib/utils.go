package lib

import (
	"strconv"

	"gorm.io/gorm"
)

// Struct for sending custom messages
type Message struct {
	Message string `json:"message"`
}

// Parsing QueryParams
func QueryParamToInt(value string, defaultValue int) (int, error) {
	if value == "" {
		return defaultValue, nil
	}
	val, err := strconv.ParseInt(value, 10, 16)
	return int(val), err
}

// Database pagination
func Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
