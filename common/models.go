package common

// Source is the arguments coming in
type Source struct {
	Repository string            `json:"repository"`
	SecretName string            `json:"secret_name"`
	Options    map[string]string `json:"options"`
}

// Version is version
type Version struct {
	VersionID string `json:"version_id"`
}

type Request struct {
	Source  `json:"source"`
	Version `json:"version"`
}
