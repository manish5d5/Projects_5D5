package models


type ProductFilter struct {
    CategoryID []int
    Gender     []string
    Color      []string
    Size       []string
    Brand      []string
    MinPrice   int
    MaxPrice   int
    Rating     float32
    Page       int
    Limit      int
}


type ProductResult struct {
	ProductID int
	Name      string
	Brand     string
	Gender    string
	Category  string
	Rating    float32
	Reviews    []string
	BasePrice int
	Color     string
	VariantID int
	SKU       string
}