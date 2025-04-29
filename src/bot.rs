use sqlx::postgres::PgPool;
use teloxide::{prelude::*, utils::command::BotCommands};

#[derive(BotCommands, Clone)]
#[command(
    rename_rule = "lowercase",
    description = "These commands are supported:"
)]
pub enum Command {
    #[command(description = "Set your username.")]
    Username(String),
}

pub async fn answer(bot: Bot, msg: Message, cmd: Command, pool: PgPool) -> ResponseResult<()> {
    match cmd {
        Command::Username(username) => {
            let record = sqlx::query(
                "
                SELECT telegram_chat_id
                FROM users
                WHERE jellyfin_username = $1
                ",
            )
            .bind(&username)
            .fetch_optional(&pool)
            .await;

            if let Ok(maybe_row) = record {
                match maybe_row {
                    Some(_) => {
                        let _ = sqlx::query(
                            "
                            UPDATE users
                            SET telegram_chat_id = $1
                            WHERE jellyfin_username = $2
                            ",
                        )
                        .bind(msg.chat.id.0)
                        .bind(username)
                        .fetch_optional(&pool)
                        .await;
                    }
                    None => {
                        let _ = sqlx::query(
                            "
                            INSERT INTO users (telegram_chat_id, jellyfin_username)
                            VALUES ($1, $2)
                            ",
                        )
                        .bind(msg.chat.id.0)
                        .bind(&username)
                        .fetch_optional(&pool)
                        .await;
                    }
                }
            }

            bot.send_message(msg.chat.id, "Username set".to_string())
                .await?
        }
    };

    Ok(())
}
