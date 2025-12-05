package handler
import (
	"encoding/json"
	"net/http"		
	"filtering_pagination/models"
	"filtering_pagination/repositories"
	"strconv"
    "log"
    "strings"
)

type FashionHandler struct {
    repo *repositories.FashionStoreRepository
}
func NewFashionHandler(repo *repositories.FashionStoreRepository) *FashionHandler {
	return &FashionHandler{repo: repo}
}



func (h *FashionHandler) Filter(w http.ResponseWriter, r *http.Request) {
    var filter models.ProductFilter

// CATEGORY IDs (multiple)
        categoryVals := r.URL.Query()["category_id"]
        var catIDs []int
        for _, v := range categoryVals {
            if id, err := strconv.Atoi(v); err == nil {
                catIDs = append(catIDs, id)
            }
        }
        filter.CategoryID = catIDs
        
        // GENDERS (multiple)
        genders := r.URL.Query()["gender"] // []string
        for i := range genders {
            genders[i] = strings.ToLower(genders[i])
        }
        filter.Gender = genders
        
        // COLORS (multiple)
        filter.Color = r.URL.Query()["color"]
        
        // SIZES (multiple)
        filter.Size = r.URL.Query()["size"]
        
        // BRANDS (multiple)
        filter.Brand = r.URL.Query()["brand"]
        
        // PRICE RANGE
        filter.MinPrice, _ = strconv.Atoi(r.URL.Query().Get("min_price"))
        filter.MaxPrice, _ = strconv.Atoi(r.URL.Query().Get("max_price"))
        
        // RATING
        ratingStr := r.URL.Query().Get("rating")
        if ratingStr != "" {
            r64, _ := strconv.ParseFloat(ratingStr, 64)
            filter.Rating = float32(r64)
        }
        log.Println("Filter Params at handler:", filter)
        
        // PAGINATION
        filter.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
        filter.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
            products, err := h.repo.FilterProducts(r.Context(), filter)
            if err != nil {
                http.Error(w, "Error fetching products", http.StatusInternalServerError)
                return
            }
        
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(products)
}
