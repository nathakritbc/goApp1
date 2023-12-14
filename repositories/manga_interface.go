package repositories

import (
	model "go_app1/models"
)

type MangaRepositoryInterface interface {
	InsertManga(model.PostManga) bool
	GetAllManga() []model.Manga
	GetOneManga(uint) model.Manga
	UpdateManga(uint, model.PostManga) model.Manga
	DeleteManga(uint) bool
}
