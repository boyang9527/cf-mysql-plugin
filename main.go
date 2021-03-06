package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/andreasf/cf-mysql-plugin/cfmysql"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "This executable is a cf plugin. "+
			"Run `cf install-plugin %s` to install it\nand `cf mysql service-name` "+
			"to use it.\n",
			os.Args[0])
		os.Exit(1)
	}

	mysqlPlugin := newPlugin()
	plugin.Start(mysqlPlugin)

	os.Exit(mysqlPlugin.GetExitCode())
}

func newPlugin() *cfmysql.MysqlPlugin {
	httpClientFactory := cfmysql.NewHttpClientFactory()
	http := cfmysql.NewHttpWrapper(httpClientFactory)
	apiClient := cfmysql.NewApiClient(http)

	sshRunner := cfmysql.NewSshRunner()
	netWrapper := cfmysql.NewNetWrapper()
	waiter := cfmysql.NewPortWaiter(netWrapper)
	randWrapper := cfmysql.NewRandWrapper()
	cfService := cfmysql.NewCfService(apiClient, sshRunner, waiter, http, randWrapper)

	execWrapper := cfmysql.NewExecWrapper()
	runner := cfmysql.NewMysqlRunner(execWrapper)

	portFinder := cfmysql.NewPortFinder()

	return cfmysql.NewMysqlPlugin(cfmysql.PluginConf{
		In:          os.Stdin,
		Out:         os.Stdout,
		Err:         os.Stderr,
		CfService:   cfService,
		PortFinder:  portFinder,
		MysqlRunner: runner,
	})
}
