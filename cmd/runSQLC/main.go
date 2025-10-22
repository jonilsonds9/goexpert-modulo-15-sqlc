package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jonilsonds9/goexpert-modulo-15-sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Backend Description", Valid: true},
	//})
	//if err != nil {
	//	panic(err)
	//}

	//categories, err := queries.ListCategories(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//for _, category := range categories {
	//	println("Category:", category.ID, category.Name, category.Description.String)
	//}

	//err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	//	ID:          "17a5a842-1488-433a-8b30-ee2a3e301ac7",
	//	Name:        "Backend Updated",
	//	Description: sql.NullString{String: "Backend Description Updated", Valid: true},
	//})

	err = queries.DeleteCategory(ctx, "17a5a842-1488-433a-8b30-ee2a3e301ac7")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println("Category:", category.ID, category.Name, category.Description.String)
	}
}
