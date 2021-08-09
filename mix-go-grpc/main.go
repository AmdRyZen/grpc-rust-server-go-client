package main

import (
	"mix-grpc/commands"
	_ "mix-grpc/configor"
	_ "mix-grpc/di"
	_ "mix-grpc/dotenv"

	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
