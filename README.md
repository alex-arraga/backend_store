## Impact the database using GORM

1. Open `internal/database/gorm_models`
2. Create a new file with the name of the table to be created. Example `new_table.go`
3. Define the new model fields

```go
   package gorm_models
   
   import "gorm.io/gorm"

   // NewTable represents a new table of the database.
   type NewTable struct {
       gorm.Model
       Name string `gorm:"size:255"`
   }
```

4. Open `internal/database/migrations` and add the new model to the existing models list. This function will be impact at Database.

```go

package migrations

// {...} - Rest of logic

// Executes migrations and impacts the database
func ExecMigrations() {
	migrate(
		&gorm_models.User{},
		&gorm_models.NewGormModel{},
	)
}
```

5. Open ``cmd/api`` and run the next command: `go build && ./api.exe`
6. Verify the new changes in the database using `pgAdmin` or other database manager (GUI)
