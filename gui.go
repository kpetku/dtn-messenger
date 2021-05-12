package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	app fyne.App
	w   fyne.Window

	Roster *Roster
}

func NewGUI() *GUI {
	gui := new(GUI)
	gui.app = app.New()
	gui.w = gui.app.NewWindow("dtn-messenger v0.0.1")
	gui.w.Resize(fyne.NewSize(800, 600))

	gui.Roster = gui.NewRoster()
	return gui
}

// NewRoster returns a new Roster
func (g *GUI) NewRoster() *Roster {
	r := new(Roster)
	r.rMap = make(map[string][]string)
	r.tree = widget.NewTreeWithStrings(r.rMap)
	r.tree.OnSelected = func(uid string) {
		g.w.SetContent(g.ShowConversation(uid))
		g.w.Show()
	}

	r.tree.Show()
	r.tree.Refresh()

	return r
}

// NewMessage renders a new message of content for contactName (the message can have an Inbound or Outbound direction)
func (g *GUI) NewMessage(contactName, direction, content string) {
	for num, contact := range g.Roster.contacts {
		if contact.name == contactName {
			g.Roster.contacts[num].Conversation.AppendString(direction, content)
			g.Roster.contacts[num].Conversation.sb.Refresh()
			g.Roster.contacts[num].Conversation.sb.ScrollToBottom()
		}
	}
}

// RefreshRoster reloads a roster from the existing list of contacts
func (g *GUI) RefreshRoster(r *Roster) *Roster {
	r.rMap = make(map[string][]string)
	r.tree = widget.NewTreeWithStrings(r.rMap)
	r.tree.OnSelected = func(uid string) {
		g.w.SetContent(g.ShowConversation(uid))
		g.w.Show()
	}
	for _, contact := range r.contacts {
		r.rMap[""] = append(r.rMap[""], contact.name)
	}
	return r
}

func (g *GUI) ShowConversation(name string) *fyne.Container {
	for num, contact := range g.Roster.contacts {
		if contact.name == name {
			back := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
				g.RefreshRoster(g.Roster)
				g.Roster.tree.Show()
				g.w.SetContent(g.Roster.tree)
			})
			top := container.New(layout.NewGridLayout(5), back, layout.NewSpacer(), widget.NewLabel(name))
			bottom := g.footerEntry(name)
			content := container.New(layout.NewBorderLayout(top, bottom, nil, nil), top, g.Roster.contacts[num].Conversation.sb, bottom)
			return content
		}
	}
	return container.New(layout.NewVBoxLayout())
}

func (g *GUI) footerEntry(name string) *widget.Form {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Message")
	entry.OnSubmitted = func(s string) {
		if entry.Text != "" {
			g.NewMessage(name, Outbound, s)
			entry.SetText("")
		}
	}
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: entry}},
		OnSubmit: func() {
			if entry.Text != "" {
				g.NewMessage(name, Outbound, entry.Text)
				entry.SetText("")
			}
		},
	}
	form.SubmitText = "Message"
	return form
}
