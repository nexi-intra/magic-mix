package main

import (
	"runtime/debug"
	"strings"

	"github.com/magicbutton/magic-mix/magicapp"
	"github.com/magicbutton/magic-mix/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: magic-mix
description: Describe the main purpose of this kitchen
---

# magic-mix
`
	utils.Setup(".env")
	magicapp.RegisterServeCmd("magic-mix", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.RegisterServiceCmd()
	magicapp.RegisterConvertCmd()
	magicapp.RegisterDownloadCmd()
	magicapp.RegisterUploadCmd()
	magicapp.RegisterCombineCmd()
	magicapp.RegisterMoveCmd()
	magicapp.RegisterCopyCmd()
	magicapp.RegisterSQLCmd()
	magicapp.RegisterParserCmd()
	magicapp.RegisterFlowCmd()
	utils.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "magic-mix", "")
}
