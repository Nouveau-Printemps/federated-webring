# Fediring protocol

Documentation of the Fediring protocol.

## General

### Request

All requests must have the `Accept` http header set to `application/json`.

All requests use an endpoint starting with `/api/`.

If the method is not specified, the method used must be `GET`.

### Response

A response is always this type:
```json
{
  "status": 0,
  "message": "",
  "data": null
}
```
`status` is the status of the query.
`message` is a message describing the result.
`data` is the data linked with the query.

### Data

#### Server information

*HelloData* contains all important information about a server.
```json
{
  "name": "",
  "version": "",
  "application_name": "",
  "description": ""
}
```
`name` is the name of the webring.
`version` is the version of the protocol used.
`application_name` is the name of the application using Fediring.
`description` is the description of the webring.

#### Website information

*SiteData* contains all important information about a website.
```json
{
  "name": "",
  "url": "",
  "description": "",
  "type": ""
}
```
`name` is the name of the website.
`url` is the URL of the website. It must be a valid HTTP URL (e.g., `https://example.org/`)
`description` is the description of the website.
`type` is type of the website (e.g., blog, portfolio...). Each type must be separated by one space and must be in lowercase.

#### Federation

*FederationData* is used to send information about the federation.
```json
{
  "type": "",
  "message": "",
  "origin": "",
  "uuid": ""
}
```
`type` is the type of request.
`message` is the message containing information about the request.
`origin` is the origin server (e.g., `https://example.org/`).
`uuid` is the unique UUID of the request. 

## Server communication

### Information

To get all important information about the server, you have to query the endpoint `/api/hello`.

The type of data returned is *HelloData*.

### Sites

To query sites in the webring, you have to call the endpoint `/api/sites`.

The type of data returned is a list of "SiteData".

#### Random website

You can get a random website with the endpoint `/api/site-random`.

The type of data returned is *SiteData*.

#### Information about a website

You can get every information about a website with its name or its URL with the endpoint `/api/site`.
You have to set the parameter `url` with the URL or `name` with its name.
e.g., `/api/site?url=https://example.org/`

If the website is not found, the server send a 404.
If the request is mal formed (`url` or `name` is not set), the server send a 400.

The type of data returned is *SiteData*.

## Federation

To federate with a server, you have to send a federation request. 
Then, the server may accept it (or not) and send you a federation response.

When a server is federating with another server, they share their server list.
Each server must update others servers' list each four hours. 

### Checks validity of a request: *validity check*

To validate a request, the server has to send a new request to the endpoint `/api/federation-indox` of the origin server
using the `POST` method and the data "FederationData".
The type must indicate the origin request's type (e.g., `valid/federation-request`). 
This type is called *validity type*.
The message must contain the UUID of the first request.
Then, the server which sent the request has to validate it with an HTTP status 200 or invalidate it with an HTTP status 403.

e.g., Server A send the request to Server B
```json
{
  "type": "federation/request",
  "message": "I want to federate with you :)",
  "origin": "https://a.example.org/",
  "uuid": "19d2b596-48d4-42c5-8d3e-64c270e3e641"
}
```
Server B must validate this request and then send a request to Server A.
```json
{
  "type": "valid/federation-request",
  "message": "19d2b596-48d4-42c5-8d3e-64c270e3e641",
  "origin": "https://b.example.org/",
  "uuid": "51e19824-e84b-4943-8ff7-2a28b08ab2e2"
}
```
Server A must check the validity of this request by checking if the UUID in the message is valid and if it is linked with
a federation request to Server B.
If it is, it sends an HTTP status 200.
If not, it sends an HTTP status 403.

### Request federation

To federate with a server, use the endpoint `/api/federation-inbox` and send with the `POST` method the data "FederationData".
The type must be `federation/request`.
The message is the reason of the request.

The server which receives the request has to validate it (using the *validity check*).
The *validity type* is `valid/federation-request`.

If the request is not valid, the server has to ignore it.

### Response of a federation request

When the server accepts or not the request, it has to send a request to the origin server.
This request has the data "FederationData".
The type must be `federation/response`.
The message is the justification of the response.

The server which receives the request has to validate it (using the *validity check*).
The *validity type* is `valid/federation-response`.

If the request is not valid, the server has to ignore it.
