package main

import "juniormayhe.com/finfollow/pkg/models"

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
	Asset  *models.Asset
	Assets []*models.Asset
}
