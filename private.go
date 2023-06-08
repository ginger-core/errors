package errors

import (
	"context"
	"fmt"
)

type private struct {
	error
	ctx         context.Context
	Id          string
	Type        Type
	Description string
	Trace       string
	// MessageOne is the content of the message for the CLDR plural form "one".
	messageOne string
	// values set of message values to set into message while being rendered
	values     map[string]interface{}
	source     string
	detail     Detail
	properties map[string]interface{}
	// pluralCount determines which plural form of the message is used.
	pluralCount int
	// isIdCustomized determines if id has been set by command or still using default message
	isIdCustomized bool
	// isMessageCustomized determines if message has been set by command or still using default message
	isMessageCustomized bool
}

func (p *private) clone() private {
	return private{
		error:               p.error,
		ctx:                 p.ctx,
		Id:                  p.Id,
		Type:                p.Type,
		Description:         p.Description,
		Trace:               p.Trace,
		messageOne:          p.messageOne,
		values:              p.values,
		source:              p.source,
		detail:              p.detail,
		properties:          p.properties,
		pluralCount:         p.pluralCount,
		isIdCustomized:      p.isIdCustomized,
		isMessageCustomized: p.isMessageCustomized,
	}
}

func (p *private) String() string {
	var r string
	if p.ctx != nil {
		if id := p.ctx.Value("reqId"); id != nil {
			r += "reqId: " + id.(string) + ", "
		}
	}
	if p.source != "" {
		r += "source: " + p.source + ", "
	}
	if p.Trace != "" {
		r += "trace: " + p.Trace + ", "
	}
	if p.Id != "" {
		r += "id: " + p.Id + ", "
	}
	if p.error != nil {
		r += "error: " + p.error.Error() + ", "
	}
	if p.Description != "" {
		r += "Description: " + p.Description + ", "
	}
	if p.detail != nil {
		r += "detail:\n"
		for k, v := range p.detail {
			r += " " + k + ": " + fmt.Sprint(v) + "\n"
		}
	}
	return r
}
