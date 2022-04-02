package firestoredb

import (
	"context"
	"log"
	"strings"
	"time"

	"juniormayhe.com/finfollow/pkg/models"
)

// This update user balance in database
func (m *FirestoreModel) UpdateBalance(username string, value float64) (bool, error) {

	balance, err := m.GetBalance(username)
	if err != nil && err != models.ErrNoRecord {
		return false, err
	}

	var previousBalanceValue float64
	if balance != nil {
		previousBalanceValue = balance.Value
	}

	_, dbErr := m.Client.Collection("balances").Doc(username).Set(context.Background(), map[string]interface{}{
		"value":   previousBalanceValue + value,
		"updated": time.Now(),
	})

	if dbErr != nil {
		return false, err
	}

	return true, nil
}

// This will return a specific balance based on its username.
func (m *FirestoreModel) GetBalance(username string) (*models.Balance, error) {
	ds, err := m.Client.Collection("balances").Doc(username).Get(context.Background())

	// return our own models.ErrNoRecord error if no document was found
	if err != nil && strings.Contains(err.Error(), "NotFound") {
		return nil, models.ErrNoRecord
	}

	// Initialize a pointer to a new zeroed Snippet struct.
	balance := &models.Balance{}
	ds.DataTo(&balance)

	log.Printf("balance.Value: %f", balance.Value)

	return balance, err
}
