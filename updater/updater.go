package updater

import (
	"context"
	"fmt"
	"time"

	"github.com/machinebox/graphql"
)

func RunUpdater(dysonUrl string, selfName string, selfUrl string, selfSubUrl string) {
	fmt.Println("Updater initiated")
	client := graphql.NewClient(dysonUrl)
	ctx := context.Background()

	req := graphql.NewRequest(fmt.Sprintf(`
			mutation register {
				register(serviceName: "%s", apiEndpoint: "%s", subEndpoint: "%s")
		  	}
		`, selfName, selfUrl, selfSubUrl))

	if err := client.Run(ctx, req, nil); err != nil {
		fmt.Println("Error: ", err)
	}

	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := client.Run(ctx, req, nil); err != nil {
					fmt.Println("Error: ", err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
