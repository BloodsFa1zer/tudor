CREATE TABLE "advertisements" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "provider_id" BIGINT NOT NULL,
    "attachment" VARCHAR NOT NULL,
    "experience" INT NOT NULL, 
    "category_id" BIGINT NOT NULL,
    "time" INT NOT NULL,
    "price" INT NOT NULL,
    "format" VARCHAR NOT NULL,
    "language" VARCHAR NOT NULL,
    "description" TEXT NOT NULL,
    "mobile_phone" VARCHAR,
    "email" VARCHAR,
    "telegram" VARCHAR,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);