// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/tabbed/pqtype"
)

type Cart struct {
	CartID    int64  `json:"cart_id"`
	CartState string `json:"cart_state"`
	// active, completed, fail, pending, lock
	Enums            sql.NullString  `json:"enums"`
	CartProducts     json.RawMessage `json:"cart_products"`
	CartCountProduct sql.NullInt64   `json:"cart_count_product"`
	CartUserId       int64           `json:"cart_UserId"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        sql.NullTime    `json:"updated_at"`
}

type Comment struct {
	CommentID        int64          `json:"comment_id"`
	CommentProductId int64          `json:"comment_productId"`
	CommentUserId    sql.NullInt64  `json:"comment_userId"`
	CommentContent   sql.NullString `json:"comment_content"`
	CommentRight     sql.NullInt64  `json:"comment_right"`
	CommentLeft      sql.NullInt64  `json:"comment_left"`
	CommentParentId  int64          `json:"comment_parentId"`
	IsDeleted        sql.NullBool   `json:"isDeleted"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
}

type Discount struct {
	DiscountID            int64                 `json:"discount_id"`
	DiscountName          string                `json:"discount_name"`
	DiscountDescription   string                `json:"discount_description"`
	DiscountType          sql.NullString        `json:"discount_type"`
	DiscountValue         int64                 `json:"discount_value"`
	DiscountCode          string                `json:"discount_code"`
	DiscountStartDate     time.Time             `json:"discount_start_date"`
	DiscountEndDate       time.Time             `json:"discount_end_date"`
	DiscountMaxUses       int64                 `json:"discount_max_uses"`
	DiscountUsesCount     int64                 `json:"discount_uses_count"`
	DiscountUsersUsed     pqtype.NullRawMessage `json:"discount_users_used"`
	DiscountMinOrderValue int64                 `json:"discount_min_order_value"`
	DiscountShopId        int64                 `json:"discount_shopId"`
	DiscountIsActive      sql.NullBool          `json:"discount_is_active"`
	// all, specific
	DiscountAppliesTo  string                `json:"discount_applies_to"`
	DiscountProductIds pqtype.NullRawMessage `json:"discount_product_ids"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          sql.NullTime          `json:"updated_at"`
}

type Inventory struct {
	InvenProductId    int64                 `json:"inven_productId"`
	InvenLocation     sql.NullString        `json:"inven_location"`
	InvenStock        int64                 `json:"inven_stock"`
	InvenShopId       int64                 `json:"inven_shopId"`
	InvenReservations pqtype.NullRawMessage `json:"inven_reservations"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         sql.NullTime          `json:"updated_at"`
}

type Notification struct {
	NotiID int64 `json:"noti_id"`
	// ORDER-001, ORDER-002, ORDER-003, SHOP-001, PROMOTION-001
	NotiType       string                `json:"noti_type"`
	NotiSenderId   int64                 `json:"noti_senderId"`
	NotiReceivedId int64                 `json:"noti_receivedId"`
	NotiContent    string                `json:"noti_content"`
	NotiOptions    pqtype.NullRawMessage `json:"noti_options"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      sql.NullTime          `json:"updated_at"`
}

type Order struct {
	OrderID             int64                 `json:"order_id"`
	OrderUserId         int64                 `json:"order_userId"`
	OrderCheckout       pqtype.NullRawMessage `json:"order_checkout"`
	OrderShipping       pqtype.NullRawMessage `json:"order_shipping"`
	OrderPayment        pqtype.NullRawMessage `json:"order_payment"`
	OrderProducts       json.RawMessage       `json:"order_products"`
	OrderTrackingNumber sql.NullString        `json:"order_trackingNumber"`
	// pending, confirmed, shipped, cancelled, delivered
	OrderStatus sql.NullString `json:"order_status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Product struct {
	ProductID          int64          `json:"product_id"`
	ProductName        string         `json:"product_name"`
	ProductThumb       string         `json:"product_thumb"`
	ProductDescription sql.NullString `json:"product_description"`
	ProductSlug        sql.NullString `json:"product_slug"`
	ProductPrice       int64          `json:"product_price"`
	ProductQuantity    int64          `json:"product_quantity"`
	// Electronics, Clothing, Furniture
	ProductType           string                `json:"product_type"`
	ProductShop           int64                 `json:"product_shop"`
	ProductAttributes     json.RawMessage       `json:"product_attributes"`
	ProductRatingsAverage sql.NullString        `json:"product_ratingsAverage"`
	ProductVariations     pqtype.NullRawMessage `json:"product_variations"`
	IsDraft               sql.NullBool          `json:"isDraft"`
	IsPublished           sql.NullBool          `json:"isPublished"`
	CreatedAt             time.Time             `json:"created_at"`
	UpdatedAt             sql.NullTime          `json:"updated_at"`
}

type User struct {
	UserID            int64        `json:"user_id"`
	Username          string       `json:"username"`
	FullName          string       `json:"full_name"`
	Email             string       `json:"email"`
	Phone             string       `json:"phone"`
	Avatar            string       `json:"avatar"`
	Country           string       `json:"country"`
	City              string       `json:"city"`
	StreetAddress     string       `json:"street_address"`
	Zip               string       `json:"zip"`
	Status            sql.NullBool `json:"status"`
	HashedPassword    string       `json:"hashed_password"`
	PasswordChangedAt time.Time    `json:"password_changed_at"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         sql.NullTime `json:"updated_at"`
}
