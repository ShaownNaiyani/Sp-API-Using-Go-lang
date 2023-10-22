package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	godotenv.Load()

	c, err := spapi.NewClient(
		&spapi.Config{
			RefreshToken:  os.Getenv("refresh_token"),
			ClientID:      os.Getenv("lwa_app_id"),
			ClientSecret:  os.Getenv("lwa_client_secret"),
			AccessKey:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretKey:     os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Endpoint:      "sellingpartnerapi-na.amazon.com",
			MarketplaceID: "A2EUQ1WTGCTBG2",
			Region:        "us-east-1",
		},
	)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	res, err := c.GetCatalogItem(context.Background(), "B01N0O7QKJ")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	endTime := time.Now()
	responseTime := endTime.Sub(startTime).Seconds()
	fmt.Printf("Response time: %.2f seconds\n", responseTime)
}
