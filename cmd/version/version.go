package version

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

var CurrentVersion = Version{1, 0, 1}

func ParseVersion(version string) Version {
	vreg := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	vstr := vreg.FindString(version)
	vstrs := strings.Split(vstr, ".")
	if len(vstrs) != 3 {
		return Version{0, 0, 0}
	}
	major, _ := strconv.Atoi(vstrs[0])
	minor, _ := strconv.Atoi(vstrs[0])
	patch, _ := strconv.Atoi(vstrs[0])
	return Version{major, minor, patch}
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) Cmp(o Version) int {
	if v.Major > o.Major {
		return 1
	}
	if v.Major < o.Major {
		return -1
	}

	if v.Minor > o.Minor {
		return 1
	}
	if v.Minor < o.Minor {
		return -1
	}

	if v.Patch > o.Patch {
		return 1
	}
	if v.Patch < o.Patch {
		return -1
	}

	return 0
}
