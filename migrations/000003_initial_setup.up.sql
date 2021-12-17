ALTER TABLE todos DROP COLUMN id;
ALTER TABLE todos ADD COLUMN id UUID not null primary key;
