package container

import (
	"example-goka/app/config/database"
	"example-goka/app/config/kafka"
	"example-goka/app/repository"
	"example-goka/app/usecase"
	"log"

	"gitlab.mncbank.co.id/backend/lockey"
	"gitlab.mncbank.co.id/backend/response/v2"
)

type Container struct {
	Usecase  usecase.Usecase
	Response response.ServerResponse
}

func SetupContainer() Container {
	mysql := database.SetupMySQL()
	log.Println("lewattt")
	res := response.New(mysql.Sql)
	repository.NewRepository(mysql.DB)

	configKafka := kafka.ConsumerConfig()
	producer := configKafka.Producer

	usecase := usecase.NewUsecase(producer, res)

	lockey.Init()

	return Container{
		Usecase:  usecase,
		Response: res,
	}
}
