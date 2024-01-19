package testdata

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"sort"
)

type Theme struct {
	ID   int
	Name string
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	CreatedDate time.Time
	Themes      []Theme
	Weight      float64
}

type User struct {
	ID              int
	FollowedThemes  []Theme
}

var (
	themes = []Theme{
		{ID: 1, Name: "summerDelights"},
		{ID: 2, Name: "funkyTeens"},
		{ID: 3, Name: "boldWoman"},
		{ID: 4, Name: "colourfulBliss"},
		{ID: 5, Name: "under500"},
		{ID: 6, Name: "oldIsGold"},
	}

	users = map[int]User{
		1: {
			ID: 1,
			FollowedThemes: []Theme{
				{ID: 1, Name: "summerDelights"},
				{ID: 2, Name: "funkyTeens"},
				{ID: 3, Name: "boldWoman"},
			},
		},
		2: {
			ID: 2,
			FollowedThemes: []Theme{
				{ID: 4, Name: "colourfulBliss"},
				{ID: 5, Name: "under500"},
				{ID: 6, Name: "oldIsGold"},
			},
		},
		3: {
			ID: 3,
			FollowedThemes: []Theme{
				{ID: 1, Name: "summerDelights"},
				{ID: 3, Name: "boldWoman"},
			},
		},
		4: {
			ID: 4,
			FollowedThemes: []Theme{
				{ID: 4, Name: "colourfulBliss"},
				{ID: 5, Name: "under500"},
				{ID: 6, Name: "oldIsGold"},
				{ID: 6, Name: "boldWoman"},
			},
		},
		5: {
			ID: 5,
			FollowedThemes: []Theme{
			},
		},
		6: {
			ID: 6,
			FollowedThemes: []Theme{
				{ID: 4, Name: "colourfulBliss"},
				{ID: 5, Name: "under500"},
				{ID: 6, Name: "oldIsGold"},
			},
		},
	}

	dummyProducts = []Product{
		{ID: 1, Name: "Product 1", Price: 100, CreatedDate: time.Now().AddDate(0, -1, 0), Themes: []Theme{{ID: 1, Name: "summerDelights"}}, Weight: 0.0},
		{ID: 2, Name: "Product 2", Price: 200, CreatedDate: time.Now().AddDate(0, -2, 0), Themes: []Theme{{ID: 2, Name: "funkyTeens"}}, Weight: 0.0},
		{ID: 3, Name: "Product 3", Price: 300, CreatedDate: time.Now().AddDate(0, -3, 0), Themes: []Theme{{ID: 3, Name: "boldWoman"}}, Weight: 0.0},
		{ID: 4, Name: "Product 4", Price: 400, CreatedDate: time.Now().AddDate(0, -4, 0), Themes: []Theme{{ID: 4, Name: "colourfulBliss"}}, Weight: 0.0},
		{ID: 5, Name: "Product 5", Price: 500, CreatedDate: time.Now().AddDate(0, -5, 0), Themes: []Theme{{ID: 5, Name: "under500"}}, Weight: 0.0},
		{ID: 6, Name: "Product 6", Price: 600, CreatedDate: time.Now().AddDate(0, -6, 0), Themes: []Theme{{ID: 6, Name: "oldIsGold"}}, Weight: 0.0},
		{ID: 7, Name: "Product 7", Price: 700, CreatedDate: time.Now().AddDate(0, -7, 0), Themes: []Theme{{ID: 7, Name: "summerDelights"}}, Weight: 0.0},
		{ID: 8, Name: "Product 8", Price: 800, CreatedDate: time.Now().AddDate(0, -8, 0), Themes: []Theme{{ID: 8, Name: "funkyTeens"}}, Weight: 0.0},
		{ID: 9, Name: "Product 9", Price: 900, CreatedDate: time.Now().AddDate(0, -9, 0), Themes: []Theme{{ID: 9, Name: "boldWoman"}}, Weight: 0.0},
		{ID: 10, Name: "Product 10", Price: 1000, CreatedDate: time.Now().AddDate(0, -10, 0), Themes: []Theme{{ID: 10, Name: "colourfulBliss"}}, Weight: 0.0},
		{ID: 11, Name: "Product 11", Price: 1100, CreatedDate: time.Now().AddDate(0, -11, 0), Themes: []Theme{{ID: 11, Name: "under500"}}, Weight: 0.0},
		{ID: 12, Name: "Product 12", Price: 1200, CreatedDate: time.Now().AddDate(0, -12, 0), Themes: []Theme{{ID: 12, Name: "oldIsGold"}}, Weight: 0.0},
		{ID: 13, Name: "Product 13", Price: 1300, CreatedDate: time.Now().AddDate(0, -13, 0), Themes: []Theme{{ID: 13, Name: "summerDelights"}}, Weight: 0.0},
		{ID: 14, Name: "Product 14", Price: 1400, CreatedDate: time.Now().AddDate(0, -14, 0), Themes: []Theme{{ID: 14, Name: "funkyTeens"}}, Weight: 0.0},
		{ID: 15, Name: "Product 15", Price: 1500, CreatedDate: time.Now().AddDate(0, -15, 0), Themes: []Theme{{ID: 15, Name: "boldWoman"}}, Weight: 0.0},
		{ID: 16, Name: "Product 16", Price: 1600, CreatedDate: time.Now().AddDate(0, -16, 0), Themes: []Theme{{ID: 16, Name: "colourfulBliss"}}, Weight: 0.0},
		{ID: 17, Name: "Product 17", Price: 1700, CreatedDate: time.Now().AddDate(0, -17, 0), Themes: []Theme{{ID: 17, Name: "under500"}}, Weight: 0.0},
		{ID: 18, Name: "Product 18", Price: 1800, CreatedDate: time.Now().AddDate(0, -18, 0), Themes: []Theme{{ID: 18, Name: "oldIsGold"}}, Weight: 0.0},
		{ID: 19, Name: "Product 19", Price: 1900, CreatedDate: time.Now().AddDate(0, -19, 0), Themes: []Theme{{ID: 19, Name: "summerDelights"}}, Weight: 0.0},
		{ID: 20, Name: "Product 20", Price: 2000, CreatedDate: time.Now().AddDate(0, -20, 0), Themes: []Theme{{ID: 20, Name: "funkyTeens"}}, Weight: 0.0},
	}
)

