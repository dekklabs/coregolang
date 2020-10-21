package products

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//CreateCategory modelo para crear una categoria de producto
func CreateCategory(category *entities.CategoryProduct) (err error) {
	db, _ := db.Conexion()

	results, err := db.Exec("INSERT INTO category_product(name) VALUES(?)", category.Name)
	if err != nil {
		return err
	}

	category.ID, err = results.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//UpdateCategory actualizar categoría
func UpdateCategory(category *entities.CategoryProduct) (int64, error) {
	db, _ := db.Conexion()

	row, err := db.Exec("update category_product SET name = ? where id = ?", category.Name, category.ID)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
