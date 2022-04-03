package firestoredb

import (
	"context"
	"log"
	"time"

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
	log.Printf("hashedPassword: %+v", hashedPassword)

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
		"created":  time.Now(),
	})

	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *FirestoreModel) AuthenticateUser(email, password string) (string, error) {
	// Retrieve the id and hashed password associated with the given email. If
	// matching email exists, we return the ErrInvalidCredentials error.
	ds, dbErr := m.Client.Collection("users").Where("email", "==", email).Limit(1).Documents(context.Background()).Next()
	if dbErr == iterator.Done {
		log.Println("******** iterator error")
		return "", models.ErrInvalidCredentials
	} else if dbErr != nil {
		log.Println("******** general error")
		return "", dbErr
	}

	log.Printf("\n\r>>>> Document data: %#v\n", ds.Data())

	user := &models.User{}
	user.Id = ds.Ref.ID
	user.HashedPassword = ToByteSlice(ds.Data()["password"].([]uint8))
	ds.DataTo(&user)
	log.Printf("\n>>>> hashed: %+v\n", user.HashedPassword)
	log.Printf("\n>>>> pass: %v\n", []byte(password))

	// Check whether the hashed password and plain-text password provided match
	// If they don't, we return the ErrInvalidCredentials error.
	bCryptErr := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if bCryptErr == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("******** bcrypt error")
		return "", models.ErrInvalidCredentials
	} else if bCryptErr != nil {
		log.Println("******** other bcrypt error")
		return "", bCryptErr
	}

	// Otherwise, the password is correct. Return the user ID.
	return user.Id, nil
}

func ToByteSlice(b []byte) []byte {
	return b
}

// We'll use the Get method to fetch details for a specific user based
// on their user ID.
func (m *FirestoreModel) GetUser(id int) (*models.User, error) {
	return nil, nil
}
