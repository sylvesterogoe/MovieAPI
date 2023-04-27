package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	runtimeJson := fmt.Sprintf("%d mins", r)

	quotedRuntimeJson := strconv.Quote(runtimeJson)

	return []byte(quotedRuntimeJson), nil
}