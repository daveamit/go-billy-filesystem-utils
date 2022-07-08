package utils

import "errors"

// ErrResourceIsNotAFile is returned if resource is not a file, but a dir or something.
var ErrResourceIsNotAFile = errors.New("the resource specified is not a file")
