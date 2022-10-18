package main

import (
	"go-web/database"
	"go-web/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	/*
		//Veritabanı bağlantı testi -------------------
		start := time.Now()
		var a time.Duration
		timer := time.Now()
		var i int = 0
		for true {
			if time.Now().Sub(timer).Seconds() >= 10 {
				break
			}
			start = time.Now()
			database.DB.Find(&models.User{})
			a += time.Now().Sub(start)
			i++
		}
		fmt.Println("Ortalama : ", a/time.Duration(i))
		fmt.Println(i)
		//---------------------------------------------
	*/

	database.DB.AutoMigrate(&models.User{})

	app := fiber.New(fiber.Config{
		Concurrency: 16 * 640,
	})

	app.Static("/", "./public/")

	/*
		//Gzip, deflate ve brotli kullanarak sıkıştırma sağlar.
		app.Use(compress.New(compress.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Path() == "/dont_compress"
			},

			//LevelBestCompression: 8.9kb'lık veriyi 2.5kb'ya düşürdü 60-65ms de açıldı.
			//LevelBestSpeed: 8.9kb'lık veriyi 3.6kb'ye düşürdü 43-50ms de açıldı.
			//Eğer yüksek sıkıştırma istiyorsak LevelBestCompression kullanmalıyız.
			//Eğer yine yüksek ama LevelBestCompression'a göre daha az bir sıkıştırma olsun sayfamız hızlı yüklensin istersek
			//LevelBestSpeed'i tercih etmeliyiz.
			Level: compress.LevelBestSpeed,
		}))

				//Expirationda belirtilen süre içerisinde gelebilecek maksimumum isteği belirler. Eğer bu limit aşılırsa tekrardan
				//expirationda belirtilen süre kadar kullanıcıyı bekletir.
				app.Use(limiter.New(limiter.Config{
					Max:        5, //
					Expiration: 1 * time.Second,
					LimitReached: func(c *fiber.Ctx) error {
						return c.SendFile("./public/errors/429.html")
					},
				}))



			//Favicon belirler.
			app.Use(favicon.New(favicon.Config{
				File: "./public/img/favicon.ico",
			}))


		//İçerik değişmediyse web sunucusu yanıtı daha hızlı gönderebilmesi için önbelleğin daha verimli kullanılması ve
		//bant genişliğinin daha verimli kullanılması için kullanılır.
		app.Use(etag.New(etag.Config{
			Weak: true,
		}))
	*/
	app.Get("/users", func(c *fiber.Ctx) error {

		offset, _ := strconv.Atoi(c.Query("offset"))

		if offset <= 0 {
			offset = 0
		}

		err := c.Render("./public/index.html", fiber.Map{
			"Users":  models.SelectUser(50, offset),
			"Offset": offset,
		})
		return err
	})

	app.Get("/dont_compress", func(c *fiber.Ctx) error {

		offset, _ := strconv.Atoi(c.Query("offset"))

		if offset <= 0 {
			offset = 0
		}

		err := c.Render("./public/index.html", fiber.Map{
			"Users":  models.SelectUser(50, offset),
			"Offset": offset,
		})
		return err
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Listen(":5000")

}
