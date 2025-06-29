package repository_test

import (
	"command-service/internal/infrastructure/sqlboiler/handler"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepImplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../config/database.toml")
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := handler.DBConnect()
	Expect(err).To(BeNil(), "Database connection should be established successfully")
})
