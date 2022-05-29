package version

import "strconv"

const (
	major = 1
	minor = 1
)

func HTTP() string {
	return "HTTP/" + strconv.Itoa(major) + "." + strconv.Itoa(minor)
}

func IsValid(httpVersion string) bool {
	return httpVersion == HTTP()
}
