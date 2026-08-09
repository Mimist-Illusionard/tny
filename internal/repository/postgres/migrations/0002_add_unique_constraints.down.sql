ALTER TABLE urls
    DROP CONSTRAINT IF EXISTS urls_short_unique,
    DROP CONSTRAINT IF EXISTS urls_original_unique;
