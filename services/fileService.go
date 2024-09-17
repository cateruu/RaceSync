package fileService

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileService struct {
	ctx context.Context
}

func New() *FileService {
	return &FileService{}
}

func (s *FileService) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *FileService) OpenFile() (string, error) {
	file, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		Title: "Select app",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Exe Files (*.exe)",
				Pattern:     "*.exe",
			},
		},
	})

	fmt.Println(file)

	return file, err
}
