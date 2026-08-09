ALTER TABLE urls
    ADD CONSTRAINT urls_short_unique UNIQUE (short),
    ADD CONSTRAINT urls_original_unique UNIQUE (original);
