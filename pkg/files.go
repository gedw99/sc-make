package pkg

import (
	//"github.com/gedw99/gio-make"

	_ "github.com/gedw99/gio-make"
)

func GetAllFilenames() (out []string, err error) {

	/*
		rootAssets, err := debme.FS(assets, ".")
		if err != nil {
			return nil, err
		}

		return getFilenames(&rootAssets.FS, ".")
	*/

	files, _ := getFilenames(&assets, ".")
	if err != nil {
		return nil, err
	}

	return files, nil

}
