//+build windows

package main

import (
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonutz/ccc"
	"github.com/gonutz/wui/v2"
)

func runGui(inputFilePath string) {
	windowFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	window := wui.NewWindow()
	window.SetFont(windowFont)
	window.SetInnerSize(402, 143)
	window.SetTitle("ccc - xor file with random numbers")
	window.SetHasMinButton(false)
	window.SetHasMaxButton(false)
	window.SetResizable(false)

	icon, _ := wui.NewIconFromExeResource(10)
	window.SetIcon(icon)

	seed := wui.NewIntUpDown()
	seed.SetHorizontalAnchor(wui.AnchorCenter)
	seed.SetBounds(161, 59, 80, 22)
	window.Add(seed)

	seedCaption := wui.NewLabel()
	seedCaption.SetHorizontalAnchor(wui.AnchorCenter)
	seedCaption.SetBounds(106, 62, 50, 13)
	seedCaption.SetText("Seed")
	seedCaption.SetAlignment(wui.AlignRight)
	window.Add(seedCaption)

	okButton := wui.NewButton()
	okButton.SetAnchors(wui.AnchorCenter, wui.AnchorMax)
	okButton.SetBounds(161, 103, 80, 25)
	okButton.SetText("Save As...")
	okButton.SetOnClick(func() {
		save := wui.NewFileSaveDialog()

		if strings.HasSuffix(strings.ToLower(inputFilePath), ".ccc") {
			f := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath))
			f, _ = filepath.Abs(f)
			save.SetInitialPath(f)
		} else {
			save.AddFilter("ccc file", ".ccc")
			f, _ := filepath.Abs(inputFilePath)
			f += ".ccc"
			save.SetInitialPath(f)
		}

		if ok, outputFilePath := save.Execute(window); ok {
			fail := func(msg string, err error) {
				wui.MessageBoxError("Error", msg+": "+err.Error())
			}

			in, err := os.Open(inputFilePath)
			if err != nil {
				fail("Unable to open input file", err)
				return
			}
			defer in.Close()

			out, err := os.Create(outputFilePath)
			if err != nil {
				fail("Unable to create output file", err)
				return
			}
			defer out.Close()

			random := rand.New(rand.NewSource(int64(seed.Value())))
			xor := ccc.NewXORReader(in, ccc.NewFuncReader(func() byte {
				return byte(random.Intn(256))
			}))
			_, err = io.Copy(out, xor)
			if err != nil {
				fail("Failed to write output file", err)
				return
			}

			window.Close()
		}
	})
	window.Add(okButton)

	inputCaption := wui.NewLabel()
	inputCaption.SetHorizontalAnchor(wui.AnchorMinAndMax)
	inputCaption.SetBounds(5, 25, 392, 13)
	inputCaption.SetText(inputFilePath)
	inputCaption.SetAlignment(wui.AlignCenter)
	window.Add(inputCaption)

	window.SetOnShow(func() {
		seed.Focus()
		seed.SelectAll()
	})

	window.SetShortcut(window.Close, wui.KeyEscape)
	window.SetShortcut(okButton.OnClick(), wui.KeyReturn)

	window.Show()
}
