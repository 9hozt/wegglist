package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	Code    string `json:"code"`
	Regex   string `json:"regex"`
	Unique  bool   `json:"unique"`
	Comment string `json:"comment"`
}

type Theme struct {
	Name     string    `json:"name"`
	Short    string    `json:"short"`
	Commands []Command `json:"commands"`
}

func getWeggliPath() (string, error) {
	var whichCmd string
	if runtime.GOOS == "windows" {
		whichCmd = "where"
	} else {
		whichCmd = "which"
	}

	cmd := exec.Command(whichCmd, "weggli")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func runWeggli(out io.Writer, weggliPath string, command string, path string, regex string, unique bool) error {
	var cmdArgs []string

	if regex != "" {
		cmdArgs = append(cmdArgs, "-R", regex)
	}

	if unique {
		cmdArgs = append(cmdArgs, "--unique")
	}

	cmdArgs = append(cmdArgs, command, path)
	cmd := exec.Command(weggliPath, cmdArgs...)
	cmd.Stdout = out
	cmd.Stderr = out
	return cmd.Run()
}

func loadJSON(jsonFile string) ([]Theme, error) {
	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	var Themes []Theme
	if err := json.Unmarshal(jsonData, &Themes); err != nil {
		return nil, err
	}

	return Themes, nil
}

func listThemes(themes []Theme, detailed bool) {
	fmt.Println("Available themes:")
	for _, theme := range themes {
		fmt.Printf("- %s (%s)\n", theme.Name, theme.Short)
		if detailed {
			fmt.Println("  Commands:")
			for _, cmd := range theme.Commands {
				fmt.Printf("    - %s : %s\n", cmd.Code, cmd.Comment)
			}
		}
	}
}

func main() {
	var jsonFile string
	var path string
	var selectedThemes string
	var outputFile string
	var listThemesFlag bool
	var listDetailed bool

	flag.StringVar(&jsonFile, "json", "cmd.json", "Path to json rules file")
	flag.StringVar(&path, "path", ".", "Path to source code")
	flag.StringVar(&selectedThemes, "theme", "all", "Comma-separated list of themes to analyze. Use 'all' to analyze all themes.")
	flag.StringVar(&outputFile, "output", "", "Path to a file to also write results to, in addition to stdout")
	flag.BoolVar(&listThemesFlag, "list", false, "List available themes and exit")
	flag.BoolVar(&listDetailed, "list-detailed", false, "List available themes with detailed information")
	flag.Parse()

	themes, err := loadJSON(jsonFile)
	if listThemesFlag || listDetailed {
		if err != nil {
			fmt.Println("Error while loading JSON file:", err)
			os.Exit(1)
		}
		listThemes(themes, listDetailed)
		os.Exit(0)
	}

	if err != nil {
		fmt.Println("Error while loading JSON file:", err)
		os.Exit(1)
	}

	// Listing themes doesn't need weggli, so this check only runs once we know
	// an actual scan is requested.
	weggliPath, err := getWeggliPath()
	if err != nil {
		fmt.Println("Error: weggli not found, check installation - https://github.com/weggli-rs/weggli")
		os.Exit(1)
	}

	var out io.Writer = os.Stdout
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			os.Exit(1)
		}
		defer f.Close()
		out = io.MultiWriter(os.Stdout, f)
	}

	selectedThemeSet := make(map[string]bool)
	if selectedThemes != "all" {
		selectedThemeList := strings.Split(selectedThemes, ",")
		for _, theme := range selectedThemeList {
			selectedThemeSet[theme] = true
		}
	}

	hadError := false
	fmt.Fprintln(out, "Security code analysis:")
	for _, theme := range themes {
		if selectedThemes != "all" && !selectedThemeSet[theme.Short] {
			continue
		}
		fmt.Fprintf(out, "Current theme - %s:\n", theme.Name)
		for _, cmd := range theme.Commands {
			fmt.Fprintf(out, "%s\n", cmd.Comment)
			if err := runWeggli(out, weggliPath, cmd.Code, path, cmd.Regex, cmd.Unique); err != nil {
				fmt.Fprintf(out, "Error while running pattern %q for theme %s: %v\n", cmd.Code, theme.Name, err)
				hadError = true
			}
		}
	}
	fmt.Fprintln(out, "Done!")

	if hadError {
		os.Exit(1)
	}
}
