package main

import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    //"fyne.io/fyne/v2/container"
    //"fyne.io/fyne/v2/widget"
    "os"
)

var baseURL string

func init() {
    baseURL = os.Getenv("SERVER_URL")
    if baseURL == "" {
        baseURL = "http://10.0.10.6:4000"
    }
}

type Application struct {
    service *Service
    token   string
    user    *User
}

func main() {
    // Crear la aplicación Fyne
    myApp := app.New()
    myApp.Settings().SetTheme(&PastelTheme{}) // Aplicar el tema personalizado

    // Crear la ventana principal
    myWindow := myApp.NewWindow("Red Social")
    myWindow.Resize(fyne.NewSize(800, 600)) // Tamaño inicial de la ventana

    // Inicializar la aplicación
    app := Application{service: NewService()}

    // Mostrar la pantalla de inicio (login/registro)
    showLoginScreen(myWindow, &app)

    // Mostrar la ventana
    myWindow.ShowAndRun()
}

