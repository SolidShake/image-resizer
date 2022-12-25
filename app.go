package main

import (
	"context"
	"fmt"

	"github.com/SolidShake/image-resizer/internal/watermark"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.uber.org/zap"
)

// App struct
type App struct {
	ctx    context.Context
	logger *zap.SugaredLogger
	dir    string
}

// NewApp creates a new App application struct
func NewApp(logger *zap.SugaredLogger, dir string) *App {
	return &App{logger: logger, dir: dir}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	selection, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.jpg)",
				Pattern:     "*.jpg",
			},
		},
	})
	if err != nil {
		a.logger.Errorf("open multiple files dialog: %w", err)
		return ""
	}

	// debug
	a.logger.Debug(selection, err)

	//
	message := "Успех"
	result := watermark.AddWatermark(a.logger, a.dir, selection)
	if !result {
		message = "Не удалось обработать фото"
	}

	_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		// Title:   "Select File Success",
		Message: message,
	})

	return fmt.Sprintf("Hello 123 %s, It's show time!", "name")
}

// Greet returns a greeting for the given name
func (a *App) SavePath() string {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Folder",
	})

	fmt.Println(selection, err)

	if err != nil {
		return ""
	}

	fmt.Println("directory " + a.dir)

	a.dir = selection

	fmt.Println("new directory " + a.dir)

	return fmt.Sprintf("Hello 123 %s, It's show time!", "name")
}
