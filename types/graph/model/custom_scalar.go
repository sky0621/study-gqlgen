package model

import (
	"errors"
	"fmt"
	"io"
)

type CustomScalar struct {
	str string
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (s *CustomScalar) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("no target")
	}
	s.str = str
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (s *CustomScalar) MarshalGQL(w io.Writer) {
	if _, err := w.Write([]byte(s.str)); err != nil {
		fmt.Println(err)
	}
}
