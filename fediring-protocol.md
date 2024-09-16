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