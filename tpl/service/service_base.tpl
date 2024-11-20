package service

import (
    "{{.daoPackageName}}"
)

type Service struct {
    dao *dao.Dao
}

func Get() *Service {
    return &Service{}
}