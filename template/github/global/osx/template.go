package osx

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.DS_Store",
		".AppleDouble",
		".LSOverride",
		"",
		"# Icon must end with two \\r",
		"Icon",
		"",
		"# Thumbnails",
		"._*",
		"",
		"# Files that might appear in the root of a volume",
		".DocumentRevisions-V100",
		".fseventsd",
		".Spotlight-V100",
		".TemporaryItems",
		".Trashes",
		".VolumeIcon.icns",
		".com.apple.timemachine.donotpresent",
		"",
		"# Directories potentially created on remote AFP share",
		".AppleDB",
		".AppleDesktop",
		"Network Trash Folder",
		"Temporary Items",
		".apdisk",
		"",
	}
	template.Add("github/global/osx", strings.Join(ignore, "\n"))
}
