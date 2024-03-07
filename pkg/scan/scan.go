package scan

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"time"

	"github.com/Jarover/crf/internal/config"
	"github.com/Jarover/crf/pkg/utils"
	log "github.com/sirupsen/logrus"
)

var lastModTime time.Time

// scanRules - ищем папки из конфига
func ScanRules() error {
	curFileTime := GetLastModTime()
	for _, v := range config.GetConfig().Rules {

		for _, p := range v.Paths {
			curTime, err := ScanFolders(p.InPath, p.OutPath)
			if err != nil {
				return err
			}
			if curFileTime.Before(curTime) {
				log.Info(lastModTime, curFileTime, curTime)
				curFileTime = curTime
			}
		}
	}
	SetLastModTime(curFileTime)
	return nil

}

// ScanFolders - сканируем папку
func ScanFolders(inpath, outpath string) (time.Time, error) {
	log.Info("Сканируем корневую папку - " + inpath)

	curFileTime := GetLastModTime()
	files, err := ioutil.ReadDir(inpath)
	if err != nil {

		return curFileTime, err
	}
	for _, file := range files {
		if file.IsDir() {

			curFileTime, err = scanFolder(inpath+file.Name(), file, outpath)
			if err != nil {
				return curFileTime, err
			}
			if curFileTime.Before(file.ModTime()) {
				log.Info(lastModTime, curFileTime, file.ModTime())
				curFileTime = file.ModTime()
			}
		}
	}

	return curFileTime, nil

}

// scanFolder - проверяем папку
func scanFolder(path string, folder fs.FileInfo, outpath string) (time.Time, error) {
	curFileTime := lastModTime
	if lastModTime.Before(folder.ModTime()) {
		log.Info(folder.Name(), " - ", folder.ModTime())

		files, err := os.ReadDir(path)
		if err != nil {

			return curFileTime, err
		}
		for _, entry := range files {
			if !entry.IsDir() {

				file, err := entry.Info()
				if err != nil {
					return curFileTime, err
				}
				if lastModTime.Before(file.ModTime()) {
					log.Info("From : ", path+"\\"+file.Name(), " To: ", outpath+file.Name())
					utils.Copy(path+"\\"+file.Name(), outpath+file.Name())
					//utils.CopyFile(file.Name(), path, outpath)
					if err != nil {
						return curFileTime, err
					}

					//  Check File exist

					if _, err := os.Stat(outpath + file.Name()); errors.Is(err, os.ErrNotExist) {
						// file does not exist
						return curFileTime, err
					}

					if curFileTime.Before(file.ModTime()) {
						log.Info(lastModTime, curFileTime, file.ModTime())
						curFileTime = file.ModTime()
					}

				}

			}
		}

	}
	return curFileTime, nil
}

func SetLastModTime(tm time.Time) {

	lastModTime = tm
}

func GetLastModTime() time.Time {

	return lastModTime
}

func InitLastModTime() {
	// Читем последний ID
	lastTime, err := config.ReadState(utils.GetBaseFile() + "_config.json")
	if err != nil {
		log.Info("Not Config file")
		lastModTime = time.Now()
	}

	lastModTime = lastTime

}
