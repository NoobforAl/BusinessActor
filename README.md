# Business Actor

This Project For WeConnect company! ( interview question )

## How work ?

At startup, this program first checks in mongodb whether it has data or not.  
If not, the program reads the csv file and saves the data to the mongodb database.  
The application then listens on your custom port to receive requests.

## How Run ?

Two way have you can run this project:

1 - run docker compose
> docker compose up

And server run in port 8080.

For more environment option see this table:

| environment |                              what's for                                            |
| ----- | :--------------------------------------------------------------------------------------: |
|  DSN        | your mongodb database url  |
|  CSV_PATH   | csv data path file |
| LISTEN_PORT | your port when you want run app |
| LISTEN_IP   | your ip when you want run app |
| GIN_MODE    | gin app mod run (debug, release) |

You can setup this value in docker-compose.yml

2 - Run with golang and run mongodb:

First setup database mongodb.
Secund Install dependency:
> go mod tidy

Setup .env file like .env.example.
And run code :
> go run main.go

You can run with air for debug project:
> air

How install air see this [Doc](https://github.com/cosmtrek/air).

## Help Api

For request to rest api you need postman and you can load postman file in folder postman ( src/postman ).

### Get Business Actor

This is very simple query for get data with objectId.  

Endpoint: (GET)
> 0.0.0.0:8080/api/get/:id

id(string): is a hex objectId.

Response: (200 OK)

```json
{
    "id": "6488b3434ecb572ccabb6a59",
    "series_reference": "BDCQ.SF1AA2CA",
    "period": "2016-09-01T00:00:00Z",
    "data_value": 1070.874,
    "suppressed": false,
    "status": "F",
    "units": "Dollars",
    "magnitude": 6,
    "subject": "Business Data Collection - BDC",
    "group": "Industry by financial variable (NZSIOC Level 2)",
    "series_title_1": "Sales (operating income)",
    "series_title_2": "Forestry and Logging",
    "series_title_3": "Current prices",
    "series_title_4": "Unadjusted",
    "series_title_5": ""
}
```

### Get Many Business Actor

This end point for get data with range size.  
for example: get page 2 and size 10, return range of data 10 to 20.

Endpoint: (GET)
> 0.0.0.0:8080/api/getMany?size=< num >&page=< num >

size(int64): how many fetch data on database.

page(int64): which one page.

Response: (200 OK)

```json
[
    {
        "id": "6488b3434ecb572ccabb6a59",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2016-09-01T00:00:00Z",
        "data_value": 1070.874,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a5b",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2017-06-01T00:00:00Z",
        "data_value": 1233.7,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a5c",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2017-09-01T00:00:00Z",
        "data_value": 1282.436,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a5d",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2017-12-01T00:00:00Z",
        "data_value": 1290.82,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a5e",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2018-03-01T00:00:00Z",
        "data_value": 1412.007,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a5f",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2018-06-01T00:00:00Z",
        "data_value": 1488.055,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a60",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2018-09-01T00:00:00Z",
        "data_value": 1497.678,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a61",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2018-12-01T00:00:00Z",
        "data_value": 1570.507,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a62",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2019-03-01T00:00:00Z",
        "data_value": 1393.749,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    },
    {
        "id": "6488b3434ecb572ccabb6a63",
        "series_reference": "BDCQ.SF1AA2CA",
        "period": "2019-06-01T00:00:00Z",
        "data_value": 1517.143,
        "suppressed": false,
        "status": "F",
        "units": "Dollars",
        "magnitude": 6,
        "subject": "Business Data Collection - BDC",
        "group": "Industry by financial variable (NZSIOC Level 2)",
        "series_title_1": "Sales (operating income)",
        "series_title_2": "Forestry and Logging",
        "series_title_3": "Current prices",
        "series_title_4": "Unadjusted",
        "series_title_5": ""
    }
]
```

### Create Business Actor

This endpoint for create new Business Actor.

Endpoint: (POST)
> 0.0.0.0:8080/api/create

You need send json data like this:

```json
{
    "series_reference": "BDCQ.SF1AA2CA",
    "period": "2016-06-01T00:00:00Z",
    "data_value": 1116.386,
    "suppressed": false,
    "status": "F",
    "units": "Dollars",
    "magnitude": 6,
    "subject": "Business Data Collection - BDC",
    "group": "Industry by financial variable (NZSIOC Level 2)",
    "series_title_1": "Sales (operating income)",
    "series_title_2": "Forestry and Logging",
    "series_title_3": "Current prices",
    "series_title_4": "Unadjusted",
    "series_title_5": "test3343"
}
```

Response: (200 OK)

```json
{
    "id": "6488866fcadba500478d8067",
    "series_reference": "BDCQ.SF1AA2CA",
    "period": "2016-06-01T00:00:00Z",
    "data_value": 1116.386,
    "suppressed": false,
    "status": "F",
    "units": "Dollars",
    "magnitude": 6,
    "subject": "Business Data Collection - BDC",
    "group": "Industry by financial variable (NZSIOC Level 2)",
    "series_title_1": "Sales (operating income)",
    "series_title_2": "Forestry and Logging",
    "series_title_3": "Current prices",
    "series_title_4": "Unadjusted",
    "series_title_5": "test3343"
}
```

### Update Business Actor

This endpoint for update Business Actor data.  
Note: if your send empty filed on request, filed empty save on database. (be carful)

Endpoint: (put)
> 0.0.0.0:8080/api/update/:id

id(string): is a hex objectId.

You need send json data like this:

```json
{
    "series_reference": "BDCQ.SF1AA2CA",
    "period": "2016-06-01T00:00:00Z",
    "data_value": 1116.386,
    "suppressed": false,
    "status": "F",
    "units": "Dollars",
    "magnitude": 6,
    "subject": "Business Data Collection - BDC",
    "group": "Industry by financial variable (NZSIOC Level 2)",
    "series_title_1": "Sales (operating income)",
    "series_title_2": "Forestry and Logging",
    "series_title_3": "Current prices",
    "series_title_4": "Unadjusted",
    "series_title_5": "test3343"
}
```

Response: (200 OK)

```json
{
    "id": "6488866fcadba500478d8067",
    "series_reference": "BDCQ.SF1AA2CA",
    "period": "2016-06-01T00:00:00Z",
    "data_value": 1116.386,
    "suppressed": false,
    "status": "F",
    "units": "Dollars",
    "magnitude": 6,
    "subject": "Business Data Collection - BDC",
    "group": "Industry by financial variable (NZSIOC Level 2)",
    "series_title_1": "Sales (operating income)",
    "series_title_2": "Forestry and Logging",
    "series_title_3": "Current prices",
    "series_title_4": "Unadjusted",
    "series_title_5": "test3343"
}
```

### Delete Business Actor

This endpoint for delete on record in database.

Endpoint: (delete)
> 0.0.0.0:8080/api/delete/:id

id(string): is a hex objectId.

Response: (200 OK)

```json
{
    "id": "",
    "series_reference": "",
    "period": "0001-01-01T00:00:00Z",
    "data_value": 0,
    "suppressed": false,
    "status": "",
    "units": "",
    "magnitude": 0,
    "subject": "",
    "group": "",
    "series_title_1": "",
    "series_title_2": "",
    "series_title_3": "",
    "series_title_4": "",
    "series_title_5": ""
}
```

### Errors

If your program get a error, program response like this:

Status Code : (404, 400, 500)

```json
{
    "detail": "Error Message."
}
```
