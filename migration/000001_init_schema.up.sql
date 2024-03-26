-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-03-19T08:48:01.299Z

CREATE TABLE "users" (
    "user_id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "phone" varchar NOT NULL,
    "avatar" varchar NOT NULL,
    "country" varchar NOT NULL,
    "city" varchar NOT NULL,
    "street_address" varchar NOT NULL,
    "zip" varchar NOT NULL,
    "status" bool DEFAULT (true),
    "hashed_password" varchar NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "products" (
    "product_id" bigserial PRIMARY KEY,
    "product_name" varchar NOT NULL,
    "product_thumb" varchar NOT NULL,
    "product_description" varchar,
    "product_slug" varchar,
    "product_price" bigint NOT NULL,
    "product_quantity" bigint NOT NULL,
    "product_type" varchar NOT NULL,
    "product_shop" bigint NOT NULL,
    "product_attributes" JSON NOT NULL,
    "product_ratingsAverage" NUMERIC DEFAULT 4.5,
    "product_variations" JSON DEFAULT '[]'::json,
    "isDraft" Boolean DEFAULT (true),
    "isPublished" Boolean DEFAULT (false),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "discounts" (
    "discount_id" bigserial PRIMARY KEY,
    "discount_name" varchar NOT NULL,
    "discount_description" varchar NOT NULL,
    "discount_type" varchar DEFAULT 'fixed_amount',
    "discount_value" bigint NOT NULL,
    "discount_code" varchar NOT NULL,
    "discount_start_date" timestamptz NOT NULL DEFAULT (now()),
    "discount_end_date" timestamptz NOT NULL,
    "discount_max_uses" bigint NOT NULL,
    "discount_uses_count" bigint NOT NULL,
    "discount_users_used" JSON DEFAULT '[]'::json,
    "discount_min_order_value" bigint NOT NULL,
    "discount_shopId" bigint NOT NULL,
    "discount_is_active" bool DEFAULT (true),
    "discount_applies_to" varchar NOT NULL,
    "discount_product_ids" JSON DEFAULT '[]'::json,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "carts" (
    "cart_id" bigserial PRIMARY KEY,
    "cart_state" varchar NOT NULL,
    "enums" varchar DEFAULT 'active',
    "cart_products" JSON NOT NULL DEFAULT '[]'::json,
    "cart_count_product" bigint DEFAULT (0),
    "cart_UserId" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "orders" (
    "order_id" bigserial PRIMARY KEY,
    "order_userId" bigint NOT NULL,
    "order_checkout" JSON DEFAULT '[]'::json,
    "order_shipping" JSON DEFAULT '[]'::json,
    "order_payment" JSON DEFAULT '[]'::json,
    "order_products" JSON NOT NULL DEFAULT '[]'::json,
    "order_trackingNumber" varchar DEFAULT '#0000103112023',
    "order_status" varchar DEFAULT 'pending',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "inventory" (
    "inven_productId" bigserial NOT NULL,
    "inven_location" varchar DEFAULT 'unKnown',
    "inven_stock" bigint NOT NULL,
    "inven_shopId" bigint NOT NULL,
    "inven_reservations" JSON DEFAULT '[]'::json,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "comments" (
    "comment_id" bigserial PRIMARY KEY,
    "comment_productId" bigint NOT NULL,
    "comment_userId" bigint DEFAULT 1,
    "comment_content" varchar DEFAULT 'text',
    "comment_right" bigint DEFAULT (0),
    "comment_left" bigint DEFAULT (0),
    "comment_parentId" bigint NOT NULL,
    "isDeleted" bool DEFAULT (false),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE TABLE "notifications" (
    "noti_id" bigserial PRIMARY KEY,
    "noti_type" varchar NOT NULL,
    "noti_senderId" bigint NOT NULL,
    "noti_receivedId" bigint NOT NULL,
    "noti_content" varchar NOT NULL,
    "noti_options" JSON DEFAULT '{}'::json,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);

CREATE INDEX ON "products" ("isDraft");

CREATE INDEX ON "products" ("isPublished");

CREATE UNIQUE INDEX ON "products" ("isDraft", "isPublished", "product_shop");

CREATE INDEX ON "discounts" ("discount_shopId");

CREATE INDEX ON "inventory" ("inven_productId");

CREATE INDEX ON "inventory" ("inven_shopId");

CREATE UNIQUE INDEX ON "inventory" ("inven_productId", "inven_shopId");

CREATE INDEX ON "comments" ("comment_productId");

CREATE INDEX ON "comments" ("comment_parentId");

CREATE UNIQUE INDEX ON "comments" ("comment_productId", "comment_parentId");

CREATE INDEX ON "notifications" ("noti_senderId");

COMMENT ON COLUMN "products"."product_type" IS 'Electronics, Clothing, Furniture';

COMMENT ON COLUMN "discounts"."discount_applies_to" IS 'all, specific';

COMMENT ON COLUMN "carts"."enums" IS 'active, completed, fail, pending, lock';

COMMENT ON COLUMN "orders"."order_status" IS 'pending, confirmed, shipped, cancelled, delivered';

COMMENT ON COLUMN "notifications"."noti_type" IS 'ORDER-001, ORDER-002, ORDER-003, SHOP-001, PROMOTION-001';

ALTER TABLE "products" ADD FOREIGN KEY ("product_shop") REFERENCES "users" ("user_id");

ALTER TABLE "discounts" ADD FOREIGN KEY ("discount_shopId") REFERENCES "users" ("user_id");

ALTER TABLE "inventory" ADD FOREIGN KEY ("inven_productId") REFERENCES "products" ("product_id");

ALTER TABLE "inventory" ADD FOREIGN KEY ("inven_shopId") REFERENCES "users" ("user_id");

ALTER TABLE "comments" ADD FOREIGN KEY ("comment_productId") REFERENCES "products" ("product_id");

ALTER TABLE "comments" ADD FOREIGN KEY ("comment_parentId") REFERENCES "comments" ("comment_id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("noti_senderId") REFERENCES "users" ("user_id");
