## ✅ Questions

**1. What is polymorphism?**

- ○ When two things share properties in common.
- ○ When the definition of a class changes over time.
- ○ When multiple objects have distinct methods.
- ○ When one thing can have multiple forms.

**2. Which of the following statements is true?**

- ○ Inheritance and overriding are required for polymorphism.
- ○ Inheritance and overriding enable polymorphism.
- ○ Overriding is necessary for inheritance.
- ○ Inheritance is necessary for overriding.

**3. If a type satisfies an interface, which of the following statements is true?**

- ○ The type defines all methods specified in the interface.
- ○ The type defines all data specified in the interface.
- ○ The type defines a method specified in the interface.
- ○ The interface includes a definition of the type.

**4. Which of the following statements is true?**

- ○ A concrete type is always a dynamic type.
- ○ An interface always has a dynamic value.
- ○ An interface always has a dynamic type.
- ○ An interface type is the same as a dynamic type.

**5. Which of the following statements is/are true?**  
I. Interfaces can support abstraction by concealing differences between types.  
II. Type assertions can reveal differences between type satisfying an interface.  
iii. Type assertions return two values.

- ○ I and II but NOT III.
- ○ II and III but NOT I.
- ○ I and III but NOT II.
- ○ I, II, and III

**6. What is a use for an empty interface?**

- ○ It allows two interfaces to be merged into one.
- ○ It allows a function to accept a variable number of arguments.
- ○ It can be used to allow a function to accept any type as a parameter.
- ○ An empty interface cannot exist in Go.

**7. After executing the expression below, what is the value of `err` if there is no error?**

```go
f, err := os.Open("/harris/test.txt")
```

- ○ nil
- ○ 0
- ○ -1
- ○ -2

---

## ✅ Answers with Explanations

### Q1. What is polymorphism?

**Answer:** When one thing can have multiple forms.  
**Explanation:** Polymorphism allows the same interface to be used for different underlying forms (data types).

---

### Q2. Which of the following statements is true?

**Answer:** Inheritance and overriding enable polymorphism.  
**Explanation:** These mechanisms allow one method name to be used with different implementations, enabling polymorphism.

---

### Q3. If a type satisfies an interface, which of the following statements is true?

**Answer:** The type defines all methods specified in the interface.  
**Explanation:** A type must implement all the methods declared in an interface to satisfy it.

---

### Q4. Which of the following statements is true?

**Answer:** An interface always has a dynamic type.  
**Explanation:** In Go, interfaces contain a dynamic type and a value, representing the actual concrete type that implements the interface.

---

### Q5. Which of the following statements is/are true?

**Answer:** I, II, and III  
**Explanation:** Interfaces support abstraction, type assertions reveal type info, and assertions return both the value and a success indicator.

---

### Q6. What is a use for an empty interface?

**Answer:** It can be used to allow a function to accept any type as a parameter.  
**Explanation:** An empty interface (`interface{}`) in Go can hold any type, providing maximum flexibility.

---

### Q7. After executing the expression below, what is the value of `err` if there is no error?

**Answer:** nil  
**Explanation:** If there is no error, Go sets the `err` variable to `nil` to indicate successful execution.
