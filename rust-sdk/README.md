# QuickNode Rust SDK

A simple way to interact with QuickNode using Rust language

## Dependencies

```toml
error-stack = "0.2.4"
futures = "0.3.25"
graphql_client = { version = "0.11.0", features = ["reqwest"] }
reqwest = "0.11.12"
serde = "1.0.147"
tokio = { version = "1.21.2", features = ["full"] }
```

## Usage

```rust
[tokio::main]
async fn main() -> Result<()> {
    let sdk = QuickNodeSDK::new("");
    let res = sdk
        .get_collection_details("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")
        .await;

    match res {
        Ok(collection_details) => println!("{:?}", collection_details),
        Err(err) => println!("{:?}", err),
    }
}
```

