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
  "description": "",
  "update_endpoint": ""
}
```
`name` is the name of the webring.
`version` is the version of the protocol used.
`application_name` is the name of the application using Fediring.
`description` is the description of the webring.
`update_endpoint` is the endpoint to call when a federated server updates its information.

### Update

When a server updates its information, it has to send these modifications to others on the update endpoint.
This endpoint is specified by the `/api/hello` endpoint.

The data sent has the same type as the data sent by the `/api/hello` endpoint.
