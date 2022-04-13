package tree

import (
	"errors"
	"fmt"
)

//Hello World, added to debug using packages from github
func He() {
	fmt.Println("hello world")
}

//default node, uses string as root
type Node struct {
	root   string
	l_node *Node
	r_node *Node
}

//creates a new node
func create_node(root string) *Node {
	new_node := Node{root, nil, nil}
	return &new_node
}

//edits the left node of a given node
func Edit_l_node(node *Node, new_val *Node) {
	node.l_node = new_val
}

//edits the right node of a given node
func Edit_r_node(node *Node, new_val *Node) {
	node.r_node = new_val
}

//Basic tree insertion
func insert_node(parent, child, insertee *Node) error {
	if child == parent.l_node {
		insertee.l_node, parent.l_node = child, insertee
		return nil
	} else if child == parent.r_node {
		insertee.r_node, parent.r_node = child, insertee
		return nil
	} else {
		return errors.New("node not child of root")
	}
}

//Basic node deletion
func delete_node(parent, child *Node) error {
	if child == parent.l_node {
		parent.l_node = nil
		return nil
	}
	if child == parent.r_node {
		parent.r_node = nil
		return nil
	}
	return errors.New("node not child of root")
}

//Deletes node and all children
func delete_all_nodes(parent, child *Node) error {

	//end case
	if child == nil {
		return nil
	}

	//Delete child
	err := delete_node(parent, child)
	if err != nil {
		return err
	}

	//Delete left node and all its children
	err = delete_all_nodes(child, child.l_node)
	if err != nil {
		return err
	}

	//Delete right node and all its children
	err = delete_all_nodes(child, child.l_node)
	if err != nil {
		return err
	}
	return nil
}
