package server_test

import (
	"command-service/internal/application"
	"command-service/internal/presentation/adapter"
	"command-service/internal/presentation/server"
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/watarui-go-cqrs/pb/pb"
	"go.uber.org/fx"
)

var _ = Describe("categoryServer構造体", Ordered, Label("メソッドのテスト"), func() {
	var srv pb.CategoryCommandServer
	var category *pb.Category
	var ctx context.Context
	var container *fx.App
	BeforeAll(func() {
		ctx = context.Background()
		container = fx.New(
			application.SrvDepend,
			fx.Provide(
				adapter.NewCategoryAdapterImpl,
				server.NewCategoryServer,
			),
			fx.Populate(&srv),
		)
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	AfterAll(func() {
		err := container.Stop(context.Background())
		Expect(err).NotTo(HaveOccurred())
	})
	Context("Add()メソッドのテスト", Label("Add"), func() {
		It("カテゴリ登録が成功し、CategoryUpResultが返る", func() {
			param := pb.CategoryUpParam{Crud: pb.CRUD_INSERT, Id: "", Name: "飲料水_テスト1"}
			result, _ := srv.Create(ctx, &param)
			category = result.Category
			Expect(result.Error).To(BeNil())
		})
		It("カテゴリ登録が失敗し、bp.Errorを保持したCategoryUpResultが返る", func() {
			param := pb.CategoryUpParam{Crud: pb.CRUD_INSERT, Id: category.GetId(), Name: category.GetName()}
			result, _ := srv.Create(ctx, &param)
			e := pb.Error{Type: "CRUD Error", Message: "category 飲料水_テスト1 already exists"}
			Expect(result.Error).To(Equal(&e))
		})
	})
	Context("Update()メソッドのテスト", Label("Update"), func() {
		It("カテゴリの更新が成功し、CategoryUpResultが返る", func() {
			param := pb.CategoryUpParam{Crud: pb.CRUD_UPDATE, Id: category.GetId(), Name: "衣料品_更新テスト"}
			result, _ := srv.Update(ctx, &param)
			Expect(result.Error).To(BeNil())
		})
		It("カテゴリの更新が失敗し、CategoryUpResultが返る", func() {
			id := "b1524011-b6af-417e-8bf2-f449dd58b5c1"
			param := pb.CategoryUpParam{Crud: pb.CRUD_UPDATE, Id: id, Name: "衣料品_更新テスト"}
			result, _ := srv.Update(ctx, &param)
			e := pb.Error{Type: "CRUD Error", Message: fmt.Sprintf("Failed to find category with ID: %s", id)}
			Expect(result.Error).To(Equal(&e))
		})
	})
	Context("Delete()メソッドのテスト", Label("Delete"), func() {
		It("カテゴリの削除が成功し、CategoryUpResultが返る", func() {
			param := pb.CategoryUpParam{Crud: pb.CRUD_DELETE, Id: category.GetId(), Name: category.GetName()}
			result, _ := srv.Delete(ctx, &param)
			Expect(result.Error).To(BeNil())
		})
		It("カテゴリの削除が失敗し、CategoryUpResultが返る", func() {
			id := "b1524011-b6af-417e-8bf2-f449dd58b5c1"
			param := pb.CategoryUpParam{Crud: pb.CRUD_DELETE, Id: id, Name: "衣料品_削除テスト"}
			result, _ := srv.Delete(ctx, &param)
			e := pb.Error{Type: "CRUD Error", Message: fmt.Sprintf("Failed to find category with ID: %s", id)}
			Expect(result.Error).To(Equal(&e))
		})
	})
})
