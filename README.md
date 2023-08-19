# Vehicle Rental
Vehicle rental app allows people to rent vehicles when they' re on vacation or while prepare for a vacation from their current location. Not only for travelers, this app is available for everyone who wants to rent vehicles they need.

##
## Application Installation steps
Install go modules.
``` bash
    go mod tidy
```

Set environment variables.
``` bash
    APP_ENV=dev
    PORT=3021
    BASE_URL=http://localhost:3021
    JWT_SECRET=Th1515Jwt53Cr3t

    DB_HOST=
    DB_NAME=
    DB_USER=
    DB_PASS=

    # Create cloudinary account first and get cloudinary name, api key, api secret and then paste into the keys below.
    CLOUD_NAME=
    CLOUD_KEY=
    CLOUD_SECRET=

    # get google smtp host and port, and then get google application password from your account
    SMTP_HOST=
    SMTP_PORT=
    MAIL_USER=
    MAIL_PASSWORD=
```

There are 3 CLIs you can use in this app.
``` bash

    # to create all tables
    go run . migrate -u

    # to drop all tables
    go run . migrate -d

    # to seed datas as records of the tables
    go run . seed -u

    # to delete all records of the tables
    go run . seed -d

    # to run application
    go run . serve
```