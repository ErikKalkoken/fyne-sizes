package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Sizes")
	c := container.NewVBox(
		makeTitle("Caption", theme.SizeNameCaptionText),
		makeLabelWithSize(lorem, theme.SizeNameCaptionText),
		makeTitle("Normal", theme.SizeNameText),
		makeLabelWithSize(lorem, theme.SizeNameText),
		makeTitle("Sub Heading", theme.SizeNameSubHeadingText),
		makeLabelWithSize(lorem, theme.SizeNameSubHeadingText),
		makeTitle("Heading", theme.SizeNameHeadingText),
		makeLabelWithSize(lorem, theme.SizeNameHeadingText),
		widget.NewLabel(fmt.Sprintf("Icon %ddp", int(theme.Size(theme.SizeNameInlineIcon)))),
		container.NewHBox(widget.NewLabel("Icon: "), widget.NewIcon(theme.ComputerIcon())),
	)
	w.SetContent(container.NewVScroll(c))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func makeTitle(text string, sizeName fyne.ThemeSizeName) *widget.RichText {
	return makeLabelWithSize(
		fmt.Sprintf("%s %ddp", text, int(theme.Size(sizeName))),
		sizeName,
	)
}

func makeLabelWithSize(text string, sizeName fyne.ThemeSizeName) *widget.RichText {
	x := widget.NewRichText(&widget.TextSegment{
		Style: widget.RichTextStyle{
			ColorName: theme.ColorNameForeground,
			SizeName:  sizeName,
			Inline:    true,
		},
		Text: text,
	})
	x.Truncation = fyne.TextTruncateEllipsis
	return x
}
