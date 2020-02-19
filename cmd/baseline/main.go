package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

// Based on this https://towardsdatascience.com/linked-list-implementation-guide-16ed67be18e4

// Option ...
type Option struct {
	Name string
	Set  flag.FlagSet
}

// DocumentNode ...
type DocumentNode struct {
	Val  string
	Type string
	Next *DocumentNode
}

func (dn *DocumentNode) String() string {
	return fmt.Sprintf("Node(%s,%s)", dn.Type, dn.Val)
}

// Document ...
type Document struct {
	Name string
	buf  bytes.Buffer
	size int
	Root *DocumentNode
}

// IsEmpty ...
func (d *Document) IsEmpty() bool {
	return d.Root == nil
}

// InsertFront ...
func (d *Document) InsertFront(typ, val string) {
	if d.IsEmpty() {
		d.size++
		node := new(DocumentNode)
		node.Type = typ
		node.Val = val
		node.Next = nil
		d.Root = node
	} else {
		node := new(DocumentNode)
		node.Type = typ
		node.Val = val
		node.Next = d.Root
		d.Root = node
	}
}

// NewDocument ...
func NewDocument(name string) *Document {
	return &Document{Name: name, Root: nil}
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
	d.InsertFront("#", title)

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

	d := Document{Name: "Example", Root: nil}
	d.Heading(h1, "Hello World").Heading(h2, "Hej")
	fmt.Println(d.Root)
	fmt.Println(d.Root.Next)
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
