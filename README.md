# Finch and managing integrations is easy
<p align="center">  
    <img src="https://github.com/unm4sked/finch/blob/main/static/finch.webp" width=40% height=40%>
</p>

**Finch** is an application for managing configurations and integrations along with validation and verification.

Normally you probably have these credentials or configurations stored somewhere in db or other type storage, entered manually by someone, but what if your customer wants to update them? What if he made some kind of mistake during the update and the integration is messed up?

Finch is an application that your customer has access to can complete the data for integration validate it and verify it through an API request. When the integration is ready your backend system can fetch the data for integration, you are assured that it is correct checked and fresh.

## How to run locally
1. install packages using `go mod download`
1. create postgres db locally using docker-compose `docker-compose up`
1. run service using `go run cmd/api/main.go`
