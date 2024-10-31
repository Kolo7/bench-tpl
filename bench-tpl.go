package main

import (
	"embed"

	"github.com/Kolo7/bench-tpl/cmd"
)

//go:embed tpl/*
var Content embed.FS

func main() {
	cmd.Execute(Content)
}
