pragma solidity ^0.6.6;

contract EV3SpeedMonitor {
    string public deviceId;
    bool public running = false;
    uint32 public speed = 0;

    constructor(string memory _deviceId) public {
        deviceId = _deviceId;
    }

    function start(string calldata _deviceId) public {
        deviceId = _deviceId;
        running = true;
    }

    function stop() public {
        running = false;
    }

    function update(uint32 _speed) public {
        speed = _speed;
    }
}
