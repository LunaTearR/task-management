use axum::body::to_bytes;
use axum::{
    Router,
    body::Body,
    extract::Request,
    http::{Method, Response, StatusCode},
    routing::{delete, get, post},
};
use reqwest::Client;
use std::net::SocketAddr;
use tower_http::cors::{Any, CorsLayer};

#[tokio::main]
async fn main() {
    let cors = CorsLayer::new()
        .allow_origin(Any)
        .allow_methods([Method::GET, Method::POST, Method::DELETE])
        .allow_headers(Any);

    let app = Router::new()
        // Routes for user-service
        .route("/api/v1/users", get(proxy_service_user))
        .route("/api/v1/users", post(proxy_service_user))
        .route("/api/v1/users", delete(proxy_service_user))
        // Routes for task-service
        .route("/api/v1/tasks", get(proxy_service_task))
        .route("/api/v1/tasks", post(proxy_service_task))
        .route("/api/v1/tasks", delete(proxy_service_task))
        .layer(cors);

    let addr = SocketAddr::from(([0, 0, 0, 0], 8084));
    println!("🚀 API Gateway running on http://{}", addr);

    axum::serve(
        tokio::net::TcpListener::bind(addr).await.unwrap(),
        app.into_make_service(),
    )
    .await
    .unwrap();
}

// Handler สำหรับ task-service
async fn proxy_service_task(req: Request<Body>) -> Result<Response<Body>, StatusCode> {
    proxy_request(req, "http://task-service:3001").await
}

// Handler สำหรับ user-service
async fn proxy_service_user(req: Request<Body>) -> Result<Response<Body>, StatusCode> {
    proxy_request(req, "http://user-service:3002").await
}

// ฟังก์ชันกลางสำหรับส่งต่อ request
async fn proxy_request(
    req: Request<Body>,
    backend_url: &str,
) -> Result<Response<Body>, StatusCode> {
    let client = Client::new();

    let (parts, body) = req.into_parts();
    let method = parts.method.clone();
    let uri = parts.uri.clone();
    let headers = parts.headers.clone();

    let path_and_query = uri.path_and_query().map(|x| x.as_str()).unwrap_or("");
    let url = format!("{}{}", backend_url, path_and_query);
    let body_bytes = to_bytes(body, usize::MAX)
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let mut request_builder = client.request(method, &url).body(body_bytes);

    // ใส่ headers เดิม
    request_builder = request_builder.headers(headers.clone());

    let response = request_builder
        .send()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let status = response.status();
    let bytes = response
        .bytes()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let resp = Response::builder()
        .status(status)
        .body(Body::from(bytes))
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(resp)
}
