package main

import (
	"fmt"
	"os"

	"github.com/docker/go-plugins-helpers/volume"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Docker mfs volume plugin"
	app.Usage = "Run docker mfs plugin"
	app.Version = "latest"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "mfs-base",
			Value:  "/mnt/mfs",
			Usage:  "MooseFS base directory",
			EnvVar: "DOCKER_VOLUME_PLUGIN_MFS_BASE",
		},
		cli.StringFlag{
			Name:   "addr",
			Value:  "/run/docker/plugins/mfs.sock",
			Usage:  "Plugin unix socket file path",
			EnvVar: "DOCKER_VOLUME_PLUGIN_MFS_SOCK",
		},
		cli.IntFlag{
			Name:   "gid",
			Value:  0,
			Usage:  "Gid to serve unix socket file",
			EnvVar: "DOCKER_VOLUME_PLUGIN_GID",
		},
	}
	app.Action = func(c *cli.Context) error {
		d := newMFSDriver(c.String("mfs-base"))
		h := volume.NewHandler(d)

		err := h.ServeUnix(c.String("addr"), c.Int("gid"))
		if err != nil {
			fmt.Println("Error:", err)
		}
		return err
	}
	app.Run(os.Args)
}
