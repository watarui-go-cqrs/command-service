package impl_test

import (
	"command-service/internal/application"
	"command-service/internal/application/service"
	"command-service/internal/domain/models/categories"
	"command-service/internal/errs"
	"context"
	"fmt"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("categoryServiceImpl構造体", Ordered, Label("メソッドのテスト"), func() {
	var (
		category  *categories.Category
		service   service.CategoryService
		ctx       context.Context
		container *fx.App
	)

	BeforeAll(func() {
		ctx = context.Background()
		container = fx.New(
			application.SrvDepend,
			fx.Populate(&service),
		)
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {
		// タイムスタンプを使ってユニークなカテゴリ名を生成
		uniqueName := fmt.Sprintf("飲料水_%d", GinkgoRandomSeed())
		name, _ := categories.NewCategoryName(uniqueName)
		category, _ = categories.NewCategory(name)
	})

	AfterAll(func() {
		err := container.Stop(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Add()メソッドのテスト", Label("Add"), func() {
		It("カテゴリ登録が成功し、nilが返る", func() {
			result := service.Add(ctx, category)
			Expect(result).To(BeNil())
		})
		It("存在するカテゴリ名の場合、errs.CRUDErrorが返る", func() {
			// 新しいカテゴリを作成して追加
			uniqueName := fmt.Sprintf("重複テスト_%d", GinkgoRandomSeed()+100)
			name, _ := categories.NewCategoryName(uniqueName)
			testCategory, _ := categories.NewCategory(name)

			// 先に追加してから同じものを追加
			service.Add(ctx, testCategory)
			result := service.Add(ctx, testCategory)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("category %s already exists", testCategory.Name().Value()))))
		})
	})

	Context("Update()メソッドのテスト", Label("Update"), func() {
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 新しいカテゴリを作成、追加、更新
			uniqueName := fmt.Sprintf("更新テスト_%d", GinkgoRandomSeed()+200)
			name, _ := categories.NewCategoryName(uniqueName)
			testCategory, _ := categories.NewCategory(name)

			// まずカテゴリを追加
			service.Add(ctx, testCategory)
			result := service.Update(ctx, testCategory)
			log.Println("存在するobj_idを指定すると、nilが返る", result)
			Expect(result).To(BeNil())
		})
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			// 新しいカテゴリ（DBに存在しない）で更新を試す
			uniqueName := fmt.Sprintf("存在しない_%d", GinkgoRandomSeed()+300)
			name, _ := categories.NewCategoryName(uniqueName)
			nonExistentCategory, _ := categories.NewCategory(name)
			result := service.Update(ctx, nonExistentCategory)
			log.Println("存在しないobj_idを指定すると、errs.CRUDErrorが返る", result)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("Failed to find category with ID: %s", nonExistentCategory.Id().Value()))))
		})
	})

	Context("Delete()メソッドのテスト", Label("Delete"), func() {
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 新しいカテゴリを作成、追加、削除
			uniqueName := fmt.Sprintf("削除成功テスト_%d", GinkgoRandomSeed()+400)
			name, _ := categories.NewCategoryName(uniqueName)
			testCategory, _ := categories.NewCategory(name)

			// まずカテゴリを追加してから削除
			service.Add(ctx, testCategory)
			result := service.Delete(ctx, testCategory.Id())
			Expect(result).To(BeNil())
		})
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			// 新しいカテゴリ（DBに存在しない）で削除を試す
			uniqueName := fmt.Sprintf("削除失敗テスト_%d", GinkgoRandomSeed()+500)
			name, _ := categories.NewCategoryName(uniqueName)
			nonExistentCategory, _ := categories.NewCategory(name)
			result := service.Delete(ctx, nonExistentCategory.Id())
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("Failed to find category with ID: %s", nonExistentCategory.Id().Value()))))
		})
	})
})
