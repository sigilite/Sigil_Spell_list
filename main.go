package main

import "fmt"

func main() {
	fmt.Println("Sigil in Go!")
}

type Sigil interface {
	Name() string
	Size() int
	Text() string
}

type Charm struct {
	name string
	text string
}

func (c *Charm) Size() int {
	return 1
}

func (c *Charm) Name() string {
	return c.name
}

func (c *Charm) Text() string {
	return c.text
}

type MinorSpell struct {
	name string
	text string
}

func (s *MinorSpell) Size() int {
	return 3
}

func (s *MinorSpell) Name() string {
	return s.name
}

func (s *MinorSpell) Text() string {
	return s.text
}

type MajorSpell struct {
	name string
	text string
}

func (s *MajorSpell) Size() int {
	return 5
}

func (s *MajorSpell) Name() string {
	return s.name
}

func (s *MajorSpell) Text() string {
	return s.text
}
