package version

import "strconv"

const (
	major = 1
	minor = 1
)

func HTTPVersion() string {
	return "HTTP/" + strconv.Itoa(major) + "." + strconv.Itoa(minor)
}
