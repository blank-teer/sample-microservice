package helpers

import (
	"encoding/json"
	"fmt"
	"io"

	"order-manager/pkg/api/v3/errs"
)

func ParsePayload(r io.Reader, into interface{}) error {
	if err := json.NewDecoder(r).Decode(&into); err != nil {
		return fmt.Errorf("%w: payload: cannot decode: %s", errs.ErrNotValid, err)
	}
	return nil
}
