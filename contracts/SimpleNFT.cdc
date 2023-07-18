import FungibleToken from "./FungibleToken.cdc"
import NonFungibleToken from "./NonFungibleToken.cdc"

// Define the NFT contract structure
pub contract MyNFT: NonFungibleToken.Contract {
    pub let metadata: {String: String}
    pub let collection: &{String: NonFungibleToken.Collection}

    init() {
        self.metadata = {}
        self.collection = {}
    }

    pub fun createToken(metadataURI: String): UInt64 {
        let tokenID = self.collection[self.account.id]?.supply ?? 0
        let newTokenID = tokenID + 1

        let token = NonFungibleToken.Token(
            id: newTokenID,
            metadata: {"name": "My NFT", "description": "A custom NFT"},
            owner: self.account.owner
        )

        self.mintNonFungibleToken(self.account, token: token)

        self.metadata[newTokenID.toString()] = metadataURI

        return newTokenID
    }

    pub fun getTokenMetadata(tokenID: UInt64): String {
        return self.metadata[tokenID.toString()] ?? ""
    }
}