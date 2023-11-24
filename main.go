package main

import (
	"fmt"
	"sort"
)

// SortMode represents the sorting mode for documents.
type DisplayMode string

// Define constants for each sort mode.
const (
	ByDate      DisplayMode = "date"
	BySpecialty DisplayMode = "specialty"
	ByType      DisplayMode = "type"
)

type TreeNodeBuilder struct {
    Documents   []Document
    TreeNode    *TreeNode
    DisplayMode DisplayMode
}

func NewTreeNodeBuilder(documents []Document, displayMode DisplayMode) *TreeNodeBuilder {
    return &TreeNodeBuilder{
        Documents: documents,
        TreeNode: NewTreeNode("root"),
        DisplayMode: displayMode,
    }
}

func (n *TreeNodeBuilder) SetDisplayMode(displayMode DisplayMode) {
    n.DisplayMode = displayMode
}

func (n *TreeNodeBuilder) Build() {

    switch n.DisplayMode {
	case ByDate:
	    n.ByDate()
	case BySpecialty:
	    n.BySpecialty()
	case ByType:
	    n.ByType()
	default:
		panic(fmt.Sprintf("Unknown display mode: %s", string(n.DisplayMode)))
	}
}

func (n *TreeNodeBuilder) ToString() string {
    return n.TreeNode.ToString("")
}

func (n *TreeNodeBuilder) ByDate() {

    // Sort the documents by date in descending order (most recent first)
    sort.Sort(byDate(n.Documents))

    // Create a map to store child nodes by year
	yearNodes := make(map[int]*TreeNode)

	// Iterate through the documents and organize them into the tree structure
	for _, document := range n.Documents {
		year := document.Date.Year()

		// Check if a child node for the year exists, otherwise create one
		yearNode, exists := yearNodes[year]
		if !exists {
			yearNode = NewTreeNode(fmt.Sprintf("%d", year))
			yearNodes[year] = yearNode
			n.TreeNode.AddChild(yearNode)
		}

		// Add the document to the child node for the year
		yearNode.AddDocument(document)
	}
}

func (n *TreeNodeBuilder) BySpecialty() {
    // Sort documents by specialty and type for efficient traversal
    sort.Slice(n.Documents, func(i, j int) bool {
        if n.Documents[i].Specialty == n.Documents[j].Specialty {
            return n.Documents[i].Type < n.Documents[j].Type
        }
        return n.Documents[i].Specialty < n.Documents[j].Specialty
    })

    for _, document := range n.Documents {
        specialtyNode := n.TreeNode.GetOrCreateChild(document.Specialty)
        typeNode := specialtyNode.GetOrCreateChild(document.Type)
        typeNode.AddDocument(document)
    }
}

func (n *TreeNodeBuilder) ByType() {
    // Sort documents by type and specialty for efficient traversal
    sort.Slice(n.Documents, func(i, j int) bool {
        if n.Documents[i].Type == n.Documents[j].Type {
            return n.Documents[i].Specialty < n.Documents[j].Specialty
        }
        return n.Documents[i].Type < n.Documents[j].Type
    })

    for _, document := range n.Documents {
        typeNode := n.TreeNode.GetOrCreateChild(document.Type)
        specialtyNode := typeNode.GetOrCreateChild(document.Specialty)
        specialtyNode.AddDocument(document)
    }
}
