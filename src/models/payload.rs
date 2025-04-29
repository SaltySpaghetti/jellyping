#![allow(warnings)]

use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Payload {
    pub subject: String,
    pub image: String,
    notification_type: String,
    event: String,
    message: String,
    pub media: Media,
    pub request: Request,
    issue: Option<serde_json::Value>,
    comment: Option<serde_json::Value>,
    extra: Vec<Option<serde_json::Value>>,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Media {
    #[serde(rename = "media_type")]
    pub media_type: String,
    tmdb_id: String,
    tvdb_id: String,
    status: String,
    #[serde(rename = "status4k")]
    status4_k: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Request {
    request_id: String,
    #[serde(rename = "requestedBy_email")]
    requested_by_email: String,
    #[serde(rename = "requestedBy_username")]
    pub requested_by_username: String,
    #[serde(rename = "requestedBy_avatar")]
    requested_by_avatar: String,
    #[serde(rename = "requestedBy_settings_discordId")]
    requested_by_settings_discord_id: String,
    #[serde(rename = "requestedBy_settings_telegramChatId")]
    requested_by_settings_telegram_chat_id: String,
}
