

## A Quick Look at Trait Objects in Rust

原文: [A Quick Look at Trait Objects in Rust](https://tratt.net/laurie/blog/entries/a_quick_look_at_trait_objects_in_rust.html)



### 基础

- 函数调用的分发方式

  - 静态分发(static dispatch): 编译期已经确定

  - 动态分发(dynamic dispatch): 运行时才能确定

    如 OO 类语言中的类和子类

- Rust 中, trait **可能**会引起动态分发

  - 这个场景中, s 有确定的类型 S, 而 S 实现了 T, 因而拥有 `T::m` 这个方法. 

    这些在编译期即可确定, 因此是静态分发.

    ```
    trait T {
      fn m(&self) -> u64;
    }
    
    struct S {
      i: u64
    }
    
    impl T for S {
      fn m(&self) -> u64 { self.i }
    }
    
    fn main() {
      let s = S{i : 100};
      println!("{}", s.m());
    }
    ```

  - 下面的代码是无法编译的, 因为编译期无法确定 x 的大小.

    ```
    fn f(x: T) {
      println!("{}", x.m())
    }
    ```

    ```
    error[E0277]: the size for values of type `(dyn T + 'static)` cannot be known at compilation time
      --> src/main.rs:21:6
       |
    21 | fn f(x: T) {
       |      ^ doesn't have a size known at compile-time
       |
       = help: the trait `std::marker::Sized` is not implemented for `(dyn T + 'static)`
       = note: to learn more, visit <https://doc.rust-lang.org/book/second-edition/ch19-04-advanced-types.html#dynamically-sized-types-and-the-sized-trait>
       = note: all local variables must have a statically known size
       = help: unsized locals are gated as an unstable feature
    ```

  - 对于这个场景, rust 会为每种使用到的 X 编译一个对应的 f, 因此仍然是静态分发的.

    但是这种方式会导致编译结果的膨胀.

    ```
    fn f<X: T>(x: X) {
      println!("{}", x.m())
    }
    ```

  - 这个场景将会触发动态分发.

    由于`&T` 的大小可以确定, 因此 `f` 只会被编译成一个版本.

    此时, 调用的具体路径将由运行时根据附带的额外信息来确定.

    ```
    trait T {
      fn m(&self) -> u64;
    }
    
    struct S1 {
      i: u64
    }
    
    impl T for S1 {
      fn m(&self) -> u64 { self.i * 2 }
    }
    
    struct S2 {
      j: u64
    }
    
    impl T for S2 {
      fn m(&self) -> u64 { self.j * 4 }
    }
    
    fn f(x: &T) {
      println!("{}", x.m())
    }
    
    fn main() {
      let s1 = S1{i : 100};
      f(&s1);
      let s2 = S2{j : 100};
      f(&s2);
    }
    ```

    但是对于代码阅读者来说, 如果他事先不知道 `T` 是一个 trait 而非 struct, 那么他不会意识到运行时的差异.

    因此 rust 会让你添加一定的语法标识来注明, 变成:

    ```
    fn f(x: &dyn T) {
      println!("{}", x.m())
    }
    ```

    `dyn` 标识没有语义方面的影响, 而没有这个标识并不代表不会发生动态分发.

  - 另一个不使用引用而有同样效果的实现方法是使用`Box`:

    ```
    fn f2(x: Box<T>) {
      println!("{}", x.m())
    }
    
    fn main() {
      let b: Box<S1> = Box::new(S1{i: 100});
      f2(b);
    }
    ```



### 指针 (fat pointers vs. inner vpointers)

- A `fat pointer` is simply a pointer-plus-some-other-stuff, so is at least two machine words big.

- a object’s `vtable` is a list of pointers to a struct’s dynamically dispatched functions

- a pointer to a vtable is a `vpointer`

- 对于 `vpointers`, 存在两种处理方式:

  1. 由对象携带, 即文中所称的 `inner vpointers`
  2. 由指针携带.

  因此, 用于动态分发的指针中额外的 machine word 存在两种方案:

  1. to pointers (fat pointers) 
  2. to objects (inner vpointers)

- 如果选择使用 inner vpointers, 每个对象都需要携带. 而在 rust 的场景中, 可能绝大多数情况这些 inner vpointers 都不会被使用到.

  因此 rust 选择使用 fat pointers, 这样只会在确定使用动态分发的情况下带来额外的开销.



### 不同指针实现方式带来的性能差异

略

