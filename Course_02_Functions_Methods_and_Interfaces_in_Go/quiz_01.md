## ✅ Questions

**1. Using functions in code has which of the following impacts?**

- ○ Improve code understandability
- ○ Facilitate code reuse
- ○ Support abstraction
- ○ All of the above

**2. Which of these statements is true?**

- ○ A function can have only one return value.
- ○ A function cannot have more than two parameters.
- ○ A function can have parameters of different types.
- ○ The type of the arguments do not need to be specified.

**3. Let’s say that you are writing a function which takes a structure as an argument. What is a good reason to rewrite the function to take a pointer to the structure as an argument, instead of the structure itself?**

- ○ The function needs to modify the structure.
- ○ The function needs to read data inside the structure.
- ○ The function needs to copy the structure.
- ○ The structure uses very little space in memory.

**4. Which of the features of functions listed below improve code understandability?**  
I. Low number of arguments  
II. Performs several distinct tasks  
III. Low control-flow complexity

- ○ I and II
- ○ I and III
- ○ II and III
- ○ I, II, and III

**5. What is a difference between passing a slice as an argument and passing an array as an argument?**

- ○ Passing a slice passes a copy of all the data in the slice.
- ○ Passing an array is faster than passing a slice.
- ○ There is no difference.
- ○ Passing an array passes a copy of the entire array.

---

## ✅ Answers with Explanations

### Q1. Using functions in code has which of the following impacts?

**Answer:** All of the above  
**Explanation:** Functions improve code readability, allow reuse of logic in different places, and help break the code into abstract logical blocks. Hence, all the listed options are valid impacts.

---

### Q2. Which of these statements is true?

**Answer:** A function can have parameters of different types.  
**Explanation:** Functions are designed to handle diverse input, so they can accept multiple parameters of varying types (e.g., `int`, `float`, `string`, etc.) depending on the function signature.

---

### Q3. Let’s say that you are writing a function which takes a structure as an argument. What is a good reason to rewrite the function to take a pointer to the structure as an argument, instead of the structure itself?

**Answer:** The function needs to modify the structure.  
**Explanation:** Passing by pointer allows the function to directly access and modify the original structure. If passed by value, only a copy is modified, leaving the original unchanged.

---

### Q4. Which of the features of functions listed below improve code understandability?

**Answer:** I and III  
**Explanation:** Having fewer arguments and low control-flow complexity makes functions easier to understand and maintain. Performing several distinct tasks in a single function usually reduces clarity.

---

### Q5. What is a difference between passing a slice as an argument and passing an array as an argument?

**Answer:** Passing an array passes a copy of the entire array.  
**Explanation:** In many languages, arrays are passed by value, resulting in a full copy, while slices are typically references to the original data, making them more efficient and avoiding unnecessary copying.
