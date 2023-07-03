package main

import (
	c "Prisma/utils/config"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	c.Config()
	fmt.Println(c.Configuration.CLIENT_SECRET)
	fmt.Println(viper.Get("GOPATH"))
}
