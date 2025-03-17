package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
)

// Config struct holds the structure of the configuration
type Config struct {
	APIKey string `json:"api_key"`
}

// LoadConfig loads the configuration from a JSON file
func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

const API_URL = "https://api.quantumnumbers.anu.edu.au"

// GetQuantumNumbers makes an API request to fetch quantum numbers
func GetQuantumNumbers(apiKey string, length int, dtype string, blockSize int) ([]byte, error) {
	url := fmt.Sprintf("%s?length=%d&type=%s&size=%d", API_URL, length, dtype, blockSize)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the error message from the response body
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed: %s", body)
	}

	return ioutil.ReadAll(resp.Body)
}

// Password Generation Logic
func processPasswordData(data []byte, startUppercase, includeSpecialChar, includeNumbers bool, length int) string {
	password := ""
	quantumNumbers := make([]int, length)

	// Parse the quantum numbers from the API response
	for i := 0; i < length; i++ {
		quantumNumbers[i] = int(data[i]) // Convert byte to int
	}

	for i := 0; i < length; i++ {
		var char rune

		// Use the quantum number to decide the type of character
		charType := quantumNumbers[i] % 4 // 0: lowercase, 1: uppercase, 2: number, 3: special

		switch charType {
		case 0: // Lowercase letter
			char = rune(quantumNumbers[i]%26 + 97)
		case 1: // Uppercase letter
			char = rune(quantumNumbers[i]%26 + 65)
		case 2: // Number
			if includeNumbers {
				char = rune(quantumNumbers[i]%10 + 48)
			} else {
				char = rune(quantumNumbers[i]%26 + 97) // Fallback to lowercase if numbers are disabled
			}
		case 3: // Special character
			if includeSpecialChar {
				specialChars := "!@#$%^&*"
				char = rune(specialChars[quantumNumbers[i]%len(specialChars)])
			} else {
				char = rune(quantumNumbers[i]%26 + 97) // Fallback to lowercase if special chars are disabled
			}
		}

		password += string(char)
	}

	// Handle special case: first character must be uppercase
	if startUppercase && len(password) > 0 {
		// Replace the first character with a random uppercase letter
		password = string(rune(quantumNumbers[0]%26 + 65)) + password[1:]
	}

	return password
}

func main() {
	// Create Fyne app
	a := app.New()

	// Load the icon.ico file
	icon, err := fyne.LoadResourceFromPath("icon.ico")
	if err != nil {
		log.Println("Failed to load icon:", err)
		// Continue without the icon if it fails to load
	} else {
		// Set the application icon
		a.SetIcon(icon)
	}

	w := a.NewWindow("quantumpass")

	// Set window size (wider and taller)
	w.Resize(fyne.NewSize(600, 400))

	// Length entry
	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("Enter password length (0-255)")

	// Password options
	uppercaseCheckbox := widget.NewCheck("Start with Uppercase", nil)
	specialCharCheckbox := widget.NewCheck("Include Special Characters", nil)
	numbersCheckbox := widget.NewCheck("Include Numbers", nil)

	// Label to display the generated password (with wrapping enabled)
	passwordLabel := widget.NewLabel("")
	passwordLabel.Wrapping = fyne.TextWrapWord

	// Button to generate the password
	generateButton := widget.NewButton("Generate Password", func() {
	// Get the length and options from the UI
	lengthStr := lengthEntry.Text
	length, err := strconv.Atoi(lengthStr)
	if err != nil || length < 0 || length > 255 {
		passwordLabel.SetText("Invalid length. Please enter a number between 0 and 255.")
		return
	}

	startUppercase := uppercaseCheckbox.Checked
	includeSpecialChar := specialCharCheckbox.Checked
	includeNumbers := numbersCheckbox.Checked

	// Load the config file
	config, err := loadConfig("config.json")
	if err != nil || config.APIKey == "" {
		// Stop execution and show the error message if API key is missing or config load fails
		passwordLabel.SetText("Add API key to config.json")
		return
	}

	// Fetch quantum numbers from the API
	data, err := GetQuantumNumbers(config.APIKey, length, "uint8", 100)
	if err != nil {
		// Stop execution and show the error message if API request fails
		passwordLabel.SetText("Error fetching quantum numbers: Add API key to config.json")
		return
	}

	// If everything is successful, then process the data into a password
	password := processPasswordData(data, startUppercase, includeSpecialChar, includeNumbers, length)
	passwordLabel.SetText(password) // Only display the password
})



	// Button to copy the password to clipboard
	copyButton := widget.NewButtonWithIcon("", fyne.NewStaticResource("copy", []byte("📋")), func() {
		clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
		clipboard.SetContent(passwordLabel.Text) // Copy only the password
		dialog.ShowInformation("Copied", "Password copied to clipboard!", w)
	})

	// Hyperlink to Patreon (bottom right corner)
	patreonLink := widget.NewHyperlink("Magic Industries", nil)
	patreonLink.SetURLFromString("https://patreon.com/magicindustriessoftware")

	// Layout with the hyperlink in the bottom right corner
	content := container.NewBorder(
		nil, // Top
		container.NewHBox(
			container.NewHBox(), // Spacer to push the hyperlink to the right
			patreonLink,
		), // Bottom
		nil, // Left
		nil, // Right
		container.NewVBox(
			lengthEntry,
			uppercaseCheckbox,
			specialCharCheckbox,
			numbersCheckbox,
			generateButton,
			container.NewHBox(
				container.NewScroll(passwordLabel), // Wrap the label in a scroll container
				copyButton,
			),
		),
	)

	// Set the content for the window
	w.SetContent(content)

	// Show and run the app
	w.ShowAndRun()
}