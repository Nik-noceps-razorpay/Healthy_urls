package main

import (
	"DbConn"
	"Models"
	"Routes"
	"fmt"
	"github.com/robfig/cron"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {

	//open a db connection

	//  xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-- recurring health checkup every minute the server is running --xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	c := cron.New()
	c.AddFunc("*/1 * * * *", healthCheckUp)
	c.Start()

}

// ------------------------------------------------------------ setting up Routes ----------------------------------------------------------------------
func main() {

	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-- Initializing database --xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	DbConn.InitDB()

	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-- starting a server that accepts url via postman for regular health monitoring --xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	Routes.Router()

	defer DbConn.Db.Close()

	fmt.Println("Connected to database")
}


// -------------------------------------------------- concurrent health checkups of urls --------------------------------------------------

func healthCheckUp() {
	var urls []Models.Url

	DbConn.Db.Find(&urls)

	if len(urls) <= 0 {
		fmt.Println( "No urls found!, kindly insert some urls.")
		return
	}
	ct := len(urls)
	fmt.Printf("type %T value %d ",ct,ct)
	wg.Add(ct)
	for i := 0 ; i < len(urls) ; i++ {

		go pingUrl(urls[i])
	}


}


//-------------------------------------------------- updates status of health checkups to the url_hits table ------------------------------

func pingUrl(url Models.Url) {
	defer wg.Done()


	for i := 0; i < url.Failure_threshold ; i++ {

		var hit Models.UrlHits

		hit.UrlId = url.ID
		hit.Hit_number = i + 1

		client := http.Client{
			Timeout: time.Duration(url.Crawl_timeout) * time.Second,
		}

		resp, err := client.Get(url.UrlName)

		if err != nil {
			fmt.Println("Bhai error aarela hai, code dhang se dekh, error nichu likhela hai ")
			fmt.Println(err)
			hit.Status = 2
			DbConn.Db.Save(&hit)
			time.Sleep(time.Duration(url.Frequency) * time.Second)
		} else {

			if resp.StatusCode >= 200 && resp.StatusCode < 300 {

				fmt.Println("site to chalreli hai, database mai update karne ko bheja hai code ko, dekh lena")

				hit.Status = 1

				DbConn.Db.Save(&hit)

				break

			} else {

				hit.Status = 0

				DbConn.Db.Save(&hit)

				time.Sleep(time.Duration(url.Frequency) * time.Second)
			}

		}
	}
}
