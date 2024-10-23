package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "31"
	gitVersion   = "v1.31.2-k3s1"
	gitCommit    = "de7e66ca723e786caa1f01f769dbc5b70afb766a"
	gitTreeState = "clean"
	buildDate = "2024-10-23T15:29:46Z"
)
