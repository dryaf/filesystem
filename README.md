# filesystem

## What is it about?

This is a sample for embedding files for serving via http
speaking of public files and templates

## Usage

### 1. Copy the file 

next to your main.go

### 2. Sample Usage:
```go
    // (embeded) public files and templates
	fileSystem := getFileSystem("files", in_dev_mode)            // filesystem.go
	templates := templates.New(fileSystem, "./templates", nil)   // github.com/dryaf/templates
	templates.AlwaysReloadAndParseTemplates = in_dev_mode
	templates.MustParseTemplates()
	e.Renderer = echo_wrappers.Renderer(templates)  // github.com/dryaf/templates/echo_wrappers
```

## Todos
[ ] Write Tests