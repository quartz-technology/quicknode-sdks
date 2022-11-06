use graphql_client::{reqwest::post_graphql, GraphQLQuery, Response};
use std::error::Error;
use reqwest;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/ExampleQuery.gql",
    response_derives = "Debug",
)]
pub struct ExampleQuery;

pub async fn send_request() -> Result<(), Box<dyn Error>> {
    let client = reqwest::Client::new();
    let vars = example_query::Variables {
        address: String::from("0x60E4d786628Fea6478F785A6d7e704777c86a7c6"),
    };

    let request_body = ExampleQuery::build_query(vars);

    let res = client.post("https://graphql.icy.tools/graphql")
        .json(&request_body)
        .send()
        .await?;
    let response_body: Response<example_query::ResponseData> = res.json().await?;
    println!("{:#?}", response_body);
    

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn it_works() {
        match send_request().await {
            Ok(_) => (),
            Err(e) => { println!("{:?}", e.as_ref()) }
        }
    }
}
