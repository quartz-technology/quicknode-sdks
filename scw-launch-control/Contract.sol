/*
 * SPDX-License-Identifier: MIT
 * Midpoint Sample Contract v2.0.0
 *
 * This is a contract generated at 2022-11-06 00:17:51 for making requests to and receiving responses from midpoint 449.
 * For more information on setting up a midpoint and using this contract see docs.midpointapi.com
 */

pragma solidity>=0.8.0;

interface IMidpoint {
    function callMidpoint(uint64 midpointId, bytes calldata _data) external returns(uint256 requestId);
}

contract VolumeLaunchControl {
    // These events are for demonstration purposes only; they can be removed without effect.
    event RequestMade(uint256 requestId, string zone, string server_name, string volume_type, uint256 volume_size);
    event ResponseReceived(uint256 requestId, string volume_id, string volume_name, string volume_state);

    address constant startpointAddress = 0x0000000000000000000000000000000000000000;
    address constant whitelistedCallbackAddress = 0xC0FFEE4a3A2D488B138d090b8112875B90b5e6D9;

    // Midpoint ID
    uint64 constant midpointID = 449;

    // Owner -> User who deployed the contract
    address owner;

    constructor () {
        owner = msg.sender;
    }

    /*
     * This function makes a call to a midpoint with on-chain variables specified as function inputs.
     *
     * Note that this is a public function and will allow any address or contract to call midpoint 449.
     * The contract whitelist permits this entire contract to call your midpoint; calls to 'callMidpoint'
     * must be additionally restricted to intended callers.
     * Any call to 'callMidpoint' from a whitelisted contract will make a call to the midpoint;
     * there may be multiple places in this contract that call the midpoint or multiple midpoints called by the same contract.
     */

    function callMidpoint(string memory zone, string memory server_name, string memory volume_type, uint256 volume_size) public {
        require(msg.sender == owner);

        // Argument String
        bytes memory args = abi.encodePacked(zone, bytes1(0x00), server_name, bytes1(0x00), volume_type, bytes1(0x00), volume_size);

        // Call Your Midpoint
        uint256 requestId = IMidpoint(startpointAddress).callMidpoint(midpointID, args);

        // For Demonstration Purposes Only
        emit RequestMade(requestId, zone, server_name, volume_type, volume_size);
    }

    /*
     * This function is the callback target specified in the prebuilt function in the midpoint response workflow.
     * The callback does not need to be defined in the same contract as the request.
     */

    function callback(uint256 _requestId, uint64 _midpointId, string memory volume_id, string memory volume_name, string memory volume_state) public {
        // Only allow the verified callback address to submit information for your midpoint.
        require(tx.origin == whitelistedCallbackAddress, "Invalid callback address");
        require(midpointID == _midpointId, "Invalid Midpoint ID");

        // Your callback function here

        // For Demonstration Purposes Only
        emit ResponseReceived(_requestId, volume_id, volume_name, volume_state);
    }
}
