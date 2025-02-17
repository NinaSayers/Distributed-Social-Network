package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func displayPosts(messages []Message) {
	fmt.Println("Posts:")
	fmt.Println(strings.Repeat("=", 50))
	if len(messages) == 0 {
		fmt.Println("No hay mensajes")
	}
	for i, message := range messages {
		if i == 5 {
			break
		}
		fmt.Printf("\n%s: %s\n", message.UserID, message.Content)
		fmt.Printf("  📅 %s  🆔 %s\n ", message.CreatedAt, message.MessageID)
	}
	fmt.Println(strings.Repeat("-", 50))
}

func displayPost(message Message, user User) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("\n%s: %s\n", user.UserName, message.Content)
	fmt.Printf("  📅 %s\n", message.CreatedAt)
	fmt.Println(strings.Repeat("-", 50))
}

func displayRepost(message Message, reposter User, original User) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("  🔄 Reposteado por: %s\n", reposter.UserName)
	fmt.Println(strings.Repeat(".", 50))
	fmt.Printf("\n%s: %s\n", reposter.UserName, message.Content)
	fmt.Printf("  📅 %s\n", message.CreatedAt)
	fmt.Println(strings.Repeat("-", 50))
}

func displayProfileHeader(user User) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("ID: %s  %s (%s)\n", user.UserID, user.UserName, user.Email)
	fmt.Printf("  🗓️  Joined %s\n", user.CreatedAt.Local())
	// fmt.Printf("  📊 %d Following   👥 %d Followers\n", user.Following, user.Followers)
	fmt.Println(strings.Repeat("=", 50))
}

func displayUser(user User) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("ID: %d - %s (%s)\n", user.UserID, user.UserName, user.Email)
	fmt.Println(strings.Repeat("-", 50))
}

func displayUsers(users []User) {
	fmt.Println(strings.Repeat("=", 50))
	if len(users) == 0 {
		fmt.Println("No hay usuarios")
	}
	for _, user := range users {
		displayUser(user)
	}
	fmt.Println(strings.Repeat("=", 50))
}

func displayProfile(user User, messages []Message) {
	displayProfileHeader(user)
	displayPosts(messages)
}

func pressKeyToContinue() {
	fmt.Println("Presiona cualquier tecla para continuar...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadByte()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
}
