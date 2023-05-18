# gorm-uuid

Datatype UUID stored as binary(16) by GORM

## About

GORM store UUID as string or uuid type, it's more efficient and use less
resources to store UUID as a binary, it would be a bummer to not use this
option because we're using an ORM.

## Usage

Use the type `GormUUID` on the model ID as following :

````go
import "github.com/euphoria-laxis/gorm-uuid/uuid"

type MyModel struct {
	ID          uuid.GormUUID   `gorm:"primary_key;default:(UUID_TO_BIN(UUID()));"`
    // following fields are used to match gorm.Model interface
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}
````

And that's it! Gorm will do all the work to and you will have nothing to do.

## Compatibility

| Driver     |     Compatible     | Versions  |
|:-----------|:------------------:|:---------:|
| MySQL      | :heavy_check_mark: | V5 and V8 |
| MariaDB    | :heavy_check_mark: |   V10+    |
| Postgresql |        :x:         |           |
| SQLite     |        :x:         |           |

## Roadmap

- Add Postgres support
- Add SQLite support

|  Feature   |      Planned       |      Assigned      | WIP | Ready for tests | Validated | Released |
|:----------:|:------------------:|:------------------:|:---:|:---------------:|:---------:|:--------:|
| Postgresql | :heavy_check_mark: | :heavy_check_mark: |     |                 |           |          |
|   SQLite   | :heavy_check_mark: | :heavy_check_mark: |     |                 |           |          |

## License

This project is under [MIT License](./LICENSE).