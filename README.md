## globber

A small go utility for visualizing globs (and maybe regex file searches later) directly in the terminal.

![globber screenshot](https://raw.githubusercontent.com/jcquinlan/globber/master/screenshot.png)

![globber screenshot with warning](https://raw.githubusercontent.com/jcquinlan/globber/master/screenshot_warning.png)

## API

#### --glob
Default: `"*"`
Accepts any glob you'd like to test

#### --root
Default: `"."`
The starting location of the command (`glob`) will be relative to this location

#### --max-depth
Default: `20`
The maximum recursion level when traversing file structures.

#### --include-hidden-files
Default: `false`
Whether or not hidden directories and files should be considered in the visualization

#### --max-files-to-scan
Default: `200`
How many files to scan before stopping (useful to avoid scanning every file from root or something)

## Basic Usage
```bash
globber --glob=**/*.json
```
