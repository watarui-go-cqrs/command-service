package repository

import (
	"command-service/internal/domain/models/categories"
	"command-service/internal/errs"
	"command-service/internal/infrastructure/sqlboiler/handler"
	"command-service/internal/infrastructure/sqlboiler/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
)

type categoryRepositorySQLBoiler struct {
}

func NewCategoryRepositorySQLBoiler() categories.CategoryRepository {
	models.AddCategoryHook(boil.AfterInsertHook, CategoryAfterInsertHook)
	models.AddCategoryHook(boil.AfterUpdateHook, CategoryAfterUpdateHook)
	models.AddCategoryHook(boil.AfterDeleteHook, CategoryAfterDeleteHook)
	return &categoryRepositorySQLBoiler{}
}

func (r *categoryRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	cond := models.CategoryWhere.Name.EQ(category.Name().Value())
	if exists, err := models.Categories(cond).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists {
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("category %s already exists", category.Name().Value()))
	}
}

func (r *categoryRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	model := &models.Category{
		ID:    0, // ID will be auto-generated by the database
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}

	if err := model.Insert(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (rep *categoryRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if model == nil {
		return errs.NewCRUDError(fmt.Sprintf("Failed to find category with ID: %s", category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	model.ObjID = category.Id().Value()
	model.Name = category.Name().Value()

	if _, err = model.Update(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *categoryRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, id *categories.CategoryId) error {
	model, err := models.Categories(qm.Where("obj_id = ?", id.Value())).One(ctx, tran)
	if model == nil {
		return errs.NewCRUDError(fmt.Sprintf("Failed to find category with ID: %s", id.Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}

	if _, err = model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func CategoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, m *models.Category) error {
	log.Printf("CategoryAfterInsertHook: Category with ID: %s and Name: %s was inserted", m.ObjID, m.Name)
	return nil
}

func CategoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, m *models.Category) error {
	log.Printf("CategoryAfterUpdateHook: Category with ID: %s and Name: %s was updated", m.ObjID, m.Name)
	return nil
}

func CategoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, m *models.Category) error {
	log.Printf("CategoryAfterDeleteHook: Category with ID: %s and Name: %s was deleted", m.ObjID, m.Name)
	return nil
}
