
# Delivery System Example

## Description
Within this project, you can handle deliveries in correct Delivery Points.

### Delivery Points and Shipments
~~~~
Shipment Types:
* Package: Refers to a single good.
* Sack: Refers to shipment type that consists more than one good.
~~~~

~~~~
There are three different delivery points in the system.
* Branch: Only packages can be unloaded. Sacks and packages in sacks can not be unloaded.
* Distribution Center: Sacks, packages in sacks and packages that are not assigned to a sack can be unloaded.
* Transfer Center: Only sacks and packages in sacks can be unloaded.
~~~~
## Folder Structure

- **/cfg**: Database connection and database initialization
- **/cmd**: Root file of project
- **/docs**: Swagger documentation file
- **/domain**: Contains entity files, enums, interfaces etc.
- **/internal**: Repositories, response and services
## Service Design

![service design](https://i.hizliresim.com/axluc7d.jpg)

## Technologies

**Development**: Golang, MongoDB, Docker

**API Documentation**: Swagger



## Endpoints

#### Check API Health

```curl
 GET /
 
 curl --location --request GET 'http://localhost:3000/'
```



#### Handle Deliveries

```curl
 POST /deliver

  curl --location --request POST 'http://localhost:3000/deliver' \
--header 'Content-Type: application/json' \
--data-raw '{
  "vehicle": "34 TL 34",
  "route": [
    {
      "deliveryPoint": 1,
      "deliveries": [
        {"barcode": "P7988000121"},
        {"barcode": "P7988000122"},
        {"barcode": "P7988000123"},
        {"barcode": "P8988000121"},
        {"barcode": "C725799"}
      ]
    },
    {
      "deliveryPoint": 2,
      "deliveries": [
        {"barcode": "P8988000123"},
        {"barcode": "P8988000124"},
        {"barcode": "P8988000125"},
        {"barcode": "C725799"}
      ]
    },
    {
      "deliveryPoint": 3,
      "deliveries": [
        {"barcode": "P9988000126"},
        {"barcode": "P9988000127"},
        {"barcode": "P9988000128"},
        {"barcode": "P9988000129"},
        {"barcode": "P9988000130"}
      ]
    }
  ]
}'
```

## Run The Tests

To check test results
```bash
  go test ./...
```
To check coverage rate

```bash
  go test -v -coverprofile cover.out ./...
  go tool cover -html="cover.out" -o cover.html
```
Then you can check html file and see uncovered lines

## Run The Application

Set variables in .env file
~~~~
MONGO_URI=
Database=
PORT=
APP_ENV=
~~~~

### Development

```bash
  $ cd workspace

  $ git clone PROJECT_URL

  $ cd delivery-system
  
  $ go mod tidy
  
  $ go run .
```
### Build
```bash
  $ go build -o /app
```

### Docker
```bash
  $ cd workspace

  $ git clone PROJECT_URL

  $ cd delivery-system
  
  $ docker image build -t delivery-system.
  
  $ docker compose up -d
```




  
## Thanks

- Thank you for this challenging opportunity

  