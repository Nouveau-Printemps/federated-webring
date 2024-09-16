# Fediring protocol

Documentation on the Fediring protocol.

## General

### Request

All your requests must have the `Accept` http header set to `application/json`.

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

## Server communication

### Information

To get all important information about the server, you have to query the endpoint `/api/hello`.

The data returned is always this type:
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

### Update

When a server updates its information, it has to send these modifications to others on the endpoint `/api/update`.

The data sent has the same type as the data sent by the `/api/hello` endpoint.

### Sites

To query sites in the webring, you have to call the endpoint `/api/sites`.

The data sent is always this type:
```json
[
  {
    "name": "",
    "url": "",
    "description": "",
    "type": ""
  }
]
```
`name` is the name of the website.
`url` is the URL of the website. It must be a valid HTTP URL (e.g., `https://example.org/`)
`description` is the description of the website.
`type` is type of the website (e.g., blog, portfolio...). Each type must be separated by one space and must be in lowercase.

#### Random website

You can get a random website with the endpoint `/api/site-random`.

#### Information about a website

You can get every information about a website with its name or its URL with the endpoint `/api/site`.
You have to set the parameter `url` with the URL or `name` with its name.
e.g., `/api/site?url=https://example.org/`

If the website is not found, the server send a 404.
If the request is mal formed (`url` or `name` is not set), the server send a 400.
