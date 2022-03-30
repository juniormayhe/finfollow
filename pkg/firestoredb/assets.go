package firestoredb

import (
	"context"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"juniormayhe.com/finfollow/pkg/models"
)

// This will insert a new asset into the database.
func (m *FirestoreModel) Insert(username string, name string, value float32, currency string, custody string, created time.Time, finished time.Time, active bool) (string, error) {

	docRef, _, err := m.Client.Collection("accounts").Doc(username).Collection("assets").Add(context.Background(), map[string]interface{}{
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
func (m *FirestoreModel) Get(username string, id string) (*models.Asset, error) {
	ds, err := m.Client.Collection("accounts").Doc(username).Collection("assets").Doc(id).Get(context.Background())

	// return our own models.ErrNoRecord error if no document was found
	if err != nil && strings.Contains(err.Error(), "NotFound") {
		return nil, models.ErrNoRecord
	}

	// Initialize a pointer to a new zeroed Snippet struct.
	asset := &models.Asset{}
	asset.Id = ds.Ref.ID
	ds.DataTo(&asset)

	log.Printf("asset: %+v", asset)

	return asset, err
}

// This will return the 10 most recently created assets.
func (m *FirestoreModel) Latest(username string) ([]*models.Asset, error) {
	// Initialize an empty slice to hold the models.Asset objects.
	assets := []*models.Asset{}
	iter := m.Client.Collection("accounts").Doc(username).Collection("assets").Query.OrderBy("created", firestore.Desc).Limit(10).Documents(context.Background())
	for {
		// Use iter.Next to iterate through the docs in the DocumentIterator.
		ds, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		asset := &models.Asset{}
		asset.Id = ds.Ref.ID
		ds.DataTo(&asset)
		assets = append(assets, asset)
	}

	return assets, nil
}

func Sum(assets []*models.Asset) float32 {
	sum := float32(0)
	for _, asset := range assets {
		sum += asset.Value
	}
	return sum
}
