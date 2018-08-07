package main

import (
	"context"
	"server/audit-srv/db"
	"server/audit-srv/handler"
	audit_proto "server/audit-srv/proto/audit"
	"server/common"
	pubsub_proto "server/static-srv/proto/pubsub"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-os/config"
	"github.com/micro/go-os/config/source/file"
	"github.com/micro/go-os/metrics"
	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/metrics/telegraf"
	_ "github.com/micro/go-plugins/transport/nats"
	log "github.com/sirupsen/logrus"
)

func main() {
	configFile, _ := common.PopParameter("config")
	// Create a config instance
	conf := config.NewConfig(
		// poll every hour
		config.PollInterval(time.Hour),
		// use file as a config source
		config.WithSource(file.NewSource(config.SourceName(configFile))),
	)

	defer conf.Close()

	// create new metrics
	m := telegraf.NewMetrics(
		metrics.Namespace(conf.Get("service", "name").String("audit")),
		metrics.Collectors(
			// telegraf/statsd address
			common.MetricAddress(),
		),
	)
	defer m.Close()

	common.AuditSrv = conf.Get("service", "audit").String("go.micro.srv.audit")
	version := conf.Get("service", "version").String("latest")
	descr := conf.Get("service", "description").String("Micro service")
	service := micro.NewService(
		micro.Name(common.AuditSrv),
		micro.Version(version),
		micro.Metadata(map[string]string{"Description": descr}),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*10),
		micro.Flags(
			cli.BoolFlag{
				Name: "debug",
			},
		),
	)
	var debugEnabled bool
	service.Init(
		micro.Action(func(c *cli.Context) {
			debugEnabled = c.Bool("debug")
		}),
	)
	if debugEnabled {
		log.SetLevel(log.DebugLevel)
	}

	brker := service.Client().Options().Broker
	brker.Connect()
	auditService := &handler.AuditService{
		Broker: brker,
	}
	audit_proto.RegisterAuditServiceHandler(service.Server(), auditService)
	db.Init(service.Client())

	// subscribe in the first
	if err := auditService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.AUDIT_ACTION}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
