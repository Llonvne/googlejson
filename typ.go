package googlejson

import (
	"encoding/json"
	"time"
)

// Params represents the params object in the schema
type Params struct {
	ID string `json:"id,omitempty"`
}

// Data represents the data object in the schema
type Data struct {
	Kind             string          `json:"kind,omitempty"`
	Fields           string          `json:"fields,omitempty"`
	Etag             string          `json:"etag,omitempty"`
	ID               string          `json:"id,omitempty"`
	Lang             string          `json:"lang,omitempty"`
	Updated          *time.Time      `json:"updated,omitempty"` // RFC 3339 format
	Deleted          *bool           `json:"deleted,omitempty"`
	CurrentItemCount *int            `json:"currentItemCount,omitempty"`
	ItemsPerPage     *int            `json:"itemsPerPage,omitempty"`
	StartIndex       *int            `json:"startIndex,omitempty"`
	TotalItems       *int            `json:"totalItems,omitempty"`
	PageIndex        *int            `json:"pageIndex,omitempty"`
	TotalPages       *int            `json:"totalPages,omitempty"`
	PageLinkTemplate string          `json:"pageLinkTemplate,omitempty"`
	Next             json.RawMessage `json:"next,omitempty"`
	NextLink         string          `json:"nextLink,omitempty"`
	Previous         json.RawMessage `json:"previous,omitempty"`
	PreviousLink     string          `json:"previousLink,omitempty"`
	Self             json.RawMessage `json:"self,omitempty"`
	SelfLink         string          `json:"selfLink,omitempty"`
	Edit             json.RawMessage `json:"edit,omitempty"`
	EditLink         string          `json:"editLink,omitempty"`
	Items            []interface{}   `json:"items,omitempty"`
}

// ErrorInfo represents the error object in the schema
type ErrorInfo struct {
	Code    *int          `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
}

// ErrorDetail represents an individual error detail in the errors array
type ErrorDetail struct {
	Domain       string `json:"domain,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Message      string `json:"message,omitempty"`
	Location     string `json:"location,omitempty"`
	LocationType string `json:"locationType,omitempty"`
	ExtendedHelp string `json:"extendedHelp,omitempty"`
	SendReport   string `json:"sendReport,omitempty"`
}

// NewParams creates a new Params instance
func NewParams() *Params {
	return &Params{}
}

// WithID sets the ID field and returns the Params
func (p *Params) WithID(id string) *Params {
	p.ID = id
	return p
}

// NewData creates a new Data instance
func NewData() *Data {
	return &Data{}
}

// NewErrorInfo creates a new ErrorInfo instance
func NewErrorInfo() *ErrorInfo {
	return &ErrorInfo{}
}

// NewErrorDetail creates a new ErrorDetail instance
func NewErrorDetail() *ErrorDetail {
	return &ErrorDetail{}
}

// WithKind sets the Kind field and returns the Data
func (d *Data) WithKind(kind string) *Data {
	d.Kind = kind
	return d
}

// WithFields sets the Fields field and returns the Data
func (d *Data) WithFields(fields string) *Data {
	d.Fields = fields
	return d
}

// WithEtag sets the Etag field and returns the Data
func (d *Data) WithEtag(etag string) *Data {
	d.Etag = etag
	return d
}

// WithID sets the ID field and returns the Data
func (d *Data) WithID(id string) *Data {
	d.ID = id
	return d
}

// WithLang sets the Lang field and returns the Data
func (d *Data) WithLang(lang string) *Data {
	d.Lang = lang
	return d
}

// WithUpdated sets the Updated field and returns the Data
func (d *Data) WithUpdated(updated *time.Time) *Data {
	d.Updated = updated
	return d
}

// WithDeleted sets the Deleted field and returns the Data
func (d *Data) WithDeleted(deleted *bool) *Data {
	d.Deleted = deleted
	return d
}

// WithCurrentItemCount sets the CurrentItemCount field and returns the Data
func (d *Data) WithCurrentItemCount(count *int) *Data {
	d.CurrentItemCount = count
	return d
}

// WithItemsPerPage sets the ItemsPerPage field and returns the Data
func (d *Data) WithItemsPerPage(itemsPerPage *int) *Data {
	d.ItemsPerPage = itemsPerPage
	return d
}

