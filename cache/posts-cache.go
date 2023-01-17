package cache

import "github.com/charagmz/CrashCourse/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
