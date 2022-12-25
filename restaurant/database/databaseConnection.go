package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:{port}"

	fmt.Print(MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	// ApplyURI를 통해서 url을 적용 시켜 준다.
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	// -> 적용을 안한다면 아래쪽에서는 이런식으로 사용을 해야 한다.
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// context는 작업 시간을 제어 하는데 사용한다.
	// withTimeout은 작업 흐름을 얼마나 유지 할지를 결정한다.
	// 즉 현재는 10초간의 유예기간을 주는 코드와 같고.
	// defer를 통해서 만약 함수가 문제 없이 성공하였다면 -> 10초 이내에
	// context를 통료 시킨다.
	// -> 보통은 ctx.Done를 통해서 어떠한 처리를 하지만 현재에는 쓸 만한 곳이 없기 떄문에 없는 것이고
	// -> 끝내기 전에 루틴을 하나 없애기 위해서 defer cancel를 호출
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect to mongoDB")

	return client

}

var Client *mongo.Client = DBinstance()


func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.database("DB이름").collection(collectionName)

	return collection
}