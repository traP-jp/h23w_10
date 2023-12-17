package imggen

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// cleanerは指定されたディレクトリ内のファイルを定期的にクリーンアップします。
// maxAgeよりも古いファイルは削除されます。
// dir: クリーンアップ対象のディレクトリパス
// maxAge: 削除するファイルの最大経過時間
func cleaner(dir string, maxAge time.Duration) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		<-ticker.C

		files, err := os.ReadDir(dir)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, file := range files {
			fi, err := file.Info()
			if err != nil {
				log.Println(err)
				continue
			}
			if !fi.IsDir() && time.Since(fi.ModTime()) > maxAge {
				if err := os.Remove(filepath.Join(dir, file.Name())); err != nil {
					log.Println(err)
				}
			}
		}
	}
}
