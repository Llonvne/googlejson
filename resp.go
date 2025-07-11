package googlejson

import (
	"encoding/json"
)

// Response represents the top-level JSON structure as defined in the schema
type Response struct {
	APIVersion string     `json:"apiVersion,omitempty"`
	Context    string     `json:"context,omitempty"`
	ID         string     `json:"id,omitempty"`
	Method     string     `json:"method,omitempty"`
	Params     *Params    `json:"params,omitempty"`
	Data       *Data      `json:"data,omitempty"`
	Error      *ErrorInfo `json:"error,omitempty"`
}

// MarshalResponse converts a Response struct to JSON bytes
func MarshalResponse(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// UnmarshalResponse parses JSON bytes into a Response struct
func UnmarshalResponse(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// NewResponse creates a new Response instance
func NewResponse() *Response {
	return &Response{}
}

// WithAPIVersion sets the APIVersion field and returns the Response
func (r *Response) WithAPIVersion(apiVersion string) *Response {
	r.APIVersion = apiVersion
	return r
}

// WithContext sets the Context field and returns the Response
func (r *Response) WithContext(context string) *Response {
	r.Context = context
	return r
}

// WithID sets the ID field and returns the Response
func (r *Response) WithID(id string) *Response {
	r.ID = id
	return r
}

// WithMethod sets the Method field and returns the Response
func (r *Response) WithMethod(method string) *Response {
	r.Method = method
	return r
}

// WithParams sets the Params field and returns the Response
func (r *Response) WithParams(params *Params) *Response {
	r.Params = params
	return r
}

// WithData sets the Data field and returns the Response
func (r *Response) WithData(data *Data) *Response {
	r.Data = data
	return r
}

// WithError sets the Error field and returns the Response
func (r *Response) WithError(err *ErrorInfo) *Response {
	r.Error = err
	return r
}

// ToRequest converts a Response to a Request
func (r *Response) ToRequest() *Request {
	return &Request{
		APIVersion: r.APIVersion,
		Context:    r.Context,
		ID:         r.ID,
		Method:     r.Method,
		Params:     r.Params,
	}
}
