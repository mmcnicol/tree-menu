package main

import (
	"fmt"
	"sort"
)

// TreeNode represents a node in the tree menu.
type TreeNode struct {
	Name      string
	Children  []*TreeNode
	Documents []Document
	DocumentCount int
}

// NewTreeNode creates a new TreeNode with the given name.
func NewTreeNode(name string) *TreeNode {
	return &TreeNode{
		Name:      name,
		Children:  make([]*TreeNode, 0),
		Documents: make([]Document, 0),
		DocumentCount: 0,
	}
}

// AddChild adds a child node to the current node.
func (n *TreeNode) AddChild(child *TreeNode) {
	n.Children = append(n.Children, child)
}

// AddDocument adds a document to the node.
func (n *TreeNode) AddDocument(doc Document) {
	n.Documents = append(n.Documents, doc)
}

// GetOrCreateChild gets an existing child node or creates a new one.
func (n *TreeNode) GetOrCreateChild(name string) *TreeNode {
	// Check if a child with the specified name already exists
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}

	// If not, create a new child node, add it to the current node, and return it
	newChild := NewTreeNode(name)
	n.AddChild(newChild)
	return newChild
}

func (n *TreeNode) SortNodeDocumentsRecursive() {
    // Sort the Documents slice by date, most recent first
    sort.Slice(n.Documents, func(i, j int) bool {
        return n.Documents[i].Date.After(n.Documents[j].Date)
    })

    for _, child := range n.Children {
        child.SortNodeDocumentsRecursive()
    }
}

// UpdateDocumentCountRecursive updates the document count for the current node and its children.
func (n *TreeNode) UpdateDocumentCountRecursive() {
    // Initialize document count for the current node
    n.DocumentCount = len(n.Documents)

    // Update document count for each child node
    for _, child := range n.Children {
        child.UpdateDocumentCountRecursive()
        // Add the child's document count to the current node's count
        n.DocumentCount += child.DocumentCount
    }
}

// ToString converts the tree structure to a string representation.
func (n *TreeNode) ToString(indent string) string {
	var result string
	result += fmt.Sprintf("%s%s (%d)\n", indent, n.Name, n.DocumentCount)

	for _, child := range n.Children {
		result += child.ToString(indent + "  ")
	}

	for _, doc := range n.Documents {
		result += fmt.Sprintf("%s  - Date: %s, Specialty: %s, Type: %s\n",
			indent, doc.Date.Format("2006-01-02"), doc.Specialty, doc.Type)
	}

	return result
}
