package recommendation

import (
	"math"

	"../testdata"
	"fmt"
	"sort"
)

type RecommendationSystem struct{}

func NewRecommendationSystem() *RecommendationSystem {
	return &RecommendationSystem{}
}

func (rs *RecommendationSystem) GetRecommendations(userId, pageLimit, page int) ([]testdata.Product, error) {
	recommendations := []testdata.Product{}

	userThemes, err := testdata.GetUserThemes(userId)
	if err != nil {
		return recommendations, err
	}
	selectedThemes := make(map[testdata.Theme]int)
	if len(userThemes) >= 3 {
		for _, userTheme := range userThemes {
			selectedThemes[userTheme] = 10
		}
	} else {
		for _, userTheme := range userThemes {
			selectedThemes[userTheme] = 10
		}

		for i := 0; i < 3-len(userThemes); i++ {
			randomTheme := testdata.FetchRandomTheme()
			selectedThemes[randomTheme] = 5
		}
	}

	totalWeight := 0
	for _, weight := range selectedThemes {
		totalWeight += weight
	}

	distributedProducts := make(map[testdata.Theme]int)
	for theme, weight := range selectedThemes {
		productCount := int(math.Round(float64(weight*pageLimit) / float64(totalWeight)))
		distributedProducts[theme] = productCount
	}

	fmt.Println(distributedProducts)

	sortedThemes := make([]testdata.Theme, 0, len(distributedProducts))
	for theme := range distributedProducts {
		sortedThemes = append(sortedThemes, theme)
	}
	sort.Slice(sortedThemes, func(i, j int) bool {
		return distributedProducts[sortedThemes[i]] > distributedProducts[sortedThemes[j]]
	})

	carry := 0

	for _, theme := range sortedThemes {
		count := distributedProducts[theme] + carry
		products := testdata.GetProductsByTheme(theme, count, page)
		recommendations = append(recommendations, products...)
		if count < len(products) {
			carry += count - len(products)
		}
	}

	if len(recommendations) > pageLimit {
		recommendations = recommendations[:pageLimit]

	}

	return recommendations, nil
}
