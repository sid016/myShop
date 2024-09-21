package service

import (
	"context"
	"myshop/internal/db"
	"myshop/internal/models"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

type (
	Service struct {
		Echo       *echo.Echo
		Handler    *handler
		Middleware *middleware
		RoutePath  *path
		poolRW     *pgxpool.Pool
		poolRO     *pgxpool.Pool
		repo       db.Repo
	}
	path struct {
		Basepath string
	}
)

func New(conn *pgxpool.Pool, repo *db.RepoConn) (*Service, error) {
	s := &Service{
		Echo:   echo.New(),
		poolRW: conn,
		poolRO: conn,
		repo:   repo,
	}

	s.Handler = newHandler(s)
	s.Middleware = newMiddleware(s)
	s.registerRoutes()

	return nil, nil
}

func (s *Service) registerRoutes() {

	api := s.Echo.Group(s.RoutePath.Basepath)

	api.Add(http.MethodGet,
		s.RoutePath.GetSeller(),
		s.Handler.GetSeller(),
		s.Middleware.GetSellerFromDB(),
	)

}

func (p *path) GetSeller() string {
	return p.Basepath + "/seller"
}

// GetSellerHandler to handle fetching seller details
func GetSellerHandler(c echo.Context) error {
	conn := c.Get("db").(*pgxpool.Pool)
	sellerID := c.Param("id")

	var seller models.Seller
	err := conn.QueryRow(context.Background(), "SELECT * FROM sellers WHERE seller_id = $1", sellerID).Scan(&seller.ID, &seller.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Seller not found"})
	}

	return c.JSON(http.StatusOK, seller)
}

// SetupRoutes defines all API routes in one block
func SetupRoutes(e *echo.Echo, conn *pgxpool.Pool) {
	// e.Use(DBMiddleware(conn))

	// Define the endpoints
	e.GET("/seller/:id", GetSellerHandler)
	e.GET("/seller/:id", GetSellerHandler) // GET method with handler and middleware
	e.GET("/product/:id", func(c echo.Context) error {
		conn := c.Get("db").(*pgxpool.Pool)
		productID := c.Param("id")

		var product models.Product
		err := conn.QueryRow(context.Background(), "SELECT * FROM products WHERE product_id = $1", productID).Scan(&product.ID, &product.ProductID, &product.Name)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}

		return c.JSON(http.StatusOK, product)
	})
	// Add more routes similarly...
}
