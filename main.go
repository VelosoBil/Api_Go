package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Clientes struct {
	gorm.Model
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/clientes", getClientes)
	r.POST("/clientes", createClientes)
	r.GET("/clientes/:id", getClientesid)
	r.PUT("/clientes/:id", updateClientes)
	r.DELETE("/clientes/:id", deleteClientes)

	return r
}

func getClientes(c *gin.Context) {
	var clientes []Clientes
	if err := db.Find(&clientes).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, clientes)
	}
}

func createClientes(c *gin.Context) {
	var clientes Clientes
	c.BindJSON(&clientes)
	if err := db.Create(&clientes).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"Novo Cliente": clientes.NOME + " " + "criado com sucesso!!"})
		c.JSON(200, clientes)
	}
}

func getClientesid(c *gin.Context) {
	id := c.Params.ByName("id")
	var clientes Clientes
	if err := db.Where("id = ?", id).First(&clientes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, clientes)
	}
}
func updateClientes(c *gin.Context) {
	var clientes Clientes
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&clientes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&clientes)
	db.Save(&clientes)
	c.JSON(200, gin.H{"Novo Cliente": clientes.NOME + " " + "Atualizado com sucesso!!"})
	c.JSON(200, clientes)
}

func deleteClientes(c *gin.Context) {
	id := c.Params.ByName("id")
	var clientes Clientes
	d := db.Where("id = ?", id).Delete(&clientes)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deletado com sucesso!!"})
}

func main() {
	dsn := "root:960102@tcp(127.0.0.1:3306)/dbgoapi?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Clientes{})

	r := setupRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
