package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	phonebook := make(map[string]string) // Initialize the phonebook map

	scanner := bufio.NewScanner(os.Stdin) // To handle user input
	for {
		// Display the menu
		fmt.Println("\nPhonebook Menu:")
		fmt.Println("1. Add a new contact")
		fmt.Println("2. Search for a contact")
		fmt.Println("3. List all contacts")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice) // Read user's menu choice

		switch choice {
		case 1:
			// Add a new contact
			fmt.Print("Enter contact name: ")
			scanner.Scan() // Read input with spaces
			name := scanner.Text()
			fmt.Print("Enter phone number: ")
			scanner.Scan()
			phone := scanner.Text()
			phonebook[name] = phone
			fmt.Printf("Contact %s added successfully!\n", name)

		case 2:
			// Search for a contact
			fmt.Print("Enter contact name to search: ")
			scanner.Scan()
			name := scanner.Text()
			phone, exists := phonebook[name]
			if exists {
				fmt.Printf("Contact found: %s - %s\n", name, phone)
			} else {
				fmt.Printf("No contact found for name: %s\n", name)
			}

		case 3:
			// List all contacts
			fmt.Println("\nAll Contacts:")
			if len(phonebook) == 0 {
				fmt.Println("No contacts in the phonebook.")
			} else {
				for name, phone := range phonebook {
					fmt.Printf("%s: %s\n", name, phone)
				}
			}

		case 4:
			// Exit the program
			fmt.Println("Exiting Phonebook. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

