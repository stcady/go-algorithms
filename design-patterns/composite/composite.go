package main

import "fmt"

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

// perform method for leaflet class
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// perform method for branch class
func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}
	for _, branch := range branch.branches {
		branch.perform()
	}
}

// add method for branch class
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// addBranch method for the branch class
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// getLeaflets method for branch class
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func main() {
	branch1 := &Branch{name: "branch 1"}
	branch2 := Branch{name: "branch 2"}
	leaf1 := Leaflet{name: "leaf 1"}
	leaf2 := Leaflet{name: "leaf 2"}
	branch1.add(leaf1)
	branch1.add(leaf2)
	branch1.addBranch(branch2)
	branch1.perform()
}
