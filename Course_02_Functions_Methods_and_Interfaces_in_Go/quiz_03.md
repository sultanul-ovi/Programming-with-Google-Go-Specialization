## ✅ Questions

**1. What is the difference between an object and a class?**

- ○ An object is a field of data inside a class.
- ○ A class is a template and an object is an instance of that template.
- ○ An object is a particular kind of class.
- ○ An object typically contains more data fields than a class.

**2. What is the difference between a struct in Go and a class in an object-oriented language?**

- ○ A struct contains only data while a class can also contain methods.
- ○ A class describes data fields but a struct does not.
- ○ A struct can only be created inside a class.
- ○ A struct cannot contain another struct.

**3. Which of the following refers to data hiding?**

- ○ Instantiation
- ○ Polymorphism
- ○ Inheritance
- ○ Encapsulation

**4. How do you associate a method with an arbitrary data type on Go?**

- ○ Define the method so that its receiver type is the data type of interest.
- ○ Define the method inside the data type definition.
- ○ Include the name of the data type in the name of the method.
- ○ Define the data type and the method in the same file.

**5. In Go, how do you hide variables or functions in a package, so that functions outside of the package cannot access them?**

- ○ Use the package keyword
- ○ Use the private keyword.
- ○ Give the variable/function a name which starts with a lower-case letter
- ○ Define the variable/function inside the package.

**6. Say that you have defined a type `t` and you have declared an object of that type called `t1`. Assume that the type `t` is the receiver type for a method called `Foo()`. Which expression shows a proper invocation of the method `Foo()`?**

- ○ Foo(t1)
- ○ Foo(t)
- ○ t1.Foo()
- ○ t.Foo(t1)

**7. Assume that the type `t` is the receiver type for a method called `Foo()`. Under what conditions would it be better to make the receiver type of `Foo()` a pointer to `t`, rather than itself?**  
I. When the receiver type `t` uses a large amount of memory.  
II. When the method `Foo()` must modify the data in the object of the receiver type.

- ○ Only I
- ○ Only II
- ○ Both I and II
- ○ Neither I nor II

---

## ✅ Answers with Explanations

### Q1. What is the difference between an object and a class?

**Answer:** A class is a template and an object is an instance of that template.  
**Explanation:** A class defines the blueprint for data and behavior; an object is a concrete instance based on that blueprint.

---

### Q2. What is the difference between a struct in Go and a class in an object-oriented language?

**Answer:** A struct contains only data while a class can also contain methods.  
**Explanation:** Go's structs are purely for data storage, unlike classes in OOP languages which can also encapsulate behavior.

---

### Q3. Which of the following refers to data hiding?

**Answer:** Encapsulation  
**Explanation:** Encapsulation involves bundling data and methods and restricting access to some components for security and integrity.

---

### Q4. How do you associate a method with an arbitrary data type on Go?

**Answer:** Define the method so that its receiver type is the data type of interest.  
**Explanation:** Go allows attaching methods to types using a receiver, which binds the function to that type.

---

### Q5. In Go, how do you hide variables or functions in a package, so that functions outside of the package cannot access them?

**Answer:** Give the variable/function a name which starts with a lower-case letter  
**Explanation:** In Go, identifiers starting with a lowercase letter are unexported and only accessible within the same package.

---

### Q6. Say that you have defined a type `t` and you have declared an object of that type called `t1`. Assume that the type `t` is the receiver type for a method called `Foo()`. Which expression shows a proper invocation of the method `Foo()`?

**Answer:** t1.Foo()  
**Explanation:** The method is invoked on the instance `t1` using dot notation, standard in Go.

---

### Q7. Assume that the type `t` is the receiver type for a method called `Foo()`. Under what conditions would it be better to make the receiver type of `Foo()` a pointer to `t`, rather than itself?

**Answer:** Both I and II  
**Explanation:** Using a pointer saves memory for large structs and allows methods to modify the original data structure.
