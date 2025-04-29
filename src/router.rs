use crate::handlers::users::notify_user;
use crate::models::state::App;

use axum::Router;
use axum::routing::post;

pub fn init() -> Router<App> {
    Router::new().nest(
        "/jellyping",
        Router::new().nest(
            "/v1",
            Router::new().route("/notify-user", post(notify_user)),
        ),
    )
}
