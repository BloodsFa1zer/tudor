CREATE TABLE "advertisements" (
    "id" BIGSERIAL NOT NULL,
    "title" VARCHAR NOT NULL,
    "provider" VARCHAR NOT NULL,
    "provider_id" BIGINT NOT NULL,
    "attachment" VARCHAR,
    "experience" INT NOT NULL, 
    "category_id" BIGINT NOT NULL,
    "time" INT NOT NULL,
    "price" INT NOT NULL,
    "format" VARCHAR NOT NULL,
    "language" VARCHAR NOT NULL,
    "description" TEXT NOT NULL,
    "mobile_phone" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "telegram" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "advertisement_pkey" PRIMARY KEY ("id")
);