package queue

import (
	"auth_service/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisEnqueue struct {
	rdb *redis.Client
}

func NewRedisEnqueue(rdb *redis.Client) *RedisEnqueue {
	return &RedisEnqueue{
		rdb: rdb,
	}
}

var ctx = context.Background()

func (r *RedisEnqueue) EnqueueInitSekolahTask(sekolahModel models.SekolahTenant) error {
	payload, err := json.Marshal(sekolahModel)
	if err != nil {
		return err
	}

	err = r.rdb.RPush(ctx, "sekolah_init_queue", payload).Err()
	if err != nil {
		return err
	}

	fmt.Println("Task ditambahkan ke queue")
	return nil
}

type InitTaskPayload struct {
	Sekolah models.SekolahTenant `json:"sekolah"`
	UserID  uint64               `json:"user_id"`
}

func (r *RedisEnqueue) EnqueueInitSCTask(sekolahModel models.SekolahTenant, userId uint64) error {
	payload := struct {
		Sekolah models.SekolahTenant `json:"sekolah"`
		UserID  uint64               `json:"user_id"`
	}{
		Sekolah: sekolahModel,
		UserID:  userId,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("gagal marshal payload: %w", err)
	}

	err = r.rdb.RPush(context.Background(), "sc_init_queue", data).Err()
	if err != nil {
		return fmt.Errorf("gagal kirim ke Redis queue: %w", err)
	}

	fmt.Println("Task ditambahkan ke queue dengan userId:", userId)
	return nil
}
