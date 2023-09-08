package main

import (
	"github.com/anthdm/weavebox"
)

func main(){
	app := weavebox.New()

	app.Get("/product", func(ctx *weavebox.Context) error {return nil})

	app.Serve(3001)
}