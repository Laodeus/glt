package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Glt")

	label := widget.NewLabel("Connexion")
	inputUsername := widget.NewEntry()
	inputUsername.PlaceHolder = "Username"
	inputPassword := widget.NewPasswordEntry()
	inputPassword.PlaceHolder = "secret"

	loginButton := widget.NewButton("Login", func() {
		// Action to perform when the "Login" button is clicked
	})

	registerButton := widget.NewButton("Register", func() {
		// Action to perform when the "Register" button is clicked
	})

	content := container.NewVBox(
		label,
		inputUsername,
		inputPassword,
		loginButton,
		registerButton,
	)

	content.Resize(content.MinSize())

	myWindow.SetContent(container.NewCenter(content))
	newSize := fyne.NewSize(content.MinSize().Width*5, content.MinSize().Height*2)
	myWindow.Resize(newSize)

	myWindow.SetFixedSize(false)
	myWindow.ShowAndRun()
}
