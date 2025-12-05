package repositories
import (
	"context"
	"fmt"	
	// "time"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"filtering_pagination/models"

)






type FashionStoreRepository struct {
	db *pgxpool.Pool
}
func NewFashionStoreRepository(db *pgxpool.Pool) *FashionStoreRepository {
	return &FashionStoreRepository{
		db: db,
	}
}



//______________filtering and pagination functions____________________//
func (r *FashionStoreRepository) FilterProducts(ctx context.Context, f models.ProductFilter) ([]models.ProductResult, error) {

    query := `
        SELECT 
            p.id AS product_id,
            p.name,
            p.brand,
            p.gender,
            p.base_price,
            p.rating,
            COALESCE(json_agg(pr.review_text) FILTER (WHERE pr.review_text IS NOT NULL), '[]') AS reviews,
            c.name AS category,
            v.id AS variant_id,
            v.color,
            v.sku
        FROM products p
        JOIN categories c ON p.category_id = c.id
        LEFT JOIN product_variants v ON v.product_id = p.id
        LEFT JOIN product_reviews pr ON pr.product_id = p.id
    `

    where := "WHERE 1=1"
    args := []interface{}{}
    idx := 1

    // CATEGORY (multi-select)
    if len(f.CategoryID) > 0 {
        where += fmt.Sprintf(" AND p.category_id = ANY($%d)", idx)
        args = append(args, f.CategoryID)
        idx++
    }

    // BRAND (multi-select)
    if len(f.Brand) > 0 {
        where += fmt.Sprintf(" AND p.brand = ANY($%d)", idx)
        args = append(args, f.Brand)
        idx++
    }

    // GENDER (multi-select)
    if len(f.Gender) > 0 {
        where += fmt.Sprintf(" AND p.gender = ANY($%d)", idx)
        log.Print("--------------------------------------------")
        log.Println("Gender filter at repo:", f.Gender)
        log.Print("--------------------------------------------")
        args = append(args, f.Gender)
        idx++
    }

    // COLOR (multi-select)
    if len(f.Color) > 0 {
        where += fmt.Sprintf(" AND v.color = ANY($%d)", idx)
        args = append(args, f.Color)
        idx++
    }

    // SIZE (multi-select)
    if len(f.Size) > 0 {
        query += " JOIN variant_sizes vs ON vs.variant_id = v.id "
        where += fmt.Sprintf(" AND vs.size = ANY($%d)", idx)
        args = append(args, f.Size)
        idx++
    }

    // PRICE RANGE
    if f.MinPrice > 0 {
        where += fmt.Sprintf(" AND p.base_price >= $%d", idx)
        args = append(args, f.MinPrice)
        idx++
    }

    if f.MaxPrice > 0 {
        where += fmt.Sprintf(" AND p.base_price <= $%d", idx)
        args = append(args, f.MaxPrice)
        idx++
    }

    // RATING
    if f.Rating > 0 {
        where += fmt.Sprintf(" AND p.rating >= $%d", idx)
        args = append(args, f.Rating)
        idx++
    }

    query += where + `
        GROUP BY 
            p.id, p.name, p.brand, p.gender, p.base_price, p.rating,
            c.name, v.id, v.color, v.sku
    `

    // Pagination
    if f.Page == 0 { f.Page = 1 }
    if f.Limit == 0 { f.Limit = 20 }

    offset := (f.Page - 1) * f.Limit

    query += fmt.Sprintf(" ORDER BY p.rating DESC LIMIT $%d OFFSET $%d", idx, idx+1)
    args = append(args, f.Limit, offset)

    log.Println("QUERY => ", query)
    log.Println("ARGS  => ", args)



    rows, err := r.db.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []models.ProductResult

    for rows.Next() {
        var pr models.ProductResult
        err = rows.Scan(
            &pr.ProductID,
            &pr.Name,
            &pr.Brand,
            &pr.Gender,
            &pr.BasePrice,
            &pr.Rating,
            &pr.Reviews,
            &pr.Category,
            &pr.VariantID,
            &pr.Color,
            &pr.SKU,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, pr)
    }

    return results, nil
}
