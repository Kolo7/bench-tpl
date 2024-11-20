package service

import (
    "{{.daoPackageName}}"
)

type Service struct {
    d *dao.Dao
}

func Get() *Service {
    return &Service{}
}