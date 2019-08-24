package cmd

const (
	appName  = "scp-cmd"
	appDescr = "Splunk Cloud Platform command line"
)

// version infomation block set via ldflags
var (
	semVersion string
	commitHash string
	buildDate  string
)
