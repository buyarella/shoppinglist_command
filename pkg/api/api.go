package api

import (
	"context"
)

func New() *API {
	return &API{}
}

func (it *API) NewShoppingList(_ context.Context, _ *NewShoppingListRequest) (*ShoppingList, error) {
	return nil, nil
}

func (it *API) SetActiveShoppingList(_ context.Context, _ *SetActiveShoppingListRequest) (*ShoppingList, error) {
	return nil, nil
}

func (it *API) FinishShopping(_ context.Context, _ *FinishShoppingRequest) (*FinishShoppingResponse, error) {
	return nil, nil
}

func (it *API) AddNewItem(_ context.Context, _ *AddNewItemRequest) (*Item, error) {
	return nil, nil
}

func (it *API) AddItemToShoppingList(_ context.Context, _ *AddItemToShoppingListRequest) (*ShoppingList, error) {
	return nil, nil
}

type API struct {
}
