package adapter_test

import (
	"command-service/internal/domain/models/categories"
	"command-service/internal/errs"
	"command-service/internal/presentation/adapter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/watarui-go-cqrs/pb/pb"
)

var _ = Describe("categoryAdapeterImpl構造体", Ordered, Label("メソッドのテスト"), func() {
	var adapt adapter.CategoryAdapter
	// 前処理
	BeforeAll(func() {
		// インスタンス生成
		adapt = adapter.NewCategoryAdapterImpl()
	})
	// ToEntity()メソッドのテスト
	Context("ToEntity()メソッド", Label("ToEntity"), func() {
		It("IdとNameフィールドに値を渡すと、entity.Categoryを返す", func() {
			param := pb.CategoryUpParam{Crud: pb.CRUD_UPDATE, Id: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: "文房具"}
			result, _ := adapt.ToEntity(&param)
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			name, _ := categories.NewCategoryName("文房具")
			Expect(result).To(Equal(categories.BuildCategory(id, name)))
		})
		It("Idのみを渡すと、entity.Categoryを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_DELETE,
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				Name: ""}
			result, _ := adapt.ToEntity(&param)
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			Expect(result).To(Equal(categories.BuildCategory(id, nil)))
		})
		It("Nameのみを渡すと、entity.Categoryを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_INSERT,
				Id:   "",
				Name: "文房具"}
			result, _ := adapt.ToEntity(&param)
			Expect(result.Name().Value()).To(Equal("文房具"))
		})
		It("36文字より短いID値を渡すと、errs.DomainErrorを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_UPDATE,
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c",
				Name: "文房具"}
			_, err := adapt.ToEntity(&param)
			Expect(err).To(Equal(errs.NewDomainError("category ID must be 36 characters long")))
		})
		It("36文字より大きいID値を渡すと、errs.DomainErrorを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_UPDATE,
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5cac",
				Name: "文房具"}
			_, err := adapt.ToEntity(&param)
			Expect(err).To(Equal(errs.NewDomainError("category ID must be 36 characters long")))
		})
		It("UUID形式ではないID値を渡すと、errs.DomainErrorを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_UPDATE,
				Id:   "aaaaaaaaaabbbbbbbbbbccccccccccdddddd",
				Name: "文房具"}
			_, err := adapt.ToEntity(&param)
			Expect(err).To(Equal(errs.NewDomainError("category ID must be a valid UUID format")))
		})
		It("2文字未満のName値を渡すと、errs.DomainErrorを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_UPDATE,
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				Name: "文"}
			_, err := adapt.ToEntity(&param)
			Expect(err).To(Equal(errs.NewDomainError("category name must be between 3 and 50 characters long")))
		})
		It("20文字より大きいName値を渡すと、errs.DomainErrorを返す", func() {
			param := pb.CategoryUpParam{
				Crud: pb.CRUD_UPDATE,
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				Name: "aaaaaaaaaabbbbbbbbbbb"}
			_, err := adapt.ToEntity(&param)
			Expect(err).To(Equal(errs.NewDomainError("category name must be between 3 and 50 characters long")))
		})
	})
	// ToResult()メソッドのテスト
	Context("ToResult()メソッド", Label("ToResult"), func() {
		It("entity.Categoryを渡すと、pb.Categoryを保持したCategoryUpResultを返す", func() {
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			name, _ := categories.NewCategoryName("文房具")
			category := categories.BuildCategory(id, name)
			result := adapt.ToResult(category)
			r := pb.Category{Id: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: "文房具"}
			Expect(result.Category).To(Equal(&r))
		})
		It("errs.DomainErrorを渡すと、pb.Errorを保持したCategoryUpResultを返す", func() {
			err := errs.NewDomainError("文房具は既に登録されています。")
			result := adapt.ToResult(err)
			Expect(result.Error.Type).To(Equal("Domain Error"))
			Expect(result.Error.Message).To(Equal("文房具は既に登録されています。"))
		})
		It("errs.CRUDErrorを渡すと、pb.Errorを保持したCategoryUpResultを返す", func() {
			err := errs.NewCRUDError("文房具は既に登録されています。")
			result := adapt.ToResult(err)
			e := pb.Error{Type: "CRUD Error", Message: "文房具は既に登録されています。"}
			Expect(result.Error).To(Equal(&e))
		})
		It("errs.InternalErrorを渡すと、pb.Errorを保持したCategoryUpResultを返す", func() {
			err := errs.NewInternalError("文房具は既に登録されています。")
			result := adapt.ToResult(err)
			e := pb.Error{Type: "Internal Error", Message: "文房具は既に登録されています。"}
			Expect(result.Error).To(Equal(&e))
		})
	})
})
