package firestoredb

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
	"juniormayhe.com/finfollow/pkg/models"
)

// We'll use the Insert method to add a new record to the users table.
func (m *FirestoreModel) InsertUser(name, email, password string) (string, error) {
	// Create a bcrypt hash of the plain-text password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	// check if user is duplicate
	ds, dbErr := m.Client.Collection("users").Where("email", "==", email).Limit(1).Documents(context.Background()).Next()
	if dbErr != nil && dbErr != iterator.Done {
		return "", dbErr
	}

	if ds != nil && ds.Ref.ID != "" {
		return "", models.ErrDuplicateEmail
	}

	docRef, _, err := m.Client.Collection("users").Add(context.Background(), map[string]interface{}{
		"name":     name,
		"email":    email,
		"password": hashedPassword,
	})

	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *FirestoreModel) AuthenticateUser(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Get method to fetch details for a specific user based
// on their user ID.
func (m *FirestoreModel) GetUser(id int) (*models.User, error) {
	return nil, nil
}
