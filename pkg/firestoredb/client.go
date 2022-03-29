package firestoredb

import (
	"cloud.google.com/go/firestore"
)

// define a struct to wrap the firestore client
type FirestoreModel struct {
	Client *firestore.Client
}
