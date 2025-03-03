package core

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileService struct {
	uploadPath string
}

func NewFileService(uploadPath string) *FileService {
	return &FileService{
		uploadPath: uploadPath,
	}
}

func (fs *FileService) SaveFile(file *multipart.FileHeader, folder string) (string, error) {
	// Asegúrate de que el directorio existe
	uploadDir := filepath.Join(fs.uploadPath, folder)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("error creando directorio: %v", err)
	}

	// Genera un nombre único para el archivo
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	filePath := filepath.Join(uploadDir, filename)

	// Abre el archivo fuente
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("error abriendo archivo: %v", err)
	}
	defer src.Close()

	// Crea el archivo destino
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creando archivo: %v", err)
	}
	defer dst.Close()

	// Copia el contenido
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("error guardando archivo: %v", err)
	}

	// Retorna la ruta relativa del archivo
	return fmt.Sprintf("/%s/%s", folder, filename), nil
}
