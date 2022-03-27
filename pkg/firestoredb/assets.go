package firestoredb

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"juniormayhe.com/finfollow/pkg/models"
)

// define a struct to wrap the firestore client
type AssetModel struct {
	Client *firestore.Client
}

// This will insert a new asset into the database.
func (m *AssetModel) Insert(name string, value float32, currency string, custody string, created time.Time, finished time.Time, active bool) (*firestore.DocumentRef, *firestore.WriteResult, error) {

	/*
		Name     string
		Value    float32
		Currency string
		Custody  string
		Created  time.Time
		Finished time.Time
		Active   bool
	*/
	docRef, result, err := m.Client.Collection("assets").Add(context.Background(), map[string]interface{}{
		"name":     name,
		"value":    value,
		"currency": currency,
		"custody":  custody,
		"created":  created,
		"finished": finished,
		"active":   active,
	})
	if err != nil {
		return nil, nil, err
	}

	return docRef, result, nil
}

// This will return a specific asset based on its id.
func (m *AssetModel) Get(id int) (*models.Asset, error) {
	return nil, nil
}

// This will return the 10 most recently created assets.
func (m *AssetModel) Latest() ([]*models.Asset, error) {
	return nil, nil
}
