# wegglist

![demoe](demo.gif)

## Introduction

wegglist is a small Go wrapper around [weggli](https://github.com/weggli-rs/weggli) meant to speed up the first pass of a manual C code audit. It organizes a batch of semantic search patterns into themes (buffer overflows, heap issues, format strings, ...) and runs them all against a codebase in one command instead of invoking `weggli` pattern by pattern, printing out the lines it considers potentially vulnerable so you know where to start looking.

It is not a vulnerability scanner and does not confirm anything on its own: every hit is a pattern match, not a proven bug, and it will miss anything that doesn't fit its patterns. Treat the output as a shortlist to review by hand, not a verdict.

Most patterns are sourced from [0xdea/weggli-patterns](https://github.com/0xdea/weggli-patterns). The project was also inspired by [this article](https://dustri.org/b/playing-with-weggli.html).

## Requirements

- Go 1.21 or later
- [weggli](https://github.com/weggli-rs/weggli) installed and available on `PATH`

## Installation

```bash
go build .
```

## Usage

```bash
Usage of ./wegglist:
  -json string
        Path to json rules file (default "cmd.json")
  -list
        List available themes and exit
  -list-detailed
        List available themes with detailed information
  -output string
        Path to a file to also write results to, in addition to stdout
  -path string
        Path to source code (default ".")
  -theme string
        Comma-separated list of themes to analyze. Use 'all' to analyze all themes. (default "all")
```

Exit code is non-zero if at least one pattern failed to run (for example, due to invalid syntax), which makes it suitable for use in scripts or CI.

## Rule file format

```json
[
  {
    "name": "Theme Name",
    "short": "themeShortName",
    "commands": [
      {
        "code": "pattern",
        "regex": "optional regex filter, adds the weggli -R option",
        "unique": false,
        "comment": "pattern description"
      }
    ]
  }
]
```

See [cmd.json](cmd.json) for a complete example.

## Contributing

Contributions are welcome. Fork the repository, add themes and patterns to `cmd.json`, and open a pull request. See [LICENSE](LICENSE.md) for license terms.

## Roadmap

- Add more patterns
- Add C++ support
- Add a `--extensions` option
- Add more test coverage
