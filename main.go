package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	v := validator.New()
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well , hello there !!!")
	})

	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found !")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate :"required,min=4"`
			// Vendor  string `json:"vendor" validate :"min=3,max=15"`
			// Country string `json:"country" validate :"len=2"`
		}

		var reqBody body
		e.Validator = &ProductValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}

		products = append(products, product)

		return c.JSON(http.StatusOK, product)
	})

	e.PUT("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found !")
		}

		type body struct {
			Name string `json:"product_name" validate :"required,min=4"`
		}

		var reqBody body
		e.Validator = &ProductValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product[pID] = reqBody.Name
		return c.JSON(http.StatusOK, product)
	})

	e.DELETE("/products/:id", func(c echo.Context) error {
		var product map[int]string
		var index int
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for i, p := range products {
			for k := range p {
				if pID == k {
					product = p
					index = i
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found !")
		}
		splice := func(s []map[int]string, index int) []map[int]string {
			return append(s[:index], s[index+1:]...)
		}
		products = splice(products, index)
		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

}
