pragma solidity ^0.6.6;

contract RandomNumber {
    uint public randomNumber = 10;

    function setRandomNumber(uint _randomNumber) public {
        randomNumber = _randomNumber;
    }
}
