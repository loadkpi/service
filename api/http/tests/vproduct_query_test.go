package tests

import (
	"net/http"
	"sort"

	"github.com/ardanlabs/service/app/api/page"
	"github.com/ardanlabs/service/app/domain/vproductapp"
	"github.com/google/go-cmp/cmp"
)

func vproductQuery200(sd seedData) []table {
	prds := toAppVProducts(sd.Admins[0].User.User, sd.Admins[0].Products)
	prds = append(prds, toAppVProducts(sd.Users[0].User.User, sd.Users[0].Products)...)

	sort.Slice(prds, func(i, j int) bool {
		return prds[i].ID <= prds[j].ID
	})

	table := []table{
		{
			Name:       "basic",
			URL:        "/v1/vproducts?page=1&rows=10&orderBy=product_id,ASC",
			Token:      sd.Admins[0].Token,
			StatusCode: http.StatusOK,
			Method:     http.MethodGet,
			Resp:       &page.Document[vproductapp.Product]{},
			ExpResp: &page.Document[vproductapp.Product]{
				Page:        1,
				RowsPerPage: 10,
				Total:       len(prds),
				Items:       prds,
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
	}

	return table
}
