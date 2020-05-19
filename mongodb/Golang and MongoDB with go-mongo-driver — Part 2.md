# Golang and MongoDB with go-mongo-driver — Part 2

原文地址: https://medium.com/glottery/golang-and-mongodb-with-go-mongo-driver-part-2-1d0aa64cdbf4

# Requirements

- [MongoDB](https://www.mongodb.com/) version 3 or higher. Alternatively, a cloud service such as [Mongodb Atlas](https://www.mongodb.com/cloud/atlas).
- [Go](https://golang.org/) version 1.10 or higher.
- [go-mongo-driver](https://github.com/mongodb/mongo-go-driver/) version 1.0.3

# Should I read this post?

The short answer is yes. OK, now seriously, this post assumes that you has used [Go](https://golang.org/) (Golang) and you have some knowledge of [MongoDB](https://www.mongodb.com/). This is the second part of the post about [go-mongo-driver](https://github.com/mongodb/mongo-go-driver/), an alternative to the well-known [mgo](https://labix.org/mgo). Well, this time we’re going to test the [MongoDB Indexes](https://docs.mongodb.com/manual/indexes/). If this interests you, well, go ahead!

# What is a Index?

According to [MongoDB](https://docs.mongodb.com/manual/indexes/):

> Indexes are special data structures that store a small portion of the collection’s data set in an easy to traverse form. The index stores the value of a specific field or set of fields, ordered by the value of the field. The ordering of the index entries supports efficient equality matches and range-based query operations. In addition, MongoDB can return sorted results by using the ordering in the index.

# Prepare the workspace

If you followed the [first part](https://medium.com/glottery/golang-and-mongodb-with-go-mongo-driver-part-1-1c43aba25a1) of this publication, you should know this; otherwise, simply follow the steps below.

For this project, we will use the [go modules](https://blog.golang.org/using-go-modules), if you do not know what I am talking about, you should probably look for it. But you can work on your GOPATH, ignore the “*mod commands*” and follow the post.

In the terminal, type the following command, or you can create a directory with a GUI tool, but it is not cool, right?.

```
mkdir mongo-golang-indexes && cd $_
```

Inside the new directory, we execute the following command.

```
go mod init github.com/<username>/mongo-golang-idexes
```

This creates the **go.mod** file, but do not worry about that now. With that, we have initialized the module and we can download the dependency we need for the project, as below.

```
go get -v go.mongodb.org/mongo-driver/mongo@v1.0.3
```

# Starting

For a fast advance and considering that we saw the configuration in the [last publication,](https://medium.com/glottery/golang-and-mongodb-with-go-mongo-driver-part-1-1c43aba25a1) we will start with the base code that is shown below.

```
package mainimport (
        "context"
        "fmt"
        "time"        "go.mongodb.org/mongo-driver/bson"
        "go.mongodb.org/mongo-driver/bson/primitive"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
)const (
        DBName          = "glottery"
        notesCollection = "notes"
        URI = "mongodb://<user>:<password>@<host>/<name>"
)type Note struct {
        ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
        Title     string             `json:"title"`
        Body      string             `json:"body"`
        CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
        UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}func main() {
        // Base context.
        ctx := context.Background()
        clientOpts := options.Client().ApplyURI(URI)
        client, err := mongo.Connect(ctx, clientOpts)
        if err != nil {
                fmt.Println(err)
                return
        }        db := client.Database(DBName)
        coll := db.Collection(notesCollection)
}
```

# Adding an index

Probably, the most important (and confusing) of the indexes with the go-mongo driver is the need to import the [bsonx](https://godoc.org/go.mongodb.org/mongo-driver/x/bsonx) package. Initially, this was confusing to me, in fact, it’s not very intuitive if you ask me. But when you get it, the rest is very easy. Go to code.

```
[...]
import (
        "context"
        "fmt"
        "time"        "go.mongodb.org/mongo-driver/bson/primitive"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"        "go.mongodb.org/mongo-driver/x/bsonx"
)
[...]
```

## Index Options

We can set several options for the Index such as, for example, the unique restriction, but this is very specific for [MongoDB](https://docs.mongodb.com/manual/indexes/#b-tree). If you want more information, I recommend consulting the [documentation](https://docs.mongodb.com/manual/core/index-creation/#index-creation-background).

```
func main() {
       [...]
        
       // Options
       indexOptions := options.Index().SetUnique(true)
}
```

## Index Keys

Now, let’s define the key (or keys) for the index. Again, this is very specific for MongoDB and I invite you to visit the [Documentation](https://docs.mongodb.com/manual/indexes/#index-types). For this example, we create a [Single field index](https://docs.mongodb.com/manual/core/index-single/) for the *Title* field, the key on this map (Yes, the **bsonx.MDoc** type is a map) represents the field and the value, in this case, can be a positive number (order ascending) or a negative number (descending order).

```
[...]
func main() {
        [...]
        indexOptions := options.Index().SetUnique(true)
        indexKeys := bsonx.MDoc{
                "title": bsonx.Int32(-1),
        }
}
```

## The index model

```
[...]
func main() {
        [...]
        indexOptions := options.Index().SetUnique(true)
        indexKeys := bsonx.MDoc{
                "title": bsonx.Int32(-1),
        }        noteIndexModel := mongo.IndexModel{
                Options: indexOptions,
                Keys:    indexKeys,
        }
}
```

Create index to the collection

```
[...]
func main() {
        ctx := context.Background()
        clientOpts := options.Client().ApplyURI(URI)
        client, err := mongo.Connect(ctx, clientOpts)
        if err != nil {
                fmt.Println(err)
                return
        }        db := client.Database(DBName)
        coll := db.Collection(notesCollection)

        indexOptions := options.Index().SetUnique(true)
        indexKeys := bsonx.MDoc{
                "title": bsonx.Int32(-1),
        }        noteIndexModel := mongo.IndexModel{
                Options: indexOptions,
                Keys:    indexKeys,
        }        indexName, err := coll.Indexes().CreateOne(ctx, noteIndexModel)
        if err != nil {
                fmt.Println(err)
                return
        }        fmt.Println(indexName) // Output: title_-1
}
```

Simple, right? Now the title field should be unique and the queries by this field will be more fast.

I consider important to mention that it is possible to create several Indices at once using the *CreateMany* method instead of *CreateOne*, this receives an array of *mongo.IndexModel*.

# Text Indexes

> MongoDB provides [text indexes](https://docs.mongodb.com/manual/core/index-text/#index-feature-text) to support text search queries on string content. `text` indexes can include any field whose value is a string or an array of string elements.

Another interesting type of index is the text index, very useful for text search. Like the previous one, we make a text index in the following way:

```
[...]
func main() {
        [...]
        textIndexModel := mongo.IndexModel{
                Options: options.Index().SetBackground(true),
                Keys: bsonx.MDoc{
                        "title": bsonx.String("text"),
                        "body":  bsonx.String("text"),
                },
        }        indexName, err = coll.Indexes().CreateOne(ctx, textIndexModel)
        if err != nil {
                fmt.Println(err)
                return
        }        fmt.Println(indexName)  // Output: title_text_body_text
}
```

## Making queries

First, we are going to insert some dummy notes to the collection.

```
[...]
var notes = []interface{}{
        Note{
                ID:        primitive.NewObjectID(),
                Title:     "First note",
                Body:      "Concurrency is not parallelism",
                CreatedAt: time.Now(),
        }, Note{
                ID:        primitive.NewObjectID(),
                Title:     "Second note",
                Body:      "A little copying is better than a little dependency",
                CreatedAt: time.Now(),
        }, Note{
                ID:        primitive.NewObjectID(),
                Title:     "Third note",
                Body:      "Don't communicate by sharing memory, share memory by communicating",
                CreatedAt: time.Now(),
        }, Note{
                ID:        primitive.NewObjectID(),
                Title:     "Fourth note",
                Body:      "Don't just check errors, handle them gracefully",
                CreatedAt: time.Now(),
        },
}
[...]
```

OK, lets to insert dummy notes.

```
[...]
func main() {
        [...]
        _, err = coll.InsertMany(ctx, notes)
        if err != nil {
                fmt.Println(err)
                return
        }
}
```

Now, the queries!

```
[...]
func main() {
        [...]
        n := Note{}
        fmt.Println("First example")
        cursor, err := coll.Find(ctx, bson.M{
                "$text": bson.M{
                        "$search": "note",
                },
        })        for cursor.Next(ctx) {
                cursor.Decode(&n)
                fmt.Println(n)
        }        fmt.Println("Second example")
        cursor, err = coll.Find(ctx, bson.M{
                "$text": bson.M{
                        "$search": "gracefully",
                },
        })        for cursor.Next(ctx) {
                cursor.Decode(&n)
                fmt.Println(n)
        }
}
```

**NOTE:** If we run the binary twice or more, an error is returned. This is due to the attempt to insert again the same notes, which violates the unique constraint that we did before.

The full code of this publication is available on [GitHub](https://github.com/orlmonteverde/mongo-golang-indexes/edit/master/README.md). Soon we will explore more about [go-mongo-driver](https://github.com/mongodb/mongo-go-driver). Thanks for reading, see you next time.