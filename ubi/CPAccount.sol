// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CPAccount {
    address public owner;
    uint8 public windowPoStProofType;
    string public peerId;
    string[] public multiAddresses;

    struct Beneficiary {
        address beneficiaryAddress;
        uint256 quota;
        uint256 expiration;
    }

    Beneficiary public beneficiary;

    struct Task {
        string taskId;
        uint8 taskType;
        string proof;
        bool isSubmitted;
    }

    mapping(string => Task) public tasks;

    constructor(
        address _owner,
        uint8 _windowPoStProofType,
        string memory _peerId,
        string[] memory _multiAddresses
    ) {
        owner = _owner;
        windowPoStProofType = _windowPoStProofType;
        peerId = _peerId;
        multiAddresses = _multiAddresses;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }

    function changeMultiaddrs(string[] memory newMultiaddrs) public onlyOwner {
        multiAddresses = newMultiaddrs;
    }

    function changeOwnerAddress(address newOwner) public onlyOwner {
        owner = newOwner;
    }

    function changeBeneficiary(
        address newBeneficiary,
        uint256 newQuota,
        uint256 newExpiration
    ) public onlyOwner {
        beneficiary = Beneficiary({
        beneficiaryAddress: newBeneficiary,
        quota: newQuota,
        expiration: newExpiration
        });
    }

    function submitUBIProof(string memory _taskId, uint8 _taskType, string memory _proof) public onlyOwner {
        require(!tasks[_taskId].isSubmitted, "Proof for this task is already submitted.");
        tasks[_taskId] = Task({
        taskId: _taskId,
        taskType: _taskType,
        proof: _proof,
        isSubmitted: true
        });
    }
}