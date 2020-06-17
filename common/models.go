package common

// Source is the arguments coming in
type Source struct {
	Repository string            `json:"repository"`
	Options    map[string]string `json:"options"`
	Tags       []string          `json:"tags"`
	Host       string            `json:"host"`
}

// Version is version
type Version struct {
	VersionID string `json:"version_id"`
}

type Request struct {
	Source  `json:"source"`
	Version `json:"version"`
}
