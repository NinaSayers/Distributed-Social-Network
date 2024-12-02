package main

import (
	"fmt"
	"strings"
)

func displayPosts(tweets []Tweet) {
	fmt.Println("Posts:")
	for i, tweet := range tweets {
		if i == 5 {
			break
		}
		fmt.Printf("\n%s: %s\n", tweet.UserId, tweet.Content)
		fmt.Printf("  📅 %s\n", tweet.CreatedAt)
	}
	fmt.Println(strings.Repeat("-", 50))
}
