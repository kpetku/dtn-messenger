package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const hardWrapLength = 45

const (
	Inbound  = "inbound"
	Outbound = "outbound"
)

// Conversation contains a list of inbound or outbound messages in a Conversation stored in a fyne Vbox container
type Conversation struct {
	vbox *fyne.Container
	sb   *container.Scroll
}

// NewConversation returns a new Conversation
func NewConversation() *Conversation {
	h := new(Conversation)
	h.vbox = container.NewVBox()
	h.sb = container.NewVScroll(h.vbox)
	return h
}

// AppendString accepts a direction (Inbound or Outbound) along with message content and adds it to a Conversation
func (h *Conversation) AppendString(direction, content string) {
	switch direction {
	case Inbound:
		hbox := container.NewHBox()
		pic := canvas.NewImageFromResource(theme.HomeIcon())
		pic.SetMinSize(fyne.NewSize(32, 32))
		pic.FillMode = canvas.ImageFillContain
		hbox.Add(pic)
		wrappedContent := widget.NewTextGridFromString(wrap(content, hardWrapLength))
		for num := range wrappedContent.Rows {
			wrappedContent.SetRowStyle(num, &widget.CustomTextGridStyle{FGColor: theme.ForegroundColor(), BGColor: theme.PrimaryColor()})
		}
		hbox.Add(wrappedContent)
		h.vbox.Add(hbox)
	case Outbound:
		pic := canvas.NewImageFromResource(theme.HomeIcon())
		pic.SetMinSize(fyne.NewSize(32, 32))
		pic.FillMode = canvas.ImageFillContain
		wrappedContent := widget.NewTextGridFromString(wrap(content, hardWrapLength))
		for num := range wrappedContent.Rows {
			wrappedContent.SetRowStyle(num, &widget.CustomTextGridStyle{FGColor: theme.ForegroundColor(), BGColor: theme.ErrorColor()})
		}
		h.vbox.Add(container.New(layout.NewHBoxLayout(), widget.NewLabel(""), layout.NewSpacer(), wrappedContent, pic))
	}
}

// https://gist.github.com/kennwhite/f3881b815f43e0d9d7bd3ef8166e5d1b
func wrap(text string, colBreak int) string {
	if colBreak < 1 {
		return text
	}
	text = strings.TrimSpace(text)
	wrapped := ""
	var i int
	for i = 0; len(text[i:]) > colBreak; i += colBreak {
		wrapped += text[i:i+colBreak] + "\n"
	}
	wrapped += text[i:]
	return wrapped
}
