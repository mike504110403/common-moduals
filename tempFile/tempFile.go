package tempfile

import (
	"os"
)

// 可以在執行檔被執行時，將被指定的檔案byte暫時寫入指定的磁碟空間，以便適用於純路徑取用的方法，例如os.Open

type Config struct {
	// 暫存檔存放的位址
	Path string
	// 暫存檔的前贅字
	Prefix string
}

type TempFileInfo struct {
	Name  string
	Path  string
	IsDel bool
}

var cfg = Config{
	Path:   "",
	Prefix: "tempFile",
}

func Init(initCfg Config) {
	cfg = initCfg
}

func TempFile(data []byte) (TempFileInfo, error) {
	tempFileInfo := TempFileInfo{}
	tmpFile, err := os.CreateTemp(cfg.Path, cfg.Prefix)
	if err != nil {
		return tempFileInfo, err
	}
	tempFileInfo.Name = tmpFile.Name()
	tempFileInfo.Path = cfg.Path + tmpFile.Name()

	if _, err := tmpFile.Write(data); err != nil {
		return tempFileInfo, err
	}
	if err := tmpFile.Close(); err != nil {
		return tempFileInfo, err
	}

	return tempFileInfo, nil
}

func (tempFileInfo TempFileInfo) Delete() error {
	return os.Remove(tempFileInfo.Path)
}
