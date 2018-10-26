package pathtree

import (
	"fmt"
	"io"
	"strings"
)

// Node represents a node of a PathTree
type Node struct {
	// Children are the child nodes
	// using slice instead of a map to preserve order
	Children []*Node
	// Name is the basename of the node
	Name string
}

// insert Inserts a new child to the node and returns the child
func (n *Node) insert(child string) *Node {
	for _, c := range n.Children {
		if c.Name == child {
			return c
		}
	}
	c := &Node{
		Name:     child,
		Children: []*Node{},
	}
	n.Children = append(n.Children, c)
	return c
}

// print recursively prints a Node and its children
func (n *Node) print(w io.Writer, indent string) {
	fmt.Fprintf(w, "%s%s\n", indent, n.Name)
	indent = strings.Repeat(indent, 2)
	for _, c := range n.Children {
		c.print(w, indent)
	}
}

// PathTree represents a root node
type PathTree Node

// New returns a new PathTree
func New(paths ...string) *PathTree {
	p := &PathTree{
		Name:     "/",
		Children: []*Node{},
	}
	for _, path := range paths {
		if err := p.Add(path); err != nil {
			return nil
		}
	}
	return p
}

// Root returns the root of the PathTree
func (p *PathTree) Root() string {
	return p.Name
}

// Add adds a given path to the path tree
func (p *PathTree) Add(path string) error {
	if !strings.HasPrefix(path, "/") {
		return fmt.Errorf("invalid root - paths must be absolute")
	}
	components := strings.Split(path, "/")

	curr := (*Node)(p)
	for _, c := range components[1:] {
		// c == "" if multiple slashes are specified in a path
		if c == "" {
			continue
		}
		curr = curr.insert(c)
	}

	return nil
}

// Print prints the PathTree to a given io.Writer
func (p *PathTree) Print(w io.Writer) {
	(*Node)(p).print(w, " ")
}
