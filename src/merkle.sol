// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract Merkle {

    bytes32 public MerkleRoot = 0x0000000000000000000000000000000000000000000000000000000000000100 ;
    mapping(address => bool) public whiteListClaimed;

    function whiteListMint(bytes32[] calldata _merkleProof) public{

        require( !whiteListClaimed[msg.sender], "address has already claimed");
        bytes32 leaf = keccak256(abi.encodePacked(msg.sender));
        require(MerkleProof.verify(_merkleProof,MerkleRoot,leaf),"invalid proof");

        whiteListClaimed[msg.sender] = true;
    } 
}