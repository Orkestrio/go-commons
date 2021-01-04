package updater

import (
	"context"
	"fmt"
	"time"

	"github.com/machinebox/graphql"
)

func RunUpdater(dysonUrl string, selfName string, selfUrl string, selfSubUrl string) {
	fmt.Println("Updater init")
	// client := graphql.NewClient("http://dyson:4000/graphql")
	client := graphql.NewClient(dysonUrl)
	ctx := context.Background()

	// make a request
	// req := graphql.NewRequest(fmt.Sprintf(`
	// 		mutation register {
	// 			register(serviceName: "ServiceRegistry", apiEndpoint: "http://service-registry:3004/graphql", subEndpoint: "http://service-registry:3004/subs")
	// 	  	}
	// 	`, selfName, selfUrl, selfSubUrl))
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
