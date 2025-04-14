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