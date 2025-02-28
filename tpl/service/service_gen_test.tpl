package service

import (
	"context"
	"errors"
	"testing"

	"git.imgo.tv/ft/go-lib2/ecode"
	"github.com/Kolo7/bench-tpl/output/api"
	"github.com/Kolo7/bench-tpl/output/internal/dao"
	"github.com/Kolo7/bench-tpl/output/internal/model"
	. "github.com/bytedance/mockey"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestService_{{.upperTableName}}Get(t *testing.T) {
	mockRecord := &model.{{.upperTableName}}{Id: 1}

	PatchConvey("Test Success", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return mockRecord, nil
		}).Build()

		service := Service{d: &dao.Dao{}}
		resp, code := service.{{.upperTableName}}Get(context.Background(), api.{{.upperTableName}}Req{Id: 1})
		So(code, ShouldEqual, ecode.OK)
		So(resp.Id, ShouldEqual, mockRecord.Id)
	})

	PatchConvey("Test Record Not Found Error", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return nil, gorm.ErrRecordNotFound
		}).Build()

		service := Service{d: &dao.Dao{}}
		resp, code := service.{{.upperTableName}}Get(context.Background(), api.{{.upperTableName}}Req{Id: 1})
		So(code, ShouldEqual, ecode.RequestErr)
		So(resp, ShouldBeNil)
	})

	PatchConvey("Test Server Error", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return nil, errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		resp, code := service.{{.upperTableName}}Get(context.Background(), api.{{.upperTableName}}Req{Id: 1})
		So(code, ShouldEqual, ecode.ServerErr)
		So(resp, ShouldBeNil)
	})
}

func TestService_{{.upperTableName}}GetAll(t *testing.T) {
	mockRecords := []*model.{{.upperTableName}}{
		{Id: 1},
		{Id: 2},
	}
	mockResp := api.{{.upperTableName}}GetAllResp{
		List:  make([]*api.{{.upperTableName}}Resp, 0),
		Total: 2,
	}
	for _, record := range mockRecords {
		mockResp.List = append(mockResp.List, &api.{{.upperTableName}}Resp{
			Id:      record.Id,
		})
	}

	PatchConvey("Test Success", t, func() {
		Mock((*dao.Dao).GetAll{{.upperTableName}}).To(func(_ *dao.Dao, _ context.Context, pageNum int, pageSize int, order string) ([]*model.{{.upperTableName}}, int64, error) {
			return mockRecords, 2, nil
		}).Build()

		service := Service{d: &dao.Dao{}}
		resp, code := service.{{.upperTableName}}GetAll(context.Background(), api.{{.upperTableName}}GetAllReq{PageNum: 1, PageSize: 10, Order: "id"})
		So(code, ShouldEqual, ecode.OK)
		So(resp.Total, ShouldEqual, mockResp.Total)
		So(resp.List, ShouldResemble, mockResp.List)
	})

	PatchConvey("Test Server Error", t, func() {
		Mock((*dao.Dao).GetAll{{.upperTableName}}).To(func(_ *dao.Dao, _ context.Context, pageNum int, pageSize int, order string) ([]*model.{{.upperTableName}}, int64, error) {
			return nil, 0, errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		resp, code := service.{{.upperTableName}}GetAll(context.Background(), api.{{.upperTableName}}GetAllReq{PageNum: 1, PageSize: 10, Order: "id"})
		So(code, ShouldEqual, ecode.ServerErr)
		So(resp, ShouldBeNil)
	})
}

func TestService_{{.upperTableName}}Insert(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}Insert).To(func(_ *dao.Dao, _ context.Context, record *model.{{.upperTableName}}) error {
			return nil
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Insert(context.Background(), api.{{.upperTableName}}InsertReq{})
		So(code, ShouldEqual, ecode.OK)
	})

	PatchConvey("Test Server Error", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}Insert).To(func(_ *dao.Dao, _ context.Context, record *model.{{.upperTableName}}) error {
			return errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Insert(context.Background(), api.{{.upperTableName}}InsertReq{})
		So(code, ShouldEqual, ecode.ServerErr)
	})

	PatchConvey("Test ErrUniqueConflict", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}Insert).To(func(_ *dao.Dao, _ context.Context, record *model.{{.upperTableName}}) error {
			return dao.ErrUniqueConflict
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Insert(context.Background(), api.{{.upperTableName}}InsertReq{})
		So(code.Code(), ShouldEqual, ecode.RequestErr)
	})
}

