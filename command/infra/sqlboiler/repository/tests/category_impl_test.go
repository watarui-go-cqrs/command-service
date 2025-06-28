package repository_test

import (
	"command-service/command/domain/models/categories"
	"command-service/command/errs"
	"command-service/command/infra/sqlboiler/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/aarondl/sqlboiler/v4/boil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("categoryRepositorySQLBoiler構造体", Ordered, Label("CategoryRepositoryインターフェイスメソッドのテスト"), func() {
	var rep categories.CategoryRepository
	var ctx context.Context
	var tran *sql.Tx
	// 前処理
	BeforeAll(func() {
		// リポジトリの生成
		rep = repository.NewCategoryRepositorySQLBoiler()
	})

	// テスト毎の前処理
	BeforeEach(func() {
		ctx = context.Background()       // Contextの生成
		tran, _ = boil.BeginTx(ctx, nil) // トランザクションの開始
	})

	// テスト毎の後処理
	AfterEach(func() {
		tran.Rollback() // トランザクションのロールバック
	})

	// Exists()メソッドのテスト
	Context("同名の商品カテゴリが存在確認結果を返す", Label("Exists"), func() {
		It("存在しない商品の場合nilが返る", func() {
			name, _ := categories.NewCategoryName("食品")
			category, _ := categories.NewCategory(name)
			result := rep.Exists(ctx, tran, category)
			Expect(result).To(BeNil())
		})
		It("存在するカテゴリ名の場合、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(name)
			result := rep.Exists(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("category %s already exists", category.Name().Value()))))
		})
	})
	// Create()メソッドのテスト
	Context("新しい商品カテゴリを永続化する", Label("Create"), func() {
		It("カテゴリが登録成功し、nilが返る", func() {
			name, _ := categories.NewCategoryName("食品")
			category, _ := categories.NewCategory(name)
			result := rep.Create(ctx, tran, category)
			Expect(result).To(BeNil())
		})
		It("obj_idが同じカテゴリを追加すると、errs.CRUDErrorが返る", func() {
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			name, _ := categories.NewCategoryName("文房具")
			category := categories.BuildCategory(id, name)
			result := rep.Create(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError("duplicate entry error")))
		})
	})
	// UpdateById()メソッドのテスト
	Context("商品カテゴリを変更する", Label("UpdateById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
			name, _ := categories.NewCategoryName("文房具")
			category := categories.BuildCategory(id, name)
			result := rep.UpdateById(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("Failed to find category with ID: %s", category.Id().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			name, _ := categories.NewCategoryName("文房具1")
			category := categories.BuildCategory(id, name)
			result := rep.UpdateById(ctx, tran, category)
			Expect(result).To(BeNil())
		})
	})
	// DeleteById()メソッドのテスト
	Context("商品カテゴリを削除する", Label("DeleteById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
			name, _ := categories.NewCategoryName("文房具1")
			category := categories.BuildCategory(id, name)
			result := rep.DeleteById(ctx, tran, category.Id())
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("Failed to find category with ID: %s",
					category.Id().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 削除対象のカテゴリを登録する
			name, _ := categories.NewCategoryName("食品")
			category, _ := categories.NewCategory(name)
			rep.Create(ctx, tran, category)
			// 登録したカテゴリを削除する
			result := rep.DeleteById(ctx, tran, category.Id())
			Expect(result).To(BeNil())
		})
	})
})
