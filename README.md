# Sushi: a series of Go utilities
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Open Source Love 
svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

![alt text](assets/sushi.png?raw=true)

## Function

`Function` represents any generic transformation applied on a generic type `T` into another (or the same) type `V`.
The transformation can generate an error, which is returned in output, as its declaration shows:

```go
type Function[T any, V any] func(t T) (*V, error)
```

To implement a `Function` type you can simply so something like this:

```go
func myfunc(input int) (*string, error) {
    if input < 10 {
        return nil, fmt.Errorf("input is less than 10!")
    } 
    output := fmt.Sprintf(" -- %d -- ", input-5)
    return &output, nil
}
```

## Either

`Either` represents the result of an operation which can either be a successful result or an error as shown by its declaration:

```go
type Either[V any] struct {
        result *V
        err    error
}
```

You can force the generation of an `Either` leaving aside any computation using these two utility methods:

```go
successEither := functional.EitherFromResult[int](10)
failureEither := functional.EitherFromError[string](fmt.Error("error!"))
```

After a computation, you can check if an `Either` wraps a success result or an error through the following functions:

```go
successEither.IsResult() // -> true
successEither.IsError()  // -> false

failureEither.IsResult() // -> false
failureEither.IsError()  // -> true
```

And then, you can get the result or the error accordingly:

```go
result := successEither.getResult()
fmt.Println(*result)

err := failureEither.getError()
fmt.Println(err)
```

If you use the function `GetResult()` in presence of an error, the goroutine panics, so better check before attempting to get a result.
If you want to get both the values from an `Either` you can:

```go
result, err := failureEither.Get()
```

If you want to establish a fallback result in case of error you can use the following:

```go
result := failureEither.GetOrElse("fallbackValue") // it will return a pointer to "fallbackValue"
```

## Future

It is the implementation of an async task running on a different goroutine. 
It takes advantage of both `Function` and `Either` types.
`Future` implementation is under `goutils/functional` package. To import it:

```go
import (
    "github.com/gyozatech/sushi/functional"
)
```

To use a `Future` you need to define a `Function` which represents the task to be implemented asynchronously and the input to that task:

```go

func task(input int) (*string, error) {
    if input < 10 {
        return nil, fmt.Errorf("input is less than 10!")
    } 
    output := fmt.Sprintf(" -- %d -- ", input-5)
    return &output, nil
}

func main() {
    // create and process a task asynchronously
    future := functional.ProcessAsync(task, 10)
    // block the current goroutine until the result is ready
    either := future.WaitForResult()
    // get the result
    result := either.GetResult()
    fmt.Printf("%T : '%v' \n", result, *result)
}
```

You can execute heavy tasks in parallel by simply doing:

```go
future1 := functional.ProcessAsync(task, 10)
future2 := functional.ProcessAsync(task, 100)
future3 := functional.ProcessAsync(task, 65)

res1, err1 := future1.WaitForResult().Get()
res2, err2 := future2.WaitForResult().Get()
res3, err3 := future3.WaitForResult().Get()
```

The function `WaitForResult()` is a blocking operation, so that the time to wait in the main goroutine is the max time among the executions of the three tasks.

You can also choose not to start the tasks directly this way:

```go
// prepare the future without starting its execution
input := 10
future := functional.NewFuture(task, input)

// execute asynchronously
future.Process()

// wait and get the result
res, err := future.WaitForResult().Get()
```
----
## Utils

### `EnrichContextWithValue` and `FetchContextValue`

These are useful to enrich the conxtext and fetch values from it:

```go
import 
(
    "github.com/gyozatech/sushi/utils"
    "context"
)
func LetsTry() {
    ctx := context.TODO()
    ctx = utils.EnrichContextWithValue(ctx, "name", "Alessandro")
    ctx = utils.EnrichContextWithValue(ctx, "age", 35)
    ctx = utils.EnrichContextWithValue(ctx, "subscribed", true)
    ctx = utils.EnrichContextWithValue(ctx, "book", Book{Title: "Jurassic Park", Author: "Michael Chricton"})

    var name *string = utils.FetchContextValue[string](ctx, "name")
    var age *int = utils.FetchContextValue[int](ctx, "age")
    var subscribed *bool = utils.FetchContextValue[bool](ctx, "subscribed")
    var book *Book = utils.FetchContextValue[Book](ctx, "book")
    var unexisting *string= utils.FetchContextValue[string](ctx, "unexisting")
}

```

