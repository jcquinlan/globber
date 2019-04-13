## globber

A small go utility for visualizing globs (and maybe regex file searches later) directly in the terminal.

![globber screenshot](https://raw.githubusercontent.com/jcquinlan/globber/master/screen_shot.png)

## API

#### -glob
Default: `"*"`
Accepts any glob you'd like to test

#### -root
Default: `"."`
The starting location of the command (`glob`) will be relative to this location

#### -maxDepth
Default: `20`
The maximum recursion level when traversing file structures.

#### -includeHiddenFiles
Default: `false`
Whether or not hidden directories and files should be considered in the visualization

## Basic Usage
```bash
globber -glob=**/*.json
```
