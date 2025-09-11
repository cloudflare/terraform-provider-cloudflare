---
title: Reentrancy, no assignment before
---

No field assignment before a transfer.

See case 1 here: https://github.com/runtimeverification/amp/issues/39#issuecomment-1137314683

tags: #reentrancy, #vulnerability

```grit
and {
    [ ... contains EtherTransfer($amount) ],
    not [
        ...,
        `this.$x = $y`,
        ...,
        EtherTransfer($amount),
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
