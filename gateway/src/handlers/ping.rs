use axum::{Json};
use serde_json::json;

pub async fn ping_handler() -> Json<serde_json::Value> {
    Json(json!({
        "message": "pong from Rust gateway"
    }))
}
