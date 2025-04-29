use axum::{Json, extract::State, http::StatusCode};
use serde_json::json;
use sqlx::Row;

use crate::models::{payload::Payload, state::App};

pub async fn notify_user(
    State(state): State<App>,
    Json(payload): Json<Payload>,
) -> Result<(), (StatusCode, String)> {
    let bot_api_url = format!(
        "https://api.telegram.org/bot{}/sendPhoto",
        state.config.bot_id
    );

    let record = sqlx::query(
        "
        SELECT telegram_chat_id
        FROM users
        WHERE jellyfin_username = $1
        ",
    )
    .bind(payload.request.requested_by_username)
    .fetch_optional(&state.pool)
    .await;

    let telegram_chat_id = match record {
        Ok(row) => match row {
            Some(row) => row.get::<Option<i64>, _>("telegram_chat_id"),
            None => {
                return Err((StatusCode::NOT_FOUND, "User not found".to_string()));
            }
        },
        Err(err) => {
            return Err((
                StatusCode::INTERNAL_SERVER_ERROR,
                format!("Error fetching telegram chat id: {err}"),
            ));
        }
    };

    let Some(telegram_chat_id) = telegram_chat_id else {
        return Err((
            StatusCode::NOT_FOUND,
            "Telegram chat ID not found for user".to_string(),
        ));
    };

    let body = json!(
           {
               "chat_id": telegram_chat_id,
               "photo": payload.image,
               "caption": format!("New {} available: {}", payload.media.media_type, payload.subject),
           }
    );

    let res = state.http_client.post(bot_api_url).json(&body).send().await;

    match res {
        Ok(response) => {
            if response.status().is_success() {
                Ok(())
            } else {
                Err((
                    StatusCode::INTERNAL_SERVER_ERROR,
                    format!("Failed to send notification: {}", response.status()),
                ))
            }
        }
        Err(err) => Err((
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Error sending notification: {err}"),
        )),
    }
}
