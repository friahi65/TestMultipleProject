# API core

## Sample queries and expectations

Get a list of existing/pre-inserted objects:
```sh
curl -sX GET http://127.0.0.1:8181/v1.0/employee | jq
```
Expected output:
```
{
  "result": [
    {
      "age": 42,
      "id": "6f8c0464-3054-445f-a7b6-3744bd4709a3",
      "location": "SiteA",
      "name": "John"
    }
  ],
  "status_code": 0
}
```

Request a specified by ```id``` object:
```sh
curl -sX GET http://127.0.0.1:8181/v1.0/employee/6f8c0464-3054-445f-a7b6-3744bd4709a3 | jq
```
Expected output:
```
{
  "result": [
    {
      "age": 42,
      "id": "6f8c0464-3054-445f-a7b6-3744bd4709a3",
      "location": "SiteA",
      "name": "John"
    }
  ],
  "status_code": 0
}

```

Insert another object:
```sh
curl -sX POST http://127.0.0.1:8181/v1.0/employee -d '{"name":"Name_111","location":"SiteA"}' | jq
```
Expected output:
```
{
  "result": {
    "id": "f0ac0eee-c235-4ec1-97a7-232e81033f84"
  },
  "status_code": 0
}
```

Get a list of objects:
```sh
curl -sX GET http://127.0.0.1:8181/v1.0/employee | jq
```
Expected output:
```
{
  "result": [
    {
      "age": 42,
      "id": "6f8c0464-3054-445f-a7b6-3744bd4709a3",
      "location": "SiteA",
      "name": "John"
    },
    {
      "age": 42,
      "id": "f0ac0eee-c235-4ec1-97a7-232e81033f84",
      "location": "SiteA",
      "name": "Name_111"
    }
  ],
  "status_code": 0
}
```

Find object by exact field values:
```sh
curl -sX GET http://127.0.0.1:8181/v1.0/employee?name="Name_111" | jq
```
Expected output:
```
{
  "result": [
    {
      "age": 42,
      "id": "f0ac0eee-c235-4ec1-97a7-232e81033f84",
      "location": "SiteA",
      "name": "Name_111"
    }
  ],
  "status_code": 0
}

```

Update a specified by ```id``` object:
```sh
curl -sX PUT http://127.0.0.1:8181/v1.0/employee/f0ac0eee-c235-4ec1-97a7-232e81033f84 -d '{"name":"Name_111_Updated","location":"SiteB"}'
```
Expected output:
```
{
   "result": {
      "id": "f0ac0eee-c235-4ec1-97a7-232e81033f84"
   },
   "status_code": 0
}
```

Request a specified, just updated object:
```sh
curl -sX GET http://127.0.0.1:8181/v1.0/employee/f0ac0eee-c235-4ec1-97a7-232e81033f84 | jq
```
Expected output:
```
{
  "result": [
    {
      "age": 42,
      "id": "f0ac0eee-c235-4ec1-97a7-232e81033f84",
      "location": "SiteB",
      "name": "Name_111_Updated"
    }
  ],
  "status_code": 0
}
```

Delete a specified object:
```sh
curl -sX DELETE http://127.0.0.1:8181/v1.0/employee/f0ac0eee-c235-4ec1-97a7-232e81033f84 | jq
```
Expected output:
```
{
   "result": {
      "id": "f0ac0eee-c235-4ec1-97a7-232e81033f84"
   },
   "status_code": 0
}
```

Get a list of objects to make sure the specified one is really deleted:
```sh
curl -sX GET http://127.0.0.1:8181/v1.0/employee | jq
```
Expected output:
```
{
  "result": [
    {
      "age": 42,
      "id": "6f8c0464-3054-445f-a7b6-3744bd4709a3",
      "location": "SiteA",
      "name": "John"
    }
  ],
  "status_code": 0
}

```

Try to insert incorrect object:
```sh
curl -sX POST http://127.0.0.1:8181/v1.0/employee -d '{"name":"Name_222","location":"Wrong"}' | jq
```
Expected output:
```
{
  "result": {},
  "status_code": 1,
  "status_detail": {
    "message": "invalid enum Wrong. Valid values: [HQ SiteA SiteB]"
  }
}
```

## Support for TenantID:

To execute a request in a context of specific customer (or tenant), put the related TenantID value into HTTP request headers:

```sh
curl --header "Tenant-Id: user1" -sX POST http://127.0.0.1:8181/v1.0/employee -d '{"name":"Name_222","location":"Wrong"}' | jq
```

## Supported make targets:

Fetching and setting up the required dependencies through glide:
```make```

Cleanup:
```make clean```


The following test targets must be run with the following order:

Building and running the tests: 
```make tests```

Generating coverage statistics:
```make test-coverage```

Getting the test output in XML format:
```make test-xml```
