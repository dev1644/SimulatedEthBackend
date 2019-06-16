pragma solidity ^0.5.7;

contract Challenge {

    address public a;
    uint256 public b;

    function setA() public returns (bool) {
        a = msg.sender;
        return (a == msg.sender);
    }

   function setB(uint256 _b) public returns (bool) {
        b = _b;
        return (b == _b);
    }

    function getA() public view returns (address) {
        return a;
    }

    function getB() public view returns (uint256) {
        return b;
    }
}