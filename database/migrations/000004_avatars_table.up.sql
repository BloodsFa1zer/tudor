CREATE TABLE avatars (
    id BIGSERIAL PRIMARY KEY,
    filename TEXT NOT NULL,
    file_address TEXT NOT NULL,
    data BLOB NOT NULL,
    provider_id BIGINT NOT NULL UNIQUE
);

CREATE INDEX avatars_provider_id_idx ON avatars (provider_id);
