package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/jmoiron/sqlx"
)

// NewProductRepository is the constructor for productRepository
func NewProductRepository(db *sqlx.DB) product.ProductRepository {
	return &productRepository{db: db}
}

// productRepository is the implementation of product.ProductRepository
type productRepository struct {
	db *sqlx.DB
}

func (r *productRepository) Save(product *product.Product) error {
	_ = product.BeforeCreate()

	q := `INSERT INTO Products (
			Name, 
			Slug, 
			Description, 
			Private, 
			CreatedAt, 
			UpdatedAt, 
			VersionFormat, 
			SetupGuideCompleted
		) VALUES (:Name, :Slug, :Description, :Private, :CreatedAt, :UpdatedAt, :VersionFormat, :SetupGuideCompleted)`

	exec, err := r.db.NamedExec(q, product)
	if err != nil {
		return err
	}

	insertId, _ := exec.LastInsertId()
	product.ID = int(insertId)

	_ = product.AfterCreate()
	return nil
}

func (r *productRepository) Find() ([]product.Product, error) {
	// Execute the database query
	rows, err := r.db.Queryx("SELECT * FROM Products")
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the products
	var products []product.Product

	// Iterate over the result set
	for rows.Next() {
		var p product.Product
		// Map the row's data to the product struct
		if err := rows.StructScan(&p); err != nil {
			return nil, err // Return an error if mapping fails
		}
		p.FormatAuditable()
		products = append(products, p) // Add the product to the slice
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil // Return the list of products
}

func (r *productRepository) FindBySlug(slug values.Slug) (product.Product, error) {
	var p product.Product
	err := r.db.QueryRowx("SELECT * FROM Products WHERE Slug = $1 LIMIT 1", slug).StructScan(&p)
	if err != nil {
		return product.Product{}, err
	}

	return p, err
}

func (r *productRepository) Update(product product.Product) error {
	_ = product.BeforeUpdate()

	q := `UPDATE Products SET 
			Name = :Name, 
			Slug = :Slug, 
			Description = :Description,
			Private = :Private,
           	VersionFormat = :VersionFormat, 
			SetupGuideCompleted = :SetupGuideCompleted WHERE ID = :ID`

	_, err := r.db.NamedExec(q, product)
	if err != nil {
		return err
	}

	_ = product.AfterUpdate()
	return nil
}

func (r *productRepository) Delete(product product.Product) error {
	//TODO implement me
	panic("implement me")
}

func (r *productRepository) SaveSetupGuide(guide product.SetupGuide) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	q := "UPDATE Products SET VersionFormat = $1, SetupGuideCompleted = true WHERE ID = $2"
	_, err = tx.Exec(q, guide.VersionFormat, guide.ProductID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	q = "INSERT INTO Channels (Name, ProductID, Closed) VALUES (:Name, :ProductID, :Closed)"
	for _, channel := range guide.Channels {
		_, err := tx.NamedExec(q, channel)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if tx.Commit() != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (r *productRepository) GetPlatformAvailability(product *product.Product) error {
	q := `SELECT 
    EXISTS (
        SELECT 1 
        FROM Apps a 
        WHERE a.ProductID = p.ID AND a.Platform = 'Android'
    ) AS HasAndroid,
    EXISTS (
        SELECT 1 
        FROM Apps a 
        WHERE a.ProductID = p.ID AND a.Platform = 'iOS'
    ) AS HasIOS,
    EXISTS (
        SELECT 1 
        FROM Apps a 
        WHERE a.ProductID = p.ID AND a.Platform = 'Windows'
    ) AS HasWindows,
    EXISTS (
        SELECT 1 
        FROM Apps a 
        WHERE a.ProductID = p.ID AND a.Platform = 'Linux'
    ) AS HasLinux,
    EXISTS (
        SELECT 1 
        FROM Apps a 
        WHERE a.ProductID = p.ID AND a.Platform = 'macOS'
    ) AS HasMacOS,
    EXISTS (
        SELECT 1 
        FROM Apps a 
        WHERE a.ProductID = p.ID AND a.Platform = 'Other'
    ) AS HasOther FROM Products p WHERE ID = $1`

	row := r.db.QueryRow(q, product.ID)

	err := row.Scan(
		&product.HasAndroid,
		&product.HasIOS,
		&product.HasWindows,
		&product.HasLinux,
		&product.HasMacOS,
		&product.HasOther,
	)
	if err != nil {
		return err
	}

	return nil
}
