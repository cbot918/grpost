package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGO_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client := newMongo(uri)
	col := client.Database("myFirstDatabase").Collection("posts")
	res := findAllFrom(col)

	res0, err := bson.MarshalExtJSON(res[0], false, false)
	if err != nil {
		panic(err)
	}

	writeFileCover("posts.json", []byte("["))
	writeContinue("posts.json", res0)
	writeContinue("posts.json", []byte(","))
	for _, item := range res {
		jsonData, err := bson.MarshalExtJSON(item, false, false)
		if err != nil {
			fmt.Println("marshal bson to json failed")
			panic(err)
		}
		writeContinue("posts.json", jsonData)
		writeContinue("posts.json", []byte(","))
	}
	writeContinue("posts.json", []byte("]"))

}

func findAllFrom(col *mongo.Collection) []bson.D {
	var results = []bson.D{}
	cursor, err := col.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err != nil {
		panic(err)
	}

	// fmt.Println(results)
	fJson(results)
	return results
}

func newMongo(uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("mongo connect failed")
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return client
}

func findOneExample(col *mongo.Collection) {
	var result bson.M
	fullName := "User 1"
	err := col.FindOne(context.TODO(), bson.D{{"fullName", fullName}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", fullName)
		return
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the title %s\n", fullName)
			return
		}
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func fJson(target []bson.D) {
	jsonData, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func writeFileCover(fileName string, target []byte) {
	err := os.WriteFile(fileName, target, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func writeContinue(fileName string, target []byte) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(string(target))
	if err != nil {
		fmt.Println(err)
	}
}
