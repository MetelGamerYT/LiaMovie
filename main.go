package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed all:templates/*
var templates embed.FS

type Movie struct {
	Name          string
	UnderlineName string
	Thumbnail     string
}

var MoviePath = `Path/To/Movies`
var ThumbnailPath = `Path/To/Thumbnails`

func main() {
	fmt.Println("LiaMovie is starting...")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	templ := template.Must(template.ParseFS(templates, "templates/*"))
	router.SetHTMLTemplate(templ)
	router.Static("/css", "./css")
	router.Static("/thumbnails", ThumbnailPath)
	router.Static("/movies", MoviePath)

	router.GET("/", func(c *gin.Context) {
		movieList := getMovies()
		if len(movieList) > 10 {
			movieList = movieList[:10]
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Movies": movieList,
		})
	})

	router.GET("/search", func(c *gin.Context) {
		movieList := getMovies()
		c.HTML(http.StatusOK, "search.html", gin.H{
			"Movies": movieList,
		})
	})

	router.GET("/watch/:moviename", func(c *gin.Context) {
		movieName := c.Param("moviename")

		movieExtension, err := findMovieExtension(movieName)
		if err != nil {
			c.String(http.StatusNotFound, "Movie was not Found!")
			return
		}

		c.HTML(http.StatusOK, "watch.html", gin.H{
			"MovieName": replaceLines(movieName),
			"MoviePath": movieName + movieExtension,
		})
	})

	fmt.Println("LiaMovie started!")
	router.Run(":8080")
}

func findMovieExtension(movieName string) (string, error) {
	var movieExtension string
	err := filepath.Walk(MoviePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())) == movieName {
			// Filmdatei gefunden, extrahiere die Endung
			movieExtension = filepath.Ext(info.Name())
			return filepath.SkipDir // Überspringe den Rest des Ordners
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if movieExtension == "" {
		return "", fmt.Errorf("")
	}

	return movieExtension, nil
}

func getMovies() []Movie {
	var movies []Movie

	err := filepath.Walk(MoviePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isMovieFile(path) {
			movieName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			// fmt.Println("Ext. Name Movie: " + movieName)
			thumbnailPath := filepath.Join(ThumbnailPath, movieName+".jpg")
			// fmt.Println("Image Found: " + thumbnailPath)
			movies = append(movies, Movie{
				Name:          replaceLines(movieName),
				UnderlineName: movieName,
				Thumbnail:     thumbnailPath,
			})
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error, getting Movies: ", err)
	}

	return movies
}

func isMovieFile(filename string) bool {
	movieExtensions := []string{".mp4", ".mkv", ".avi", ".mov"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, movieExt := range movieExtensions {
		if ext == movieExt {
			return true
		}
	}
	return false
}

func replaceLines(inputstring string) string {
	newmovie := strings.Replace(inputstring, "_", " ", -1)
	newmovie = strings.Replace(newmovie, "-", " ", -1)

	return newmovie
}
