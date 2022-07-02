package pkg

import (
	"log"
	"test/internal/pkg/database"
	"test/internal/pkg/database/entity"
)

type FileService struct {
	dao *database.Dao
}

func CreateNewFileService(dao *database.Dao) *FileService {
	return &FileService{dao: dao}
}

func (fs *FileService) GetAllFiles() ([]entity.File, error) {
	return fs.dao.GetAllFiles()
}

func (fs *FileService) GetFile(id string) (*entity.File, error) {
	file, err := fs.dao.GetFile(id)
	if err != nil {
		log.Fatalf("There is error during getting file %v with id %v", err, id)
		return nil, err
	}
	return file, nil
}

func (fs *FileService) SaveFile(f entity.File) error {
	err := fs.dao.SaveFile(f)
	if err != nil {
		log.Fatalf("There is error during saving file %v with id %v", err, f.Id)
		return err
	}
	return nil
}
