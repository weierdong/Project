package main

import (
	"seckill/engine"
	"seckill/mysql"
)

func main() {
	mysql.InitDB()

	engine := engine.EngineStart()
	engine.Run(":80")

}
