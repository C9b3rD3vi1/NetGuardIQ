package utils

import (
	"fmt"
)

func SendPhishEmail(name, email, link string) error {
	fmt.Printf("Sending phishing email to %s (%s) with link %s\n", name, email, link)
	return nil
}