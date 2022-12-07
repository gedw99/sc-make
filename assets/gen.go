package assets

import (
	"embed"
	"os"
	"path/filepath"
	"strings"

	"github.com/leaanthony/debme"
	"github.com/leaanthony/gosod"
)

var (
	//go:embed version/version.txt
	version string

	Version string = strings.TrimSpace(version)

	//go:embed all:make/* all:project/*
	assets embed.FS
)

type assetData struct {
	Name string
}

func GetVersion() string {
	return version
}

func GetAllFilenames() (out []string, err error) {

	files, _ := getFilenames(&assets, ".")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func getFilenames(fs *embed.FS, path string) (out []string, err error) {
	if len(path) == 0 {
		path = "."
	}
	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res, err := getFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}

func CreateProjectFiles(targetDir string, projectName string, git string, template string) error {

	// step 1: Extract the assets to the FS.
	templateDir := gosod.New(assets)
	err := templateDir.Extract(targetDir, &assetData{Name: projectName})
	if err != nil {
		return err
	}

	return err

}

/// Authours functions .... Let here for hinting

// Install will install all default project assets
func Install(targetDir string, projectName string) error {
	templateDir := gosod.New(assets)
	err := templateDir.Extract(targetDir, &assetData{Name: projectName})
	if err != nil {
		return err
	}

	// Rename the manifest file
	windowsDir := filepath.Join(targetDir, "build", "windows")
	manifest := filepath.Join(windowsDir, "wails.exe.manifest")
	targetFile := filepath.Join(windowsDir, projectName+".exe.manifest")
	err = os.Rename(manifest, targetFile)
	if err != nil {
		return err
	}

	return nil
}

func RegenerateManifest(target string) error {
	a, err := debme.FS(assets, "build")
	if err != nil {
		return err
	}
	return a.CopyFile("windows/wails.exe.manifest", target, 0644)
}

func RegenerateAppIcon(target string) error {
	a, err := debme.FS(assets, "build")
	if err != nil {
		return err
	}
	return a.CopyFile("appicon.png", target, 0644)
}

func RegeneratePlist(targetDir string, projectName string) error {
	darwinAssets, err := debme.FS(assets, "build/darwin")
	if err != nil {
		return err
	}
	templateDir := gosod.New(darwinAssets)
	err = templateDir.Extract(targetDir, &assetData{Name: projectName})
	if err != nil {
		return err
	}

	return nil
}
