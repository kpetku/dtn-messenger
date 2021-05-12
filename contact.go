package main

import (
	"image/color"
)

// Contact stores a Conversation by name
type Contact struct {
	name         string
	profilePic   string
	textColor    color.Color
	Conversation *Conversation
}

type GroupContact struct {
	name         string
	Conversation *Conversation
}
