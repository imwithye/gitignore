package emacs

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# -*- mode: gitignore; -*-",
		"*~",
		"\\#*\\#",
		"/.emacs.desktop",
		"/.emacs.desktop.lock",
		"*.elc",
		"auto-save-list",
		"tramp",
		".\\#*",
		"",
		"# Org-mode",
		".org-id-locations",
		"*_archive",
		"",
		"# flymake-mode",
		"*_flymake.*",
		"",
		"# eshell files",
		"/eshell/history",
		"/eshell/lastdir",
		"",
		"# elpa packages",
		"/elpa/",
		"",
		"# reftex files",
		"*.rel",
		"",
		"# AUCTeX auto folder",
		"/auto/",
		"",
		"# cask packages",
		".cask/",
		"dist/",
		"",
		"# Flycheck",
		"flycheck_*.el",
		"",
		"# server auth directory",
		"/server/",
		"",
		"# projectiles files",
		".projectile",
	}
	template.Add("github/global/emacs", strings.Join(ignore, "\n"))
}
