ALTER TABLE todos ADD COLUMN created_at timestamp default current_timestamp;
ALTER TABLE todos ADD COLUMN updated_at timestamp;
ALTER TABLE todos ADD COLUMN deleted_at timestamp;
