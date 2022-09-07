pragma solidity ^0.8.15;

contract ZkBNB {
    function depositBNB(string calldata _accountName) external payable {}

    function depositBEP20(address _token, uint104 _amount, string calldata _accountName) external  {}

    function depositNft(
        string calldata _accountName,
        address _nftL1Address,
        uint256 _nftL1TokenId
    ) external {}

    function registerZNS(string calldata _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) external payable {}

    function createPair(address _tokenA, address _tokenB) external {}

    function requestFullExit(string calldata _accountName, address _asset) public {}

    function requestFullExitNft(string calldata _accountName, uint32 _nftIndex) public {}

    struct PairInfo {
        address tokenA;
        address tokenB;
        uint16 feeRate;
        uint32 treasuryAccountIndex;
        uint16 treasuryRate;
    }

    function updatePairRate(PairInfo memory _pairInfo) external {}
}
