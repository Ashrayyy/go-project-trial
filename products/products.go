package products

import (
	"project/constants"
	"project/products/amazon"
	"project/products/flipkart"
)

func GetProducts(filter ...int) (constants.Resp, error) {
	category := constants.ALL_PRODUCTS
	if len(filter) > 0 {
		category = filter[0]
	}
	var resp constants.Resp

	k, _ := amazon.GetProducts(category)
	resp.Items = append(resp.Items, k.Items...)
	k, _ = flipkart.GetProducts(category)
	resp.Items = append(resp.Items, k.Items...)

	return resp, nil
}
