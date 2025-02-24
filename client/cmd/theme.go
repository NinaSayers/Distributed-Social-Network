package main

import (
    "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
)

// PastelTheme define un tema personalizado con colores pasteles.
type PastelTheme struct{}

func (t PastelTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
    switch name {
    case theme.ColorNameBackground:
        return color.NRGBA{R: 173, G: 216, B: 230, A: 255} // Azul pastel
    case theme.ColorNameForeground:
        return color.NRGBA{R: 147, G: 112, B: 219, A: 255} // Morado pastel
    case theme.ColorNamePrimary:
        return color.NRGBA{R: 102, G: 205, B: 170, A: 255} // Verde aqua pastel
    default:
        return theme.DefaultTheme().Color(name, variant)
    }
}

func (t PastelTheme) Font(style fyne.TextStyle) fyne.Resource {
    return theme.DefaultTheme().Font(style)
}

func (t PastelTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
    return theme.DefaultTheme().Icon(name)
}

func (t PastelTheme) Size(name fyne.ThemeSizeName) float32 {
    return theme.DefaultTheme().Size(name)
}