package model

import (
	"fmt"
	"io"
)

type YesNo bool

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (y *YesNo) UnmarshalGQL(v interface{}) error {
	yes, ok := v.(string)
	if !ok {
		return fmt.Errorf("points must be strings")
	}

	if yes == "yes" {
		*y = true
	} else {
		*y = false
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (y YesNo) MarshalGQL(w io.Writer) {
	if y {
		if _, err := w.Write([]byte(`"yes"`)); err != nil {
			fmt.Println(err)
		}
	} else {
		if _, err := w.Write([]byte(`"no"`)); err != nil {
			fmt.Println(err)
		}
	}
}
