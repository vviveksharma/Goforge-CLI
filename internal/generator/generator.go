package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

type ProjectConfig struct {
	ProjectName string
	ProjectPath string
	ModulePath  string
}

type Generator struct {
	config ProjectConfig
}

func NewGenerator(config ProjectConfig) *Generator {
	return &Generator{config: config}
}

func (g *Generator) Generate() error {
	// Create directory structure
	if err := g.createDirectories(); err != nil {
		return err
	}

	// Generate files from templates
	if err := g.generateFiles(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) createDirectories() error {
	dirs := []string{
		"cmd/api",
		"internal/config",
		"internal/handler",
		"internal/middleware",
		"internal/server",
		"pkg/logger",
		"pkg/database",
		"pkg/cache",
		"deployments",
		"scripts",
	}

	for _, dir := range dirs {
		path := filepath.Join(g.config.ProjectPath, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

func (g *Generator) generateFiles() error {
	files := map[string]string{
		"templates/go.mod.tmpl":                          "go.mod",
		"templates/README.md.tmpl":                       "README.md",
		"templates/Makefile.tmpl":                        "Makefile",
		"templates/.env.example.tmpl":                    ".env.example",
		"templates/.gitignore.tmpl":                      ".gitignore",
		"templates/docker-compose.yml.tmpl":              "docker-compose.yml",
		"templates/cmd/api/main.go.tmpl":                 "cmd/api/main.go",
		"templates/internal/config/config.go.tmpl":       "internal/config/config.go",
		"templates/internal/handler/health.go.tmpl":      "internal/handler/health.go",
		"templates/internal/middleware/security.go.tmpl": "internal/middleware/security.go",
		"templates/internal/middleware/logger.go.tmpl":   "internal/middleware/logger.go",
		"templates/internal/middleware/recovery.go.tmpl": "internal/middleware/recovery.go",
		"templates/internal/server/server.go.tmpl":       "internal/server/server.go",
		"templates/pkg/logger/logger.go.tmpl":            "pkg/logger/logger.go",
		"templates/pkg/database/postgres.go.tmpl":        "pkg/database/postgres.go",
		"templates/pkg/cache/redis.go.tmpl":              "pkg/cache/redis.go",
		"templates/deployments/Dockerfile.tmpl":          "deployments/Dockerfile",
		"templates/scripts/goswitch.tmpl":                "scripts/goswitch",
	}

	tmpl := template.New("project").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
	})

	for templatePath, outputPath := range files {
		if err := g.generateFile(tmpl, templatePath, outputPath); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generateFile(tmpl *template.Template, templatePath, outputPath string) error {
	// Read template
	content, err := templatesFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", templatePath, err)
	}

	// Parse template
	t, err := tmpl.Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	// Create output file
	outputFilePath := filepath.Join(g.config.ProjectPath, outputPath)
	file, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputPath, err)
	}
	defer file.Close()

	// Execute template
	if err := t.Execute(file, g.config); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	// Make shell scripts executable
	if strings.HasPrefix(outputPath, "scripts/") {
		if err := os.Chmod(outputFilePath, 0755); err != nil {
			return fmt.Errorf("failed to make script executable %s: %w", outputPath, err)
		}
	}

	return nil
}
