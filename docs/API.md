# Policy Server External API

## Purpose: 

The policy server API is used for creating, deleting and listing policies and tags.

## API Authentication
In order to communicate with the policy server API, a UAA oauth token with valid `network.admin` scope is required.

### Option 1: cf curl
By using the `cf curl`, the token is automatically passed in the request

Example
```sh
$ cf curl /networking/v0/external/policies
{"total_policies":2,"policies":[{"source":{...}]}
```

### Option 2: curl
When using curl the token must be explicitly provided in the `Authorization` header.

Example
```sh
$ curl http://api.bosh-lite.com/networking/v0/external/policies -H "Authorization: `cf oauth-token`"
{"total_policies":2,"policies":[{"source":{...}]}
```

## API Documentation

| Method | Path | Arguments | Request Body | Description|
| :----- | :--- | :-------- | :----------- | :----------- |
| GET | /networking/v0/external/policies | [see below](#get-networkingv0externalpolicies) | - | List Policies |
| POST | /networking/v0/external/policies | - | [see below](#post-networkingv0externalpolicies)| Create Policies |
| DELETE | /networking/v0/external/policies | - | [see below](#delete-networkingv0externalpolicies)| Delete Policies |
| GET | /networking/v0/external/tags | - | - | List all tag and `id` mappings |

Notes:
A unique tag is assigned to a policy_group_id when policies are created.

### GET /networking/v0/external/policies
#### Arguments:

[optionally] `id`: comma-separated policy_group_id values

Will return only the policies which include the given policy_group_id either as source id or destination id.

#### Response Body:

```json
{
  "total_policies": 2,
  "policies": [
    {
      "source": {
        "id": "1081ceac-f5c4-47a8-95e8-88e1e302efb5"
      },
      "destination": {
        "id": "38f08df0-19df-4439-b4e9-61096d4301ea",
        "protocol": "tcp",
        "port": 1234
      }
    },
    {
      "source": {
        "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36"
      },
      "destination": {
        "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36",
        "protocol": "tcp",
        "port": 1234
      }
    }
  ]
}
```

### POST /networking/v0/external/policies

#### Request Body:

```json
{
  "policies": [
    {
      "source": {
        "id": "1081ceac-f5c4-47a8-95e8-88e1e302efb5"
      },
      "destination": {
        "id": "38f08df0-19df-4439-b4e9-61096d4301ea",
        "protocol": "tcp",
        "port": 1234
      }
    },
    {
      "source": {
        "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36"
      },
      "destination": {
        "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36",
        "protocol": "tcp",
        "port": 1234
      }
    }
  ]
}
```

| Field | Required? | Description |
| :---- | :-------: | :------ |
| source.id | Y | The source `policy_group_id`
| destination.id | Y | The destination `policy_group_id`
| destination.protocol | Y | The protocol (tcp or udp)
| destination.port | Y | The destination port (1 - 65535)

#### Response Status Codes:
- 200 (successful)
- 400 (invalid request)

### DELETE /networking/v0/external/policies

#### Request Body:

```json
{
  "policies": [
    {
      "source": {
        "id": "1081ceac-f5c4-47a8-95e8-88e1e302efb5"
      },
      "destination": {
        "id": "38f08df0-19df-4439-b4e9-61096d4301ea",
        "protocol": "tcp",
        "port": 1234
      }
    },
    {
      "source": {
        "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36"
      },
      "destination": {
        "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36",
        "protocol": "tcp",
        "port": 1234
      }
    }
  ]
}
```

| Field | Required? | Description |
| :---- | :-------: | :------ |
| source.id | Y | The source `policy_group_id`
| destination.id | Y | The destination `policy_group_id`
| destination.protocol | Y | The protocol (tcp or udp)
| destination.port | Y | The destination port (1 - 65535)

#### Response Status Codes:
- 200 (successful)
- 400 (invalid request)

### GET /networking/v0/external/tags

#### Response Body:

```json
{
  "tags": [
    {
      "id": "1081ceac-f5c4-47a8-95e8-88e1e302efb5",
      "tag": "0001"
    },
    {
      "id": "308e7ef1-63f1-4a6c-978c-2e527cbb1c36",
      "tag": "0002"
    },
    {
      "id": "38f08df0-19df-4439-b4e9-61096d4301ea",
      "tag": "0003"
    }
  ]
}
```