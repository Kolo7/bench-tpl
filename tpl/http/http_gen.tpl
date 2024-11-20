package http

import (
    "git.imgo.tv/ft/go-lib2/xlog"
    "github.com/gin-gonic/gin"
    "{{.apiPackageName}}"
    "{{.fqdn}}/internal/service"
    "git.imgo.tv/ft/go-ceres/pkg/net/http"
)

{{template "http_find_one.tpl" .}}

{{template "http_find_all.tpl" .}}

{{template "http_insert.tpl" .}}

{{template "http_update.tpl" .}}

{{template "http_delete.tpl" .}}
