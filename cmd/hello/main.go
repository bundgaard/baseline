package main

import "fmt"

var (
	// CommitID ...
	CommitID string
	// Branch ...
	Branch string
)

func main() {
	if len(CommitID) > 0 {
		fmt.Printf("CommitID %s\n", CommitID)
	}
	if len(Branch) > 0 {
		fmt.Printf("Branch %s\n", Branch)
	}
	fmt.Println("Hello, World!")
}
