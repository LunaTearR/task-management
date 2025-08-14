use axum::{extract::Request, body::{Body, to_bytes}, http::{Response, StatusCode}};
use reqwest::Client;

pub async fn proxy_service_task(req: Request<Body>) -> Result<Response<Body>, StatusCode> {
    proxy_request(req, "http://task-service:3001").await
    // proxy_request(req, "http://localhost:3001").await
}

pub async fn proxy_service_user(req: Request<Body>) -> Result<Response<Body>, StatusCode> {
    proxy_request(req, "http://user-service:3002").await
    // proxy_request(req, "http://localhost:3002").await
}

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
    request_builder = request_builder.headers(headers);

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
