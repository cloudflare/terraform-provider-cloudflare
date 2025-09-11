---
title: Reentrancy, not last line
---

Member assignemnt just before the transfer, but transfer not on the last line

See case 3 here: https://github.com/runtimeverification/amp/issues/39#issuecomment-1137314683

tags: #reentrancy, #vulnerability, #lowrisk

```grit
and {
  [
      ...,
      `this.$_ = $_` as $memberAccessBefore,
      ...,
      EtherTransfer($amount) as $theTransfer,
      $anotherLine
  ],
  // just a guard so only the ReentrancyBeforeAndAfter matches
  not [
      ...,
      `this.$_ = $_`,
      ...,
      EtherTransfer($amount) as $theTransfer,
      ...,
      `this.$_ = $_`,
      ...
  ]
}
```

## Example

```Solidity
function claim(
    uint256 numPasses,
    uint256 amount,
    uint256 mpIndex,
    bytes32[] calldata merkleProof
) external payable {
    require(isValidClaim(numPasses,amount,mpIndex,merkleProof));

    //return any excess funds to sender if overpaid
    uint256 excessPayment = msg.value.sub(numPasses.mul(mintPasses[mpIndex].mintPrice));
    (bool returnExcessStatus, ) = _msgSender().call{value: excessPayment}("");

    mintPasses[mpIndex].claimedMPs[msg.sender] = mintPasses[mpIndex].claimedMPs[msg.sender].add(numPasses);
    _mint(msg.sender, mpIndex, numPasses, "");
    emit Claimed(mpIndex, msg.sender, numPasses);
}

```
