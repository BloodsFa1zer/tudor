CREATE TABLE "users" (
    "id" BIGSERIAL NOT NULL,
    "name" VARCHAR,
    "email" VARCHAR NOT NULL UNIQUE,
    "photo" VARCHAR,
    "verified" BOOLEAN NOT NULL DEFAULT FALSE,
    "password" VARCHAR,
    "role" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");