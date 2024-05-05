package utils

import "github.com/gosimple/slug"

func Slugify(name string) string {
	return slug.Make(name)
}
