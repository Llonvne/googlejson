package googlejson

import (
	"encoding/json"
)

// Request represents the top-level JSON structure for API requests
type Request struct {
	APIVersion string  `json:"apiVersion,omitempty"`
	Context    string  `json:"context,omitempty"`
	ID         string  `json:"id,omitempty"`
	Method     string  `json:"method,omitempty"`
	Params     *Params `json:"params,omitempty"`
}

// MarshalRequest converts a Request struct to JSON bytes
func MarshalRequest(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// UnmarshalRequest parses JSON bytes into a Request struct
func UnmarshalRequest(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// NewRequest creates a new Request instance
func NewRequest() *Request {
	return &Request{}
}

// WithAPIVersion sets the APIVersion field and returns the Request
func (r *Request) WithAPIVersion(apiVersion string) *Request {
	r.APIVersion = apiVersion
	return r
}

// WithContext sets the Context field and returns the Request
func (r *Request) WithContext(context string) *Request {
	r.Context = context
	return r
}

// WithID sets the ID field and returns the Request
func (r *Request) WithID(id string) *Request {
	r.ID = id
	return r
}

// WithMethod sets the Method field and returns the Request
func (r *Request) WithMethod(method string) *Request {
	r.Method = method
	return r
}

// WithParams sets the Params field and returns the Request
func (r *Request) WithParams(params *Params) *Request {
	r.Params = params
	return r
}

// ToResponse converts a Request to a Response
func (r *Request) ToResponse() *Response {
	return &Response{
		APIVersion: r.APIVersion,
		Context:    r.Context,
		ID:         r.ID,
		Method:     r.Method,
		Params:     r.Params,
	}
}
