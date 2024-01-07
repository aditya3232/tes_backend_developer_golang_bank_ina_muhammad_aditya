# API Spec

## 1 Auth

### 1.1 Google Login

Request :
- Method : GET
- URL : `{{local}}:3636/google_login`
- Response :

```json 
{
    "login_url": "https://accounts.google.com/o/oauth2/auth?client_id=707621798670-iklkri7gbd2gc5av4scudopsfmmakgi7.apps.googleusercontent.com&redirect_uri=http%3A%2F%2F127.0.0.1%3A3636%2Fgoogle_callback&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=random"
}
```
### 1.2 Logout

Request :
- Method : GET
- URL : `{{local}}:3636/logout`
- Header : 
    - Authorization : string
- Response :

```json 
{
    "meta": {
        "message": "Logout Success",
        "code": 200
    },
    "data": null
}
```

## 2 Users

### 2.1 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/users`
- Response :

```json 
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": [
        {
            "id": 7,
            "name": "Irma Fitri",
            "email": "IrmaFit@gmail.com",
            "password": "$2axxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "created_at": "2024-01-07T07:45:21+07:00",
            "updated_at": "2024-01-07T07:45:21+07:00"
        },
        {
            "id": 8,
            "name": "Muhammad Aditya",
            "email": "m.aditya3232@gmail.com",
            "password": "$2axxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "created_at": "2024-01-07T08:27:28+07:00",
            "updated_at": "2024-01-07T08:27:28+07:00"
        }
    ]
}
```

### 2.2 Get By ID

Request :
- Method : GET
- URL : `{{local}}:3636/users/:id`
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 8,
        "name": "Muhammad Aditya",
        "email": "m.aditya3232@gmail.com",
        "password": "$2a$04$3Wy4MvxBS1D8FVu3at8LPefPdq73YECFZeMl5ovZ6snbA7UDwKwwS",
        "created_at": "2024-01-07T08:27:28+07:00",
        "updated_at": "2024-01-07T08:27:28+07:00"
    }
}
```

### 2.3 Create

Request :
- Method : POST
- URL : `{{local}}:3636/users`
- Body (form-data) :
    - name : string, required
    - email : string, required
    - password : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 9,
        "name": "Ichsan Ashiddiqi",
        "email": "iashiddiqi@gmail.com",
        "password": "$2a$04$96l39ugW9mnbRCP/mTnrPOWnIO6wREbf6dkUJVH1nWF4X6bJhcu1S",
        "created_at": "2024-01-07T15:04:16.699+07:00",
        "updated_at": "2024-01-07T15:04:16.699+07:00"
    }
}
```

### 2.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/users/:id`
- Body (form-data) :
    - name : string
    - email : string
    - password : string
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 9,
        "name": "",
        "email": "",
        "password": "$2a$04$gU3Yx2o8XzX6vqTR8K2ZsOU42ORUz./VpcoXgw0KeOZOzu5pQyVuy",
        "created_at": null,
        "updated_at": "2024-01-07T15:05:16.789+07:00"
    }
}
```

### 2.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/users/:id`
- Response :

```json
{
    "meta": {
        "message": "Successfully deleted data.",
        "code": 200
    },
    "data": null
}
```

## 3 Tasks

### 3.1 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/tasks`
- Header :
    - Authorization : Bearer string
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": [
        {
            "id": 1,
            "user_id": 7,
            "title": "tes_1",
            "description": "tes_1",
            "status": "working",
            "created_at": "2024-01-07T08:26:23+07:00",
            "updated_at": "2024-01-07T08:26:23+07:00"
        },
        {
            "id": 3,
            "user_id": 8,
            "title": "tes_2",
            "description": "tes_2",
            "status": "pending",
            "created_at": "2024-01-07T08:27:50+07:00",
            "updated_at": "2024-01-07T08:30:41+07:00"
        }
    ]
}
```

### 3.2 Get By ID

Request :
- Method : GET
- URL : `{{local}}:3636/tasks/:id`
- Header :
    - Authorization : Bearer string
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 3,
        "user_id": 8,
        "title": "tes_2",
        "description": "tes_2",
        "status": "pending",
        "created_at": "2024-01-07T08:27:50+07:00",
        "updated_at": "2024-01-07T08:30:41+07:00"
    }
}
```

### 3.3 Create

Request :
- Method : POST
- URL : `{{local}}:3636/tasks`
- Header :
    - Authorization : Bearer string
- Body (form-data) :
    - user_id : int, required
    - title : string, required
    - description : string, required
    - status : string
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 4,
        "user_id": 8,
        "title": "tes_3",
        "description": "tes_3",
        "status": "pending",
        "created_at": "2024-01-07T15:08:16.699+07:00",
        "updated_at": "2024-01-07T15:08:16.699+07:00"
    }
}
```

### 3.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/tasks/:id`
- Header :
    - Authorization : Bearer string
- Body (form-data) :
    - user_id : int
    - title : string
    - description : string
    - status : string
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 4,
        "user_id": 8,
        "title": "tes_3",
        "description": "tes_3",
        "status": "pending",
        "created_at": "2024-01-07T15:08:16.699+07:00",
        "updated_at": "2024-01-07T15:08:16.699+07:00"
    }
}
```

### 3.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/tasks/:id`
- Header :
    - Authorization : Bearer string
- Response :
    
```json
{
    "meta": {
        "message": "Successfully deleted data.",
        "code": 200
    },
    "data": null
}
```

# Migrations
- Perintah Migrate
    - migrate -database "mysql://root:root_password@tcp(127.0.0.1:3306)/bank_ina" -path db/migrations up
    

