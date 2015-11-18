package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/fixtures/bugs/84/models"
)

// PutEventByIDParams contains all the bound params for the put event by id operation
// typically these are obtained from a http.Request
//
// swagger:parameters putEventById
type PutEventByIDParams struct {
	// Existing event
	// Required: true
	// In: body
	Event *models.Event
	// Existing event id.
	// Required: true
	// In: path
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutEventByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := route.Consumer.Consume(r.Body, o.Event); err != nil {
		res = append(res, errors.NewParseError("event", "body", "", err))
	} else {
		if err := o.Event.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

	}

	if err := o.bindID(route.Params.Get("id"), route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutEventByIDParams) bindID(raw string, formats strfmt.Registry) error {

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}