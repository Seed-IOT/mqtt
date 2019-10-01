package main

import (
	"fmt"
	"mqtt/config"
	"mqtt/log"
	"mqtt/model"
	"mqtt/web"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"

	"os"
	"os/signal"
)

func main() {

	log.Init()

	cfg, err := config.New()
	if err != nil {
		log.GlobalLog.Fatalf("Failed to reading config file, %s\n", err)
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.GlobalLog.Fatalf("Failed to initialize model for operating all service, %s\n", err)
	}

	server := web.NewServer(cfg, service)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// 日志中间件
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.GlobalLog.Fatalf("Failed to listen for http server, %s\n", err)
		}
	}()

	// 启动mqtt客户端
	go func() {
		mqttAdaptor := mqtt.NewAdaptor("tcp://0.0.0.0:1883", "pinger")

		work := func() {
			mqttAdaptor.On("wifiDevice", func(msg mqtt.Message) {
				fmt.Println(msg)
			})
		}

		robot := gobot.NewRobot("mqttBot",
			[]gobot.Connection{mqttAdaptor},
			work,
		)
		robot.Start()
		defer robot.Stop()

	}()
	// 程序结束 关闭db 关闭mqtt
	defer service.DB.Close()

	log.GlobalLog.Println("mqtt is running")
	<-quit
	log.GlobalLog.Println("mqtt is stopped")
}
