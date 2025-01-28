# Confirm

# Init
```sh
git clone https://github.com/mtvy/confirm.git
cd confirm
docker compose up --build
```

# Usage

> POST `/send`
> ```yaml
> request: {
>     "content": "Hello World!"
> }
> responce: 
> 200 -> {
>     "id": "4dbd3732-3cf1-44bb-944d-c02eb3050e70",
>     "status": "sent for approval"
> }
> ...
> ```


> GET `/:id` 
> ```yaml
> responce: 
> 200 -> {
>     "message": {
>         "id": "c4c201a4-6d0b-4456-8217-9b87c63f0c1f",
>         "content": "Hello World!",
>         "status": "pending"
>     }
> }
> ...
> ```


> POST `/approve/:id`
> ```yaml
> responce: 
> 200 -> {
>     "message": {
>         "id": "191d4da9-01a0-45fc-b8f6-938a86ae66fb",
>         "content": "Hello World!",
>         "status": "approved"
>     },
>     "status": "approved"
> }
> 400 -> {
>     "status": "message approved"
> }
> ...
> ```


> POST `/reject/:id`
> ```yaml
> responce: 
> 200 -> {
>     "message": {
>         "id": "020dfcb9-c536-4cb7-b5a7-ce6ecfba9b10",
>         "content": "Hello World!",
>         "status": "rejected"
>     },
>     "status": "rejected"
> }
> 400 -> {
>     "status": "message rejected"
> }
> ...
> ```