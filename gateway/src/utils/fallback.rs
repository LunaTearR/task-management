use axum::{
    http::StatusCode,
    response::IntoResponse,
    Json,
};
use serde::Serialize;

#[derive(Serialize)]
struct ErrorResponse {
    status: u16,
    message: String,
}

pub async fn not_found() -> impl IntoResponse {
    let error = ErrorResponse {
        status: StatusCode::NOT_FOUND.as_u16(),
        message: "API endpoint not found".to_string(),
    };

    (StatusCode::NOT_FOUND, Json(error))
}
