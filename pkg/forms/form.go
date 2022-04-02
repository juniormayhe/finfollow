package forms

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Create a custom Form struct, which anonymously embeds a url.Values object
// (to hold the form data) and an Errors field to hold any validation errors
// for the form data.
type Form struct {
	url.Values
	Errors errors
}

// Define a New function to initialize a custom Form struct. Notice that
// this takes the form data as the parameter?
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Implement a Required method to check that specific fields in the form
// data are present and not blank. If any fields fail this check, add the
// appropriate message to the form errors.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) ValidNumber(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			f.Errors.Add(field, "This number is invalid")
			return
		}

		if number < 0 {
			f.Errors.Add(field, "This number must be greater or equal to 0")
			return
		}
	}
}

func (f *Form) GetDate(field string) time.Time {
	created, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(f.Get(field)), time.UTC)
	if err != nil {
		f.Errors.Add(field, "This date is invalid")
	}
	return created
}

func (f *Form) GetNumber(field string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(f.Get(field)), 64)
	if err != nil {
		f.Errors.Add(field, "This value is invalid")
	}
	return value
}

// Implement a MaxLength method to check that a specific field in the form
// contains a maximum number of characters. If the check fails then add the
// appropriate message to the form errors.
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

// Implement a PermittedValues method to check that a specific field in the form
// matches one of a set of specific permitted values. If the check fails
// then add the appropriate message to the form errors.
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Implement a Valid method which returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
