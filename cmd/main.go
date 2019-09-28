package main

import (
	"fmt"
	"mqtt/config"
	"mqtt/model"
	"mqtt/web"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"

	"log"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to reading config file, %s\n", err)
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize model for operating all service, %s\n", err)
	}

	server := web.NewServer(cfg, service)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen for http server, %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// 启动mqtt客户端
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

	// 放去子线程
	go func() {
		robot.Start()
	}()

	log.Println("mqtt is running")
	<-quit
	log.Println("mqtt is stopped")

	// 程序结束 关闭db 关闭mqtt
	defer service.DB.Close()
	defer robot.Stop()
}
