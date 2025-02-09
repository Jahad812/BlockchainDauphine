// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {DirectFundingConsumer} from "./DirectFundingConsumer.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Coinflip is Ownable {
    mapping(address => uint256) public playerRequestID;
    mapping(address => uint8[3]) public bets;
    mapping(uint256 => address) public requestToPlayer;
    mapping(uint256 => uint256[3]) public randomResults;

    DirectFundingConsumer private vrfRequestor;
    address private constant LINK_TOKEN = 0x779877A7B0D9E8603169DdbD7836e478b4624789;
    
    constructor() Ownable(msg.sender) {
        vrfRequestor = new DirectFundingConsumer();
    }
    
    function fundOracle() external returns (bool) {
        IERC20 link = IERC20(LINK_TOKEN);
        require(link.balanceOf(address(this)) >= 5 ether, "Not enough LINK tokens");
        require(link.transfer(address(vrfRequestor), 5 ether), "Funding failed");
        return true;
    }
    
    function userInput(uint8[3] calldata Guesses) external {
        require(Guesses[0] <= 1 && Guesses[1] <= 1 && Guesses[2] <= 1, "Invalid guesses");
        bets[msg.sender] = Guesses;
        uint256 requestId = vrfRequestor.requestRandomWords(false);  // Request randomness
        playerRequestID[msg.sender] = requestId;
        requestToPlayer[requestId] = msg.sender;
    }
    
    function checkStatus() external view returns (bool) {
        uint256 requestId = playerRequestID[msg.sender];
        require(requestId != 0, "No request found");

        ( , bool fulfilled, ) = vrfRequestor.getRequestStatus(requestId);  // Use getRequestStatus()
        return fulfilled;
    }
    
    function determineFlip() external view returns (bool) {
        uint256 requestId = playerRequestID[msg.sender];
        (, bool fulfilled, uint256[] memory randomWords) = vrfRequestor.getRequestStatus(requestId);
        require(fulfilled, "Randomness not fulfilled");

        require(randomWords.length == 3, "Invalid random numbers");
        uint8[3] memory userGuesses = bets[msg.sender];
        
        for (uint8 i = 0; i < 3; i++) {
            uint8 flipResult = randomWords[i] % 2 == 0 ? 0 : 1;
            if (flipResult != userGuesses[i]) {
                return false;
            }
        }
        return true;
    }
}
