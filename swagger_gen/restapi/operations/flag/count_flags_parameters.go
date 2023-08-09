// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCountFlagsParams creates a new CountFlagsParams object
//
// There are no default values defined in the spec.
func NewCountFlagsParams() CountFlagsParams {

	return CountFlagsParams{}
}

// CountFlagsParams contains all the bound params for the count flags operation
// typically these are obtained from a http.Request
//
// swagger:parameters countFlags
type CountFlagsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*count only deleted flags
	  In: query
	*/
	Deleted *bool
	/*count only flags exactly matching given description
	  In: query
	*/
	Description *string
	/*count only flags partially matching given description
	  In: query
	*/
	DescriptionLike *string
	/*count only flags having given enabled status
	  In: query
	*/
	Enabled *bool
	/*include all deleted flags in the count
	  In: query
	*/
	IncludeDeleted *bool
	/*count only flags matching given key
	  In: query
	*/
	Key *string
	/*count only flags with the given tags (comma separated)
	  In: query
	*/
	Tags *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCountFlagsParams() beforehand.
func (o *CountFlagsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDeleted, qhkDeleted, _ := qs.GetOK("deleted")
	if err := o.bindDeleted(qDeleted, qhkDeleted, route.Formats); err != nil {
		res = append(res, err)
	}

	qDescription, qhkDescription, _ := qs.GetOK("description")
	if err := o.bindDescription(qDescription, qhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	qDescriptionLike, qhkDescriptionLike, _ := qs.GetOK("description_like")
	if err := o.bindDescriptionLike(qDescriptionLike, qhkDescriptionLike, route.Formats); err != nil {
		res = append(res, err)
	}

	qEnabled, qhkEnabled, _ := qs.GetOK("enabled")
	if err := o.bindEnabled(qEnabled, qhkEnabled, route.Formats); err != nil {
		res = append(res, err)
	}

	qIncludeDeleted, qhkIncludeDeleted, _ := qs.GetOK("include_deleted")
	if err := o.bindIncludeDeleted(qIncludeDeleted, qhkIncludeDeleted, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDeleted binds and validates parameter Deleted from query.
func (o *CountFlagsParams) bindDeleted(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("deleted", "query", "bool", raw)
	}
	o.Deleted = &value

	return nil
}

// bindDescription binds and validates parameter Description from query.
func (o *CountFlagsParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Description = &raw

	return nil
}

// bindDescriptionLike binds and validates parameter DescriptionLike from query.
func (o *CountFlagsParams) bindDescriptionLike(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.DescriptionLike = &raw

	return nil
}

// bindEnabled binds and validates parameter Enabled from query.
func (o *CountFlagsParams) bindEnabled(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("enabled", "query", "bool", raw)
	}
	o.Enabled = &value

	return nil
}

// bindIncludeDeleted binds and validates parameter IncludeDeleted from query.
func (o *CountFlagsParams) bindIncludeDeleted(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("include_deleted", "query", "bool", raw)
	}
	o.IncludeDeleted = &value

	return nil
}

// bindKey binds and validates parameter Key from query.
func (o *CountFlagsParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Key = &raw

	return nil
}

// bindTags binds and validates parameter Tags from query.
func (o *CountFlagsParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Tags = &raw

	return nil
}
