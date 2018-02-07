Git-ignore
===
[![Build Status](https://travis-ci.org/imwithye/gitignore.svg?branch=master)](https://travis-ci.org/imwithye/gitignore)

Are you tired of managing your git ignore files? You probably work in this way

```Bash
mkdir myproject && cd myproject
cat /path-to-ignore-template >> .gitignore
git init
```

Or even worse, you edit your git ignore file every time. You may use `git init --template` a lot. But here is a more pretty solution! Try `git ignore`!

Git-ignore is a git plugin which allows you manage your git ignore files. Add ignore file to current project by using

```Bash
git ignore add python c java c++ objective-c
```

Ignore files relate to Python C and Java will be added into `.gitignore` file. Git-ignore use [GitHub gitignore template](http://github.com/github/gitignore) as submodule. There are lots of language ignore files in this repository. You don't  have to create your own template. Just start using it!

## Install

Simply download the binary file from [release](https://github.com/imwithye/gitignore/releases), rename it to `git-ignore` and add it to your system path. You are ready to go.

macOS user can install `git-ignore` via `Homebrew`
```bash
brew install imwithye/formula/git-ignore
```

## Uninstall

Simply delete the binary from your system.

## Usage

Try `git ignore -h`. It will out put some useful information.

### Add ignore

```Bash
git ignore add python java c
```

This command will add multi git gitignore templates to `.gitignore`.

### See which file will be added

```Bash
git ignore which python java
```

### See what will be added

```Bash
git ignore show python java
```
