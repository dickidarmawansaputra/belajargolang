package repository

import "dasar/unit-test/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
