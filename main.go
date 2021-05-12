package main

func main() {
	gui := NewGUI()

	gui.Roster.AppendContact("Foo")

	gui.NewMessage("Foo", Inbound, "Hello World")
	gui.NewMessage("Foo", Inbound, "bar")

	gui.Roster.AppendContact("Bar")
	gui.NewMessage("Bar", Inbound, "baz")

	gui.Roster.AppendContact("qux")

	gui.w.SetContent(gui.Roster.tree)
	gui.w.ShowAndRun()
}
