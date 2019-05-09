# Advertising System

This is a final project of Cloud Computing Course.

# How to run

# Introduction

![Framework](./img/1.png)

# APIs

## Ad Delivery System

The services are running on port 7000

### 1. Create an advertiser account

| Key  | Value |
| ---- | ----- |
| url  | /ad_user/create  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| username | string |
| password1 | string |
| password2 | string |

#### Example
Request POST http://localhost:7000/ad_user/create
```
{
	"username":"test user",
	"password1": "123",
	"password2": "123"
}
```
Response 
```
{
    "code": 0,
    "message": "success",
    "data": "5cd473ccc4cc785f0c3ce9b1"
}
```
### 2. Create an advertising plan

| Key  | Value |
| ---- | ----- |
| url  | /ad_plan/create  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| user_id | string |
| name | string |
| start_time | int |
| end_time | int |

#### Example
Request POST http://localhost:7000/ad_plan/create
```
{
	"user_id":"5cd473ccc4cc785f0c3ce9b1",
	"name": "test plan",
	"start_time": 1000000,
	"end_time":2000000
}
```
Response 
```
{
    "code": 0,
    "message": "success",
    "data": "5cd474a8c4cc785f0c3ce9b2"
}
```
### 3. Get advertising plans by user id

| Key  | Value |
| ---- | ----- |
| url  | /ad_plans  |
| method  | GET  |

| Param  | Type |
| ---- | ----- |
| user_id | string |

#### Example

Request GET http://localhost:7000/ad_plans?user_id=5cd473ccc4cc785f0c3ce9b1

Response 
```
{
    "code": 0,
    "message": "success",
    "data": [
        {
            "id": "5cd474a8c4cc785f0c3ce9b2",
            "user_id": "5cd473ccc4cc785f0c3ce9b1",
            "name": "test plan",
            "start_time": "1970-01-12 08:46:40",
            "end_time": "1970-01-23 22:33:20",
            "create_time": "2019-05-09 14:42:48",
            "update_time": "2019-05-09 14:42:48"
        }
    ]
}
```
### 4. Update advertising plan

| Key  | Value |
| ---- | ----- |
| url  | /ad_plan/update  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| plan_id | string |
| name | string |
| start_time | int |
| end_time | int |

#### Example

Request POST http://localhost:7000/ad_plan/update

```
{
	"plan_id":"5cd473ccc4cc785f0c3ce9b2",
	"name": "test plan1",
	"start_time": 1000000,
	"end_time":2000000
}
```
Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd474a8c4cc785f0c3ce9b2"
}
```

### 4. Delete advertising plan

| Key  | Value |
| ---- | ----- |
| url  | /ad_plan/delete  |
| method  | DELETE  |

| Param  | Type |
| ---- | ----- |
| plan_id | string |

#### Example

Request Delete http://localhost:7000/ad_plan/delete?plan_id=5cd474a8c4cc785f0c3ce9b2

Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd474a8c4cc785f0c3ce9b2"
}
```

### 5. Create Advertising unit

| Key  | Value |
| ---- | ----- |
| url  | /ad_unit/create  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| plan_id | string |
| name | string |
| position_type | int |
| budget | int |

#### Example

Request POST http://localhost:7000/ad_unit/create

```
{
	"plan_id": "5cd47adec4cc786178214aea",
	"name":"test ad unit",
	"position_type": 1,
	"budget":100000
}
```
Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd474a8c4cc785f0c3ce9b2"
}
```
### 6. Create advertising unit keyword limit

| Key  | Value |
| ---- | ----- |
| url  | /ad_unit/keyword  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| unit_id | string |
| keyword | string |


#### Example

Request POST http://localhost:7000/ad_unit/keyword
```
{
	"unit_id": "5cd47b84c4cc786178214aec",
	"keyword": "guitar"
}
```
Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd47bd5c4cc786178214aed"
}
```

### 7. Create advertising unit interest limit

| Key  | Value |
| ---- | ----- |
| url  | /ad_unit/interest  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| unit_id | string |
| tag | string |


#### Example

Request POST http://localhost:7000/ad_unit/interest
```
{
	"unit_id": "5cd47b84c4cc786178214aec",
	"tag": "music"
}
```
Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd47bd5c4cc786178214aed"
}
```

### 8. Create advertising unit district limit

| Key  | Value |
| ---- | ----- |
| url  | /ad_unit/district  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| unit_id | string |
| state | string |
| city | string |

#### Example

Request POST http://localhost:7000/ad_unit/district
```
{
	"unit_id": "5cd47b84c4cc786178214aec",
	"state": "NY",
	"city": "New York"
}
```
Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd47bd5c4cc786178214aed"
}
```

### 9. Create advertising creativity

| Key  | Value |
| ---- | ----- |
| url  | /ad_innovation/create  |
| method  | POST  |

| Param  | Type |
| ---- | ----- |
| unit_id | string |
| innovation_id | string |

#### Example

| Key  | Value |
| ---- | ----- |
| url  | /unit_innovation/create  |
| method  | POST  |

Request POST http://localhost:7000/ad_innovation/create
```
{
	"unit_id": "5cd473ccc4cc785f0c3ce9b1",
	"innovation_id": "5cd474a8c4cc785f0c3ce9b2"
}
```
Response
```
{
    "code": 0,
    "message": "success",
    "data": "5cd47f88c4cc78629494e131"
}
```
### 10. Create advertising unit and advertising creativity relation