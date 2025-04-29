use crate::models::config::Config;
use sqlx::postgres::PgPool;

#[derive(Clone)]
pub struct App {
    pub pool: PgPool,
    pub http_client: reqwest::Client,
    pub config: Config,
}
