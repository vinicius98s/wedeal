package loader

import (
	"context"
	"fmt"
	"strconv"

	c "wedeal/modules/category"

	"github.com/graph-gophers/dataloader"
)

func (client *Client) ListCategories(categoryIDs []uint) ([]c.Category, error) {
	var categories []c.Category
	for _, categoryID := range categoryIDs {
		c := c.Category{ID: categoryID, Name: fmt.Sprintf("name#%v", categoryID)}
		categories = append(categories, c)
	}
	return categories, nil
}

func GetCategoryBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		results = append(results, &result)
		return results
	}
	var categoryIDs []uint
	for _, key := range keys {
		id, err := strconv.ParseUint(key.String(), 10, 32)
		if err != nil {
			return handleError(err)
		}
		categoryIDs = append(categoryIDs, uint(id))
	}
	categories, err := keys[0].(*ResolverKey).GetClient().ListCategories(categoryIDs)
	if err != nil {
		return handleError(err)
	}
	var categoryMap = make(map[uint]c.Category, len(categoryIDs))
	for _, category := range categories {
		categoryMap[category.ID] = category
	}
	for _, category := range categories {
		category = categoryMap[category.ID]
		result := dataloader.Result{
			Data:  category,
			Error: nil,
		}
		results = append(results, &result)
	}

	return results
}
