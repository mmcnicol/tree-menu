package main

import (
	"time"
)

// Document represents a clinical document in the tree menu.
type Document struct {
    Date            time.Time
    Specialty       string
    Type            string
    DocumentSource  string
}

// Implement the sort.Interface for sorting by date.
type byDate []Document

func (bd byDate) Len() int           { return len(bd) }
func (bd byDate) Swap(i, j int)      { bd[i], bd[j] = bd[j], bd[i] }
func (bd byDate) Less(i, j int) bool { return bd[i].Date.After(bd[j].Date) }

/*
// Documents is a slice of Document.
type Documents []Document

// Len returns the length of the slice.
func (docs Documents) Len() int {
	return len(docs)
}

// Swap swaps the elements with indexes i and j.
func (docs Documents) Swap(i, j int) {
	docs[i], docs[j] = docs[j], docs[i]
}

// Less returns whether the document at index i should sort before the document at index j.
func (docs Documents) Less(i, j int) bool {
	return docs[i].Date.After(docs[j].Date)
}
*/
