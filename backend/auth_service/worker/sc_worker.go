package worker

import (
	"auth_service/models"
	"auth_service/services"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func StartSCInitWorker(rdb *redis.Client) {
	ctx := context.Background()

	for {
		val, err := rdb.BLPop(ctx, 0*time.Second, "sc_init_queue").Result()
		if err != nil {
			log.Printf("Gagal ambil task dari queue: %v", err)
			continue
		}

		var model *models.TaskPayload
		if err := json.Unmarshal([]byte(val[1]), &model); err != nil {
			log.Printf("Gagal parsing task: %v", err)
			continue
		}

		if err := initSCService(model); err != nil {
			log.Printf("Gagal eksekusi initSekolahService: %v", err)
		} else {
			log.Println("initSekolahService berhasil dijalankan.")
		}
	}
}

func initSCService(taskPayload *models.TaskPayload) error {
	scClient, err := services.NewSCServiceClient()
	if err != nil {
		return err
	}
	if err := scClient.RegistrasiSekolahTenant(&taskPayload.SekolahTenant, taskPayload.UserId); err != nil {
		return err
	}

	return nil
}
