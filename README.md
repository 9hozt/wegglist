# wegglist

## Introduction

This is a simple [weggli](https://github.com/weggli-rs/weggli) wrapper written un Go.  
The idea is to centralize and sort you patterns.
Most of the patterns comes from https://github.com/0xdea/weggli-patterns  
Also inspired by https://dustri.org/b/playing-with-weggli.html

## Usage

```bash
Usage of ./wegglist:
  -json string
        Path to json rules file (default "cmd.json")
  -list
        List available themes and exit
  -list-detailed
        List available themes with detailed information
  -path string
        Path to source code (default ".")
  -theme string
        Comma-separated list of themes to analyze. Use 'all' to analyze all themes. (default "all")
```

## Format

JSON is formatted this way

```json
[
  {
    "name": "Theme Name",
    "short": "themeShortName",
    "commands": [
      {
        "code": "pattern",
        "regex" : "regex use in pattern - this add -R option",
        "comment": "pattern description"
      }
    ]
  },
  {
    "name": "Theme Name",
    "short": "themeShortName",
    "commands": [
      {
        "code": "pattern",
        "regex" : "regex use in pattern - this add -R option",
        "comment": "pattern description"
      }
    ]
  }
]
```

See [base file](cmd.json)

## How to contribute

Do not hesitate to fork and add you themes and patterns to the cmd.json :) PR are welcome 
You can also fork and do WTF you want (see [Licence](LICENSE.md))

## TODO

 - Add --uniq option
 - Add more patterns :)
 - Add C++ support
 - Add --extensions option
 - Add more test code
 - Add an output to file option