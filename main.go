package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type Clipboard struct {
	Date   string `json:"date"`
	Digest string `json:"digest"`
	Short  string `json:"short"`
	Size   int    `json:"size"`
	URL    string `json:"url"`
	Status string `json:"status"`
	UUID   string `json:"uuid"`
}

var clipboards map[string]Clipboard

func init() {
	clipboards = make(map[string]Clipboard)
}

func main() {
	e := echo.New()

	e.POST("/clips", CreateClipboard)
	e.GET("/clips/:uuid", GetClipboard)
	e.PUT("/clips/:uuid", UpdateClipboard)
	fmt.Println("fab")
	e.DELETE("/clips/:uuid", DeleteClipboard)

	e.Logger.Fatal(e.Start(":8080"))
}

func CreateClipboard(c echo.Context) error {
	content := c.FormValue("content")
	uuid := fmt.Sprintf("%d", time.Now().UnixNano())
	digest := fmt.Sprintf("%x", content)
	short := uuid[:4]
	size := len(content)
	url := "http://example.com/" + short

	clipboards[uuid] = Clipboard{
		Date:   time.Now().Format(time.RFC3339),
		Digest: digest,
		Short:  short,
		Size:   size,
		URL:    url,
		Status: "created",
		UUID:   uuid,
	}

	jsonData, _ := json.MarshalIndent(clipboards[uuid], "", "  ")
	// if err != nil {
    // // 处理错误
    // return err
	// }
	return c.String(http.StatusOK, string(jsonData))
}

func GetClipboard(c echo.Context) error {
	uuid := c.Param("uuid")
	if clipboard, ok := clipboards[uuid]; ok {
		clipboard.Status = "found"
		jsonData, _ := json.MarshalIndent(clipboard, "", "  ")
		
		return c.String(http.StatusOK, string(jsonData))
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Clipboard not found"})
}

func UpdateClipboard(c echo.Context) error {
	uuid := c.Param("uuid")
	if _, ok := clipboards[uuid]; ok {
		content := c.FormValue("content")
		digest := fmt.Sprintf("%x", content)
		size := len(content)

		clipboards[uuid] = Clipboard{
			Date:   time.Now().Format(time.RFC3339),
			Digest: digest,
			Size:   size,
			Status: "updated",
			UUID:   uuid,
		}

		jsonData, _ := json.MarshalIndent(clipboards[uuid], "", "  ")
		
		return c.String(http.StatusOK, string(jsonData))

		// return c.JSON(http.StatusOK, clipboards[uuid])
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Clipboard not found"})
}

func DeleteClipboard(c echo.Context) error {
	uuid := c.Param("uuid")
	if _, ok := clipboards[uuid]; ok {
		delete(clipboards, uuid)

		jsonData, _ := json.MarshalIndent(map[string]string{"status": "deleted\n" , "uuid": uuid}, "", "  ")
		
		return c.String(http.StatusOK, string(jsonData))

		// return c.JSON(http.StatusOK, map[string]string{"status": "deleted\n" , "uuid": uuid})
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Clipboard not found"})
}
