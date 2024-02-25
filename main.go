package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Platform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Game struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	PlatformID  int     `json:"platform_id"`
	Url         string  `json:"url"`
	Price       float64 `json:"price"`
	ReleaseDate string  `json:"release_date"`
}

// temp data array
var platforms = []Platform{
	{ID: 1, Name: "PlayStation 4", Url: "https://www.playstation.com/en-us/"},
	{ID: 2, Name: "Xbox One", Url: "https://www.xbox.com/en-US/"},
	{ID: 3, Name: "Nintendo Switch", Url: "https://www.nintendo.com/"},
	{ID: 4, Name: "Windows Store", Url: "https://www.microsoft.com/en-us/store/b/xbox?activetab=pivot:overviewtab"},
	{ID: 5, Name: "Steam", Url: "https://store.steampowered.com/"},
	{ID: 6, Name: "GOG", Url: "https://www.gog.com/"},
	{ID: 7, Name: "Origin", Url: "https://www.origin.com/"},
	{ID: 8, Name: "Uplay", Url: "https://store.ubi.com/us/"},
	{ID: 9, Name: "Epic Games Store", Url: "https://www.epicgames.com/store/en-US/"},
	{ID: 10, Name: "Battle.net", Url: "https://us.shop.battle.net/en-us"},
	{ID: 11, Name: "itch.io", Url: "https://itch.io/"},
	{ID: 12, Name: "Humble Store", Url: "https://www.humblebundle.com/store"},
	{ID: 13, Name: "Google Play", Url: "https://play.google.com/store?hl=en_US"},
	{ID: 14, Name: "App Store", Url: "https://www.apple.com/ios/app-store/"},
}

var games = []Game{
	{ID: 1, Name: "The Last of Us Part II", PlatformID: 1, Url: "https://www.playstation.com/en-us/games/the-last-of-us-part-ii-ps4/", Price: 59.99, ReleaseDate: "2020-06-19"},
	{ID: 2, Name: "Cyberpunk 2077", PlatformID: 2, Url: "https://www.cyberpunk.net/us/en/", Price: 59.99, ReleaseDate: "2020-09-17"},
	{ID: 3, Name: "The Legend of Zelda: Breath of the Wild", PlatformID: 3, Url: "https://www.zelda.com/breath-of-the-wild/", Price: 59.99, ReleaseDate: "2017-03-03"},
	{ID: 4, Name: "Halo: The Master Chief Collection", PlatformID: 4, Url: "https://www.xbox.com/en-US/games/halo-the-master-chief-collection", Price: 39.99, ReleaseDate: "2019-12-03"},
	{ID: 5, Name: "Half-Life: Alyx", PlatformID: 5, Url: "https://store.steampowered.com/app/546560/HalfLife_Alyx/", Price: 59.99, ReleaseDate: "2020-03-23"},
	{ID: 6, Name: "The Witcher 3: Wild Hunt", PlatformID: 6, Url: "https://www.gog.com/game/the_witcher_3_wild_hunt", Price: 39.99, ReleaseDate: "2015-05-19"},
	{ID: 7, Name: "Star Wars Jedi: Fallen Order", PlatformID: 7, Url: "https://www.origin.com/usa/en-us/store/star-wars/star-wars-jedi-fallen-order", Price: 59.99, ReleaseDate: "2019-11-15"},
	{ID: 8, Name: "Assassin's Creed Odyssey", PlatformID: 8, Url: "https://store.ubi.com/us/assassins-creed-odyssey/5b06a3984e0165fa45ffdc9d.html", Price: 59.99, ReleaseDate: "2018-10-05"},
	{ID: 9, Name: "Fortnite", PlatformID: 9, Url: "https://www.epicgames.com/fortnite/en-US/home", Price: 0.00, ReleaseDate: "2017-07-25"},
	{ID: 10, Name: "Call of Duty: Modern Warfare", PlatformID: 10, Url: "https://us.shop.battle.net/en-us/product/call-of-duty-modern-warfare", Price: 59.99, ReleaseDate: "2019-10-25"},
	{ID: 11, Name: "Celeste", PlatformID: 11, Url: "https://mattmakesgames.itch.io/celeste", Price: 19.99, ReleaseDate: "2018-01-25"},
	{ID: 12, Name: "Hollow Knight", PlatformID: 12, Url: "https://www.humblebundle.com/store/hollow-knight", Price: 14.99, ReleaseDate: "2017-02-24"},
	{ID: 13, Name: "Minecraft", PlatformID: 13, Url: "https://play.google.com/store/apps/details?id=com.mojang.minecraftpe&hl=en_US", Price: 6.99, ReleaseDate: "2011-10-07"},
	{ID: 14, Name: "Among Us", PlatformID: 14, Url: "https://apps.apple.com/us/app/among-us/id1351168404", Price: 0.00, ReleaseDate: "2018-06-15"},
	{ID: 15, Name: "God of War", PlatformID: 1, Url: "https://www.playstation.com/en-us/games/god-of-war-ps4/", Price: 19.99, ReleaseDate: "2018-04-20"},
	{ID: 16, Name: "Red Dead Redemption 2", PlatformID: 2, Url: "https://www.rockstargames.com/reddeadredemption2/", Price: 59.99, ReleaseDate: "2018-10-26"},
	{ID: 17, Name: "Super Mario Odyssey", PlatformID: 3, Url: "https://www.nintendo.com/games/detail/super-mario-odyssey-switch/", Price: 59.99, ReleaseDate: "2017-10-27"},
	{ID: 18, Name: "Forza Horizon 4", PlatformID: 4, Url: "https://www.xbox.com/en-US/games/forza-horizon-4", Price: 59.99, ReleaseDate: "2018-10-02"},
	{ID: 19, Name: "Helldivers", PlatformID: 5, Url: "https://store.steampowered.com/app/394510/HELLDIVERS/", Price: 19.99, ReleaseDate: "2015-12-07"},
	{ID: 20, Name: "God of War Ragnarok", PlatformID: 1, Url: "https://www.playstation.com/en-us/games/god-of-war-ragnarok-ps5/", Price: 59.99, ReleaseDate: "2021-12-31"},
}

// GetPlatforms Get all platforms
func GetPlatforms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, platforms)
}

// GetGames Get all games
func GetGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

// GetGamesByPlatformGET Get games by platform
func GetGamesByPlatformGET(c *gin.Context) {
	platformID, _ := strconv.Atoi(c.Param("id"))
	var platformGames []Game

	for _, game := range games {
		if game.PlatformID == platformID {
			platformGames = append(platformGames, game)
		}
	}
	c.IndentedJSON(http.StatusOK, platformGames)
}

// GetGamesByPlatformPOST Get games by platform to gett all games for a specific platform implemented as a POST request
func GetGamesByPlatformPOST(c *gin.Context) {
	var requestData struct {
		PlatformID int `json:"platform_id"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	var platformGames []Game
	for _, game := range games {
		if game.PlatformID == requestData.PlatformID {
			platformGames = append(platformGames, game)
		}
	}
	c.IndentedJSON(http.StatusOK, platformGames)
}

// 4. Get games by price

// 5. Get games by release date

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.GET("/platforms", GetPlatforms)
	router.GET("/games", GetGames)
	router.GET("/platforms/:id/games", GetGamesByPlatformGET)
	router.POST("/games", GetGamesByPlatformPOST)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
