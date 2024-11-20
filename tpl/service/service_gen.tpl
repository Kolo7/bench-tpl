package service

import (
    "context"
    "errors"

    "{{.apiPackageName}}"
    "{{.modelPackageName}}"
    "{{.daoPackageName}}"
    "github.com/jinzhu/gorm"
    "git.imgo.tv/ft/go-lib2/xlog"
    "git.imgo.tv/ft/go-lib2/ecode"
)

{{template "service_find_one.tpl" .}}

{{template "service_find_all.tpl" .}}

{{template "service_insert.tpl" .}}

{{template "service_update.tpl" .}}

{{template "service_delete.tpl" .}}