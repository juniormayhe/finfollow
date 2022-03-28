package firestoredb

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"juniormayhe.com/finfollow/pkg/models"
)

// define a struct to wrap the firestore client
type AssetModel struct {
	Client *firestore.Client
}

// This will insert a new asset into the database.
func (m *AssetModel) Insert(name string, value float32, currency string, custody string, created time.Time, finished time.Time, active bool) (string, error) {

	docRef, _, err := m.Client.Collection("assets").Add(context.Background(), map[string]interface{}{
		"name":     name,
		"value":    value,
		"currency": currency,
		"custody":  custody,
		"created":  created,
		"finished": finished,
		"active":   active,
	})
	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

// This will return a specific asset based on its id.
func (m *AssetModel) Get(id string) (*models.Asset, error) {
	ds, err := m.Client.Collection("assets").Doc(id).Get(context.Background())

	// Initialize a pointer to a new zeroed Snippet struct.
	asset := &models.Asset{}
	ds.DataTo(&asset)

	log.Printf("asset.Name: %s", asset.Name)

	return asset, err
}

// This will return the 10 most recently created assets.
func (m *AssetModel) Latest() ([]*models.Asset, error) {
	return nil, nil
}
