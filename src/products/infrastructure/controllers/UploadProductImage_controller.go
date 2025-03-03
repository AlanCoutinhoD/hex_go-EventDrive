package controllers

import (
	"demo/src/core"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UploadProductImageController struct {
	fileService *core.FileService
}

func NewUploadProductImageController() *UploadProductImageController {
	fileService := core.NewFileService("uploads")
	return &UploadProductImageController{
		fileService: fileService,
	}
}

func (c *UploadProductImageController) Execute(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No se ha proporcionado ninguna imagen"})
		return
	}

	// Validar el tipo de archivo
	if !isValidImageType(file.Filename) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de archivo no válido. Solo se permiten imágenes (jpg, jpeg, png)"})
		return
	}

	// Guardar el archivo
	imagePath, err := c.fileService.SaveFile(file, "products")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Imagen subida correctamente",
		"path":    imagePath,
	})
}

func isValidImageType(filename string) bool {
	ext := filepath.Ext(filename)
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	return validExtensions[ext]
}
