package ezservice

type args struct {
	Version          bool     `short:"V" long:"version" description:"Print the version and exit."`
	ServiceInstall   struct{} `command:"install-service" description:"Installs service"`
	ServiceUninstall struct{} `command:"uninstall-service" description:"Uninstalls service"`
	Run              struct{} `command:"run" description:"Execute in-place (no service)"`
	ServiceStart     struct{} `command:"start" description:"Start service"`
	ServiceStop      struct{} `command:"stop" description:"Stop service"`
	ServiceStatus    struct{} `command:"status" description:"Get service status"`
}
