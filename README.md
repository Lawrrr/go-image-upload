# Go Image Upload

### Description

This is an image uploading app using Golang and SQLite. Every time an image is uploaded it will be saved under `public\images` folder and the image's details (path, content-type, etc.) will be inserted to database.

### Setup SQLite

To setup SQLite in your local machine. Run the ff. commands:
This will create the database. To exit sqlite terminal just input `CTRL+C`

```
1.) sqlite3 go-image.db
2.) .databases
```

Create image table with its corresponding column from the migration file.

```
sqlite3 go-image.db ".read ./database/migration.sql"
```

VOILA! You have successfully created your SQLite database.

Note: Make sure have SQLite installed in your machine.

### Run Website

1. Create `.env` file then copy the values found in `.env.example`. You can change the values to your preference.
2. To run this website, you must be in the project's root folder and run this command

```
go run ./api/main.go
```
