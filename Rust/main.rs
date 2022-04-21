use crypto::{digest::Digest, sha2};

use base64;
use serde_derive::Deserialize;
use serde_derive::Serialize;
use serde_json;

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct SolutionList {
    pub postfix: i32,
    pub hash: String,
}

fn get_hash(input: &str, padding_length: usize) -> String {
    let padding = "0".repeat(padding_length);
    let mut postfix: i32 = 0;
    loop {
        postfix += 1;
        let stri = input.to_owned() + &postfix.to_string();
        let encoded_hash = gen_sha256(&stri);
        if encoded_hash.starts_with(&padding) {
            let list = SolutionList {
                postfix: postfix,
                hash: encoded_hash,
            };
            let mut solution = [&list; 10];
            for x in 0..10 {
                solution[x] = &list;
            }
            let serialized = serde_json::to_string(&solution).unwrap();
            return base64::encode(serialized);
        }
    }
}

fn gen_sha256(hashme: &str) -> String {
    let mut hasher = sha2::Sha256::new();
    hasher.input_str(hashme);

    return hasher.result_str();
}
