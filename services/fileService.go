package fileService

import (
	"RaceSync/pkg/icon"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	if err != nil {
		return "", err
	}

	err = s.saveAppToFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return file, nil
}

func (s *FileService) saveAppToFile(file string) error {
	dirPath := "C:\\Users\\pawel\\Documents\\RaceSync"
	filePath := filepath.Join(dirPath, "data.json")

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	if file == "" {
		return fmt.Errorf("no file path specified")
	}

	fileName := filepath.Base(file)
	appName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	iconPath := filepath.Join(dirPath, appName+".png")
	fmt.Println(iconPath)

	newApp := map[string]interface{}{
		"path": file,
		"name": appName,
		"icon": iconPath,
	}

	savedData := make(map[string]interface{})

	if _, err := os.Stat(filePath); err == nil {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read saved file: %v", err)
		}

		err = json.Unmarshal(content, &savedData)
		if err != nil {
			return fmt.Errorf("failed to unmarshal saved data: %v", err)
		}
	}

	savedData[appName] = newApp

	iconBytes, err := icon.GetIconFromFile(file, true)
	if err != nil {
		return fmt.Errorf("unable to extract icon from the file")
	}

	image, err := icon.DecodeBytesToImage(iconBytes)
	if err != nil {
		return fmt.Errorf("unable to convert an image")
	}

	err = icon.SaveAsPNG(iconPath, image)
	if err != nil {
		return fmt.Errorf("unable to save a file")
	}

	jsonData, err := json.Marshal(savedData)
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %v", err)
	}

	saveFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open data file: %v", err)
	}
	defer saveFile.Close()

	_, err = saveFile.Write(jsonData)
	return err
}
