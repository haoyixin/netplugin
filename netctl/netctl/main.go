package main

import (
	"os"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/haoyixin/netplugin/netctl"
	"github.com/haoyixin/netplugin/version"
)

type byName []cli.Command

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	app := cli.NewApp()
	app.Flags = netctl.NetmasterFlags
	app.Version = "\n" + version.String()
	// TODO: use sort.Slice() in go1.8
	sort.Sort(byName(netctl.Commands))
	app.Commands = netctl.Commands
	app.Run(os.Args)
}
