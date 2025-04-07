# 📝 **List of Questions**

**Q1.** What does a compiler do?

- Generates executable machine code from source code in a high-level language
- Automatically formats source code in a readable way
- Combines multiple files into a single program
- Tests the basic functionality of the code

**Q2.** What is the scope of a variable?

- The variable declaration
- The instructions where the variable is assigned a value
- Region of a program where the variable can be accessed
- The set of values to which a variable can be assigned

**Q3.** What is Garbage Collection?

- Reorganization of source code to reduce the number of function calls
- Deallocation of objects which are no longer in use
- Reorganization of source code to improve encapsulation and understandability
- Deletion of unused segments of code

**Q4.** What does an interpreter do?

- Converts instructions in a high-level language into machine code at runtime
- Processes user inputs
- Reads files in a given data format
- Produces output data corresponding to each input region

**Q5. True or False:**  Concurrency always results in some performance improvement.

- True
- False

**Q6.** The type of a variable determines which of the following aspects of that variable? (Select ALL that are correct.)

- Size in memory
- Operations that can be performed on the variable
- The data that can be contained in the object
- The number of characters in the variable's name

**Q7.** What is the name of the package from which an executable program is generated?

- fmt
- os
- init
- main

**Q8.** Which of the following is an example of a short variable declaration?

- `x := 2.3`
- `var x int = 10`
- `var x int`
- `x = 2.3`

---

### ✅ **Answers**

**Q1.** ✔️ _Generates executable machine code from source code in a high-level language_  
**Explanation:** A compiler translates high-level programming code into machine code (binary) that the computer can execute.

---

**Q2.** ✔️ _Region of a program where the variable can be accessed_  
**Explanation:** The scope of a variable refers to the region in the code where the variable is defined and can be accessed or used.

---

**Q3.** ✔️ _Deallocation of objects which are no longer in use_  
**Explanation:** Garbage collection is an automatic memory management process that frees memory occupied by objects no longer referenced in a program.

---

**Q4.** ✔️ _Converts instructions in a high-level language into machine code at runtime_  
**Explanation:** An interpreter directly executes instructions written in a programming language without compiling them into machine-level code first.

---

**Q5.** ✔️ _False_  
**Explanation:** Concurrency **does not always** result in performance improvement. Improperly managed concurrency can lead to race conditions, deadlocks, or even overhead that reduces performance.

---

**Q6.** ✔️ _Size in memory_, _Operations that can be performed on the variable_, _The data that can be contained in the object_  
**Explanation:** The type of a variable affects:

- How much memory is allocated (e.g., `int` vs `float`)
- What operations are valid (e.g., arithmetic on `int`, not on `string`)
- What kind of data it can store (e.g., `bool` stores true/false)
  It **does not** influence the name length of the variable.

---

**Q7.** ✔️ _main_  
**Explanation:** In languages like Go, an executable program must have a `main` package, which contains the entry point (function `main()`).

---

**Q8.** ✔️ _x := 2.3_  
**Explanation:** The `:=` syntax is used for short variable declaration in languages like Go. It declares and initializes a variable without specifying its type explicitly.

---