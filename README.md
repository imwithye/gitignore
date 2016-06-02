Git-ignore
===
Are you tired of managing your git ignore files? You probably work in this way

```Bash
mkdir myproject && cd myproject
cat /path-to-ignore-template >> .gitignore
git init
```

Or even worse, you edit your git ignore file every time. You may use `git init --template` a lot. But here is a more pretty solution! Try `git ignore`!

Git-ignore is a git plugin which allows you manage your git ignore files. Add ignore file to current project by using

```Bash
git ignore xcode python c > .gitignore
```

## Install
Download Git-ignore from release page and put it under your system PATH.
