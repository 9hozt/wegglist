# wegglist

## Introduction

This is a simple weggli wrapper written un Go. The idea is to centralize and sort you patterns

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

## How to contribute

Do not hesitate to fork and add you themes and patterns to the cmd.json :) PR are welcome 
You can also fork and do WTF you want (see [Licence](LICENSE.md))

## TODO

 - Add --uniq option
 - Add more patterns :)
 - Add C++ support
 - Add --extensions option
 - Add more test code