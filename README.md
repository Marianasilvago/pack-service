# Pack Service

This allows to set up pack sizes and calculate packs for items.

## Pre requisites

- Docker
- Golang v1.21+
 
 
## Running App 

1. Build and run the app container.

`make app`

2. Inspect logs using docker 

`docker logs pack-svc-go -f`

3. Run tests in local docker using

`make test`

4. Run functional tests
   - Start the app using step 1
   - Trigger the dockerized functional tests
    `make functional-test`

## Live Website

This is the hosted link. However since it's free, it might take some to come up if it's unused.

https://pack-service.onrender.com/

Free instances spin down with inactivity

