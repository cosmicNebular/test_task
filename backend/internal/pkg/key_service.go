package pkg

import (
	"log"
	"test/internal/pkg/database"
	"test/internal/pkg/database/entity"
)

type KeyService struct {
	dao *database.Dao
}

func CreateNewKeyService(dao *database.Dao) *KeyService {
	return &KeyService{dao: dao}
}

func (ks *KeyService) GetKey(id string) (*entity.Key, error) {
	key, err := ks.dao.GetKey(id)
	if err != nil {
		log.Fatalf("There is error during getting key %v with id %v", err, id)
		return nil, err
	}
	return key, nil
}

func (ks *KeyService) SaveKey(k entity.Key) error {
	err := ks.dao.SaveKey(k)
	if err != nil {
		log.Fatalf("There is error during saving key %v with id %v", err, k.Id)
		return err
	}
	return nil
}
