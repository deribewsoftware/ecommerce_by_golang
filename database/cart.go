package database

var (
	ErrCantFindProduct   = error.New(" cannot find product")
	ErrCantDecodeProduct = error.New(" cannot decode product")
	ErrCantUserIdvalid   = error.New(" cannot user id valid")
	ErrCantUpdatedUser   = error.New(" cannot update user")
	ErrCantRemoveItem    = error.New(" cannot remove item")
	ErrCantGetItem       = error.New(" cannot get item")
	ErrCantBuyCartItem   = error.New(" cannot buy cart item")
)

func AddProductToCart() {

}
func RemoveCartItem() {}
