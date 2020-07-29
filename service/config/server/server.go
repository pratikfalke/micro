package server

import (
	"github.com/micro/cli/v2"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/go-micro/v3/store"
	"github.com/micro/micro/v3/service"
	proto "github.com/micro/micro/v3/service/config/proto"
	mustore "github.com/micro/micro/v3/service/store"
)

const (
	name = "go.micro.config"
)

var (
	// Flags specific to the config service
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "watch_topic",
			EnvVars: []string{"MICRO_CONFIG_WATCH_TOPIC"},
			Usage:   "watch the change event.",
		},
	}
)

// Run micro config
func Run(c *cli.Context) error {
	if len(c.String("watch_topic")) > 0 {
		watchTopic = c.String("watch_topic")
	}

	srv := service.New(service.Name(name))
	mustore.DefaultStore.Init(store.Table("config"))

	proto.RegisterConfigHandler(srv.Server(), new(config))
	service.RegisterSubscriber(watchTopic, new(watcher))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	return nil
}