package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"insider-assesment/pkg/domain/message"
	"reflect"
	"testing"
	"time"
)

func testNewCache() message.Cache {
	staticTime := time.Date(2025, 3, 9, 12, 0, 0, 0, time.UTC)
	return message.NewMessageCache(
		"test-id",
		202,
		staticTime,
	)
}

func testRedis() *redis.Client {
	return NewRedisClient("localhost:6379")
}

func testArgs() struct {
	redis      *redis.Client
	serializer message.CacheSerializer
} {
	return struct {
		redis      *redis.Client
		serializer message.CacheSerializer
	}{
		redis:      testRedis(),
		serializer: NewSerializer(),
	}
}

func Test_service_Save(t *testing.T) {
	type fields struct {
		redis      *redis.Client
		serializer message.CacheSerializer
	}
	type args struct {
		ctx context.Context
		msg message.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "save_1", fields: testArgs(), args: struct {
			ctx context.Context
			msg message.Cache
		}{
			ctx: context.Background(),
			msg: testNewCache(),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Service{
				redis:      tt.fields.redis,
				serializer: tt.fields.serializer,
			}
			if err := r.Save(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Message(t *testing.T) {
	type fields struct {
		redis      *redis.Client
		serializer message.CacheSerializer
	}
	type args struct {
		ctx context.Context
		msg message.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    message.Cache
		wantErr bool
	}{
		{name: "messageCache_1", fields: testArgs(), args: struct {
			ctx context.Context
			msg message.Cache
		}{
			ctx: context.Background(),
			msg: testNewCache(),
		},
			want:    testNewCache(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Service{
				redis:      tt.fields.redis,
				serializer: tt.fields.serializer,
			}
			got, err := r.Message(tt.args.ctx, tt.args.msg.ID())
			if (err != nil) != tt.wantErr {
				t.Errorf("Model() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Model() got = %v, want %v", got, tt.want)
			}
		})
	}
}
