package main

import (
    "errors" // Importa el paquete errors
    "fmt"
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

func secureHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
        w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "deny")
        w.Header().Set("X-XSS-Protection", "0")
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func (app *application) logRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
        next.ServeHTTP(w, r)
    })
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                w.Header().Set("Connection", "close")
                app.serverError(w, fmt.Errorf("%s", err))
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func (app *application) JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extraer el token del header "Authorization"
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        // Eliminar el prefijo "Bearer " del token
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == "" {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        // Parsear el token y validar los claims
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        // Manejar errores específicos de jwt/v5
        if err != nil {
            switch {
            case errors.Is(err, jwt.ErrTokenMalformed):
                http.Error(w, "Malformed token", http.StatusUnauthorized)
            case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
                http.Error(w, "Token expired or not valid yet", http.StatusUnauthorized)
            default:
                http.Error(w, "Invalid token", http.StatusUnauthorized)
            }
            return
        }

        // Verificar si el token es válido
        if !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Si el token es válido, pasar al siguiente handler
        next.ServeHTTP(w, r)
    })
}