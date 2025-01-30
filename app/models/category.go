package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetCategories() ([]Category, error) {
	query := "SELECT id, name FROM categories"
    rows, err := Database.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var categories []Category
    for rows.Next() {
        var category Category
        err = rows.Scan(&category.ID, &category.Name)
        if err != nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    return categories, nil
}

func InsertIntoCategoryPost(postId, categorieId int) error {
	query := "INSERT INTO post_categorie (post_id, categorie_id) VALUES (?,?)"
    _, err := Database.Exec(query, postId, categorieId)
    if err!= nil {
        return err
    }
    return nil
}


func CorrectCategories(id int) []string {
	query:= `SELECT c.name FROM categories c
	INNER JOIN post_categorie pc ON c.id = pc.categorie_id
	WHERE pc.post_id = ?`
	rows, err := Database.Query(query, id)
	if err != nil {
		return nil
	}
	defer rows.Close()
	categories := []string{}
	for rows.Next() {
		var categorie string
		err := rows.Scan(&categorie)
		if err != nil {
			continue
		}
		categories = append(categories, categorie)
	}
	return categories
}