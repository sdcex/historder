package operations

import (
	"github.com/sdcex/historder/pkg/models"
)

// GetMerchantOrdersOKBody get merchant orders o k body
// swagger:model GetMerchantOrdersOKBody
type GetMerchantOrdersOKBody struct {

	// result
	// Required: true
	Result []*models.MerchantOrder `json:"result"`

	// total count
	// Required: true
	TotalCount *int64 `json:"totalCount"`
}
