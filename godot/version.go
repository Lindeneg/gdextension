package godot

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var versionRegex = regexp.MustCompile(`^([0-9]\.[0-9](\.[0-9])?\.[a-z]+\.[a-z]+\.[0-9 a-z]+$)`)

type version struct {
	Major  int
	MajorS string
	Minor  int
	MinorS string
	Patch  int
	PatchS string
	// raw version string
	Raw string
	// stable release
	Stable bool
	// official distribution
	Official bool
	// git commit hash
	Hash string
}

func (v version) String() string {
	if v.Patch == 0 {
		return fmt.Sprintf("%s.%s", v.MajorS, v.MinorS)
	}
	return fmt.Sprintf("%s.%s.%s", v.MajorS, v.MinorS, v.PatchS)
}

func (v version) MajorMinor() string {
	return fmt.Sprintf("%s.%s", v.MajorS, v.MinorS)
}

func newVersionFromString(s string) (version, error) {
	var (
		v   version
		err error
	)
	versionRaw := strings.TrimSpace(s)
	matches := versionRegex.FindStringSubmatch(versionRaw)
	if len(matches) < 2 {
		return v, errors.New("godot version not recognized")
	}
	version := matches[1]
	parts := strings.Split(version, ".")
	if len(parts) < 5 {
		return v, errors.New("godot version incomplete")
	}
	v.Raw = versionRaw
	if v.Major, err = strconv.Atoi(parts[0]); err != nil {
		return v, errors.New("failed to extract major version")
	}
	v.MajorS = parts[0]
	if v.Minor, err = strconv.Atoi(parts[1]); err != nil {
		return v, errors.New("failed to extract minor version")
	}
	v.MinorS = parts[1]
	offset := 0
	if len(parts) > 5 {
		if v.Patch, err = strconv.Atoi(parts[2]); err != nil {
			return v, errors.New("failed to extract patch version")
		}
		v.PatchS = parts[2]
		offset = 1
	} else {
		v.Patch = 0
		v.PatchS = "0"
	}
	v.Stable = parts[2+offset] == "stable"
	if !v.Stable {
		return v, fmt.Errorf("unstable godot version '%s' is not supported", v.Raw)
	}
	v.Official = parts[3+offset] == "official"
	if !v.Official {
		return v, fmt.Errorf("unofficial godot release '%s' is not supported", v.Raw)
	}
	v.Hash = parts[4+offset]
	if v.Major < MinMajorVersion || v.Major == MinMajorVersion && v.Minor < MinMinorVersion {
		return v, fmt.Errorf("godot version '%s' is not supported", v.Raw)
	}
	return v, nil
}
