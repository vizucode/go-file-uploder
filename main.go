package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Configurasi folder tempat menyimpan file gambar
const uploadFolder = "./uploads"

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Endpoint untuk menerima upload file gambar
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Ambil file dari form
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("File upload error")
		}

		// Buat nama file unik dengan menambahkan timestamp
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

		// Path lengkap untuk menyimpan file
		filePath := filepath.Join(uploadFolder, filename)

		// Simpan file ke folder "uploads"
		err = c.SaveFile(file, filePath)
		if err != nil {
			fmt.Println("Error saving file:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
		}

		// Kembalikan URL untuk mengakses file yang di-upload
		fileURL := fmt.Sprintf("http://%s/uploads/%s", c.Hostname(), filename)
		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
			"url":     fileURL,
		})
	})

	// Endpoint untuk serve file gambar yang sudah di-upload
	app.Static("/uploads", uploadFolder)

	// Jalankan server
	port := ":3055"
	fmt.Printf("Server running at http://localhost%s\n", port)
	log.Fatal(app.Listen(port))
}
