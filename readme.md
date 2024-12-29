# Breaking-Consensus Change
Here breaking-consensus change is introduced to the blockchain in this version.

## What is Breaking-Consensus Change
Breaking-consensus change occurs when modification is made to the blockchain such that the nodes running in the older version and nodes running in the newer version are expected to unable to interact to one another due to different data formats or protocol logic implemented in both version.

## Current Breaking-Consensus Change Introduced
In this version, a tag field is added to the create-post function. 
If using the previous version command, e.g., 
```
cruded tx crude create-post hello "jake testing" --from alice --chain-id crude
```
the transaction will fail as the blockchain is expecting tags.
Whereas the new command, e.g.,
```
cruded tx crude create-post hello "jake testing" --tags "tag1,tag2" --from alice --chain-id crude
```
will work as tag is added.

### Why this break consensus
- Older nodes: Do not recognise or expect the tag field hence the transaction will fail.
- Newer nodes: Requires tags field for validation. Transaction without the tags field will be rejected.
Hence, older nodes can't interact with new nodes due to differences in the data format.