func GetUserThemes(userID int) ([]Theme, error) {
	if user, ok := users[userID]; ok {
		return user.FollowedThemes, nil
	}
	return nil, fmt.Errorf("User does not exist for the provided ID")
}

func GetProductsByTheme(theme Theme, count, page int) []Product {
	// Filter products for the given theme 
	// This will be replaced with SQL queries
	filteredProducts := []Product{}
	for _, product := range dummyProducts {
		for _, productTheme := range product.Themes {
			if productTheme.Name == theme.Name {
				filteredProducts = append(filteredProducts, product)
				break
			}
		}
	}

	//this will be handled as a CRON job which updated weight periodically 
	assignWeightToProducts(filteredProducts)

	// Sort the filtered products by weightage
	sort.Slice(filteredProducts, func(i, j int) bool {
		return filteredProducts[i].Weight < filteredProducts[j].Weight
	})

	startIndex := (page - 1) * count
	endIndex := startIndex + count

	if startIndex >= len(filteredProducts) {
		return []Product{}
	}
	if endIndex > len(filteredProducts) {
		endIndex = len(filteredProducts)
	}

	return filteredProducts[startIndex:endIndex]
}

func FetchRandomTheme() Theme {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(themes))
	return themes[randomIndex]
}

func assignWeightToProducts(products []Product) {
	for i := range products {
		priceWeight := products[i].Price / 100 * 0.1
		recencyWeight := 0.1 * math.Min(float64(time.Since(products[i].CreatedDate).Hours()/24/7), 52.0)
		products[i].Weight = priceWeight + recencyWeight
	}
}
