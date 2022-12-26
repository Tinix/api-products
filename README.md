HTTP requests to the different endpoint URLs using a tool like curl or a browser.

For example, you can use curl to make a GET request to the /products endpoint to get a list of products:

```bash
curl http://localhost:8080/products
```

This would return a response like this:

```json
[  {"id":"1","name":"Product 1","price":9.99},  {"id":"2","name":"Product 2","price":19.99},  {"id":"3","name":"Product 3","price":29.99}]
```

You can also use curl to make a GET request to a specific product endpoint, like /products/1, to get a single product:

```bash
curl http://localhost:8080/products/1
```

This would return a response like this:

```json
{"id":"1","name":"Product 1","price":9.99}
```

You can use curl to make a POST request to the /products endpoint to create a new product, like this:

```bash
curl -X POST -d '{"name":"New Product","price":49.99}' http://localhost:8080/products
```

This would return a response like this:

```bash
{"id":"4","name":"New Product","price":49.99}
```

You can use curl to make a PUT request to a specific product endpoint, like /products/4, to update an existing product, like this:

```bash
curl -X PUT -d '{"id":"4","name":"Updated Product","price":99.99}' 
http://localhost:8080/products/4
```
This would return a response like this:

```bash
{"id":"4","name":"Updated Product","price":99.99}
```

Finally, you can use curl to make a DELETE request to a specific product endpoint, like /products/1, to delete an existing product:

```bash
curl -X DELETE http://localhost:8080/products/1
```

This would return a response with a status code of 204 No Content, indicating that the product was successfully deleted.
