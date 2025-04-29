use bot::{Command, answer};
use dotenv::dotenv;
use models::{config::Config, state::App};
use sqlx::postgres::PgPool;
use std::error::Error;
use teloxide::prelude::*;
use tracing::info;

mod bot;
mod handlers;
mod models;
mod router;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().ok();
    tracing_subscriber::fmt::fmt().pretty().init();

    let config = Config {
        bot_id: dotenv::var("TELOXIDE_TOKEN").expect("TELOXIDE_TOKEN must be set"),
    };

    let db_url = dotenv::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = PgPool::connect(&db_url).await?;

    sqlx::migrate!().run(&pool).await?;

    let http_client = reqwest::Client::builder().build()?;

    let app_state = App {
        pool: pool.clone(),
        http_client,
        config,
    };

    tokio::spawn(async move {
        let bot = Bot::from_env();
        Command::repl(bot, move |bot, msg, cmd| {
            answer(bot, msg, cmd, pool.clone())
        })
        .await;
    });

    let router = router::init().with_state(app_state);

    let address = format!(
        ":{}",
        dotenv::var("SERVER_PORT").expect("SERVER_PORT must be set")
    );
    let listener = tokio::net::TcpListener::bind(&address).await?;

    info!("Starting server on {}", address);
    axum::serve(listener, router).await.unwrap();

    Ok(())
}
