package application

import (
	"github.com/wailsapp/wails/v3/pkg/logger"
)

type Options struct {
	Name        string
	Description string
	Icon        []byte
	Mac         MacOptions
	Bind        []any
	Logger      struct {
		Silent        bool
		CustomLoggers []logger.Output
	}
	// InjectRuntime will inject the JS runtime into all created webview windows.
	InjectRuntime bool
}
