package meta

import "time"

// FileMeta: 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt time.Time
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fMeta FileMeta) {
	fileMetas[fMeta.FileSha1] = fMeta
}

// GetFileMeta:通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// RemoveFileMeta：删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}

func GetAllFileMeta() []FileMeta {
	all := make([]FileMeta, 0)
	for _, file := range fileMetas {
		all = append(all, file)
	}
	return all
}
