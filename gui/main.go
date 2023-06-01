package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Laodeus/glt/gui/apiCall"
	"github.com/Laodeus/glt/utils"
)

func main() {

	utils.LoadEnv()

	myApp := app.New()
	myWindow := myApp.NewWindow("Glt")

	label := widget.NewLabel("Connexion")
	inputUsername := widget.NewEntry()
	inputUsername.PlaceHolder = "Username"
	inputPassword := widget.NewPasswordEntry()
	inputPassword.PlaceHolder = "secret"

	loginButton := widget.NewButton("Login", func() {
		response, err := apiCall.Login(inputUsername.Text, inputPassword.Text)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(response)
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
