# Learning java


### Strings
In java it's recomendable to use interning if we want to use == comparing two string

```bash
String s1 = "I love";
s1 += " Java";
String s2 = "I";
s2 += " love Java";


if (s1 == s2) // false

String s3 = s1.intern();
String s4 = s2.intern();

if (s3 == s4) // true
```
