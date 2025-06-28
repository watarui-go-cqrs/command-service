package products_test

import (
	"command-service/command/domain/models/products"
	"command-service/command/errs"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Entity を構成する VO", Ordered, Label("ProductId 構造体の生成"), func() {
	var (
		emptyStr   *errs.DomainError
		lengthOver *errs.DomainError
		notUuid    *errs.DomainError
		productId  *products.ProductId
		uid        string
	)

	BeforeAll(func() {
		_, emptyStr = products.NewProductId("")
		_, lengthOver = products.NewProductId("1234567890123456789012345678901234567") // 37 characters
		_, notUuid = products.NewProductId("not-valid-uuid-format-here-123456789")     // 36 characters but invalid UUID format
		id, _ := uuid.NewRandom()
		uid = id.String()
		productId, _ = products.NewProductId(id.String())
	})

	Context("文字数の検証", Label("文字数"), func() {
		It("空文字列はエラー", func() {
			Expect(emptyStr).To(HaveOccurred())
			Expect(emptyStr.Error()).To(Equal("product ID must be 36 characters long"))
		})

		It("36文字を超える文字列はエラー", func() {
			Expect(lengthOver).To(HaveOccurred())
			Expect(lengthOver.Error()).To(Equal("product ID must be 36 characters long"))
		})
	})

	Context("UUID形式の検証", Label("UUID形式"), func() {
		It("UUID形式でない文字列はエラー", func() {
			Expect(notUuid).To(HaveOccurred())
			Expect(notUuid.Error()).To(Equal("product ID must be a valid UUID format"))
		})

		It("UUID形式の文字列はエラーにならない", func() {
			id, _ := products.NewProductId(uid)
			Expect(productId).To(Equal(id))
		})
	})
})
