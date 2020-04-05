# goscorer
A HTTP scorer in GoLang using GIN and PMML files

This package illustrates the use of Golang to provide scores using a random forest.

The way it works:
<li> Offline, the IRIS dataset was used to create a random forest in R and is then exported to PMML and stored under the  `/static` path. The PMML is provided.
<li> The Golang HTTP server preloads the model when the service starts (command: `go run main.go`) 
<li>  <a href="https://github.com/gin-gonic/gin">  go-Gin </a> receives the GET request, parses it, and returns a score (probability to is a Setosa)

## Try it out:

```
curl --location --request GET 'http://localhost:8085/testing' \
--header 'Content-Type: text/plain' \
--data-raw '{
"sepal_length_cm": 5.1
,"sepal_width_cm": 3.5
,"petal_length_cm":5
,"petal_width_cm": 0.2
}'
```

should return 0.55

## Generating PMML files:

You can use `generate_pmml.R` to generate PMML objects. The output is stored in `static`.