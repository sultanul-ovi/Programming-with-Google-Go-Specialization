# Hash Table

- contains key/value pairs
  - SSN/email: ssn is key and email is value
  - GPS cords/address
- each value is associated with a UNIQUE key
- hash function is used to compute the slot for a key
  - used to process the key and returns the corresponding value

## Tradeoffs of hash tables

### Advantages

- faster lookup than lists
  - contant-time vs linear time
- arbitrary keys
  - not ints, like slices or arrays

### Disadvantages

- may have collisions
  - two keys hash to same slot
  