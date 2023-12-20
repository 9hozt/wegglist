package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	Code    string `json:"code"`
	Regex   string `json:"regex"`
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

func runWeggli(weggliPath string, command string, path string, regex string) error {
	var cmdArgs []string

	// Si isRegex est true, ajoute l'option -R avec la regex au début de la commande
	if regex != "" {
		cmdArgs = append(cmdArgs, "-R", regex)
	}

	// Ajoute la commande et le chemin à la fin des arguments
	cmdArgs = append(cmdArgs, command, path)

	cmd := exec.Command(weggliPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func loadJSON(jsonFile string) ([]Theme, error) {
	jsonData, err := ioutil.ReadFile(jsonFile)
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
	var listThemesFlag bool
	var listDetailed bool

	weggliPath, err := getWeggliPath()
	if err != nil {
		fmt.Println("Error: weggli not found, check installation - https://github.com/weggli-rs/weggli")
		os.Exit(1)
	}

	flag.StringVar(&jsonFile, "json", "cmd.json", "Path to json rules file")
	flag.StringVar(&path, "path", ".", "Path to source code")
	flag.StringVar(&selectedThemes, "theme", "all", "Comma-separated list of themes to analyze. Use 'all' to analyze all themes.")
	flag.BoolVar(&listThemesFlag, "list", false, "List available themes and exit")
	flag.BoolVar(&listDetailed, "list-detailed", false, "List available themes with detailed information")
	flag.Parse()

	themes, err := loadJSON(jsonFile)
	if listThemesFlag || listDetailed {
		if err != nil {
			fmt.Println("Error while loading JSON file:", err)
			return
		}
		listThemes(themes, listDetailed)
		os.Exit(0)
	}

	if err != nil {
		fmt.Println("Error while loading JSON file:", err)
		return
	}

	selectedThemeSet := make(map[string]bool)
	if selectedThemes != "all" {
		selectedThemeList := strings.Split(selectedThemes, ",")
		for _, theme := range selectedThemeList {
			selectedThemeSet[theme] = true
		}
	}

	fmt.Println("Security code analysis:")
	for _, theme := range themes {
		if selectedThemes == "all" || selectedThemeSet[theme.Short] {
			fmt.Printf("%s:\n", theme.Name)
			for _, cmd := range theme.Commands {

				if err := runWeggli(weggliPath, cmd.Code, path, cmd.Regex); err != nil {
					fmt.Printf("Error while analyzing %s: %v\n", theme.Name, err)
					return
				}

			}
		}
	}
	fmt.Println("Done!")
	os.Exit(0)
}