### `IsEqual`
Is equal is a strong utility to compare equalness, much more elastic than deepEqual:

```go
import (
    "github.com/gyozatech/sushi/utils"
    "fmt"
)

func LetsTry() error {
    isEqual := utils.IsEqual([]*string{ "1", "2", "3"}, []uint16{ 1, 2, 3 })
    if !isEqual {
        return fmt.Errorf("They should be equal!")
    }
}
```

### `CollectResults`
Allows you to collect all results of a function in a slice:

```go
import (
    "github.com/gyozatech/sushi/utils"
    "fmt"
)

func GetUser(id string) (email string, subscribed bool, age int, err error) {
    // ...
    return "alessandro@email.com", true, 35, nil
}

func LetsTry() error {
    age := CollectResults(GetUser("12345"))[2]
    fmt.Println("Age:", age)
}
```

### `ForEach`
Allows you to apply a mutator over a slice
```go
import (
    "github.com/gyozatech/sushi/functional"
    "fmt"
)

func LetsTry() {
    // output type
    type DB struct {
	    Type string
	    Version string
	}
	
    // input slice
	dbCodes := []string{"mysql-v8.5.6", "postgres-v5.3.1", "postgres-v4.2.4", "mysql-v11.2.3"}

    // mutator function from string to DB
    var fetchDBFromCode functional.Function[string, DB] = func(dbCode string) (*DB, error) {
	     split := strings.Split(dbCode, "-")
	     if len(split) != 2 {
	         return nil, fmt.Errorf("invalid DB version provided %s", dbCode)
	     }
	     return &DB{ Type: split[0], Version: split[1] }, nil
	}
	// apply mutator to every element of the initial slice with ForEach
	databases, err := functional.ForEach(dbCodes, fetchDBFromCode)
	if err != nil {
	    fmt.Printf("Error: %s \n", err.Error())
	} else {
	    fmt.Printf("%+v \n", databases)
	}
}

```

### `FindOne` and `FindMany`

These two functions allow filtering a slice through a condition:
```go
import (
    "github.com/gyozatech/sushi/utils"
)

func LetsTry() {

    type Book struct {
        ID int
        Title string
        Author string
    }

    books := []Book{
        {ID: 1, Title: "Jurassic Park", Author: "Michael Chricton"},
        {ID: 2, Title: "The Lost World", Author: "Michael Chricton"},
        {ID: 3, Title: "Relic", Author: "Preston & Child"},
    }

    jkBook := utils.FindOne(books, func(b Book) bool{
        return b.Title == "Jurassic Park"
    })

    chrictonBooks := utils.FindMany(books, func(b Book) bool{
        return b.Author == "Michael Chricton"
    })

    _ = jkBook
    _ = chrictonBooks

}
```

## SQL utils: the Transactor

Credits to https://github.com/giornetta for the implementation.

The `sql.Transactor` interface brings transactions at the level of services, committing all the SQL queries performed `WithinTransaction` of rollbacking.
It's using the `context.Context` to pass a DB transaction across multiple services or DAOs.

Here is an example:

```go
import sqlutil "github.com/gyozatech/sushi/sql"

func (cs *ChartServiceImpl) Checkout(ctx context.Context, chart model.Chart, optins model.Optins) error {

    cs.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {

        // a transaction object is fetched from the ctx if present, else it's created and added to the context

        if err := cs.ChartDao.FlagChart(ctx, chart.ID, "checked"); err != nil {
            return err // this will cause a rollback
        }
        if err := cs.ItemService.ReduceAvailability(ctx, chart.ItemIDs); err != nil {
            return err // this will cause a rollback
        }
        if err := cs.CheckOptinFlags(ctx, optins); err != nil {
            return err // this will cause a rollback
        }
        
        return nil // this will cause a commit if the transaction has started in this `WithinTransaction` function, else not
    })
}
```
To make it possible, the DAO layers (interacting with the DB directly), should be adapted to reuse an existing transaction if present in the ctx or to use the DB directly:
```go
import sqlutil "github.com/gyozatech/sushi/sql"

func (dao *ChartDaoImpl) FlagChart(ctx context.Context, id, flag string) error {
    query := "UPDATE charts SET flag = $1 where id = $2"
    // instead of having: dao.DB.ExecContext(ctx, query, flag, id)
    _, err := sqlutil.Executor(ctx, dao.DB).ExecContext(ctx, query, flag, id)
    return err
}
```
