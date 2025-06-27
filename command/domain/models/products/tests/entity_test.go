package products_test

import (
	"command-service/command/domain/models/categories"
	"command-service/command/domain/models/products"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Entity", Ordered, Label("Product の同一性検証"), func() {
	var (
		category *categories.Category
		product  *products.Product
	)

	BeforeAll(func() {
		categoryName, err := categories.NewCategoryName("Electronics")
		Expect(err).ToNot(HaveOccurred())
		category, err = categories.NewCategory(categoryName)
		Expect(err).ToNot(HaveOccurred())
		productName, err := products.NewProductName("Smartphone")
		Expect(err).ToNot(HaveOccurred())
		productPrice, err := products.NewProductPrice(500)
		Expect(err).ToNot(HaveOccurred())
		product, err = products.NewProduct(productName, productPrice, category)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("Equals メソッドの検証", func() {
		It("比較対象が nil", Label("nilチェック"), func() {
			By("nilを比較対象にするとエラーが返ること")
			_, err := product.Equals(nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("product is nil"))
		})
	})

	Context("比較結果の検証", func() {
		It("異なる Product インスタンスを比較", Label("異なるインスタンス"), func() {
			pName, err := products.NewProductName("Tablet")
			Expect(err).ToNot(HaveOccurred())
			pPrice, err := products.NewProductPrice(300)
			Expect(err).ToNot(HaveOccurred())
			p, err := products.NewProduct(pName, pPrice, category)
			Expect(err).ToNot(HaveOccurred())
			By("異なる Product インスタンスを比較すると false が返ること")
			result, _ := product.Equals(p)
			Expect(result).To(BeFalse())
		})
		It("同一の Product インスタンスを比較", Label("同一性"), func() {
			p := products.BuildProduct(product.Id(), product.Name(), product.Price(), category)
			By("同一の Product インスタンスを比較すると true が返ること")
			result, _ := product.Equals(p)
			Expect(result).To(BeTrue())
		})
	})
})
