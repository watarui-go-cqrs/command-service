package server_test

import (
	"command-service/internal/infrastructure/sqlboiler/handler"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHelperPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "presen/serverパッケージのテスト")
}

var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../../infrastructure/sqlboiler/config/database.toml")
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := handler.DBConnect()
	Expect(err).NotTo(HaveOccurred(), "Failed to connect to the database")
})