// WithStartIndex sets the StartIndex field and returns the Data
func (d *Data) WithStartIndex(startIndex *int) *Data {
	d.StartIndex = startIndex
	return d
}

// WithTotalItems sets the TotalItems field and returns the Data
func (d *Data) WithTotalItems(totalItems *int) *Data {
	d.TotalItems = totalItems
	return d
}

// WithPageIndex sets the PageIndex field and returns the Data
func (d *Data) WithPageIndex(pageIndex *int) *Data {
	d.PageIndex = pageIndex
	return d
}

// WithTotalPages sets the TotalPages field and returns the Data
func (d *Data) WithTotalPages(totalPages *int) *Data {
	d.TotalPages = totalPages
	return d
}

// WithPageLinkTemplate sets the PageLinkTemplate field and returns the Data
func (d *Data) WithPageLinkTemplate(template string) *Data {
	d.PageLinkTemplate = template
	return d
}

// WithNext sets the Next field and returns the Data
func (d *Data) WithNext(next json.RawMessage) *Data {
	d.Next = next
	return d
}

// WithNextLink sets the NextLink field and returns the Data
func (d *Data) WithNextLink(nextLink string) *Data {
	d.NextLink = nextLink
	return d
}

// WithPrevious sets the Previous field and returns the Data
func (d *Data) WithPrevious(previous json.RawMessage) *Data {
	d.Previous = previous
	return d
}

// WithPreviousLink sets the PreviousLink field and returns the Data
func (d *Data) WithPreviousLink(previousLink string) *Data {
	d.PreviousLink = previousLink
	return d
}

// WithSelf sets the Self field and returns the Data
func (d *Data) WithSelf(self json.RawMessage) *Data {
	d.Self = self
	return d
}

// WithSelfLink sets the SelfLink field and returns the Data
func (d *Data) WithSelfLink(selfLink string) *Data {
	d.SelfLink = selfLink
	return d
}

// WithEdit sets the Edit field and returns the Data
func (d *Data) WithEdit(edit json.RawMessage) *Data {
	d.Edit = edit
	return d
}

// WithEditLink sets the EditLink field and returns the Data
func (d *Data) WithEditLink(editLink string) *Data {
	d.EditLink = editLink
	return d
}

// WithItems sets the Items field and returns the Data
func (d *Data) WithItems(items []interface{}) *Data {
	d.Items = items
	return d
}

// WithCode sets the Code field and returns the ErrorInfo
func (e *ErrorInfo) WithCode(code *int) *ErrorInfo {
	e.Code = code
	return e
}

// WithMessage sets the Message field and returns the ErrorInfo
func (e *ErrorInfo) WithMessage(message string) *ErrorInfo {
	e.Message = message
	return e
}

// WithErrors sets the Errors field and returns the ErrorInfo
func (e *ErrorInfo) WithErrors(errors []ErrorDetail) *ErrorInfo {
	e.Errors = errors
	return e
}

// WithDomain sets the Domain field and returns the ErrorDetail
func (e *ErrorDetail) WithDomain(domain string) *ErrorDetail {
	e.Domain = domain
	return e
}

// WithReason sets the Reason field and returns the ErrorDetail
func (e *ErrorDetail) WithReason(reason string) *ErrorDetail {
	e.Reason = reason
	return e
}

// WithMessage sets the Message field and returns the ErrorDetail
func (e *ErrorDetail) WithMessage(message string) *ErrorDetail {
	e.Message = message
	return e
}

// WithLocation sets the Location field and returns the ErrorDetail
func (e *ErrorDetail) WithLocation(location string) *ErrorDetail {
	e.Location = location
	return e
}

// WithLocationType sets the LocationType field and returns the ErrorDetail
func (e *ErrorDetail) WithLocationType(locationType string) *ErrorDetail {
	e.LocationType = locationType
	return e
}

// WithExtendedHelp sets the ExtendedHelp field and returns the ErrorDetail
func (e *ErrorDetail) WithExtendedHelp(extendedHelp string) *ErrorDetail {
	e.ExtendedHelp = extendedHelp
	return e
}

// WithSendReport sets the SendReport field and returns the ErrorDetail
func (e *ErrorDetail) WithSendReport(sendReport string) *ErrorDetail {
	e.SendReport = sendReport
	return e
}
