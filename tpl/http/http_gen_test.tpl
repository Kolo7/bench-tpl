package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mghttp "git.imgo.tv/ft/go-ceres/pkg/net/http"
	"git.imgo.tv/ft/go-lib2/ecode"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/bytedance/mockey"
	"{{.apiPackageName}}"
	"{{.fqdn}}/internal/service"
)

func Test{{.upperTableName}}Get(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*service.Service).{{.upperTableName}}Get).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}Req) (*api.{{.upperTableName}}Resp, ecode.Codes) {
			return &api.{{.upperTableName}}Resp{}, ecode.OK
		}).Build()
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/{{.tableName}}/get", {{.upperTableName}}Get)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/{{.tableName}}/get?id=1", nil)
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.OK.Code())
	})

	PatchConvey("Test Error", t, func() {
		Mock((*service.Service).{{.upperTableName}}Get).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}Req) (*api.{{.upperTableName}}Resp, ecode.Codes) {
			return nil, ecode.RequestErr
		}).Build()
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/{{.tableName}}/get", {{.upperTableName}}Get)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/{{.tableName}}/get?id=1", nil)
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})

	PatchConvey("Test Validator Error", t, func() {
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/{{.tableName}}/get", {{.upperTableName}}Get)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/{{.tableName}}/get?id=", nil)
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})
}

func Test{{.upperTableName}}GetAll(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*service.Service).{{.upperTableName}}GetAll).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}GetAllReq) (*api.{{.upperTableName}}GetAllResp, ecode.Codes) {
			return &api.{{.upperTableName}}GetAllResp{}, ecode.OK
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/{{.tableName}}/get_all", {{.upperTableName}}GetAll)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/{{.tableName}}/get_all?page_num=1&page_size=10", nil)
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.OK.Code())
	})

	PatchConvey("Test Error", t, func() {
		Mock((*service.Service).{{.upperTableName}}GetAll).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}GetAllReq) (*api.{{.upperTableName}}GetAllResp, ecode.Codes) {
			return nil, ecode.RequestErr
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/{{.tableName}}/get_all", {{.upperTableName}}GetAll)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/{{.tableName}}/get_all?page_num=1&page_size=10", nil)
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})
} 

func Test{{.upperTableName}}Insert(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*service.Service).{{.upperTableName}}Insert).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}InsertReq) ecode.Codes {
			return ecode.OK
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/insert", {{.upperTableName}}Insert)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}InsertReq{}
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/insert", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.OK.Code())
	})

	PatchConvey("Test Error", t, func() {
		Mock((*service.Service).{{.upperTableName}}Insert).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}InsertReq) ecode.Codes {
			return ecode.RequestErr
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/insert", {{.upperTableName}}Insert)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}InsertReq{}
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/insert", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})
}

func Test{{.upperTableName}}Update(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*service.Service).{{.upperTableName}}Update).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}UpdateReq) ecode.Codes {
			return ecode.OK
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/update", {{.upperTableName}}Update)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}UpdateReq{Id: 1} // 假设 Id 是必需的
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/update", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.OK.Code())
	})

	PatchConvey("Test Validator Error", t, func() {
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/update", {{.upperTableName}}Update)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}UpdateReq{} // 假设缺少 Id
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/update", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})

	PatchConvey("Test Error", t, func() {
		Mock((*service.Service).{{.upperTableName}}Update).To(func(_ *service.Service, _ context.Context, _ api.{{.upperTableName}}UpdateReq) ecode.Codes {
			return ecode.RequestErr
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/update", {{.upperTableName}}Update)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}UpdateReq{}
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/update", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})
}

func Test{{.upperTableName}}Delete(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*service.Service).{{.upperTableName}}Delete).To(func(_ *service.Service, _ context.Context, req api.{{.upperTableName}}DeleteReq) error {
			return nil
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/delete", {{.upperTableName}}Delete)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}DeleteReq{Id: 1} // 假设 Id 是必需的
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/delete", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.OK.Code())
	})

	PatchConvey("Test Not Found Error", t, func() {
		Mock((*service.Service).{{.upperTableName}}Delete).To(func(_ *service.Service, _ context.Context, req api.{{.upperTableName}}DeleteReq) error {
			return ecode.RequestErr
		}).Build()

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/delete", {{.upperTableName}}Delete)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}DeleteReq{Id: 999} // 假设 Id 不存在
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/delete", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})

	PatchConvey("Test Validator Error", t, func() {
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/{{.tableName}}/delete", {{.upperTableName}}Delete)
		recorder := httptest.NewRecorder()
		req := api.{{.upperTableName}}DeleteReq{} // 假设缺少 Id
		requestBody, err := json.Marshal(req)
		So(err, ShouldBeNil)
		request, err := http.NewRequest("POST", "/{{.tableName}}/delete", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		So(err, ShouldBeNil)
		router.ServeHTTP(recorder, request)
		So(recorder.Code, ShouldEqual, http.StatusOK)
		jsonResp := mghttp.JSON{}
		err = json.Unmarshal(recorder.Body.Bytes(), &jsonResp)
		So(err, ShouldBeNil)
		So(jsonResp.Code, ShouldEqual, ecode.RequestErr.Code())
	})
}
