package utils

import (
	"database/sql"
	"github.com/go-playground/form"
)

var decoder *form.Decoder

func NewDecoder() *form.Decoder {
	decoder = form.NewDecoder()
	decoder.RegisterCustomTypeFunc(DecodeSQLNullString, sql.NullString{})
	return decoder
}

// DecodeSQLNullString is a custom type decoder for sql.NullString values
func DecodeSQLNullString(value []string) (interface{}, error) {
	if len(value) == 0 || value[0] == "" {
		return sql.NullString{String: "", Valid: false}, nil
	}
	return sql.NullString{String: value[0], Valid: true}, nil
}
