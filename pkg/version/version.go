package version

var (
	BuildTime = "unset"
	Commit    = "unset"
	Release   = "unset"
)

type Info struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Release   string `json:"release"`
}
