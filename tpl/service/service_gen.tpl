package service

import (
    "context"
    "errors"

    "{{.apiPackageName}}"
    "github.com/jinzhu/gorm"
)

{{template "service_find_one.tpl" .}}