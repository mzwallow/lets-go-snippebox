package ui

import "embed"

//go:embed "html" "static"
var Files embed.FS

// To embed multiple paths,
// "static/css" "static/img" "static/js"
// Only use forward slash (/) as a path sperator.
//
// Can also embed specific files.
// Can use wildcard (*).
//
// Finally, if a path is to a directory, then all files in that directory
// are recursively embedded — except for files with names that begin
// with . or _ characters. If you want to include those files too, then you
// should use the all: prefix at the start of the path.
