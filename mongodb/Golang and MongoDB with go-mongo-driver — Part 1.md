# Golang and MongoDB with go-mongo-driver — Part 1



原文地址: https://medium.com/glottery/golang-and-mongodb-with-go-mongo-driver-part-1-1c43aba25a1

# Requirements

- [MongoDB](https://www.mongodb.com/) version 3 or higher. Alternatively, a cloud service such as [Mongodb Atlas](https://www.mongodb.com/cloud/atlas).
- [Go](https://golang.org/) version 1.10 or higher.
- [go-mongo-driver](https://github.com/mongodb/mongo-go-driver/) version 1.0.3

# Should I read this post?

The short answer is yes. OK, now seriously, this post assumes that you has used Golang and you have some knowledge of MongoDB. The interest of this post is to be an introduction to [go-mongo-driver](https://github.com/mongodb/mongo-go-driver/). This as an alternative to the well-known [mgo](https://github.com/globalsign/mgo) package through a simple CRUD. If this interests you, well, go ahead!

# Prepare the workspace

For this project, we will use the [go modules](https://blog.golang.org/using-go-modules), if you do not know what I am talking about, you should probably look for it. But you can work on your GOPATH, ignore the “*mod commands*” and follow the post.

In the terminal, type the following command, or you can create a directory with a GUI tool, but it is not cool, right?.

```
mkdir mongo-golang-crud && cd $_
```

Inside the new directory, we execute the following command.

```
go mod init github.com/<username>/mongo-golang-crud
```

This creates the go.mod file, but do not worry about that now. With that, we have initialized the module and we can download the dependency we need for the project, as below.

```
go get -v go.mongodb.org/mongo-driver/mongo@v1.0.3
```

# Connection with MongoDB

First of all we need to make a connection with MongoDB. For this we will use the URI of the database. This is not a [MongoDB tutorial](https://docs.mongodb.com/manual/crud/), so I will not stop to explain details about this database engine.

```
package mainimport (
    "context"
    "fmt"    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)const (
    // Name of the database.
    DBName = "glottery"
    URI = "mongodb://<user>:<password>@<host>/<name>"
)func main() {
    // Base context.
    ctx := context.Background()    // Options to the database.
    clientOpts := options.Client().ApplyURI(URI)
    client, err := mongo.Connect(ctx, clientOpts)
    if err != nil {
        fmt.Println(err)
        return
    }    db := client.Database(DBName)
    fmt.Println(db.Name()) // output: glottery
}
```

Well, the connection is ready. Now we can make interactions with the database.

# A new collection

If you have previously worked with MongoDB, and I hope so because I am leaving many things out, you will know that you use document collections. We are going to use a collection named notes. This collection will store many documents of a simple structure called Note.

```
[...]
const (
    // Name of the database.
    DBName = "glottery"
    // Name of the collection.
    notesCollection = "notes"    URI = "mongodb://<user>:<password>@<host>/<name>"
)
[...]
```

Now we define the Note struct. Don’t forget to import the **time** package.

```
type Note struct {
    ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
    Title string `json:"title"`
    Body string `json:"body"`
    CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
```

Making some small changes in the **main** function.

```
func main() {
    [...]
    db := client.Database(DBName)
    coll := db.Collection(notesCollection)
    fmt.Println(coll.Name()) // output: notes
}
```

# Inserting data

## Insert one document

```
func main() {
    [...]
    note := Note{}    // An ID for MongoDB.
    note.ID = primitive.NewObjectID()
    note.Title = "First note"
    note.Body = "Some spam text"
    note.CreatedAt = time.Now()
    note.UpdatedAt = time.Now()
    result, err := coll.InsertOne(ctx, note)
    if err != nil {
        fmt.Println(err)
        return
    }    // ID of the inserted document.
    objectID := result.InsertedID.(primitive.ObjectID)
    fmt.Println(objectID) // output: ObjectID("5d100d9c23affb7006dd9cff")
}
```

## Insert many documents

```
func main() {
    [...]    // Insert Many Documents.
    notes := []interface{}{}
    note2 := Note{}
    note2.ID = primitive.NewObjectID()
    note2.Title = "First note"
    note2.Body = "Some spam text"
    note2.CreatedAt = time.Now()
    note2.UpdatedAt = time.Now()    note3 := Note{
        ID:        primitive.NewObjectID(),
        Title:     "Third note",
        Body:      "Some spam text",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }    notes = append(notes, note2, note3)
    results, err := coll.InsertMany(ctx, notes)
    if err != nil {
        fmt.Println(err)
        return
    }    fmt.Println(results.InsertedIDs)
    // output: [ObjectID("5d100faf685dbd3130446bac") ObjectID("5d100faf685dbd3130446bad")]
}
```

As you can see, the **InsertMany** method required a slice of empty interfaces, in addition to the typical context, of course.

Just like to the **InsertOne** method, this returns the ID, but InsertMany returns a slice of empty interfaces with the objectIDs of the inserted documents.

# Update one document

First, we need to import the bson package, this package comes together with go-mongo-driver. We can import it in this way *“****go.mongodb.org/mongo-driver/bson\****”*.

```
func main() {
    [...]
    // Update a document.    // Parsing a string ID to ObjectID from MongoDB.
    objID, err := primitive.ObjectIDFromHex("5d100d9c23affb7006dd9cff")
    if err != nil {
        fmt.Println(err)
        return
    }    resultUpdate, err := coll.UpdateOne(
        ctx,
        bson.M{"_id": objID},
        bson.M{
            "$set": bson.M{
            "body":       "Some updated text",
            "updated_at": time.Now(),
            },
        },
    )    fmt.Println(resultUpdate.ModifiedCount) // output: 1
}
```

# Delete one document

```
func main() {
    [...]
    // Delete one document.
    objID, err = primitive.ObjectIDFromHex("5d1017bda2dc2ce292e5f16c")
    if err != nil {
        fmt.Println(err)
        return
    }    resultDelete, err := coll.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        fmt.Println(err)
        return
    }    fmt.Println(resultDelete.DeletedCount) // output: 1
}
```

Simple, right? Like **DeleteOne** we can use the **DeleteMany** method. Well, to no one’s surprise, **DeleteMany** removes any document that matches the filter.

# Find Documents

Finally, it is time to look for documents in the database, maybe the most typical operation that we should to do.

## Find one document

```
func main() {
    [...]
    // Find documents.
    objID, err = primitive.ObjectIDFromHex("5d100d9c23affb7006dd9cff")
    if err != nil {
        fmt.Println(err)
        return
    }    findResult := coll.FindOne(ctx, bson.M{"_id": objID})
    if err := findResult.Err(); err != nil {
        fmt.Println(err)
        return
    }    n := Note{}
    err = findResult.Decode(&n)
    if err != nil {
        fmt.Println(err)
        return
    }    fmt.Println(n.Body) // output: Some updated text
}
```

## Find many documents

```
func main() {
    [...]
    notesResult := []Note{}
    cursor, err := coll.Find(ctx, bson.M{})
    if err != nil {
        fmt.Println(err)
        return
    }    // Iterate through the returned cursor.
    for cursor.Next(ctx) {
        cursor.Decode(&n)
        notesResult = append(notesResult, n)
    }    for _, el := range notesResult {
        fmt.Println(el.Title)
    }
}
```

Well, this is the basic information about the MongoDB queries, but we have not touched the surface yet. If you are interested in learning more about the MongoDB queries. I recommend that you review the [documentation](https://docs.mongodb.com/manual/tutorial/query-documents/). I recommend checking more about the [cursors](https://docs.mongodb.com/manual/tutorial/iterate-a-cursor/#read-operations-cursors).

Soon we will explore more about [go-mongo-driver](https://github.com/mongodb/mongo-go-driver). Thanks for reading, see you next time.

**NOTE**: The full code of this publication is available on [GitHub](https://github.com/orlmonteverde/mongo-golang-crud).