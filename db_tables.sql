CREATE TABLE IF NOT EXISTS public.user
(
    id            uuid PRIMARY KEY,
    name          text NOT NULL,
    telegram_name text NOT NULL,
    age           text,
    gender        text,
    created_at    text NOT NULL,
    updated_at    text,
    status        text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.book
(
    id               uuid PRIMARY KEY,
    user_id          uuid NOT NULL,
    name             text NOT NULL,
    file_path        text,
    text_file_path   text,
    book_settings_id uuid,
    loaded           text
);

CREATE TABLE IF NOT EXISTS public.book_settings
(
    id           uuid PRIMARY KEY,
    book_id      uuid NOT NULL,
    current_page int,
    page_in_day  int,
    updated_at   text,
    book_status  text
);