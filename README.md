---
# mlParser
---

This package contains the necessary code to evaluate mathematical fucntions at runtime


### Feature requests
- [ ] add support for ** power operator (probably just a regexp replace)
- [ ] add support for function arguments; ex. different log base

### Supported ast structs
- [ ] BadExpr
  - placeholder for expression containing syntax error, no need to handle this case
- [x] Ident
  - indentifier struct is used to represent package names, function names, variable names, etc.
  - would be good functionality to add for handling variables
- [ ] Ellipsis
  - not needed in mathematical formulas
- [x] BasicLit
  - simplest types, handled INT and FLOAT
- [x] FuncLit
  -  handle functions, basics in place just need to add functions
- [ ] CompositeLit
  - array, map, etc.; not sure where this would come up in formulas
- [x] ParenExpr
  - parenthesis for order of operations
- [ ] SelectorExpr
  - not sure what a selector is
- [ ] IndexExpr
  - used for indexing arrays/slices, not used in math formulas
- [ ] SliceExpr
  - represents slice, not used in math formulas
- [ ] TypeAssertExpr
  - interface type assertion, not used in math formulas
- [x] CallExpr
  - warapper around function and arguments
- [ ] StarExpr
  - for pointer or multiplication, do not need to reimplement for math formulas
- [x] UnaryExpr
  - implemented as wrapper to represent negative number, should investigate other use cases
- [x] BinaryExpr
  - simlpe operator use cases
- [ ] KeyValueExpr
  - tuple used in composite literals
