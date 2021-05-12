package main

import "fyne.io/fyne/v2/widget"

type Roster struct {
	tree     *widget.Tree
	rMap     map[string][]string
	contacts []Contact
}

// AppendContact adds a new contact "name" to the roster
func (r *Roster) AppendContact(name string) *Roster {
	r.rMap[""] = append(r.rMap[""], name)
	c := new(Contact)
	c.name = name
	c.Conversation = NewConversation()
	r.contacts = append(r.contacts, *c)
	return r
}
