package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Asset struct {
	Name     string
	Value    float32
	Currency string
	Custody  string
	Created  time.Time
	Finished time.Time
	Active   bool
}
