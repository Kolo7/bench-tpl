package main

import "embed"

//go:embed etc/config.yaml tpl/*
var Content embed.FS
