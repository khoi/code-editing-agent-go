package tools

import (
	"encoding/json"
	"os/exec"
	"strings"
)

// ExecuteShellInput defines the input structure for the execute_shell tool
type ExecuteShellInput struct {
	Command string `json:"command" jsonschema_description:"The shell command to execute."`
}

// Generate the JSON schema for the tool's input
var ExecuteShellInputSchema = GenerateSchema[ExecuteShellInput]()

// ExecuteShellDefinition defines the tool
var ExecuteShellDefinition = ToolDefinition{
	Name:        "execute_shell",
	Description: "Execute a shell command and return the output. Use this when you need to run commands in the system shell.",
	InputSchema: ExecuteShellInputSchema,
	Function:    ExecuteShell,
}

// ExecuteShell executes a shell command and returns its output
func ExecuteShell(input json.RawMessage) (string, error) {
	executeShellInput := ExecuteShellInput{}
	err := json.Unmarshal(input, &executeShellInput)
	if err != nil {
		return "", err
	}

	// Execute the command
	cmd := exec.Command("sh", "-c", executeShellInput.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output) + "\nError: " + err.Error(), err
	}

	return strings.TrimSpace(string(output)), nil
}
