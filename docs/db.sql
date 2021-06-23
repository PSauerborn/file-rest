CREATE TABLE file_metadata(
    file_id UUID NOT NULL PRIMARY KEY,
    created TIMESTAMP NOT NULL DEFAULT now(),
    size integer NOT NULL,
    metadata JSON NOT NULL DEFAULT '{}',
    archived BOOLEAN NOT NULL DEFAULT false,
    file_name TEXT NOT NULL
)