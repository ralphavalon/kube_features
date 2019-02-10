package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	input, err := ioutil.ReadFile("./swagger-ui/swagger-base.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output := input

	swaggerAPIHost := os.Getenv("API_HOST")

	if len(swaggerAPIHost) != 0 {
		fmt.Println("Setting Swagger API Host to: " + swaggerAPIHost)

		oldHost := `"host": "localhost:8081"`
		newHost := fmt.Sprintf(`"host": "%s"`, swaggerAPIHost)

		output = bytes.Replace(input, []byte(oldHost), []byte(newHost), -1)
	}

	if err = ioutil.WriteFile("./swagger-ui/swagger.json", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
