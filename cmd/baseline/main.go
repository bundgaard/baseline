package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

// Option ...
type Option struct {
	Name string
	Set  flag.FlagSet
}

// DocumentNode ...
type DocumentNode struct {
	Val  string
	Root *DocumentNode
}

// Document ...
type Document struct {
	Name string
	buf  bytes.Buffer
	Root *DocumentNode
}

// NewDocument ...
func NewDocument(name string) *Document {

	return &Document{Name: name, Root: &DocumentNode{}}
}

// DocumentHeading ...
type DocumentHeading int

const (
	h1 DocumentHeading = iota
	h2
	h3
	h4
	h5
	h6
)

// Heading ...
func (d *Document) Heading(level DocumentHeading, title string) *Document {
	return d
}

var (

	// CommitID ...
	CommitID string
	// Branch ...
	Branch string

	// Flags
	initFlags  = flag.NewFlagSet("init", flag.ExitOnError)
	initCreate = initFlags.Bool("create", false, "Create skeleton, default \"false\"")
)

const readme = `
# Badges
<img src="{{GITLAB_URL}}/{{PROJECT_PATH}}/badges/master/pipeline.svg">
<img src="{{GITLAB_URL}}/{{PROJECT_PATH}}/badges/master/coverage.svg">

# What is this ?

# How to run ?

# How to build ?



`

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initFlags.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if initFlags.Parsed() {
		fmt.Println("parsed init")

	}

}
