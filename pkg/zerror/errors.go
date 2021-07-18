package zerror

// zerror package contains all the internal errors that can be used to identify common errors

import "errors"

var ErrQuit = errors.New("user is quitting the application")
var ErrNotFound = errors.New("no results found")
