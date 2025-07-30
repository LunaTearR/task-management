use axum::{Router};
use crate::handlers::ping;

#[allow(dead_code)]
pub fn create_router() -> Router {
    Router::new().route("/ping", axum::routing::get(ping::ping_handler))
}
