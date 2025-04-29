CREATE TABLE IF NOT EXISTS users (
    jellyfin_username VARCHAR(255) NOT NULL,
    telegram_chat_id BIGINT NOT NULL,
    PRIMARY KEY (jellyfin_username, telegram_chat_id)
);