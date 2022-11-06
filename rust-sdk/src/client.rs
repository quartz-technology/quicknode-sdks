use error_stack::{IntoReport, Report, ResultExt};
use graphql_client::{GraphQLQuery, Response};
use std::error::Error;
use std::fmt;
use std::fmt::{Debug, Display};

#[derive(Debug)]
pub enum QuickNodeError {
    Deserialization,
    Request,
    GraphQL(Vec<graphql_client::Error>),
}

impl Display for QuickNodeError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.write_str("QuickNode error")
    }
}

impl Error for QuickNodeError {}

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetCollectionDetailsQuery.gql",
    response_derives = "Debug"
)]
pub struct CollectionDetails;

// #[derive(GraphQLQuery)]
// #[graphql(
//     schema_path = "src/schema.gql",
//     query_path = "src/queries/GetContractEventLogs.gql",
//     response_derives = "Debug"
// )]
// pub struct ContractEvents;


#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTDetails.gql",
    response_derives = "Debug"
)]
pub struct NFTDetails;


#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTsByENS.gql",
    response_derives = "Debug"
)]
pub struct WalletNFTsByENS;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTsByWalletAddress.gql",
    response_derives = "Debug"
)]
pub struct WalletNFTs;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTsWalletAndContracts.gql",
    response_derives = "Debug"
)]
pub struct NFTsWalletAndContract;

struct TokensFilterInputType {
    contractAddressIn: Vec<String>,
    name: String,  
}

pub struct QuickNodeSDK {
    base_url: String,
    icy_api_key: String,
    client: reqwest::Client,
}

impl QuickNodeSDK {
    pub fn new(icy_api_key: &str) -> Self {
        Self {
            base_url: String::from("https://graphql.icy.tools/graphql"),
            icy_api_key: icy_api_key.to_string(),
            client: reqwest::Client::new(),
        }
    }

    pub fn set_icy_api_key(&mut self, key: String) {
        self.icy_api_key = key;
    }

    pub async fn get_collection_details(
        self,
        address: &str,
    ) -> Result<collection_details::ResponseData, Report<QuickNodeError>> {
        let vars = collection_details::Variables {
            address: address.to_string(),
        };

        let request_body = CollectionDetails::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<collection_details::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }

    // pub async fn get_contract_event_logs(self) {}

    pub async fn get_nft_details(self, contract_address: &str, token_id: &str) {}

    pub async fn get_nfts_by_ens(self, ens_name: &str, first: u32, after: &str) {}

    pub async fn get_nfts_by_wallet_address(self, address: &str, first: u32, after: &str) {}

    pub async fn get_nfts_wallet_and_contracts(self, filter: TokensFilterInputType, address: &str, after: &str) {

    }

}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn it_works() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_collection_details("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")
            .await;

        match res {
            Ok(collection_details) => println!("{:?}", collection_details),
            Err(err) => println!("{:?}", err),
        }
    }
}
