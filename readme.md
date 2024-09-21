# Gemini AI Assistant CLI

This command-line tool allows you to interact with the Gemini AI directly from your terminal. Ask questions, get information, or engage in conversations with the AI assistant powered by Google's Gemini API.

## Features

- Interact with Gemini AI directly from your command line
- Ask questions, get information, or have conversations on various topics
- Powered by Google's Gemini AI for accurate and context-rich responses

## Prerequisites

- Go programming language installed on your system
- Gemini AI API key (you can obtain one from [Google AI Studio](https://makersuite.google.com/app/apikey))

## Installation

1. Clone this repository or download the source code.

2. Navigate to the project directory.

3. Build the project:
   ```
   go build -o gemini
   ```

4. (Optional) Move the built executable to a directory in your system's PATH to use it as a global command. For example:
   ```
   mv gemini /usr/local/bin/
   ```

## Configuration

Before using the tool, you need to set up your Gemini AI API key. Open the `main.go` file and locate this line:

```go
apiKey := "YOUR_API_KEY_HERE"
```

Replace `"YOUR_API_KEY_HERE"` with your actual API key:

```go
apiKey := "your-actual-api-key-here"
```

**Note**: Make sure not to share your API key publicly if you plan to distribute this code.

## Usage

To use the tool, simply run it followed by your question or prompt in quotes:

```
gemini "What is the capital of France?"
```

or

```
gemini "Tell me a joke about programming"
```

If you've added it to your PATH, you can run it from anywhere in your terminal.

If you haven't added it to your PATH, you'll need to run it from the directory where the executable is located or provide the full path to the executable.

## Adding to PATH

Adding the Gemini AI Assistant to your system PATH allows you to run it from any directory in your terminal. Follow the instructions below for your operating system:

### Windows

1. Press `Win + X` and select "System".
2. Click on "Advanced system settings" on the right.
3. Click the "Environment Variables" button.
4. Under "System variables", find and select the "Path" variable, then click "Edit".
5. Click "New" and add the full path to the directory containing your `gemini.exe`.
   For example: `C:\Users\YourUsername\go\bin`
6. Click "OK" to close all dialogs.
7. Restart any open command prompts for the changes to take effect.

### macOS and Linux

1. Open your shell configuration file in a text editor. This is typically one of:
   - `~/.bash_profile`
   - `~/.bashrc`
   - `~/.zshrc` (for Zsh, the default shell on newer macOS versions)

2. Add the following line to the file, replacing `/path/to/gemini` with the actual path to your Gemini AI Assistant executable:

   ```bash
   export PATH=$PATH:/path/to/gemini
   ```

   For example, if you moved the executable to `/usr/local/bin`, you would add:
   
   ```bash
   export PATH=$PATH:/usr/local/bin
   ```

3. Save the file and close the text editor.

4. To apply the changes immediately, run:
   ```bash
   source ~/.bash_profile  # or ~/.bashrc or ~/.zshrc
   ```

   Alternatively, you can restart your terminal.

### Verifying Installation
```
gemini "Are you gonna take my job"
```
Now you can use the `gemini` command from any directory in your terminal!

## Contributing

Feel free to fork this project and submit pull requests with improvements or bug fixes.

## License

This project is open source and available under the [MIT License](LICENSE).