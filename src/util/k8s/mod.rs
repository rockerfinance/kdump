pub mod kubectl;

use crate::util; // access all modules between util modules

#[derive(Debug, PartialEq)]
pub struct ApiVersion {
    name: String,
    version: String,
}

#[derive(Debug, PartialEq)]
pub struct ApiResourceType {
    name: String,
    short_names: Vec<String>,
    namespaced: bool,
    kind: String,
    verbs: Vec<String>,
    api_version: ApiVersion,
}

impl ApiResourceType {
    pub fn qualified_name(&self) -> String {
        return if self.api_version.name.is_empty() {
            self.name.clone()
        } else {
            self.name.clone() + "." + &self.api_version.name.clone()
        };
    }
}

pub fn parse_api_version(input: &str) -> ApiVersion {
    let api_version_str_parts = util::string::split_to_vec(input, "/", true);
    return if api_version_str_parts.len() > 1 {
        ApiVersion {
            name: api_version_str_parts[0].to_string(),
            version: api_version_str_parts[1].to_string(),
        }
    } else {
        ApiVersion {
            name: "".to_string(),
            version: api_version_str_parts[0].to_string(),
        }
    };
}