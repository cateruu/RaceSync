package fileService

import (
	"RaceSync/pkg/icon"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	dirPath  = "C:\\Users\\pawel\\Documents\\RaceSync"
	filePath = "C:\\Users\\pawel\\Documents\\RaceSync\\data.json"
)

type Data struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Icon string `json:"icon"`
}

type FileService struct {
	ctx context.Context
}

func New() *FileService {
	return &FileService{}
}

func (s *FileService) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *FileService) OpenFile() (*map[string]Data, error) {
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
		return nil, err
	}

	data, err := s.saveAppToFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *FileService) GetAppsData() (*map[string]Data, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file")
	}

	dataJson := make(map[string]Data)
	err = json.Unmarshal(data, &dataJson)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal data file")
	}

	return &dataJson, nil
}

func (s *FileService) LoadImage(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("unable to read image")
	}

	base64Encoded := base64.StdEncoding.EncodeToString(data)
	ext := filepath.Ext(path)
	if ext == "" {
		return "", fmt.Errorf("unable to extract extension from the file")
	}

	dataUrl := fmt.Sprintf("data:image/%s;base64,%s", ext[1:], base64Encoded)
	return dataUrl, nil
}

func (s *FileService) RemoveApp(name string) (*map[string]Data, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read data file")
	}

	var jsonData map[string]Data

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal data")
	}

	delete(jsonData, name)

	json, err := json.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal data")
	}

	saveFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer saveFile.Close()

	_, err = saveFile.Write(json)
	if err != nil {
		return nil, fmt.Errorf("unable to save to data file")
	}

	//remove png file
	err = os.Remove(dirPath + "\\" + name + ".png")
	if err != nil {
		return nil, fmt.Errorf("unable to remove app icon")
	}

	return &jsonData, nil
}

func (s *FileService) LaunchApps() error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("unable to read data file")
	}

	var savedData map[string]Data

	err = json.Unmarshal(data, &savedData)
	if err != nil {
		return fmt.Errorf("unable to unmarshal saved data")
	}

	for _, v := range savedData {
		cmd := exec.Command(v.Path)
		err := cmd.Start()
		if err != nil {
			return fmt.Errorf("error during app lanuch: %v", v.Name)
		}
	}

	return nil
}

func (s *FileService) saveAppToFile(file string) (*map[string]Data, error) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	if file == "" {
		return nil, fmt.Errorf("file you selected is not a real .exe file, find real .exe file then add it")
	}

	fileName := filepath.Base(file)
	appName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	iconPath := filepath.Join(dirPath, appName+".png")

	newApp := Data{
		Path: file,
		Name: appName,
		Icon: iconPath,
	}

	savedData := make(map[string]Data)

	if _, err := os.Stat(filePath); err == nil {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read saved file: %v", err)
		}

		err = json.Unmarshal(content, &savedData)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal saved data: %v", err)
		}
	}

	savedData[appName] = newApp

	iconBytes, err := icon.GetIconFromFile(file, true)
	if err != nil {
		width := 200
		height := 200

		scene := icon.NewScene(width, height)
		scene.PixelDraw(func(x, y int) color.RGBA {
			return color.RGBA{
				uint8(x * 255 / width),
				uint8(y * 255 / height),
				100,
				255,
			}
		})

		err = icon.SaveAsPNG(iconPath, scene.Image)
		if err != nil {
			return nil, fmt.Errorf("unable to save a file")
		}
	} else {
		image, err := icon.DecodeBytesToImage(iconBytes)
		if err != nil {
			return nil, fmt.Errorf("unable to convert an image")
		}

		err = icon.SaveAsPNG(iconPath, image)
		if err != nil {
			return nil, fmt.Errorf("unable to save a file")
		}
	}

	jsonData, err := json.Marshal(savedData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data to JSON: %v", err)
	}

	saveFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer saveFile.Close()

	_, err = saveFile.Write(jsonData)
	return &savedData, err
}
