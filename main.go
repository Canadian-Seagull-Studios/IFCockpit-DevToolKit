/*

	This file is part of the IFCockpit Developer Toolkit
	from Canadian Seagull Studios Ltd.

	This is not open source and may only be used by developers
	who are part of the IFCockpit developer programme under the
	terms of the programme and may not be shared with anyone
	outside the terms of that programme or otherwise without the
	explicit written consent of Canadian Seagull Studios Ltd.

	Copyright 2024, Canadian Seagull Studios Ltd.
	All rights reserved.

*/

package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// User config file
	configPath, _ := os.UserConfigDir()
	configDir := filepath.Join(configPath, "IFCockpitDevToolkit")
	_ = os.Mkdir(configDir, 0700)
	configFile := filepath.Join(configDir, "config.json")

	configData, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
	}
	defer configData.Close()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "IFCockpit Development Toolkit",
		Width:  1600,
		Height: 1280,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
