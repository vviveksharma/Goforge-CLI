package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/viveksharma/goforge/internal/generator"
)

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new Go project",
	Long:  `Create a new Go project with Fiber, PostgreSQL, Redis, and production-ready features.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func runCreate(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Security: Validate project name to prevent path traversal
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %w", err)
	}

	// Get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	projectPath := filepath.Join(currentDir, projectName)

	// Check if directory already exists
	if _, err := os.Stat(projectPath); err == nil {
		return fmt.Errorf("directory '%s' already exists", projectName)
	}

	// Create project directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	fmt.Printf("🚀 Creating project: %s\n", projectName)

	// Generate project
	config := generator.ProjectConfig{
		ProjectName: projectName,
		ProjectPath: projectPath,
		ModulePath:  fmt.Sprintf("github.com/yourusername/%s", projectName),
	}

	gen := generator.NewGenerator(config)
	if err := gen.Generate(); err != nil {
		// Clean up on failure
		os.RemoveAll(projectPath)
		return fmt.Errorf("failed to generate project: %w", err)
	}

	fmt.Println("\n✅ Project created successfully!")
	fmt.Printf("\n📁 Next steps:\n")
	fmt.Printf("   cd %s\n", projectName)
	fmt.Printf("   make up\n")
	fmt.Printf("   Visit http://localhost:8080/health/ready\n\n")

	return nil
}

// validateProjectName ensures project name is safe and valid
func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	// Prevent path traversal attacks
	if strings.Contains(name, "..") || strings.Contains(name, "/") || strings.Contains(name, "\\") {
		return fmt.Errorf("project name cannot contain path separators or '..'")
	}

	// Check for valid characters (alphanumeric, dash, underscore)
	validName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("project name can only contain letters, numbers, dashes, and underscores")
	}

	// Prevent reserved names
	reserved := []string{".", "..", "con", "prn", "aux", "nul", "com1", "com2", "com3", "com4", "lpt1", "lpt2"}
	for _, r := range reserved {
		if strings.EqualFold(name, r) {
			return fmt.Errorf("'%s' is a reserved name", name)
		}
	}

	return nil
}