func TestService_{{.upperTableName}}Update(t *testing.T) {
	mockreq := api.{{.upperTableName}}UpdateReq{
	{{- range $column := .tableColumns}}{{if not (inExcludedFields $column.Upper)}}{{$column.Upper}}: new({{$column.GoType}}),{{end}}
	{{end}}
	}

	PatchConvey("Test Success", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return &model.{{.upperTableName}}{Id: 1}, nil
		}).Build()
		Mock((*dao.Dao).{{.upperTableName}}Update).To(func(_ *dao.Dao, _ context.Context, record *model.{{.upperTableName}}) error {
			return nil
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Update(context.Background(), mockreq)
		So(code, ShouldEqual, ecode.OK)
	})

	PatchConvey("Test ErrUniqueConflict", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return &model.{{.upperTableName}}{Id: 1}, nil
		}).Build()
		Mock((*dao.Dao).{{.upperTableName}}Update).To(func(_ *dao.Dao, _ context.Context, record *model.{{.upperTableName}}) error {
			return dao.ErrUniqueConflict
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Update(context.Background(), mockreq)
		So(code.Code(), ShouldEqual, ecode.RequestErr)
	})

	PatchConvey("Test Server Error1", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return nil, errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Update(context.Background(), mockreq)
		So(code, ShouldEqual, ecode.ServerErr)
	})
	PatchConvey("Test Server Error2", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return &model.{{.upperTableName}}{Id: 1}, nil
		}).Build()
		Mock((*dao.Dao).{{.upperTableName}}Update).To(func(_ *dao.Dao, _ context.Context, record *model.{{.upperTableName}}) error {
			return errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Update(context.Background(), mockreq)
		So(code, ShouldEqual, ecode.ServerErr)
	})

	PatchConvey("Test Record Not Found", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return nil, dao.ErrNotFound
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Update(context.Background(), mockreq)
		So(code, ShouldEqual, ecode.RequestErr)
	})
}

func TestService_{{.upperTableName}}Delete(t *testing.T) {
	PatchConvey("Test Success", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return &model.{{.upperTableName}}{Id: 1}, nil
		}).Build()
		Mock((*dao.Dao).{{.upperTableName}}Delete).To(func(_ *dao.Dao, _ context.Context, id int64) error {
			return nil
		}).Build()

		service := Service{d: &dao.Dao{}}
		code := service.{{.upperTableName}}Delete(context.Background(), api.{{.upperTableName}}DeleteReq{Id: 1})
		So(code, ShouldEqual, ecode.OK)
	})

	PatchConvey("Test Record Not Found", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return nil, dao.ErrNotFound
		}).Build()

		service := Service{d: &dao.Dao{}}
		err := service.{{.upperTableName}}Delete(context.Background(), api.{{.upperTableName}}DeleteReq{Id: 1})
		So(err, ShouldEqual, ecode.RequestErr)
	})

	PatchConvey("Test Server Error1", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return nil, errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		err := service.{{.upperTableName}}Delete(context.Background(), api.{{.upperTableName}}DeleteReq{Id: 1})
		So(err, ShouldEqual, ecode.ServerErr)
	})

	PatchConvey("Test Server Error2", t, func() {
		Mock((*dao.Dao).{{.upperTableName}}FindOne).To(func(_ *dao.Dao, _ context.Context, id int64) (*model.{{.upperTableName}}, error) {
			return &model.{{.upperTableName}}{Id: 1}, nil
		}).Build()
		Mock((*dao.Dao).{{.upperTableName}}Delete).To(func(_ *dao.Dao, _ context.Context, id int64) error {
			return errors.New("测试错误")
		}).Build()

		service := Service{d: &dao.Dao{}}
		err := service.{{.upperTableName}}Delete(context.Background(), api.{{.upperTableName}}DeleteReq{Id: 1})
		So(err, ShouldEqual, ecode.ServerErr)
	})
} 