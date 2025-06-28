package handler_test

import (
	"command-service/command/infra/sqlboiler/handler"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Connect Suite")
}

var _ = Describe("DB 接続テスト", func() {
	It("DB 接続の確認", Label("DB 接続"), func() {
		absPath, _ := filepath.Abs("../../config/database.toml")
		os.Setenv("DATABASE_TOML_PATH", absPath)
		result := handler.DBConnect()
		Expect(result).To(BeNil())
	})
})
