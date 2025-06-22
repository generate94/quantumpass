Free pre-compiled binary available on my Gumroad: https://magicindustries.gumroad.com/l/quantumpass

# QuantumPass - Instructions for Compilation

QuantumPass is a password generator based on quantum numbers from the AQN API. This README provides instructions for compiling the application on different systems and explains the inclusion of a non-maintained `.exe` binary for Windows users.

# Rationale

QuantumPass’s core randomness source is truly quantum-generated numbers, offering high entropy. However, mapping these raw quantum values to fit typical password constraints (letters, digits, special characters) significantly reduces the effective randomness because it restricts the possible character set per position, lowering entropy per character.

Moreover, common password analyzers and crackers often expect passwords encoded in standard ways (like Base64, hex, or common alphanumeric sets). These encodings introduce patterns and structure, making the passwords appear weaker or easier to predict to automated tools.

In contrast, QuantumPass can output raw quantum numbers as digits only. While this may look unusual (only digits), it retains full entropy from the quantum source since no compression or re-encoding happens. Password analyzers and crackers rarely expect or specifically target purely numeric passwords of sufficient length generated from quantum randomness, making them effectively more secure than they might seem at first glance.

This represents a tradeoff between usability and maximum entropy:

Applying constraints (uppercase, special chars, etc.) makes passwords more compatible with typical password policies but reduces entropy.

Using raw quantum digits maximizes entropy but may not meet all password policy requirements.

Increasing password length compensates for entropy loss caused by constraints.

This approach emphasizes true randomness over conformity to typical password formats, pushing the boundaries of password security.

## Files Provided

The following files are included in the repository:
- `main.go`: The main application logic (UI and password generation).
- `quantum.go`: The logic for interacting with the AQN API.
- `config.json`: Configuration file (contains your API key).
- `icon.ico`: Icon used for the executable (Windows).
- `go.mod` & `go.sum`: Go module files for managing dependencies.
- `README.txt`: This file.

## Prerequisites

Before you compile QuantumPass, make sure you have the following installed:

- [Go](https://golang.org/dl/) (Version 1.18+)

- A working internet connection (for fetching quantum numbers from the AQN API)

## Setup and Compilation

### 1. Clone the Repository

If you haven't already, clone this repository to your local machine:

git clone https://github.com/yourusername/quantumpass.git cd quantumpass


### 2. Configure API Key

To interact with the AQN API, you'll need an API key. Here's how to set it up:

1. Visit [AQN API](https://api.quantumnumbers.anu.edu.au) and sign up for an API key.
2. Open `config.json` and paste your API key in the `APIKey` field.
   Example:

   {
     "APIKey": "your__api_key_here"
   }

3. Install Dependencies

To install the required dependencies, run the following command:


go mod tidy

4. Compile for Your System

Windows

For Windows users, a precompiled .exe binary is provided (quantumpass.exe). This binary is not maintained and may not work with future changes to the code. It is provided for convenience. To run it, simply double-click the quantumpass.exe file.

Linux & macOS

To compile QuantumPass for Linux or macOS, run the following commands:

For Linux:


GOOS=linux GOARCH=amd64 go build -o quantumpass-linux main.go
This will create an executable named quantumpass-linux in the same directory.

For macOS:


GOOS=darwin GOARCH=amd64 go build -o quantumpass-macos main.go
This will create an executable named quantumpass-macos in the same directory.

Once compiled, you can run the application using:


./quantumpass-linux

or


./quantumpass-macos

5. Run the Application

After compiling the application for your system, run it:

On Windows: Double-click quantumpass.exe.

On Linux/macOS: Run the compiled binary from the terminal:

./quantumpass-linux
or

./quantumpass-macos

6. Using the Application

Once the application is running, you can enter the desired password length and other options (uppercase, special characters, etc.). The program will generate a secure password based on quantum numbers retrieved from the AQN API.

Notes

Non-maintained Windows Binary: The quantumpass.exe binary provided is only for Windows users and is not maintained. It may not be up to date with the latest code changes. You are encouraged to compile the project from source to ensure you have the latest version.

The application requires a valid API key for the AQN API to function. Make sure to update config.json with your key before using the app.
If you encounter any issues during compilation, feel free to open an issue in the GitHub repository.

License

This project is licensed under the MIT License. See LICENSE for more details.



