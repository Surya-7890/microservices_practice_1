package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/gen"
)

type SalesReportService struct {
	gen.UnimplementedSalesReportServer
}

/* GET: /v1/report */
func (a *SalesReportService) GetReport(ctx context.Context, req *gen.SalesReportRequest) (*gen.SalesReportResponse, error) {
	res := &gen.SalesReportResponse{}
	fmt.Println("get sales report")
	return res, nil
}