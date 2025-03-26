# Well-Written Functions

## Understandability

- code is functions and data
- if you are asked to find a feature, you can find it quickly
  - "Where is the function that blurs the image?"
- if you are asked about where data is used, you know
  - "Where do you modify the record list?"

## Debugging Principles

- code crashes inside a function
- two options for the cause:
  - function is written incorrectly
    - sorts a slice in the wrong order
  - data that the function uses is incorrect
    - sorts slice correctly but slice has wrong elements in it

## Supporting Debugging

- functions need to be understandable
  - determine if actual behavior matches desired behavior
- data needs to be traceable
  - where did that data come from?
  - global variables complicate this
