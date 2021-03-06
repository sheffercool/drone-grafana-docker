package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"

	docker "github.com/grafana/drone-grafana-docker"
)

const version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "docker plugin"
	app.Usage = "docker plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "edition",
			Usage:  "Grafana edition",
			EnvVar: "PLUGIN_EDITION",
		},
		cli.BoolFlag{
			Name:   "ubuntu",
			Usage:  "Build Ubuntu variant?",
			EnvVar: "PLUGIN_UBUNTU",
		},
		cli.StringFlag{
			Name:   "archs",
			Usage:  "Architectures to build (default: amd64,armv7,arm64)",
			EnvVar: "PLUGIN_ARCHS",
		},
		cli.StringFlag{
			Name:   "directory",
			Usage:  "Specify directory to build in",
			EnvVar: "PLUGIN_DIRECTORY",
		},
		cli.BoolFlag{
			Name:   "dry-run",
			Usage:  "Don't publish images",
			EnvVar: "PLUGIN_DRY_RUN",
		},
		cli.StringFlag{
			Name:   "remote.url",
			Usage:  "git remote url",
			EnvVar: "DRONE_REMOTE_URL",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.BoolFlag{
			Name:   "squash",
			Usage:  "squash the layers at build time",
			EnvVar: "PLUGIN_SQUASH",
		},
		cli.StringFlag{
			Name:   "docker.username",
			Usage:  "docker username",
			EnvVar: "PLUGIN_USERNAME,DOCKER_USERNAME",
		},
		cli.StringFlag{
			Name:   "docker.password",
			Usage:  "docker password",
			EnvVar: "PLUGIN_PASSWORD,DOCKER_PASSWORD",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msgf("Unexpected error")
	}
}

func run(c *cli.Context) error {
	plugin := docker.Plugin{
		Dryrun:  c.Bool("dry-run"),
		Cleanup: c.BoolT("docker.purge"),
		Login: docker.Login{
			Username: c.String("docker.username"),
			Password: c.String("docker.password"),
		},
		Build: docker.Build{
			Edition:   c.String("edition"),
			Archs:     c.String("archs"),
			Ubuntu:    c.Bool("ubuntu"),
			Directory: c.String("directory"),
			Name:      c.String("commit.sha"),
			Squash:    c.Bool("squash"),
		},
	}

	return plugin.Exec()
}
