package domain

import "encoding/json"

type Request struct {
	RequestBody *json.RawMessage
	RemoteAddr  string
}
