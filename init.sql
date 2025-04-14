CREATE TABLE files (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    size BIGINT NOT NULL,
    metadata JSONB
);

CREATE TABLE live_photos (
    id UUID PRIMARY KEY,
    image_file_id UUID REFERENCES files(id),
    video_file_id UUID REFERENCES files(id),
    metadata JSONB
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);