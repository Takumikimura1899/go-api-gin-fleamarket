package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

// 解説ではFindAllのreturnが*[]models.Itemになっているが、スライスは参照型なのでポインタを使う必要はある？
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updatedItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updatedItem models.Item) (*models.Item, error) {
	for i, item := range r.items {
		if item.ID == updatedItem.ID {
			r.items[i] = updatedItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

// r.items[:i]で現在見ている要素の前までのスライスを取得
// r.items[i+1:]で現在見ている要素の次から最後までのスライスを取得
// これらをappendで結合している
func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, item := range r.items {
		if item.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
