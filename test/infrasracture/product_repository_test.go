package infrasracture

import (
	"context"
	"os"
	"product-app/common/postgresql"
	"product-app/domain"
	"product-app/persistence"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setup(ctx context.Context, dbpool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbpool)
}
func clear(ctx context.Context, dbpool *pgxpool.Pool) {
	TruncateTestData(ctx, dbpool)
}

func TestGetAllProduct(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()

		assert.Equal(t, 4, len(actualProducts))

		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}
func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")

		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}
func TestAddProduct(t *testing.T) {
	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    100.0,
		Discount: 0.0,
		Store:    "DNS",
	}

	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)

		actualProducts := productRepository.GetAllProducts()

		assert.Equal(t, 1, len(actualProducts))
	})
	clear(ctx, dbPool)
}

func TestGetById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetById", func(t *testing.T) {
		product, err := productRepository.GetById(2)

		assert.Nil(t, err)

		assert.Equal(t, domain.Product{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		}, product)
	})
	clear(ctx, dbPool)
}
func TestDeleteById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("Delete", func(t *testing.T) {
		err := productRepository.DeleteById(2)

		assert.Nil(t, err)
	})
	clear(ctx, dbPool)
}
