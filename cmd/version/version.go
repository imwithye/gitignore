package version

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imwithye/gitignore/cmd/platform"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

var CurrentVersion = Version{1, 0, 2}

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

type asset struct {
	Name                 string
	Browser_download_url string
}

type release struct {
	Tag_name string
	Assets   []asset
}

func parseVersion(version string) Version {
	vreg := regexp.MustCompile(`\d+\.\d+\.\d+`)
	vstr := vreg.FindString(version)
	vstrs := strings.Split(vstr, ".")
	if len(vstrs) != 3 {
		return Version{0, 0, 0}
	}
	major, _ := strconv.Atoi(vstrs[0])
	minor, _ := strconv.Atoi(vstrs[1])
	patch, _ := strconv.Atoi(vstrs[2])
	return Version{major, minor, patch}
}

func Latest() (Version, string, error) {
	url := "https://api.github.com/repos/imwithye/gitignore/releases/latest"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	r, err := client.Get(url)
	if err != nil {
		return Version{}, "", err
	}
	defer r.Body.Close()
	rel := release{}
	json.NewDecoder(r.Body).Decode(&rel)

	for _, a := range rel.Assets {
		if strings.Contains(a.Name, platform.Platform) {
			return parseVersion(rel.Tag_name), a.Browser_download_url, nil
		}
	}
	return Version{}, "", errors.New("latest release is not available")
}
